package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/geovanisouza92/api-factory/apifactory"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestLogger(&loggingFormatter{
	// analytics: ??
	}))
	router.Use(middleware.Recoverer)

	router.Route("/v1/{service}/{resource}", func(r chi.Router) {
		r.Get("/", getHandler)
		r.Put("/", putHandler)
		r.Post("/", postHandler)
		r.Patch("/", patchHandler)
		r.Delete("/", deleteHandler)
	})

	http.ListenAndServe(":3000", router)
}

func getHandler(w http.ResponseWriter, r *http.Request)    {}
func putHandler(w http.ResponseWriter, r *http.Request)    {}
func postHandler(w http.ResponseWriter, r *http.Request)   {}
func patchHandler(w http.ResponseWriter, r *http.Request)  {}
func deleteHandler(w http.ResponseWriter, r *http.Request) {}

type loggingFormatter struct {
	tracing apifactory.TracingProvider
}

func (f *loggingFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	ev := apifactory.TracingEvent{}
	ev.RequestID = middleware.GetReqID(r.Context())
	ev.Scheme = "http"
	if r.TLS != nil {
		ev.Scheme = "https"
	}
	ev.Host = r.Host
	ev.Method = r.Method
	ev.RequestURI = r.RequestURI
	ev.Proto = r.Proto
	ev.RemoteAddr = r.RemoteAddr

	return &loggingEntry{tracing: f.tracing, ev: ev}
}

type loggingEntry struct {
	tracing apifactory.TracingProvider
	ev      apifactory.TracingEvent
}

func (e *loggingEntry) Write(status, bytes int, elapsed time.Duration) {
	e.ev.Status = status
	e.ev.Bytes = bytes
	e.ev.Duration = int64(elapsed)
	e.tracing.Send(e.ev)
}

func (e *loggingEntry) Panic(v interface{}, stack []byte) {
	// TODO:
}
