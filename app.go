// app.go

package main

import (
	"database/sql"
	"log"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"123",
		"balance-db",
		"5432",
		"service")

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// показать баланс, OK
func (a *App) showBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	wal := wallet{ID: id}
	if err := wal.getWallet(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, wal)
}

// 	a.Router.HandleFunc("/products", a.getProducts).Methods("GET")
// 	a.Router.HandleFunc("/product", a.createProduct).Methods("POST")
// 	a.Router.HandleFunc("/product/{id:[0-9]+}", a.getProduct).Methods("GET")
// 	a.Router.HandleFunc("/product/{id:[0-9]+}", a.updateProduct).Methods("PUT")
// 	a.Router.HandleFunc("/product/{id:[0-9]+}", a.deleteProduct).Methods("DELETE")
// }

// резервация рублей
func (a *App) reserveRubles(w http.ResponseWriter, r *http.Request) {
	var res_q reserveQuery
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&res_q); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := res_q.makeReservation(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, res_q)
}

// подтверждение резервации
func (a *App) reserveAccept(w http.ResponseWriter, r *http.Request) {
	var res_q reserveQuery
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&res_q); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := res_q.confirmReservation(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, res_q)
}

// пополнение баланса
func (a *App) depositRubles(w http.ResponseWriter, r *http.Request) {

	var wal wallet
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wal); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()

	if err := wal.updateBalance(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wal)
}

func (a *App) initializeRoutes() {
	// документация TODO
	// основные запросы
	a.Router.HandleFunc("/reservation", a.reserveRubles).Methods("POST")
	a.Router.HandleFunc("/balance/deposit", a.depositRubles).Methods("PUT")
	a.Router.HandleFunc("/reservation/accept", a.reserveAccept).Methods("PUT")
	a.Router.HandleFunc("/balance/show/{id:[0-9]+}", a.showBalance).Methods("GET")
	// дополнительные запросы
}
