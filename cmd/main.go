package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/ping",pingHandler).Methods("GET").Schemes("http")

	server:= &http.Server{Addr:"localhost:8080",Handler:r,}
	go func(){
		log.Println("starting server...")
		err:=server.ListenAndServe()
		if err!=nil {
			log.Fatal("error occurred while starting server:", err)
		}
	}()

	select{}
}

func pingHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("pong"))

}