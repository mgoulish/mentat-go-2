package new


type Router struct {
    Name string
    Data map[string]interface{}
}


type Site struct {
    Name   string
    Path   string
    Router *Router
}



func NewRouter() *Router {
    return &Router{
	       Data:  make(map[string]interface{}),
           }        
}


func NewSite() *Site {
    return &Site{
        Router: NewRouter(),
    }
}
