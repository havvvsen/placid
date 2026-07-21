package models

type Soundscape struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Mood     string `json:"mood"`
	AudioUrl string `json:"audioUrl"`
}
