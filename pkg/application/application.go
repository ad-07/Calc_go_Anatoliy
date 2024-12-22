package application

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/ad-07/calc_go_anatoliy/pkg/calculation"
)

type Config struct {
	Port string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Port = os.Getenv("PORT")
	if config.Port == "" {
		config.Port = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
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
	host, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("Failed to parse RemoteAddr: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Printf("Received request: %s %s from IP: %s, Port: %s", r.Method, r.URL.Path, host, port)
	request := new(Request)
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&request)
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
