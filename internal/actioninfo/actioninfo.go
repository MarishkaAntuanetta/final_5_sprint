package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

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
