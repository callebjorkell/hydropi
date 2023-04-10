package server

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Start(ctx context.Context) {
	http.Handle("/api/v1/time", http.HandlerFunc(currentTimeHandler))

	frontend := getFrontendAssets()

	http.Handle("/", http.FileServer(http.FS(frontend)))

	server := &http.Server{Addr: ":8080"}
	serverSignal := make(chan struct{})
	go func() {
		defer close(serverSignal)
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			return
		}

		if err != nil {
			log.Errorf("Server failed: %v", err)
		}
	}()

	select {
	case <-serverSignal:
		log.Fatal("Server closed unexpectedly")
	case <-ctx.Done():
		c, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		_ = server.Shutdown(c)
	}
}

type timeResponse struct {
	Now time.Time `json:"time"`
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(timeResponse{Now: time.Now()})
	if err != nil {
		http.Error(w, "couldn't create time response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	if err != nil {
		http.Error(w, "couldn't write to response", http.StatusInternalServerError)
	}
}
