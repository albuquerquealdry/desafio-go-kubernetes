package router

import (
	kubernetes "kube-lib/module"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routers() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", kubernetes.Hello).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":5000", router))
}
