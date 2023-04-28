package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ningenMe/miiko-api/pkg/application"
	"github.com/ningenMe/miiko-api/pkg/infra"
	"github.com/ningenMe/miiko-api/proto/gen_go/v1/miikov1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	var err error

	//db
	{
		infra.ComproMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("compro").FormatDSN())
		if err != nil {
			log.Fatalln(err)
		}
		defer infra.ComproMysql.Close()
	}

	//server
	mux := http.NewServeMux()

	{
		miiko := &application.MiikoController{}
		path, handler := miikov1connect.NewMiikoServiceHandler(miiko)
		mux.Handle(path, handler)
	}
	{
		health := &application.HealthController{}
		path, handler := miikov1connect.NewHealthServiceHandler(health)
		mux.Handle(path, handler)
	}
	{
		healthcheck := func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprint(w, "OK")
		}
		mux.Handle("/healthcheck", http.HandlerFunc(healthcheck))
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://ningenme.net", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(h2c.NewHandler(mux, &http2.Server{}))
	err = http.ListenAndServe(":8081", corsHandler)
	if err != nil {
		return
	}
}
