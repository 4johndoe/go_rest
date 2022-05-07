package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi/internal/handlers"
	"restapi/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, h.GetList)
	router.POST(userUrl, h.CreateUser)
	router.GET(userUrl, h.GetUserByUuID)
	router.PUT(userUrl, h.UpdateUser)
	router.PATCH(userUrl, h.PartiallyUpdateUser)
	router.DELETE(userUrl, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is list of users"))
	if err != nil {
		return
	}
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("this is create user"))
	if err != nil {
		return
	}
}

func (h *handler) GetUserByUuID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is get user"))
	if err != nil {
		return
	}
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is update user"))
	if err != nil {
		return
	}
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is partially update user"))
	if err != nil {
		return
	}
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte("this is delete user"))
	if err != nil {
		return
	}
}
