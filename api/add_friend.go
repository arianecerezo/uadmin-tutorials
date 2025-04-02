package api

import (
	"net/http"
	"strconv"

	"github.com/arianecerezo/uadmin-tutorials/models"
	"github.com/uadmin/uadmin"
)

func AddFriendAPIHandler(w http.ResponseWriter, r *http.Request) {
	friend := models.Friend{}

	friendName := r.FormValue("name")
	friendEmail := r.FormValue("email")
	friendPassword := r.FormValue("password")
	friendNationalityRaw := r.FormValue("nationality")

	friendNationality, err := strconv.Atoi(friendNationalityRaw)

	if friendName == "" {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Name is required.",
		})
		return
	}

	friend.Name = friendName
	friend.Email = friendEmail
	friend.Password = friendPassword
	friend.Nationality = models.Nationality(friendNationality)

	err = uadmin.Save(&friend)

	if err != nil {
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "error",
			"err_msg": "Error saving the database : " + err.Error(),
		})
		return
	}
	uadmin.ReturnJSON(w, r, map[string]interface{}{
		"status": "ok",
	})
}
