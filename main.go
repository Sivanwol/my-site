package main

import (
	routes "main/routes"
	utils "main/utils"
)

// swagger middleware for Iris
// swagger embed files

func newApp() *utils.Bootstrapper {
	app := utils.New("my site", "sivan wolberg")
	//var serverCnf iris.Configuration = cnf["server"]
	app.Bootstrap()
	app.Configure(routes.Configure)
	return app
}

func main() {
	cnf := utils.GetConfiguration()
	app := newApp()
	port := cnf["server"].Other["Addr"].(string)
	app.Listen(port)
}
