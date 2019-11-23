package salary

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func runServer(system *System) {
	server := Server{system}
	router := mux.NewRouter()

	router.HandleFunc("/accounts", server.accounts)
	router.HandleFunc("/transactions/{id}", server.transaction)
	router.HandleFunc("/accounts/{id}/transactions", server.accountTransaction)
	//router.HandleFunc("/shutdown", shutdown)

	log.Printf("Running http server")
	err := http.ListenAndServe(":6543", router)
	if err != nil {
		log.Printf("http server failed with %v", err)
	}
	log.Printf("http server closed down")
}

type Server struct {
	System *System
}

func (s Server) transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		vars := mux.Vars(r)
		idString, hasID := vars["id"]
		if !hasID {
			http.Error(w, "Account id not provided", http.StatusNotAcceptable)
			return
		}
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		s.System.DeleteTransaction(id)

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
}

func (s Server) accountTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString, hasID := vars["id"]
	if !hasID {
		http.Error(w, "Account id not provided", http.StatusNotAcceptable)
		return
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	account, err := s.System.GetAccount(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		log.Printf("Creating transaction for %v", account.Name)

		data := struct {
			Description string `json:"description"`
			Amount float32 `json:"amount"`
		}{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err = decoder.Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		transaction := s.System.CreateTransaction(&account, data.Description, data.Amount)

		js, err := json.Marshal(transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	case "GET":
		log.Printf("Listing transactions for %v", account.Name)

		js, err := json.Marshal(account.Transactions())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
}

func (s Server) accounts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		log.Printf("Creating account")
		data := struct {
			Name string `json:"name"`
		}{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		account := s.System.CreateAccount(data.Name)

		js, err := json.Marshal(account)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "GET":
		log.Printf("Listing accounts")
		source := make(chan ListAccount, 128)
		go s.System.ListAccounts(source)
		accounts := make([]ListAccount, 0)
		for account := range source {
			accounts = append(accounts, account)
		}

		js, err := json.Marshal(accounts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}
}
