package main

import (
	"log"
	"real-time-forum/database"
	"real-time-forum/server"
)

func main() {
	db, erreur := database.InitDB("user.db")
	if erreur != nil {
		log.Fatal("Error to create DB")
	}
	erreur = database.CreateTableUser(db)
	if erreur != nil {
		log.Fatal("Error to create tableUser, main")
	}
	router := server.InitRoutes()
	srv := server.ServerParameters(router)

	log.Println("Server start at port 8080...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("listenAndServ :", err)
	}
}
