package main



import (
   "fmt"
   "log"
   "net/http"
   "sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main() {
	const filepathRoot = "."
	const port = "8080"
	cfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}


	mux := http.NewServeMux()
	mux.Handle("/app/", cfg.middlewareMetricsInc(
			http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("/healthz", handlerReadiness)
	//mux.Handle("/app/", cfg.middlewareMetricsInc(handler))
	mux.HandleFunc("/metrics", cfg.handlerHits)
	mux.HandleFunc("/reset", cfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}


func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerHits(w http.ResponseWriter, r *http.Request) {
   w.Header().Add("Content-Type", "application/json; charset=utf-8")
   w.WriteHeader(http.StatusOK)
   w.Write([]byte(`{` + fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load()) + `}`))
}
