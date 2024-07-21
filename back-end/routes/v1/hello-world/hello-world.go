package helloworld

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"wowcollector.io/common/data"
	battleNetHttp "wowcollector.io/services/http"
)

func GetRoutes(r chi.Router) {

	r.Route("/hello-world", func(r chi.Router) {
		r.Get("/", getHelloWorld)
		r.Get("/mounts", getMountsIndex)
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

// @summary Get all mounts
// @description Get you a nice hi
// @tags departments
// @accept json
// @produce json
// @Security Bearer
// @success 200
// @router /v1/mounts [get]
func getMountsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	oplog := httplog.LogEntry(r.Context())
	oplog.Info("We are fetching all mounts and returning them")

	mounts := battleNetHttp.GetInstance()
	b, err := json.Marshal(mounts.GetMountsIndex(data.EU))
	if err != nil {
		oplog.Error("Failed fetching mounts index")
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
