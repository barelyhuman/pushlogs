package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Exercises: ")
	exercises, _ := reader.ReadString('\n')
	fmt.Print("Enter Reps: ")
	reps, _ := reader.ReadString('\n')
	timeNow := time.Now()
	normalDateFormat := timeNow.Format("2006-01-02 15:04:05")
	fileName := "workout-log-" + timeNow.Format("20060102150405") + ".txt"
	var file *os.File
	var openError error
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			file, _ = os.Create(fileName)
		} else {
			file, openError = os.OpenFile(fileName, os.O_WRONLY, 0666)
			if openError != nil {
				log.Fatal(err)
			}
		}
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	_, err = bufferedWriter.WriteString(
		"Log Time: " + normalDateFormat + "\n" + exercises + reps,
	)
	if err != nil {
		log.Fatal(err)
	}

	bufferedWriter.Flush()
}

