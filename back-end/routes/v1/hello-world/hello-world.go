package helloworld

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func GetRoutes(r chi.Router) {

	r.Route("/hello-world", func(r chi.Router) {
		r.Get("/", getHelloWorld)
	})

}

// @summary Get a hi
// @description Get you a nice hi
// @tags departments
// @accept json
// @produce json
// @Security Bearer
// @success 200
// @router /v1/hello-world [get]
func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	oplog := httplog.LogEntry(r.Context())
	oplog.Info("We are gonna say hi")
	w.Write([]byte("hi"))
}
