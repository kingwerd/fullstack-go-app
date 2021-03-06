package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kingwerd/fullstack-go-app/api/models"
)

// Server holds the router for the API and the connection to the database.
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Intialize takes database specific configuration values to create
// a connection to the database and initializes the application routes.
func (server *Server) Intialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", Dbdriver)
			log.Fatal("This is the error: \n\t", err)
		} else {
			fmt.Printf("We are connected to the %s database.", Dbdriver)
		}
	}
	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// Run starts the server using the addr parameter.
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
