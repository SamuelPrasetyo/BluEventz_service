package service

import (
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/model"
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/repository"
)

func GetEventsService() ([]model.Event, error) {
	return repository.FetchEvents()
}
