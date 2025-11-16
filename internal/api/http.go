package api

import (
	"encoding/json"
	"go-worker-pool/internal/worker"
	"net/http"
)

type JobRequest struct {
	ID string `json:"id"`
}

func RegisterRoutes(pool *worker.WorkerPool) {
	http.HandleFunc("/job", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req JobRequest
		json.NewDecoder(r.Body).Decode(&req)

		if req.ID == "" {
			http.Error(w, "missing id", http.StatusBadRequest)
			return
		}

		pool.AddJob(req.ID)
		w.Write([]byte("Job added"))
	})

	http.ListenAndServe(":8080", nil)
}
