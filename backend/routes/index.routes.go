package routes

import (
	"net/http"

	"github.com/Joacker/Ayu1/backend/db"
	"github.com/Joacker/Ayu1/backend/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()
	db.DB.AutoMigrate(models.Agents{}, models.Users{}, models.Resenias{})

	w.Write([]byte("Hello World3"))
}
