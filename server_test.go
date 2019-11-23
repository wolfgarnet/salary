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
	}{
		{"{\"name\": \"account 1\"}"},
	}
	
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	deleteDatabase()
	system := Initialize()
	server := Server{system}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Accounts)

	for i, test := range tests {
		reader := strings.NewReader("{\"name\": \"account 1\"}")
		req, err := http.NewRequest("POST", "/accounts", reader)
		if err != nil {
			t.Fatal(err)
		}


		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		expected := `{"id":1,"name":"account 1"}`
		assert.Equal(t, expected, rr.Body.String())
	}
}
