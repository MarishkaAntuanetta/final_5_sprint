// Package daysteps предоставляет структуру и методы для обработки данных о количестве шагов за день.
// Включает парсинг строки данных и расчёт дистанции и потраченных калорий.
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

// DaySteps представляет данные о пройденных шагах за один день.
//
// Поля:
//   - Steps — количество пройденных шагов за день.
//   - Duration — продолжительность активности в формате time.Duration.
//   - Personal — встроенная структура с личными данными пользователя (имя, вес, рост).
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse реализует интерфейс DataParser. Метод парсит строку данных вида "12345,30m"
// и заполняет поля структуры.
//
// Формат входной строки:
//   - Два значения, разделённые запятой: количество шагов и длительность активности.
//   - Пример: "12000,60m", "+12000,1h", " 12000 , 60m ".
//
// Ошибки:
//   - invalid data format — неверное количество частей после разбиения.
//   - invalid data format — лишние пробелы до/после значений.
//   - invalid step value — значение шагов меньше или равно нулю.
//   - invalid activity duration — некорректная длительность или значение <= 0.
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

// ActionInfo реализует интерфейс DataParser. Метод рассчитывает и возвращает информацию о тренировке:
// дистанцию и количество сожжённых калорий.
//
// Возвращаемое значение:
//   - Строка с информацией о количестве шагов, дистанции и сожжённых калориях.
//   - Ошибка, если не удалось выполнить расчёты.
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
