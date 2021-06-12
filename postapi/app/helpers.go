package app

import (
	"encoding/json"
	"log"
	"net/http"
	"postapi/app/models"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

//allow responses to be displayed when it shows
func SendResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("cannot format json. err%v", err)
	}
}

func MapPostToJson(p *models.Post) models.JsonPost {
	return models.JsonPost{
		ID:      0,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}

}
