package salary

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_Accounts_Create(t *testing.T) {
	tests := []struct{
		body string
		expected string
		status int
	}{
		{"{\"name\": \"account 1\"}", `{"id":1,"name":"account 1"}`, http.StatusCreated},
		{"{\"naame\": \"account 1\"}", "json: unknown field \"naame\"\n", http.StatusBadRequest},
	}

	deleteDatabase()
	system := Initialize()
	server := Server{system}

	for i, test := range tests {
		reader := strings.NewReader(test.body)
		req, err := http.NewRequest("POST", "/accounts", reader)
		if err != nil {
			t.Fatalf("Failed test %v: %v", i, err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Accounts)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, test.status, rr.Code)
		assert.Equal(t, test.expected, rr.Body.String())
	}
}

func TestServer_Accounts_GET(t *testing.T) {
	deleteDatabase()
	system := Initialize()
	server := Server{system}

	req, err := http.NewRequest("GET", "/accounts", nil)
	if err != nil {
		t.Fatalf("Failed test %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Accounts)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestServer_AccountTransaction_Create(t *testing.T) {
	deleteDatabase()
	system := Initialize()
	server := Server{system}

	createAccount(t, server)
	createTransaction(t, server, 1)

	account, err := system.GetAccount(1)
	assert.Nil(t, err)
	assert.Len(t, account.Transactions(), 1)
}

func TestServer_Transaction(t *testing.T) {
	deleteDatabase()
	system := Initialize()
	server := Server{system}

	createAccount(t, server)
	createTransaction(t, server, 1)

	account, err := system.GetAccount(1)
	assert.Nil(t, err)
	assert.Len(t, account.Transactions(), 1)

	deleteTransaction(t, server, 1)
	assert.Len(t, account.Transactions(), 0)
}


func createAccount(t *testing.T, server Server) {
	reader := strings.NewReader("{\"name\": \"account 1\"}")
	req, err := http.NewRequest("POST", "/accounts", reader)
	if err != nil {
		t.Fatalf("Failed test %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Accounts)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}


func createTransaction(t *testing.T, server Server, account int) {
	reader := strings.NewReader("{\"description\": \"Description 1\", \"amount\":100}")
	url := fmt.Sprintf("/accounts/%v/transactions", account)
	log.Printf("URL: %v", url)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		t.Fatalf("Failed test %v", err)
	}

	rr := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/accounts/{id}/transactions", server.AccountTransaction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func deleteTransaction(t *testing.T, server Server, transaction int) {
	url := fmt.Sprintf("/transactions/%v", transaction)
	log.Printf("URL: %v", url)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Failed test %v", err)
	}

	rr := httptest.NewRecorder()
	handler := mux.NewRouter()
	handler.HandleFunc("/transactions/{id}", server.Transaction)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}