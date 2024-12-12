package main

import "github.com/Marlliton/go/crud-com-auth-jwt/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config)
}
