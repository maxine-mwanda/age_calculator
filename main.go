package main

import (
	"bufio"
	"fmt"
	age "github.com/bearbin/go-age"
	"log"
	"os"
	"time"
	_ "time"
)

func main () {
	sample, err := os.Open("Sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sample.Close()

	scanner := bufio.NewScanner(sample)
	for scanner.Scan() {
		line:= scanner.Text()
		ageInYears, err := GetTime(line)
		if err != nil {
			continue
		}
		fmt.Println("DOB", line, "Age : ", ageInYears, "years old")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}



func GetTime(dob string) (ageInYears int, err error) {
	format:= "2006-01-02"
	t, err := time.Parse(format, dob)
	if err!= nil{
		log.Println("an error occured", err)
		return
	}
	ageInYears = age.Age(t)
	return
}