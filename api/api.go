package api

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// Handler !
func Handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/api/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
	uadmin.Trail(uadmin.DEBUG, "api")
	if strings.HasPrefix(r.URL.Path, "/todo_list") {
		uadmin.Trail(uadmin.DEBUG, "TODO")
		TodoListAPIHandler(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/custom_list") {
		uadmin.Trail(uadmin.DEBUG, "custom")
		CustomListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/add_friend") {
		uadmin.Trail(uadmin.DEBUG, "add-rfirend")
		AddFriendAPIHandler(w, r)
		return
	}
}
