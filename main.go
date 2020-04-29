package main

import {
	"engine"
}

func main() {
	linkLimit := 25
	url := "https://www.yahoo.com"
	
	engine.Run(url, linkLimit)
}




