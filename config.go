package main

import (
	App "gorest/Application"
	"gorest/Posts"
)

type Config struct{
	Port int
	Routes App.RouteCollection
}

func newConfig() Config {
	return Config{
		Port: 80,
		Routes: []App.Route{
			{Method: "GET", Pattern: "/", Handler: func(request App.Request, response App.Response){
				response.Send(200, "It Works!")
			}},
			{Method: "GET", Pattern: "/posts", Handler: Posts.GetAllPosts},
			{Method: "POST", Pattern: "/posts", Handler: Posts.CreatePost},
			{Method: "GET", Pattern: "/posts/{id}", Handler: Posts.FindPost},
			{Method: "PUT", Pattern: "/posts/{id}", Handler: Posts.UpdatePost},
			{Method: "DELETE", Pattern: "/posts/{id}", Handler: Posts.DeletePost},
		},
	}
}
