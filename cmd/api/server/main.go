package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"grandChallenge1/grpcchat123/conf"
	"grandChallenge1/grpcchat123/internal/apiserver"
)

const (
	config = "../../../config.yaml"
)

var cnf *conf.Conf

func init() {
	cnf = &conf.Conf{}
	cnf = cnf.GetConf(config)
}
func main() {

	apiserver := &apiserver.Api{
		Cnf: cnf,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/messages", apiserver.GetMessage)
	r.Post("/messages", apiserver.PostMessage)
	tcp_port := fmt.Sprintf(":%d", cnf.ApiServerPort)
	http.ListenAndServe(tcp_port, r)
}
