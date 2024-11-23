package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(strIn string) (string, error) {
	// Place your code here.
	// Если строка пусто, то вернуть пустую строку
	if strIn == "" {
		return "", nil
	}

	// Проверяем, что первый символ это число, иначе ошибка
	_, err := strconv.Atoi(string(strIn[0]))
	if err == nil { // Если нет ошибки, значит это число
		return "", ErrInvalidString
	}

	// Преобразуем строку в массив runes, так как легче отличить литерал от экранирование символа в строке
	runes := []rune(strIn)

	var bufferStr strings.Builder
	var resultStr strings.Builder
	lastRuneIsDigital := false
	lastRuneIsLiteral := false

	for _, value := range runes {
		// проверка, является ли символ строки целым числом
		num, err := strconv.Atoi(string(value))

		// если цифра повторяется два раза подряд, то ошибка
		if err == nil && lastRuneIsDigital == true {
			return "", ErrInvalidString
		}

		// был ли сохранен предыдущий символ в буфере
		if bufferStr.Len() > 0 {
			// является ли символ числом, да то ..., иначе ...
			if err == nil {
				// проверка является ли это одинарным литералом
				if bufferStr.String() == "\\" && lastRuneIsLiteral == false {
					bufferStr.Reset()
					bufferStr.WriteRune(value)
					continue
				}

				// сохраняем в буфере с повторением num раз
				resultStr.WriteString(strings.Repeat(bufferStr.String(), num))
				bufferStr.Reset()

				lastRuneIsDigital = true  // флаг, что предыдущий символ является число
				lastRuneIsLiteral = false // флаг, что предыдущий не равен `\`
			} else {
				// обнуляем литерал, если он задвоился
				if bufferStr.String() == "\\" {
					bufferStr.Reset()
					bufferStr.WriteRune(value)
					lastRuneIsLiteral = true
					continue
				}

				// сохраняем в буфере с повторением 1 раз
				resultStr.WriteString(strings.Repeat(bufferStr.String(), 1))
				bufferStr.Reset()

				bufferStr.WriteRune(value) // сохраняем в буффер текущий символ
				lastRuneIsDigital = false  // флаг, что предыдущий символ является число
			}
		} else {
			bufferStr.WriteRune(value)
			lastRuneIsDigital = false
		}
	}

	// проверка на последний символ
	if bufferStr.Len() > 0 {
		resultStr.WriteString(strings.Repeat(bufferStr.String(), 1))
	}

	return resultStr.String(), nil
}
