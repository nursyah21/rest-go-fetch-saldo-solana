package models

import (
	"github.com/kamva/mgm/v3"
)

type API struct {
	mgm.DefaultModel `bson:",inline"`
	Api              string `json:"api" bson:"api"`
}
