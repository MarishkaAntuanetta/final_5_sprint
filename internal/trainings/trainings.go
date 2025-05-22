// Package trainings предоставляет структуру и методы для обработки данных о тренировках.
// Поддерживает парсинг строки с данными и вывод информации о дистанции, скорости и калориях.
package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training представляет данные о тренировке: количество шагов, тип активности,
// длительность и личные данные пользователя.
//
// Поля:
//   - Steps — количество пройденных шагов за тренировку.
//   - TrainingType — тип тренировки (например, "бег", "ходьба").
//   - Duration — продолжительность тренировки в формате time.Duration.
//   - Personal — встроенная структура с личными данными пользователя (имя, вес, рост).
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse реализует интерфейс DataParser. Метод парсит строку данных вида "12345,бег,60m"
// и заполняет поля структуры.
//
// Формат входной строки:
//   - Три значения, разделённые запятой: количество шагов, тип тренировки и длительность.
//   - Пример: "12000,бег,60m", "  12000 , бег , 1h ", "+12000,бег,60m".
//
// Ошибки:
//   - invalid data format — неверное количество частей после разбиения.
//   - invalid data format — лишние пробелы до/после значений.
//   - invalid step value — значение шагов меньше или равно нулю.
//   - invalid activity duration — некорректная длительность или значение <= 0.
func (t *Training) Parse(datastring string) (err error) {
	strData := strings.Split(datastring, ",")
	if len(strData) != 3 {
		return errors.New("invalid data format")
	}

	t.Steps, err = strconv.Atoi(strData[0])
	if err != nil {
		return err
	}
	if t.Steps <= 0 {
		return errors.New("invalid step value")
	}

	t.TrainingType = strData[1]
	if t.TrainingType == "" {
		return errors.New("training type is empty")
	}

	t.Duration, err = time.ParseDuration(strData[2])
	if err != nil {
		return err
	}
	if t.Duration <= 0 {
		return errors.New("invalid activity duration")
	}

	return nil
}

// ActionInfo рассчитывает и возвращает информацию о тренировке:
// тип активности, дистанцию, среднюю скорость и количество сожжённых калорий.
//
// Возвращаемое значение:
//   - Строка с информацией о тренировке.
//   - Ошибка, если не удалось выполнить расчёты.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error

	switch strings.ToLower(t.TrainingType) {
	case "бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", errors.New("unable to calculate calories burned")
	}

	return fmt.Sprintf("Тип тренировки: %s\n"+
		"Длительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\n"+
		"Скорость: %.2f км/ч\n"+
		"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		averageSpeed,
		calories), nil
}
