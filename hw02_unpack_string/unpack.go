package hw02unpackstring

import (
	"errors"
	"fmt"
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

	// Проверяем, что первый символ это число
	_, err := strconv.Atoi(string(strIn[0]))
	if err == nil { // Если нет ошибки, значит это число
		return "", ErrInvalidString
	}

	// Преобразуем строку в срез runes
	runes := []rune(strIn)

	var bufferStr strings.Builder
	var resultStr strings.Builder
	countNumericRune := 0

	// Выводим содержимое runes
	for num, value := range runes {
		fmt.Printf("Индекс: %d, Символ: %c, Код: %d\n", num, value, value)

		_, err := strconv.Atoi(string(value))
		if err != nil {
			bufferStr.WriteRune(value)
			countNumericRune = 0
		} else {
			countNumericRune++
			fmt.Printf(" num= %d, value =  %c, intValue = %d, rune = %d\n", num, value, int(value-'0'), countNumericRune)

			// если предыдущий символ был цифрой, то сообщаем об ошибке
			if countNumericRune > 1 {
				return "", ErrInvalidString
			}
			resultStr.WriteString(strings.Repeat(bufferStr.String(), int(value-'0')))
			bufferStr.Reset()
		}
	}
	return resultStr.String(), nil
}
