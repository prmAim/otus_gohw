package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Открываем исходный файл
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer func() {
		if err := fromFile.Close(); err != nil {
			fmt.Println("Failed to close file:", err)
		}
	}()

	// Получаем размер файла
	fileInfo, err := fromFile.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// Проверяем, что offset не превышает размер файла
	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	// Устанавливаем смещение
	_, err = fromFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	// Создаем файл назначения
	toFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func() {
		if err := toFile.Close(); err != nil {
			fmt.Println("Failed to create file:", err)
		}
	}()

	// Если limit равен 0, копируем весь файл
	if limit == 0 || limit > fileSize-offset {
		limit = fileSize - offset
	}

	// Копируем данные
	_, err = io.CopyN(toFile, fromFile, limit)
	return err
}
