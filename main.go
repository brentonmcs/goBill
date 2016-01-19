package main

import (
	"goBill/routers"
	"goBill/runners/jobRunner"
	"goBill/settings"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/pilu/xrequestid"
)

func main() {
	settings.Init()
	jobRunner.InitialiseJobs()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.Use(xrequestid.New(16))
	n.Use(negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		rw.Header().Set(xrequestid.DefaultHeaderKey, r.Header.Get(xrequestid.DefaultHeaderKey))
		next(rw, r)
	}))

	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
