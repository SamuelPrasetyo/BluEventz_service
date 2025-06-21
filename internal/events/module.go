package events

import (
	"github.com/SamuelPrasetyo/BluEventz_service/config"
	"github.com/SamuelPrasetyo/BluEventz_service/internal/events"
)

func GetAllEventsRaw() ([]events.Event, error) {
	var results []events.Event

	query := `
		SELECT 
			idevent, cnmevent, dtgl, cjam, cdesc,
			cnmmoderator, cjabatanmd,
			cnmnarsum1, cjabatannarsum1,
			cnmnarsum2, cjabatannarsum2,
			clink
		FROM mevents
	`

	// Jalankan query raw GORM dan scan ke slice struct
	err := config.DB.Raw(query).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
