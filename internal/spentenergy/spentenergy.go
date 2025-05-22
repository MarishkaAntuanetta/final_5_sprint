// Package spentenergy предоставляет функции для расчёта дистанции и потраченных калорий при ходьбе и беге.
package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, используемые в расчётах.
const (
	// mInKm — количество метров в километре.
	mInKm = 1000

	// minInH — количество минут в часе.
	minInH = 60

	// stepLengthCoefficient — коэффициент длины шага на основе роста пользователя.
	stepLengthCoefficient = 0.45

	// walkingCaloriesCoefficient — коэффициент для расчёта калорий при ходьбе.
	walkingCaloriesCoefficient = 0.5
)

// WalkingSpentCalories рассчитывает количество сожжённых калорий при ходьбе.
//
// Параметры:
//   - steps: количество пройденных шагов.
//   - weight: масса тела в килограммах.
//   - height: рост в метрах.
//   - duration: продолжительность активности в формате time.Duration.
//
// Возвращаемые значения:
//   - float64 — количество сожжённых калорий.
//   - error — ошибка, если входные данные некорректны.
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

// RunningSpentCalories рассчитывает количество сожжённых калорий при беге.
//
// Параметры:
//   - steps: количество пройденных шагов.
//   - weight: масса тела в килограммах.
//   - height: рост в метрах.
//   - duration: продолжительность активности в формате time.Duration.
//
// Возвращаемые значения:
//   - float64 — количество сожжённых калорий.
//   - error — ошибка, если входные данные некорректны.
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

// MeanSpeed рассчитывает среднюю скорость движения на основе шагов, роста и времени.
//
// Параметры:
//   - steps: количество пройденных шагов.
//   - height: рост пользователя в метрах.
//   - duration: продолжительность активности в формате time.Duration.
//
// Возвращаемое значение:
//   - float64 — средняя скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 || steps <= 0 {
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

// Distance рассчитывает дистанцию в километрах на основе количества шагов и роста.
//
// Параметры:
//   - steps: количество пройденных шагов.
//   - height: рост пользователя в метрах.
//
// Возвращаемое значение:
//   - float64 — дистанция в километрах.
func Distance(steps int, height float64) float64 {
	distance := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return distance
}
