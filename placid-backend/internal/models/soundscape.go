package models

type Soundscape struct {
	Id       string `json:"id"`
	Mood     string `json:"mood"`
	Name     string `json:"name"`
	AudioUrl string `json:"audioUrl"`
}
