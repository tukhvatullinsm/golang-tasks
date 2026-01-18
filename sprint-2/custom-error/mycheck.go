//go:build !solution

package mycheck

import (
	"errors"
	"log"
	"regexp"
)

const ()

type MyError []error

func (me MyError) Error() string {
	resultError := ""
	for index, err := range me {
		if index == len(me)-1 {
			resultError += err.Error()
		} else {
			resultError += err.Error() + ";"
		}

	}
	return resultError
}

func MyCheck(input string) error {
	errNum := errors.New("found numbers")
	errLen := errors.New("line is too long")
	errSpc := errors.New("no two spaces")
	tempString := []rune(input)
	stringErrors := make(MyError, 0, 3)
	reNum, err := regexp.Compile(`\d`)
	if err != nil {
		log.Fatalln("Ошибка компиляции регулярного выражения #1", err)
	}
	if ok := reNum.MatchString(input); ok {
		stringErrors = append(stringErrors, errNum)
	}
	if len(tempString) > 20 {
		stringErrors = append(stringErrors, errLen)
	}
	reSpc, err := regexp.Compile(`\s`)
	if err != nil {
		log.Fatalln("Ошибка компиляции регулярного выражения #3", err)
	}
	if count := reSpc.FindAllString(input, -1); len(count) != 2 {
		stringErrors = append(stringErrors, errSpc)
	}
	if len(stringErrors) == 0 {
		return nil
	}
	return stringErrors
}
