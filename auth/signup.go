package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "etextbook"
)

type SignUpInfo struct {
	Username string
	Password string
	Email    string
}

func doSignUp(w *http.ResponseWriter, info SignUpInfo) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		http.Error(*w, "Internal server error", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

	defer db.Close()

	insertStmt := `INSERT INTO users (name, password, email) VALUES ( $1, $2, $3 )`
	fmt.Println(insertStmt)
	_, err = db.Exec(insertStmt, info.Username, info.Password, info.Email)
	if err != nil {
		http.Error(*w, "Internal server error", http.StatusInternalServerError)
		fmt.Print(err)
		return
	}

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
