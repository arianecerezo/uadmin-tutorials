package main

import (
	"net/http"

	"github.com/arianecerezo/uadmin-tutorials/api"
	"github.com/arianecerezo/uadmin-tutorials/models"
	"github.com/arianecerezo/uadmin-tutorials/views"
	"github.com/uadmin/uadmin"
)

func main() {
	// Register Models
	uadmin.Register(
		models.Todo{},
		models.Category{},
		models.Friend{},
		models.Item{},
	)

	// API Handler
	http.HandleFunc("/api/", uadmin.Handler(api.Handler))
	http.HandleFunc("/page/", uadmin.Handler(views.PageHandler))
	// Register Inlines
	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Item{}, map[string]string{
		"Todo": "ItemID",
	})

	// Call InitializeRootURL function to change the RootURL value in the Settings model.
	InitializeRootURL()

	uadmin.StartServer()
}

func InitializeRootURL() {
	// Initialize Setting model
	setting := uadmin.Setting{}

	// Get the code
	uadmin.Get(&setting, "code = ?", "uAdmin.RootURL")

	// Assign the value as "/admin/"
	setting.ParseFormValue([]string{"/admin/"})

	// Save changes
	setting.Save()
}
