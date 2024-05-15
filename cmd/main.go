package main

import "github.com/smirnoffmg/url-shortener/internal/interfaces"

func main() {
	ginServer := interfaces.GetGinServer()
	ginServer.Run(":8000")
}
