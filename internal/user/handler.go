package user

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-tutorial/internal/apperror"
	"rest-api-tutorial/internal/handlers"
	"rest-api-tutorial/pkg/logging"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/api/v1/users"
	userURL  = "/api/v1/users/:uuid"
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
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetUserList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) GetUserList(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(200)
	//w.Write([]byte("This is list of users"))

	return apperror.ErrNotFound
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(201)
	//w.Write([]byte("This is create an user"))

	return fmt.Errorf("this is API error")
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error {
	//w.WriteHeader(200)
	//w.Write([]byte("This is user by uuid"))

	return apperror.NewAppError(nil, "test", "test", "t13")
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("This is fully update of users"))

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("This is partially update of users"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("This is delete of users"))

	return nil
}
