package server

import "os"

//Init initialize server
func Init() {
	e := NewRouter()
	port := os.Getenv("APP_PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
