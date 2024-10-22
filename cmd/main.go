package main

import (
	"github.com/assaidy/bookstore/internals/server"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	server := server.NewFiberServer()
	server.RegisterRoutes()
	port := ":" + os.Getenv("PORT")
	if err := server.Listen(port); err != nil {
		log.Fatal("couldn't start the server. error:", err)
	}
}
