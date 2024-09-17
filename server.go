package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	port    uint16
	gateway string
}

func NewServer(port uint16, gateway string) *Server {
	return &Server{port, gateway}
}

func (ws *Server) Port() uint16 {
	return ws.port
}

func (ws *Server) Gateway() string {
	return ws.gateway
}

func (ws *Server) Status(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		log.Printf("Status: OK")
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, "{'Status': 'OK' }")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (ws *Server) Run() {
	http.HandleFunc("/", ws.Status)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
