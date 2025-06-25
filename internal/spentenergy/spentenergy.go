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

	if weight <= 0 {
		return 0, fmt.Errorf("incorrect value of weight, must be positive - have: %.2f", weight)
	} else if height <= 0 {
		return 0, fmt.Errorf("incorrect value of height, must be positive - have: %.2f", height)
	} else if steps <= 0 {
		return 0, fmt.Errorf("incorrect value of steps, must be positive - have: %d", steps)
	} else if duration <= 0 {
		return 0, fmt.Errorf("incorrect value of time, must be positive - have: %v", duration)
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes() * walkingCaloriesCoefficient) / minInH, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {

	if weight <= 0 {
		return 0, fmt.Errorf("incorrect value of weight, must be positive - have: %.2f", weight)
	} else if height <= 0 {
		return 0, fmt.Errorf("incorrect value of height, must be positive - have: %.2f", height)
	} else if steps <= 0 {
		return 0, fmt.Errorf("incorrect value of steps, must be positive - have: %d", steps)
	} else if duration <= 0 {
		return 0, fmt.Errorf("incorrect value of time, must be positive - have: %v", duration)
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil

}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}
	return ((height * stepLengthCoefficient) * float64(steps)) / mInKm
}
