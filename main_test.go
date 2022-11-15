package main_test

import (
	"log"
	"os"
	"testing"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"github.com/TomFern/go-mux-api"
)

var a main.App

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM balance")
	a.DB.Exec("DELETE FROM services")
	a.DB.Exec("DELETE FROM reservations")
	a.DB.Exec("DELETE FROM transactions")
	// a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `

CREATE TABLE IF NOT EXISTS public.balance (
    user_id integer NOT NULL,
    ruble_balance numeric NOT NULL,
    CONSTRAINT check_positive CHECK ((ruble_balance >= (0)::numeric)),
	CONSTRAINT balance_pk PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS public.reservations (
    reservation_id integer NOT NULL,
    user_id integer NOT NULL,
    service_id integer NOT NULL,
    cost numeric NOT NULL,
    reservation_time timestamp without time zone,
	CONSTRAINT reservations_pk PRIMARY KEY (reservation_id),
	CONSTRAINT reservations_fk0 FOREIGN KEY (user_id) REFERENCES public.balance(user_id)
	--CONSTRAINT reservations_fk1 FOREIGN KEY (service_id) REFERENCES public.services(service_id)
);

-- CREATE TABLE IF NOT EXISTS public.services (
--     service_id integer NOT NULL,
--     service_name character varying(255) NOT NULL,
--     CONSTRAINT services_pk PRIMARY KEY (service_id)
-- );

CREATE TABLE IF NOT EXISTS public.transactions (
    transaction_id integer NOT NULL,
    date timestamp without time zone NOT NULL,
    amount numeric NOT NULL,
    user_id integer NOT NULL,
    service_id integer NOT NULL,
	CONSTRAINT transactions_pk PRIMARY KEY (transaction_id),
	CONSTRAINT transactions_fk0 FOREIGN KEY (user_id) REFERENCES public.balance(user_id)
	--CONSTRAINT transactions_fk1 FOREIGN KEY (service_id) REFERENCES public.services(service_id)
);

`

// func TestEmptyTable(t *testing.T) {
// 	clearTable()

// 	req, _ := http.NewRequest("GET", "/balance/show/1", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetNonExistentProduct(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/balance/show/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func TestCreateProduct(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"user_id":1, "ruble_balance": 11.22}`)
	req, _ := http.NewRequest("PUT", "/balance/deposit", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["ruble_balance"] != 11.22 {
		t.Errorf("Expected product price to be '11.22'. Got '%v'", m["ruble_balance"])
	}

	if m["user_id"] != 1.0 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", m["user_id"])
	}
}

func TestGetProduct(t *testing.T) {
	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/balance/show/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func addProducts(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO balance(user_id, ruble_balance) VALUES($1, $2)", strconv.Itoa(i + 1), (i+1.0)*10)
	}
}

func TestUpdateProduct(t *testing.T) {

	clearTable()
	addProducts(1)

	req, _ := http.NewRequest("GET", "/balance/show/1", nil)
	response := executeRequest(req)
	var originalWallet map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalWallet)

	var jsonStr = []byte(`{"user_id": 1, "ruble_balance": 10.0}`)
	req, _ = http.NewRequest("PUT", "/balance/deposit", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["user_id"] != originalWallet["user_id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalWallet["user_id"], m["user_id"])
	}
    
	if m["ruble_balance"] == originalWallet["ruble_balance"] {
		t.Errorf("Expected the price to change from '%v' to '%v'. Got '%v'", "20", m["ruble_balance"], m["ruble_balance"])
	}
}
