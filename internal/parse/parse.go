package parse

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
    "time"
)



func ParseRouterLog(path string) error {
    fmt.Printf ( "Read router log at %s\n", path )

    entries, err := os.ReadDir(path)
    if err != nil {
        return err
    }
    for _, entry := range entries {
        name := entry.Name()
	if strings.HasPrefix(name, "router-logs") {
	    router_log_path := path + "/" + name
            fmt.Printf ( "ParseRouterLog: %s\n", router_log_path )
	    findTopologyCalcs ( router_log_path )
	}
    }
    return nil
}



func findTopologyCalcs(filename string) error {
    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("error opening file %s: %w", filename, err)
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    lineNum := 1
    for scanner.Scan() {
	line := scanner.Text()
	if strings.Contains(line, "Computed next hops") {
	    fmt.Printf("Line %d: %s\n", lineNum, line)
	    entry, err := parseRouterLog(line)   // TODO change name of fn
	    if err != nil {
		fmt.Println("findTopologyCalcs error: %v", err)
		continue
	    }
	    fmt.Printf("    Time: %s\n", entry.Timestamp.Format(time.RFC3339Nano))
	    //fmt.Printf("    Map:  %+v\n", entry.MapData)
	    if 0 == len(entry.MapData) {
	        fmt.Printf("    Map:  empty\n")
	    } else {
	        fmt.Printf("    Map:  \n")
	        for key, value := range entry.MapData {
                    fmt.Println("        ", key, "->", value)
	        }
	    }

	    fmt.Println("------------")
	}
        lineNum++
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file %s: %w", filename, err)
    }

    return nil
}



// Grok ---------------------------------------------



type LogEntry struct {
	Timestamp time.Time
	MapData   map[string]string
	Message   string // e.g. "ROUTER_LS (info) Computed next hops"
	RawLine   string
}

// parseRouterLog now correctly handles both normal and empty maps
func parseRouterLog(line string) (*LogEntry, error) {
	// Improved regex: captures timestamp + everything after it
	re := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{6} [+-]\d{4})\s+(.*)$`)
	matches := re.FindStringSubmatch(line)
	if len(matches) != 3 {
		return nil, fmt.Errorf("failed to parse timestamp from line")
	}

	timestampStr := matches[1]
	rest := matches[2]

	// Parse timestamp
	t, err := time.Parse("2006-01-02 15:04:05.000000 -0700", timestampStr)
	if err != nil {
		return nil, fmt.Errorf("timestamp parse error: %w", err)
	}

	// Find the map by looking for the last occurrence of ": {" or ": {}"
	mapStart := strings.LastIndex(rest, ": {")
	if mapStart == -1 {
		// Try alternative format with space after colon
		mapStart = strings.LastIndex(rest, ": {}")
		if mapStart != -1 {
			mapStart += 1 // adjust to include the colon
		}
	}

	var mapContent string
	message := strings.TrimSpace(rest)

	if mapStart != -1 {
		message = strings.TrimSpace(rest[:mapStart])
		mapPart := strings.TrimSpace(rest[mapStart+1:]) // +1 to skip the ":"

		// Extract inside braces
		openBrace := strings.Index(mapPart, "{")
		closeBrace := strings.LastIndex(mapPart, "}")
		if openBrace != -1 && closeBrace != -1 && closeBrace > openBrace {
			mapContent = strings.TrimSpace(mapPart[openBrace+1 : closeBrace])
		}
	}

	// Parse the map (empty content = empty map)
	data, err := parsePythonDict(mapContent)
	if err != nil {
		return nil, fmt.Errorf("map parse error: %w", err)
	}

	return &LogEntry{
		Timestamp: t,
		MapData:   data,
		Message:   message,
		RawLine:   line,
	}, nil
}

// parsePythonDict handles empty string gracefully → returns empty map
func parsePythonDict(content string) (map[string]string, error) {
	result := make(map[string]string)
	content = strings.TrimSpace(content)
	if content == "" {
		return result, nil
	}

	parts := splitByTopLevelComma(content)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		colonIdx := strings.Index(part, ":")
		if colonIdx == -1 {
			continue
		}

		key := strings.TrimSpace(part[:colonIdx])
		value := strings.TrimSpace(part[colonIdx+1:])

		key = strings.Trim(key, `'"`)
		value = strings.Trim(value, `'"`)

		if key != "" {
			result[key] = value
		}
	}
	return result, nil
}

// splitByTopLevelComma - unchanged (handles commas inside quotes safely)
func splitByTopLevelComma(s string) []string {
	var parts []string
	var current strings.Builder
	inQuote := false
	quoteChar := rune(0)

	for _, r := range s {
		if (r == '\'' || r == '"') && !inQuote {
			inQuote = true
			quoteChar = r
		} else if r == quoteChar && inQuote {
			inQuote = false
			quoteChar = 0
		}

		if r == ',' && !inQuote {
			parts = append(parts, current.String())
			current.Reset()
		} else {
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		parts = append(parts, current.String())
	}
	return parts
}

// Test it
/*
func main() {
	testLines := []string{
		`2025-09-09 14:03:45.707503 +0000 ROUTER_LS (info) Computed next hops: {'prd-wyn-skupper-router-7bf8f6bfdf-nnbw6': 'prd-wyn-skupper-router-7bf8f6bfdf-nnbw6'}`,
		`2025-09-09 14:04:12.123456 +0000 ROUTER_LS (info) Computed next hops: {}`,
		`2025-09-09 14:05:00.000000 +0000 ROUTER_LS (info) Computed next hops: {}`,
		`2025-09-09 14:06:30.555555 +0000 ROUTER_LS (info) Some other message without map`,
	}

	for _, line := range testLines {
		entry, err := parseRouterLog(line)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Time: %s\n", entry.Timestamp.Format("2006-01-02 15:04:05.000000"))
		fmt.Printf("Message: %s\n", entry.Message)
		fmt.Printf("Map: %+v\n", entry.MapData)
		fmt.Println("---")
	}
}
*/
