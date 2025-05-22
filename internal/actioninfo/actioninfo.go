// Package actioninfo предоставляет функционал для обработки набора данных
// с помощью интерфейса DataParser. Основная цель — унифицировать парсинг и вывод информации.
package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — это интерфейс, который должен реализовывать любой тип,
// предназначенный для обработки строковых данных.
//
// Методы:
//   - Parse(datastring string) — парсит входную строку и заполняет поля структуры.
//   - ActionInfo() (string, error) — возвращает информацию о данных в виде строки.
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// Info — основная функция пакета. Она перебирает слайс строк и использует
// переданный DataParser для обработки каждой строки.
//
// Параметры:
//   - dataset []string — набор строковых данных для обработки.
//   - dp DataParser — объект, реализующий методы Parse и ActionInfo.
//
// Логика работы:
//  1. Для каждой строки вызывается Parse(), чтобы заполнить данные в объекте.
//  2. Если Parse() завершается ошибкой — выводится лог и строка пропускается.
//  3. Вызывается ActionInfo(), чтобы получить результат в виде строки.
//  4. Если ActionInfo() возвращает ошибку — выводится лог и строка пропускается.
//  5. Успешный результат выводится на экран.
func Info(dataset []string, dp DataParser) {
	for _, value := range dataset {
		err := dp.Parse(value)
		if err != nil {
			log.Println("не удалось идентифицировать данные")
			continue
		}
		result, err := dp.ActionInfo()
		if err != nil {
			log.Println("не удалось произвести расчет")
			continue
		}
		fmt.Println(result)
	}
}
