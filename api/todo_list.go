package api

import (
	"net/http"

	"github.com/arianecerezo/uadmin-tutorials/models"
	"github.com/uadmin/uadmin"
)

// TodoListAPIHandler !
func TodoListAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all records in the database
	todo := []models.Todo{}
	uadmin.All(&todo)

	// Accesses and fetches data from another model
	for t := range todo {
		uadmin.Preload(&todo[t])
	}

	// Return todo JSON object
	uadmin.ReturnJSON(w, r, todo)
}
