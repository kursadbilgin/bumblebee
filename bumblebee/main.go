package main

import (
	"bumblebee/Lib"
	"bumblebee/Routers"
)

func main() {
	// Create redis client
	redisClient := Lib.RedisClient()
	// Run implements a http.ListenAndServe() and takes in an optional Port number
	// The default port is :8080
	r := Routers.SetupRouter(redisClient)
	// running
	r.Run()
}
