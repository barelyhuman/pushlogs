package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

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

	// Base Directory Creation
	baseDir := filepath.Join(userHomeDir(), "pushlogs")
	_ = os.MkdirAll(baseDir, os.ModePerm)

	// Reader for Inputs
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Exercises: ")
	exercises, _ := reader.ReadString('\n')
	fmt.Print("Enter Reps: ")
	reps, _ := reader.ReadString('\n')

	// Log Time
	timeNow := time.Now()
	normalDateFormat := timeNow.Format("2006-01-02 15:04:05")

	// Create Log file or check if file exists
	fileName := "workout-log-" + timeNow.Format("20060102150405") + ".txt"
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

	// Create a buffered writer for the file
	bufferedWriter := bufio.NewWriter(file)
	_, err = bufferedWriter.WriteString(
		"Log Time: " + normalDateFormat + "\n" + exercises + reps,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Flush to disk
	bufferedWriter.Flush()
}
