package main

import (
	// "io"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

// This helps convert strings data to int
func parseScore(val string) (int, error) {
	return strconv.Atoi(strings.TrimSpace(val))
}

func parseCSV(filePath string) []student {
	var s []student

	// Open the file with the filePath and close the Open file once the read operation is done
	file, err := os.Open(filePath)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	// Read the open file all at once
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range records[1:] {
		test1, err := parseScore(row[3])
		if err != nil {
			log.Fatal(err)
		}
		test2, err := parseScore(row[4])
		if err != nil {
			log.Fatal(err)
		}
		test3, err := parseScore(row[5])
		if err != nil {
			log.Fatal(err)
		}
		test4, err := parseScore(row[6])
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, student{
			firstName: row[0],
			lastName: row[1],
			university: row[2],
			test1Score: test1,
			test2Score: test2,
			test3Score: test3,
			test4Score: test4,
		})
	}
	return s
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}

func main() {
	fmt.Println(parseCSV("./grades.csv"))
}
