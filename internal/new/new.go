package new


type Site struct {
    Name string
}



func NewSite() *Site {           
    return &Site{
        Name: "",
    }
}




