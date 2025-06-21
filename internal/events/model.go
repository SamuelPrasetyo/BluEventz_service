package events

import "time"

type Event struct {
	IDEvent          string    `json:"id_event" gorm:"column:idevent"`
	NamaEvent        string    `json:"nama_event" gorm:"column:cnmevent"`
	Tanggal          time.Time `json:"tanggal" gorm:"column:dtgl"`
	Jam              string    `json:"jam" gorm:"column:cjam"`
	Deskripsi        string    `json:"deskripsi" gorm:"column:cdesc"`
	NamaModerator    string    `json:"nama_moderator" gorm:"column:cnmmoderator"`
	JabatanModerator string    `json:"jabatan_moderator" gorm:"column:cjabatanmd"`
	NamaNarasumber1  string    `json:"nama_narasumber_1" gorm:"column:cnmnarsum1"`
	JabatanNarsum1   string    `json:"jabatan_narsum_1" gorm:"column:cjabatannarsum1"`
	NamaNarasumber2  string    `json:"nama_narasumber_2" gorm:"column:cnmnarsum2"`
	JabatanNarsum2   string    `json:"jabatan_narsum_2" gorm:"column:cjabatannarsum2"`
	Link             string    `json:"link" gorm:"column:clink"`
}
