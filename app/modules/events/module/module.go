package module

import (
	"github.com/SamuelPrasetyo/BluEventz_service/app/modules/events/model"
	"github.com/SamuelPrasetyo/BluEventz_service/config"
)

func GetAllEventsRaw() ([]model.Event, error) {
	var results []model.Event

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
