package Application

type Route struct {
	Method string
	Pattern string
	Handler func(Request, Response)
}

type RouteCollection []Route
