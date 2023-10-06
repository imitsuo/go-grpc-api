// package main

// import (
// 	api "grpc-api/api/app"
// 	"grpc-api/api/configs"
// )

// var settings = configs.GetSettings()

// func main() {
// 	app := api.MakeApp()

// 	portListen := ":" + settings["PORT"]
// 	err := app.Listen(portListen)
// 	if err != nil {
// 		panic(err)
// 	}
// }
