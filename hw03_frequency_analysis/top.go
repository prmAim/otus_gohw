package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type repeatText struct {
	count int
	text  string
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

	repWords := []repeatText{}

	// разбиваем на массив слов, убирая лишние пробелы
	txt := strings.Fields(strIn)
	isPresentWord := false

	for _, valTxt := range txt {
		isPresentWord = false

		for idx, valRepWord := range repWords {
			if valTxt == valRepWord.text {
				repWords[idx].count++
				isPresentWord = true
				break
			}
		}

		if !isPresentWord {
			repWords = append(repWords, repeatText{count: 1, text: valTxt})
		}
	}

	// Сортировка
	sort.Sort(ByCountAndText(repWords))

	// копируем
	repWords10 := []string{}

	for idx := 0; idx < 10 && idx < len(repWords); idx++ {
		repWords10 = append(repWords10, repWords[idx].text)
	}

	return repWords10
}
