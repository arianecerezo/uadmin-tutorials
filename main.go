package main

import (


	"github.com/arianecerezo/uadmin-tutorials/models"
	"github.com/uadmin/uadmin"
)


func main() {
	uadmin.Register(
		models.Todo{},
	)
	uadmin.StartServer()
}
