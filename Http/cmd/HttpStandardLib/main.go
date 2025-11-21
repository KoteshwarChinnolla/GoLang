package main

import (
	"ApiGateway/pkg/middleware"
	"ApiGateway/pkg/textmanipulation"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

// ----------------------------------
// Main
// ----------------------------------
func main() {
	repo := textmanipulation.NewPageRepo()
	controller := NewController(repo)
	router := http.NewServeMux()

	// Handle is for class and Handle Func is for methods
	router.HandleFunc("GET /", controller.ServeHTTP)
	router.HandleFunc("GET /allMessages", controller.getAllMessages)
	router.HandleFunc("GET /message/{Id}", controller.getMessage)
	router.HandleFunc("POST /message", controller.addMessage)
	router.HandleFunc("PUT /message/{Id}", controller.updateMessage)
	router.HandleFunc("DELETE /message/{Id}", controller.deleteMessage)

	// Strip Path implementation
	v1 := http.NewServeMux()

	v1.Handle("/v1/", http.StripPrefix("/v1", router))
	// handler := middleware.Logging(v1) this is for one middleware for multiple below is the implementation
	securedHandler := middleware.Chain(
		middleware.Logging,
		middleware.RateLimit,
	)
	// Enabling TLS
	certFile := filepath.Join("tls_keys", "tls.crt")
	keyFile := filepath.Join("tls_keys", "tls.key")

	serverTLSCer, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalln(err.Error())
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{serverTLSCer}}

	tlsServer := http.Server{
		Addr:      ":443",
		Handler:   securedHandler(v1),
		TLSConfig: tlsConfig,
	}

	// Suppurate Go routine for TLS listing
	go func() {
		tlsServer.ListenAndServeTLS("", "")
	}()

	// Also a Non TLS Path
	server := http.Server{
		Addr:    ":8080",
		Handler: securedHandler(v1),
	}

	server.ListenAndServe()
}

type Controller struct {
	repo textmanipulation.Repository // Dependency Injection
}

// Constructor Injection
func NewController(service textmanipulation.Repository) *Controller {
	return &Controller{repo: service}
}

// Request Routing
func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// Controller Implementation
func (c *Controller) getAllMessages(w http.ResponseWriter, r *http.Request) {
	messages := c.repo.GetAll()
	if messages == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	json, err := json.Marshal(messages)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(json))
}

func (c *Controller) getMessage(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("Id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	messages := c.repo.GetById(int(id))
	json, err := json.Marshal(messages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(json))
}

func (c *Controller) addMessage(w http.ResponseWriter, r *http.Request) {
	var message textmanipulation.SimpleMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	messages := c.repo.AddMessage(message)
	json, err := json.Marshal(messages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(json))
}

func (c *Controller) deleteMessage(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("Id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	messages := c.repo.DeleteMessage(int(id))
	json, err := json.Marshal(messages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(json))
}

func (c *Controller) updateMessage(w http.ResponseWriter, r *http.Request) {
	var message textmanipulation.SimpleMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	idStr := r.PathValue("Id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	messages := c.repo.UpdateMessage(int(id), message)
	json, err := json.Marshal(messages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(json))
}
