package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// Разделяем входные данные по запятой
	strData := strings.Split(datastring, ",")
	if len(strData) != 2 {
		return errors.New("invalid data format")
	}

	// Извлекаем количество шагов и продолжительность
	stepStr := strData[0]
	durationStr := strData[1]

	// Проверяем на лишние пробелы — если обрезанная версия не равна оригинальной, значит формат неверный
	if stepStr != strings.TrimSpace(stepStr) || durationStr != strings.TrimSpace(durationStr) {
		return errors.New("invalid data format — extra spaces")
	}

	// Очистка пробелов и знака "+" перед числом шагов
	stepStr = strings.TrimPrefix(strings.TrimSpace(stepStr), "+")
	durationStr = strings.TrimSpace(durationStr)

	// Преобразуем количество шагов в число
	ds.Steps, err = strconv.Atoi(stepStr)
	if err != nil {
		return err
	}
	if ds.Steps <= 0 {
		return errors.New("invalid step value")
	}
	// Преобразуем продолжительность в тип time.Duration
	ds.Duration, err = time.ParseDuration(durationStr)
	if err != nil {
		return err
	}
	if ds.Duration <= 0 {
		return errors.New("invalid activity duration")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// Рассчитываем дистанцию в метрах и километрах
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)

	// Считаем потраченные калории
	caloriesBurned, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		log.Println("unable to calculate calories burned")
		return "", err
	}

	// Формируем итоговую строку с результатами
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		caloriesBurned,
	), nil
}
