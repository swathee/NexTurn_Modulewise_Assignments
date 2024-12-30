//Exercise 4: Online Examination System

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Question struct to represent a quiz question
type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

// Question Bank
var questions = []Question{
	{"What is the keyword to define a function in Go?", [4]string{"func", "function", "def", "define"}, 1},
	{"What command is used to run a Go program?", [4]string{"go start", "go exec", "go run", "go boot"}, 3},
	{"Which keyword is used to stop a loop in Go? ", [4]string{"stop", "break", "exit", "end"}, 2},
	{"Which keyword is used to import packages in Go? ", [4]string{"package", "require", "include", "import"}, 4},
	{"How do you write a single-line comment in Go? ", [4]string{"//comment", "#comment", "/*comment*/", "--comment"}, 1},
}

// Timer duration for each question
const questionTimeLimit = 15 * time.Second

func main() {
	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("Instructions:")
	fmt.Println("- Enter the option number to select an answer.")
	fmt.Println("- Type 'exit' to quit the quiz anytime.")
	fmt.Println("- You have a limited time of 15 seconds per question.")

	takeQuiz()
}

// takeQuiz handles the quiz process
func takeQuiz() {
	reader := bufio.NewReader(os.Stdin)
	score := 0

	for i, question := range questions {
		fmt.Printf("Question %d: %s\n", i+1, question.Question)
		for idx, option := range question.Options {
			fmt.Printf("%d. %s\n", idx+1, option)
		}

		answerChan := make(chan string)
		go func() {
			fmt.Print("Your answer: ")
			input, _ := reader.ReadString('\n')
			answerChan <- strings.TrimSpace(input)
		}()

		select {
		case answer := <-answerChan:
			if strings.ToLower(answer) == "exit" {
				fmt.Println("Exiting the quiz...")
				break
			}

			choice, err := strconv.Atoi(answer)
			if err != nil || choice < 1 || choice > 4 {
				fmt.Println("Invalid input! Please enter a valid option (1-4).")
				continue
			}

			if choice == question.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Incorrect!")
			}

		case <-time.After(questionTimeLimit):
			fmt.Println("\nTime's up for this question!")
			continue
		}

		fmt.Println()
	}

	calculateScore(score, len(questions))
}

// calculateScore calculates and displays the user's score and performance
func calculateScore(score, total int) {
	fmt.Printf("\nQuiz Completed! Your score: %d/%d\n", score, total)
	percentage := (float64(score) / float64(total)) * 100

	if percentage >= 80 {
		fmt.Println("Performance: Excellent!")
	} else if percentage >= 50 {
		fmt.Println("Performance: Good!")
	} else {
		fmt.Println("Performance: Needs Improvement.")
	}
}
