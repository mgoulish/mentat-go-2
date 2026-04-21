package new


type Router struct {
    Name string
}


type Site struct {
    Name   string
    Path   string
    Router *Router
}



func NewRouter() *Router {
    return &Router{}        
}


func NewSite() *Site {
    return &Site{
        Router: NewRouter(),
    }
}
