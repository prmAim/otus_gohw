package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	// Парсим аргументы командной строки
	from := flag.String("from", "", "path to source file")
	to := flag.String("to", "", "path to destination file")
	offset := flag.Int64("offset", 0, "offset in source file")
	limit := flag.Int64("limit", 0, "number of bytes to copy")
	flag.Parse()

	// Проверяем обязательные аргументы
	if *from == "" || *to == "" {
		fmt.Println("Usage: go run main.go -from <source> -to <destination> [-offset <offset>] [-limit <limit>]")
		return
	}

	// Выполняем копирование
	err := Copy(*from, *to, *offset, *limit)
	if err != nil {
		log.Fatalf("Copy failed: %v", err)
	}

	fmt.Println("Copy completed successfully!")
}
