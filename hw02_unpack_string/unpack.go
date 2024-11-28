package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
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

	// Если последний символ заканчивается на `\`, то ошибка
	if string(strIn[len(strIn)-1]) == `\` {
		return "", ErrInvalidString
	}

	var bufferStr strings.Builder
	var resultStr strings.Builder
	lastRuneIsDigital := false
	countRepeatLiteral := 0

	for _, value := range strIn {
		// проверка, является ли символ строки целым числом
		num, err := strconv.Atoi(string(value))

		// если цифра повторяется два раза подряд, то ошибка
		if err == nil && lastRuneIsDigital {
			return "", ErrInvalidString
		}

		// если экранируется не символ или число, то ошибка
		if bufferStr.String() == "\\" && !unicode.IsLetter(value) && !unicode.IsDigit(value) && string(value) != "\\" {
			return "", ErrInvalidString
		}

		// был ли сохранен предыдущий символ в буфере
		if bufferStr.Len() == 0 {
			bufferStr.WriteRune(value)
			lastRuneIsDigital = false
			continue
		}

		// является ли символ числом, да то ..., иначе ...
		if err == nil {
			// проверка является ли это одинарным литералом
			if bufferStr.String() == "\\" && (countRepeatLiteral%2 == 0) {
				bufferStr.Reset()
				bufferStr.WriteRune(value)
				lastRuneIsDigital = false // флаг, что предыдущий символ является число
				continue
			}

			// сохраняем в буфере с повторением num раз
			resultStr.WriteString(strings.Repeat(bufferStr.String(), num))
			bufferStr.Reset()

			lastRuneIsDigital = true // флаг, что предыдущий символ является число
			countRepeatLiteral = 0
		} else {
			// обнуляем литерал, если он задвоился
			if bufferStr.String() == "\\" {
				bufferStr.Reset()
				bufferStr.WriteRune(value)
				countRepeatLiteral++ // кол-во повторений
				continue
			}

			// сохраняем в буфере с повторением 1 раз
			resultStr.WriteString(strings.Repeat(bufferStr.String(), 1))
			bufferStr.Reset()

			bufferStr.WriteRune(value) // сохраняем в буффер текущий символ
			lastRuneIsDigital = false  // флаг, что предыдущий символ является число
		}
	}

	// проверка на последний символ
	if bufferStr.Len() > 0 {
		resultStr.WriteString(strings.Repeat(bufferStr.String(), 1))
	}

	return resultStr.String(), nil
}
