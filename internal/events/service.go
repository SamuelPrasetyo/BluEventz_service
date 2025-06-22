/* Berisi logika bisnis */

package events

func GetEventsService() ([]MEvent, error) {
	return GetAllEventsRaw()
}
