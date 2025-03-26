package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin: "html"`
	TargetDate  time.Time
	Progress    int `uadmin: "progress_bar"`
}
