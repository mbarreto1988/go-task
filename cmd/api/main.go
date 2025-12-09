package main
import (
	"log"
	"net/http"
	"os"
	"time"
)


// App va a ser nuestro "contenedor" de dependencias
type App struct {
	logger *log.Logger
}

func main() {
	// Logger simple, pero prolijo
	logger := log.New(os.Stdout, "API ", log.LstdFlags|log.Lshortfile)

	app := &App{
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", app.healthHandler)

	addr := ":8080"

	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Printf("server escuchando en %s", addr)

	if err := srv.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
}

func (a *App) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}