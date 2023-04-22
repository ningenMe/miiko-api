package main

import (
	"github.com/ningenMe/miiko-api/pkg/application"
	"github.com/ningenMe/miiko-api/proto/gen_go/v1/miikov1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	//infra.NingenmeMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("ningenme").FormatDSN())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer infra.NingenmeMysql.Close()
	//
	//infra.ComproMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("compro").FormatDSN())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer infra.NingenmeMysql.Close()

	//server
	miiko := &application.MiikoController{}
	mux := http.NewServeMux()
	path, handler := miikov1connect.NewMiikoServiceHandler(miiko)
	mux.Handle(path, handler)
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://ningenme.net", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(h2c.NewHandler(mux, &http2.Server{}))
	err := http.ListenAndServe("localhost:8081", corsHandler)
	if err != nil {
		return
	}
}
