package salary

import (
	"github.com/stretchr/testify/assert"
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
