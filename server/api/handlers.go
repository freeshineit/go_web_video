package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

// user
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "CreateUser Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")
	io.WriteString(w, name+" login")
}

func UserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")
	io.WriteString(w, "get user:"+name+" info")
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")
	io.WriteString(w, "delete user:"+name)
}

// video
func Videos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")

	io.WriteString(w, "get "+name+"`s videos list")
}

func Video(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")
	vid := p.ByName("vid_id")

	io.WriteString(w, "play user:"+name+" video id:"+vid+" video")
}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("user_name")
	vid := p.ByName("vid_id")

	io.WriteString(w, "delete user:"+name+" video id:"+vid+" video")
}

// comments
func GetComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid_id")

	io.WriteString(w, "get video id:"+vid+" comments")
}

func PostComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid_id")

	io.WriteString(w, "post a comments to video id:"+vid)
}

func DeleteComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid_id")
	cid := p.ByName("comments_id")

	io.WriteString(w, "delete video id:"+vid+" comments id:"+cid)
}
