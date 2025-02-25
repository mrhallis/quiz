package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
)

type Question struct {
	Text    string   `json:"text"`
	Options []string `json:"options"`
}

func takeQuiz(cmd *cobra.Command, args []string) {
	for {
		question, err := fetchQuestion()
		if err != nil {
			fmt.Println("Error fetching question:", err)
			return
		}

		fmt.Println(question.Text)
		for i, option := range question.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		answer := promptForAnswer(len(question.Options))
		correct := submitAnswer(question.Text, answer-1)

		if correct {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect. Try again!")
		}

		fmt.Print("Continue? (y/n): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() != "y" {
			break
		}
	}
}

func fetchQuestion() (Question, error) {
	resp, err := http.Get("http://localhost:8080/question")
	if err != nil {
		return Question{}, err
	}
	defer resp.Body.Close()

	var question Question
	err = json.NewDecoder(resp.Body).Decode(&question)
	return question, err
}

func promptForAnswer(optionCount int) int {
	for {
		fmt.Print("Enter your answer (1-" + strconv.Itoa(optionCount) + "): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer, err := strconv.Atoi(scanner.Text())
		if err == nil && answer >= 1 && answer <= optionCount {
			return answer
		}
		fmt.Println("Invalid input. Please try again.")
	}
}

func submitAnswer(questionText string, answer int) bool {
	submission := struct {
		QuestionText string `json:"questionText"`
		Answer       int    `json:"answer"`
	}{
		QuestionText: questionText,
		Answer:       answer,
	}

	jsonData, _ := json.Marshal(submission)
	resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error submitting answer:", err)
		return false
	}
	defer resp.Body.Close()

	var result struct {
		Correct bool `json:"correct"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Correct
}

func main() {
	var rootCmd = &cobra.Command{Use: "quiz"}
	var takeCmd = &cobra.Command{
		Use:   "take",
		Short: "Take the SC-200 quiz",
		Run:   takeQuiz,
	}
	rootCmd.AddCommand(takeCmd)
	rootCmd.Execute()
}
