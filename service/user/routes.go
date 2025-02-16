package user

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Seemant-RajSingh/go-crud/config"
	"github.com/Seemant-RajSingh/go-crud/service/auth"
	"github.com/Seemant-RajSingh/go-crud/types"
	"github.com/Seemant-RajSingh/go-crud/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct { // can take any dependencies
	store types.UserStore // interface as a field type
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRouter(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

// -------------------------- REGISTER --------------------------------
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) { // using receivers to create handleFuncs
	// reading raw req body:

	// DONT WORK:
	//fmt.Println("req body: ", r.Body)
	//fmt.Println(r)

	// WORKS:
	bodyBytes, _ := io.ReadAll(r.Body) // 2nd param is error
	fmt.Println("Raw Request Body:", string(bodyBytes))

	// 1. get json payload
	var payload types.RegisterUserPayload
	// fmt.Println(payload)	// {  }
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists: ", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	// create if user dosent exist
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

// -------------------------- LOGIN --------------------------------
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// 1. get json payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found, invalid email or password"))
		return
	}

	if !auth.ComparePasswordS(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// req and resp writer: fmt.Println("w http.ResponseWriter: ", w, "r *http.Request: ", r) gives:
// w http.ResponseWriter:  &{0xc000140090 0xc0001f2280 0xc00006c400 0x21fe60 false false true {{} {0 0}} {{} 0} 0xc00006c440 {0xc0000001c0 map[] false false} map[] false 0 -1 0 false false false [] {{} 0} [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc00006e2a0 {{} 0}} r *http.Request:  &{POST /api/v1/register HTTP/1.1 1 1 map[Accept:[*/*] Accept-Encoding:[gzip, deflate, br] Connection:[close] Content-Length:[101] Content-Type:[application/json] User-Agent:[Thunder Client (https://www.thunderclient.com)]] 0xc00006c400 <nil> 101 [] true localhost:8080 map[] map[] <nil> map[] 127.0.0.1:59702 /api/v1/register <nil> <nil> <nil>  0xc000025aa0 <nil> [] map[]}
