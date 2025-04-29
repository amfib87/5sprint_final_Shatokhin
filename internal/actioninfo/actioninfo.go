package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, val := range dataset {
		err := dp.Parse(val)

		if err != nil {
			log.Println(err.Error())
			continue
		}

		text, err := dp.ActionInfo()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		fmt.Println(text)
	}
}
