// Package personaldata предоставляет структуру и методы для хранения и вывода личных данных пользователя.
package personaldata

import "fmt"

// Personal представляет данные о пользователе: имя, вес и рост.
//
// Поля:
//   - Name — имя пользователя (строка).
//   - Weight — масса тела в килограммах (float64).
//   - Height — рост в метрах (float64).
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print — метод, который выводит информацию о пользователе на экран.
//
// Формат вывода:
//   - Имя: [имя]
//   - Вес: [вес] кг.
//   - Рост: [рост] м.
//
// Пример вывода:
//
//	Имя: Алексей
//	Вес: 72.50 кг.
//	Рост: 1.78 м.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n\n", p.Name, p.Weight, p.Height)
}
