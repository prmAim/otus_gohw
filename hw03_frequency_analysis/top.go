package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type repeatText struct {
	text  string
	count int
}

type ByCountAndText []repeatText

func (a ByCountAndText) Len() int {
	return len(a)
}

func (a ByCountAndText) Less(i, j int) bool {
	// Сначала сортируем по count
	if a[i].count == a[j].count {
		// Если count равны, сортируем по text
		return a[i].text < a[j].text
	}
	return a[i].count > a[j].count
}

func (a ByCountAndText) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Top10(strIn string) []string {
	// Place your code here.
	// Если передана на вход пустая строка, то ошибка
	if len(strIn) == 0 {
		return nil
	}

	// разбиваем на массив слов, убирая лишние пробелы
	txt := strings.Fields(strIn)

	repWords := make(map[string]int)

	for _, valTxt := range txt {
		repWords[valTxt]++
	}

	// Сортировка
	// Создание среза для хранения пар ключ-значение
	sortedWords := make(ByCountAndText, 0, len(repWords))

	// Заполнение среза парами ключ-значение
	for k, v := range repWords {
		sortedWords = append(sortedWords, repeatText{k, v})
	}

	// Сортировка среза
	sort.Sort(sortedWords)

	// копируем
	size := 10
	if len(sortedWords) < 10 {
		size = len(sortedWords)
	}
	repWords10 := make([]string, size)

	for idx := 0; idx < len(sortedWords) && idx < 10; idx++ {
		repWords10[idx] = sortedWords[idx].text
	}

	return repWords10
}
