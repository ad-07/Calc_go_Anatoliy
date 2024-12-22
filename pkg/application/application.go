package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"calc_go_anatoliy/pkg/calculation"
)

type Config struct {
	Port string
}


type Application struct {
	config *Config
}

func New() *Application {
	cnf := new(Config)
	cnf.Port = "8080"
	return &Application{

		config: cnf,
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "error: %s", calculation.ErrMethodNotAllowed)
		return
	}
	log.Printf("Received request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr) // Лог о каждом запросе

	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		if errors.Is(err, calculation.ErrInvalidExpression) {
			http.Error(w, "", http.StatusUnprocessableEntity)
			fmt.Fprintf(w, "error: %s", err.Error())
		} else {
			http.Error(w, "", http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %s", err.Error())
		}

	} else {
		fmt.Fprintf(w, "result: %f", result)
	}
}

func (a *Application) RunServer() error {
	log.Println("Starting server...")
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Port, nil)
}
