package views

import (
	"html/template"
	"net/http"

	"github.com/arianecerezo/uadmin-tutorials/models"
	"github.com/uadmin/uadmin"
)

type Context struct {
	TodoList []map[string]interface{}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	c := Context{}
	todo := []models.Todo{}
	uadmin.All(&todo)

	for i := range todo {
		uadmin.Preload(&todo[i])

		c.TodoList = append(c.TodoList, map[string]interface{}{
			"ID":          todo[i].ID,
			"Name":        todo[i].Name,
			"Description": template.HTML(todo[i].Description),
			"Category":    todo[i].Category.Name,
			"Friend":      todo[i].Friend.Name,
			"Item":        todo[i].Item.Name,
			"TargetDate":  todo[i].TargetDate,
			"Progress":    todo[i].Progress,
		})
	}
	uadmin.RenderHTML(w, r, "templates/todo.html", c)
}
