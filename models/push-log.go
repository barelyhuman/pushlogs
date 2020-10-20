package models

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Sets - Array of Reps
type Sets []string

// PushLog - Base Workout log definition
type PushLog struct {
	ExerciseName string
	Sets         Sets
}

// AddExercise - Asks Questions regarding the addition of an exercise
func (pushLog *PushLog) AddExercise(reader *bufio.Reader) error {
	fmt.Print("Enter Exercises: ")
	exercise, _ := reader.ReadString('\n')
	if len(exercise) <= 3 {
		return errors.New("Cannot Continue without an exercise name")
	}
	pushLog.ExerciseName = exercise
	return nil
}

// AddSets - Asks questions regarding the addition of sets and reps for an exercise
func (pushLog *PushLog) AddSets(reader *bufio.Reader) {
	fmt.Print("Enter Number of Sets: ")
	noOfSets, _ := reader.ReadString('\n')
	noOfSets = strings.TrimSuffix(noOfSets, "\n")
	noOfSetsAsInt, _ := strconv.Atoi(noOfSets)
	for i := 0; i < noOfSetsAsInt; i++ {
		fmt.Print("Enter Reps: ")
		rep, _ := reader.ReadString('\n')
		pushLog.Sets = append(pushLog.Sets, rep)
	}
}
