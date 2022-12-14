package main

import (
	"fmt"
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
	"github.com/Joacker/Ayu1/backend/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	headers := handlers.AllowedHeaders([]string{"X-Request-Width", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/champions", routes.GetAgentsHandler).Methods("GET")
	r.HandleFunc("/champions/{Id}", routes.GetAgentHandler).Methods("GET")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{Id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users2", routes.GetUserHandler2).Methods("GET")
	r.HandleFunc("/users3", routes.GetUserHandler3).Methods("POST")
	r.HandleFunc("/logout", routes.Logout_db).Methods("POST")
	r.HandleFunc("/logeados", routes.Logeados).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/resenias", routes.GetReseniasHandler).Methods("POST")
	r.HandleFunc("/postresenias", routes.PostReseniaHandler).Methods("POST")
	r.HandleFunc("/putresenias", routes.PutReseniaHandler).Methods("PUT")
	r.HandleFunc("/deleteresenias", routes.DeleteReseniaHandler).Methods("DELETE")
	r.HandleFunc("/miresenia", routes.GetReseniaHandler).Methods("POST")
	r.HandleFunc("/allresenias", routes.GetAllReseniasHandler).Methods("POST")

	for {
		db.DBConnection()
		db.DB.AutoMigrate(models.Agents{}, models.Users{}, models.Resenias{})
		fmt.Println("Hello World 3")
		http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(r))

	}

}
