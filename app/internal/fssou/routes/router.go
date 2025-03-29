package routes

// Router Struct para armazenar as rotas e seus handlers
type Router struct {
	routes map[string]map[string]RouteHandler
}

// NewRouter Cria um novo roteador
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[string]RouteHandler),
	}
}

// AddRoute Adiciona uma nova rota ao roteador
func (r *Router) AddRoute(method, path string, handler RouteHandler) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]RouteHandler)
	}
	r.routes[method][path] = handler
}
