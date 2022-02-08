package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := flag.String("fpath", "problems.csv", "the file path to the csv containing problems")
	timer := flag.Int("timer", -1, "Set the timer in seconds, default -1 means disabled")
	flag.Parse()
	fmt.Printf("the filepath is %v\n", *filepath)
	fmt.Printf("the timer is %v\n", *timer)

	records, err := readData(*filepath)
	if err != nil {
		log.Fatal(err)
	}
	var countCorrect int = 0
	var countIncorrect int = 0

	for _, record := range records {
		question := string(record[0])
		answer := string(record[1])
		fmt.Printf("%v = ", question)
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if userAnswer == answer {
			fmt.Print("Correct\n")
			countCorrect++
		} else {
			fmt.Print("Incorrect\n")
			countIncorrect++
		}
	}

	fmt.Printf("You answered %v correct\n", countCorrect)
	fmt.Printf("You answered %v incorrect\n", countIncorrect)
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
