package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("error. steps <= 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("error. weight <= 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("error. height <= 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("error. duration <= 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / float64(minInH) * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("error. steps <= 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("error. weight <= 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("error. height <= 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("error. duration <= 0")
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return (weight * meanSpeed * durationInMinutes) / float64(minInH), nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)
	h := duration.Hours()
	return dist / h
}

func Distance(steps int, height float64) float64 {
	return height * stepLengthCoefficient * float64(steps) / float64(mInKm)
}
