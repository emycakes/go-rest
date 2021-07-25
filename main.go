package main

import (
	App "gorest/Application"
)

func main() {
	config := newConfig()
	router := App.NewRouter()
	router.RegisterRoutes(config.Routes)
	router.Listen(config.Port)
}