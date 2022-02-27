package main

import (
	"encoding/json"
	"net/http"

	"github.com/bxcodec/faker/v3"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	services = map[string]Service{}
	routers  = map[string]Router{}
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/healtcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Get("/nipe", func(rw http.ResponseWriter, r *http.Request) {
		services = map[string]Service{}
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(""))
	})
	r.Get("/servies/new", func(rw http.ResponseWriter, r *http.Request) {
		as := faker.Username()
		services[as] = Service{
			LB: LB{
				HealtCheck: &HealtCheck{
					Interval: "3s",
					Scheme:   "http",
					Timeout:  "3s",
					Path:     "/healtcheck",
				},
				Servers: []Server{
					{
						URL: "http://192.168.31.196:3000",
					},
					{
						URL: "http://192.168.32.196:3000",
					},
				},
			},
		}
		routers[as] = Router{
			Rule:    `Host("nipeharefa.dev")`,
			Service: as,
		}
	})
	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		data := &HTTPModel{}
		data.HTTP = HTTP{}
		data.HTTP.Services = services
		data.HTTP.Routers = routers
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(data)
	})
	http.ListenAndServe("0.0.0.0:3000", r)
}
