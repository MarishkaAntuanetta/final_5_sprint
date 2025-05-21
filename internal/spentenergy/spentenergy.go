package spentenergy

import (
	"errors"
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
		return 0, errors.New("неверное количество шагов")
	}
	if weight <= 0 {
		return 0, errors.New("некорректный вес")
	}
	if height <= 0 {
		return 0, errors.New("некорректный рост")
	}
	if duration <= 0 {
		return 0, errors.New("некорректное время тренировки")
	}
	minutes := duration.Minutes()
	speed := MeanSpeed(steps, height, duration)
	calories := ((weight * speed * minutes) / minInH) * walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("неверное количество шагов")
	}
	if weight <= 0 {
		return 0, errors.New("некорректный вес")
	}
	if height <= 0 {
		return 0, errors.New("некорректный рост")
	}
	if duration <= 0 {
		return 0, errors.New("некорректное время тренировки")
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * averageSpeed * minutes) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	if steps <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	hours := duration.Hours()
	if hours <= 0 {
		return 0
	}
	averageSpeed := distance / hours
	return averageSpeed
}

func Distance(steps int, height float64) float64 {
	distance := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return distance
}
