/* Berisi logika bisnis */

package events

import (
	"github.com/SamuelPrasetyo/BluEventz_service/internal/events"
)

func GetEventsService() ([]events.MEvent, error) {
	return events.GetAllEventsRaw()
}
