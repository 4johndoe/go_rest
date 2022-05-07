package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/users", h.GetList)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is list of users"))
}
