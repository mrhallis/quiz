package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func serveQuestion(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	question := Questions[rand.Intn(len(Questions))]

	responseQuestion := struct {
		Text    string   `json:"text"`
		Options []string `json:"options"`
	}{
		Text:    question.Text,
		Options: question.Options,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseQuestion)
}

func handleSubmission(w http.ResponseWriter, r *http.Request) {
	var submission struct {
		QuestionText string `json:"questionText"`
		Answer       int    `json:"answer"`
	}

	err := json.NewDecoder(r.Body).Decode(&submission)
	if err != nil {
		http.Error(w, "Invalid submission", http.StatusBadRequest)
		return
	}

	var correctAnswer int
	for _, q := range Questions {
		if q.Text == submission.QuestionText {
			correctAnswer = q.CorrectAnswer
			break
		}
	}

	result := struct {
		Correct bool `json:"correct"`
	}{
		Correct: submission.Answer == correctAnswer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/question", serveQuestion)
	http.HandleFunc("/submit", handleSubmission)
	http.ListenAndServe(":8080", nil)
}
