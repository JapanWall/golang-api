package app

import (
	"fmt"
	"log"
	"net/http"
	"postapi/app/models"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Post API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("cannot parse post body. err= %v \n", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		//create post
		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err%v", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		resp := MapPostToJson(p)
		SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) DeletePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("cannot parse post body. err= %v \n", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		//create post
		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		err = a.DB.DeletePost(p)
		if err != nil {
			log.Printf("delete post in DB. err%v", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		resp := MapPostToJson(p)
		SendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) UpdatePutHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("cannot parse post body. err= %v \n", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		//create post
		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		err = a.DB.UpdatePut(p)
		if err != nil {
			log.Printf("Update put in DB. err%v", err)
			SendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		resp := MapPostToJson(p)
		SendResponse(w, r, resp, http.StatusOK)
	}
}
