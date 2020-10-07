package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type pushLog struct {
	exerciseName string
	reps         string
}

func userHomeDir() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	return os.Getenv(env)
}

func main() {

	var inputLogs []pushLog

	// Base Directory Creation
	baseDir := filepath.Join(userHomeDir(), "pushlogs")
	_ = os.MkdirAll(baseDir, os.ModePerm)

	// Reader for Inputs
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter Exercises: ")
		exercise, _ := reader.ReadString('\n')
		exercise = strings.TrimSuffix(exercise, "\n")
		fmt.Print("Enter Reps: ")
		rep, _ := reader.ReadString('\n')
		fmt.Print("Add More?(n for no, anything else for yes): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")
		inputLogs = append(inputLogs, pushLog{exerciseName: exercise, reps: rep})
		if answer == "n" {
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
		totalLog += logItem.exerciseName + " - " + logItem.reps + "\n"
	}

	_, err = bufferedWriter.WriteString(totalLog)
	if err != nil {
		log.Fatal(err)
	}

}
