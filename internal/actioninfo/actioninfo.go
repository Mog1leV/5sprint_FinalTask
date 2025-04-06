package actioninfo

import (
	"errors"
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			err = errors.New("parsing error")
			fmt.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			err = errors.New("error receiving information")
			fmt.Println(err)
			continue
		}
		fmt.Println(info)
	}

}
