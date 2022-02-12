package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filepath := flag.String("fpath", "problems.csv", "the file path to the csv containing problems")
	timer := flag.Int("timer", -1, "Set the timer in seconds, default -1 means disabled")
	flag.Parse()
	fmt.Printf("the filepath is %v\n", *filepath)
	fmt.Printf("the timer is %v\n", *timer)

	records, err := readData(*filepath)
	if err != nil {
		exit(fmt.Sprintf("\nFailed to read file %s\n", *filepath), err)
	}
	questions := parseLinesIntoProblems(records)

	countCorrect := 0

	for i, p := range questions {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			countCorrect++
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