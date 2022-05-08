package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restapi/internal/apperror"
	"restapi/internal/handlers"
	"restapi/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersUrl = "/users"
	userUrl  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersUrl, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, userUrl, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userUrl, apperror.Middleware(h.GetUserByUuID))
	router.HandlerFunc(http.MethodPut, userUrl, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userUrl, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userUrl, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is list of users"))
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusCreated)
	_, err := w.Write([]byte("this is create user"))
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) GetUserByUuID(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is get user"))
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is update user"))
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("this is partially update user"))
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNoContent)
	_, err := w.Write([]byte("this is delete user"))
	if err != nil {
		return err
	}

	return nil
}
