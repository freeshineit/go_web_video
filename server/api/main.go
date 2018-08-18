package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	// user
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:user_name", UserInfo)
	router.DELETE("/user/:user_name", DeleteUser)

	// video
	router.GET("/user/:user_name/videos", Videos)
	router.GET("/user/:user_name/videos/:vid_id", Video)
	router.DELETE("/user/:user_name/videos/:vid_id", DeleteVideo)

	// comments
	router.GET("/videos/:vid_id/comments", GetComments)
	router.POST("/videos/:vid_id/comments", PostComments)
	router.DELETE("/videos/:vid_id/comments/:comments_id", DeleteComments)

	return router
}

func main() {
	r := RegisterHandlers()

	http.ListenAndServe(":8888", r)
}
