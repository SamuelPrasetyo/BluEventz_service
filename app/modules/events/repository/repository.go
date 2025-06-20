package repository

import (
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/model"
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/module"
)

func FetchEvents() ([]model.Event, error) {
	return module.GetAllEventsRaw()
}
