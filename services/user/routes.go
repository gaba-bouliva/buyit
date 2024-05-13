package user

import (
	"fmt"
	"net/http"

	"github.com/gaba-bouliva/buyit/services/user/auth"
	"github.com/gaba-bouliva/buyit/types"
	"github.com/gaba-bouliva/buyit/utils"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Get("/test", h.handleTest)
	router.Post("/login", h.handleLogin)
	router.Post("/register", h.handleRegister)
}

func (h *Handler) handleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handle Test works!")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// TODO

	// get Json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if the user exist
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

		// if user doesnt exist we create new user
		err = h.store.CreateUser(types.User{
			FirstName: payload.FirstName,
			LastName: payload.LastName,
			Email: payload.Email,
			Password: hashedPassword,
		})
}