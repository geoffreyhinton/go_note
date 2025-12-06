package api

import (
	"encoding/json"
	"net/http"

	"github.com/geoffreyhinton/go_note/common/error"
	"github.com/geoffreyhinton/go_note/store"
	"github.com/gorilla/mux"
)

type UserSignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleUserSignUp(w http.ResponseWriter, r *http.Request) {
	var userSignUp UserSignUp
	err := json.NewDecoder(r.Body).Decode(&userSignUp)
	if err != nil {
		error.ErrorHandler(w, "REQUEST_BODY_ERROR")
		return
	}
	user, err := store.CreateNewUser(userSignup.Username, userSignup.Password, "", "")

	if err != nil {
		error.ErrorHandler(w, "")
		return
	}

	json.NewEncoder(w).Encode(user)
}

type UserSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func handleUserSignIn(w http.ResponseWriter, r *http.Request) {
	var userSignin UserSignin
	err := json.NewDecoder(r.Body).Decode(&userSignin)

	if err != nil {
		error.ErrorHandler(w, "")
		return
	}

	user, err := store.GetUserByUsernameAndPassword(userSignin.Username, userSignin.Password)

	if err != nil {
		error.ErrorHandler(w, "")
		return
	}

	userIdCookie := &http.Cookie{
		Name:   "user_id",
		Value:  user.Id,
		MaxAge: 3600 * 24 * 30,
	}
	http.SetCookie(w, userIdCookie)

	json.NewEncoder(w).Encode(user)
}

func handleUserSignOut(w http.ResponseWriter, r *http.Request) {
	userIdCookie := &http.Cookie{
		Name:   "user_id",
		Value:  "",
		MaxAge: 0,
	}
	http.SetCookie(w, userIdCookie)
}

func RegisterAuthRoutes(r *mux.Router) {
	authRouter := r.PathPrefix("/api/auth").Subrouter()

	authRouter.HandleFunc("/signup", handleUserSignUp).Methods("POST")
	authRouter.HandleFunc("/signin", handleUserSignIn).Methods("POST")
	authRouter.HandleFunc("/signout", handleUserSignOut).Methods("POST")
}
