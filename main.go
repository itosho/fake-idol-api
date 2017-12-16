package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	i := Impl{}
	i.InitDB()
	i.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	rest.ErrorFieldName = "error_message"
	router, err := rest.MakeRouter(
		rest.Get("/idols", i.GetAllIdols),
		rest.Post("/idols", i.PostIdol),
		rest.Get("/idols/:id", i.GetIdol),
		rest.Put("/idols/:id", i.PutIdol),
		rest.Delete("/idols/:id", i.DeleteIdol),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	http.Handle("/v1/", http.StripPrefix("/v1", api.MakeHandler()))

	http.Handle("/doc/", http.StripPrefix("/doc", http.FileServer(http.Dir("./doc"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
