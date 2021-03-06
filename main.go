package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pushlogs/lib"
	"pushlogs/models"
	"strings"
	"time"
)

func main() {
	var inputLogs []models.PushLog
	dir := flag.String("d", filepath.Join(lib.UserHomeDir(), "pushlogs"), "Directory to save pushlogs too.")

	flag.Parse()
	// Base Directory Creation
	baseDir := *dir
	_ = os.MkdirAll(baseDir, os.ModePerm)

	// Reader for Inputs
	reader := bufio.NewReader(os.Stdin)

	for {
		pushLog, isBreak := askQuestion(reader)
		inputLogs = append(inputLogs, pushLog)
		if isBreak {
			break
		}
	}

	// Log Time
	timeNow := time.Now()
	normalDateFormat := timeNow.Format("02-01-2006 15:04:05")

	// Create Log file or check if file exists
	fileName := "workout-log-" + timeNow.Format("02-01-2006-15-04-05") + ".txt"
	var file *os.File
	var openError error
	_, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			file, _ = os.Create(filepath.Join(baseDir, fileName))
		} else {
			file, openError = os.OpenFile(filepath.Join(baseDir, fileName), os.O_WRONLY, 0666)
			if openError != nil {
				log.Fatal(err)
			}
		}
	}
	defer file.Close()

	totalLog := ""

	// Create a buffered writer for the file
	bufferedWriter := bufio.NewWriter(file)
	defer bufferedWriter.Flush()

	totalLog += "Log Time: " + normalDateFormat + "\n\n"

	for _, logItem := range inputLogs {
		totalLog += logItem.ExerciseName
		for _, set := range logItem.Sets {
			totalLog += " - " + set
		}
	}

	_, err = bufferedWriter.WriteString(totalLog)
	if err != nil {
		log.Fatal(err)
	}
}

func askQuestion(reader *bufio.Reader) (models.PushLog, bool) {
	isBreak := false
	pushLog := models.PushLog{}
	err := pushLog.AddExercise(reader)
	for err != nil {
		fmt.Println(err)
		fmt.Println()
		err = pushLog.AddExercise(reader)
	}
	pushLog.AddSets(reader)
	fmt.Print("Add More? (n for no, anything else for yes): ")
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSuffix(answer, "\n")
	if answer == "n" {
		isBreak = true
	}
	return pushLog, isBreak
}
