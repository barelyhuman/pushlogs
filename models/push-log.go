package models

// Sets - Array of Reps
type Sets []string

// PushLog - Base Workout log definition
type PushLog struct {
	ExerciseName string
	Sets         Sets
}
