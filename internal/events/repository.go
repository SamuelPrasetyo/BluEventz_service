package events

import (
	"github.com/SamuelPrasetyo/BluEventz_service/internal/events"
)

func FetchEvents() ([]events.Event, error) {
	return events.GetAllEventsRaw()
}
