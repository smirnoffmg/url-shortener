package main

import "github.com/smirnoffmg/url-shortener/interfaces"

func main() {
	ginServer := interfaces.GetGinServer()
	ginServer.Run(":8000")
}
