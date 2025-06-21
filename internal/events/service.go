package events

import (
	"github.com/SamuelPrasetyo/BluEventz_service/internal/events"
)

func GetEvent() ([]events.Event, error) {
	return events.FetchEvents()
}
