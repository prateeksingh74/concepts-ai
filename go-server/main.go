package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
)

type ExplainRequest struct {
	Concept string `json:"concept"`
}

type ExplainResponse struct {
	Definition string `json:"definition"`
	Analogy    string `json:"analogy"`
}

func explainHandler(w http.ResponseWriter, r *http.Request) {
	var req ExplainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.Concept == "" {
		http.Error(w, "Missing concept", http.StatusBadRequest)
		return
	}
	cmd := exec.Command("python3", "../ai-engine/explainer.py", req.Concept)
	output, err := cmd.Output()
	if err != nil {
		http.Error(w, "Error running explainer", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Welcome to the Concept Explainer API. Use POST /explain to get started."))
	})
	http.HandleFunc("/explain", explainHandler)
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
