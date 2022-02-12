package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filepath := flag.String("fpath", "problems.csv", "the file path to the csv containing problems")
	timeLimit := flag.Int("timer", 30, "the timer in seconds")
	flag.Parse()
	fmt.Printf("the filepath is %v\n", *filepath)
	fmt.Printf("the timer is %v\n", *timeLimit)

	lines, err := readData(*filepath)
	if err != nil {
		exit(fmt.Sprintf("\nFailed to read file %s\n", *filepath), err)
	}
	questions := parseLinesIntoProblems(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	countCorrect := 0
	problemloop:
		for i, p := range questions {
			fmt.Printf("Problem #%d: %s = ", i+1, p.q)
			ansChan := make(chan string)
			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				ansChan <- answer
			}()
			select {
			case <- timer.C:
				fmt.Println()
				break problemloop
			case answer := <- ansChan:
				if answer == p.a{
					countCorrect++
				}
			}
		}

	fmt.Printf("You answered %d correct out of %d\n", countCorrect, len(questions))
}

func readData(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func parseLinesIntoProblems(lines [][]string) (problems []problem) {
	problems = make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem {
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return
}

func exit(reason string, err error) {
	log.Fatal(reason, err)
}

type problem struct {
	q string
	a string 
}