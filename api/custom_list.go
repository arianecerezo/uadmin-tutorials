package api

import (
	"net/http"

	"github.com/arianecerezo/uadmin-tutorials/models"
	"github.com/uadmin/uadmin"
)

func CustomListAPIHandler(w http.ResponseWriter, r *http.Request) {
	todo := []models.Todo{}

	results := []map[string]interface{}{}

	uadmin.AdminPage("id", false, 0, 5, &todo, "")

	for i := range todo {
		uadmin.Preload(&todo[i])

		results = append(results, map[string]interface{}{
			"ID":          todo[i].ID,
			"Name":        todo[i].Name,
			"Description": todo[i].Description,
			"Category":    todo[i].Category.Name,
			"Friend":      todo[i].Friend.Name,
			"Item":        todo[i].Item.Name,
			"TargetDate":  todo[i].TargetDate,
			"Progress":    todo[i].Progress,
		})
	}
	uadmin.ReturnJSON(w, r, results)
}
