package auth

import (
	"net/http"
)

type SignUpInfo struct {
	Username string
	Password string
	Email    string
}

func doSignUp(w *http.ResponseWriter, info SignUpInfo) {
	// Do something with the info
}

func SignUpHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var info SignUpInfo
	info.Username = r.FormValue("username")
	if info.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}
	info.Password = r.FormValue("password")
	if info.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}
	info.Email = r.FormValue("email")
	if info.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	doSignUp(&w, info)

}
