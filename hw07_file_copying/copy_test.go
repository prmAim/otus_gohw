package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name     string
		fromPath string
		toPath   string
		offset   int64
		limit    int64
		wantErr  bool
	}{
		{
			name:     "Copy entire file (offset=0, limit=0)",
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/output1.txt",
			offset:   0,
			limit:    0,
			wantErr:  false,
		},
		{
			name:     "Copy with offset=100, limit=500",
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/output2.txt",
			offset:   100,
			limit:    500,
			wantErr:  false,
		},
		{
			name:     "Copy with offset=1000, limit=0 (copy till EOF)",
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/output3.txt",
			offset:   1000,
			limit:    0,
			wantErr:  false,
		},
		{
			name:     "Offset exceeds file size",
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/output4.txt",
			offset:   100000,
			limit:    0,
			wantErr:  true,
		},
		{
			name:     "Copy with limit=10",
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/output5.txt",
			offset:   0,
			limit:    10,
			wantErr:  false,
		},
		{
			name:     "Copy with offset=500 and limit=1000",
			fromPath: "testdata/input.txt",
			toPath:   "/tmp/output6.txt",
			offset:   500,
			limit:    1000,
			wantErr:  false,
		},
		{
			name:     "Copy non-existent file",
			fromPath: "testdata/nonexistent.txt",
			toPath:   "/tmp/output7.txt",
			offset:   0,
			limit:    0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCopyScenario(t, tt.fromPath, tt.toPath, tt.offset, tt.limit, tt.wantErr)
		})
	}
}

func readExpectedData(t *testing.T, fromPath string, offset, limit int64) []byte {
	t.Helper() // Помечаем функцию как вспомогательную

	fromFile, err := os.Open(fromPath)
	if err != nil {
		t.Fatalf("Failed to open source file: %v", err)
	}
	defer func() {
		if err := fromFile.Close(); err != nil {
			fmt.Println("Failed to close file:", err)
		}
	}()

	// Устанавливаем смещение
	_, err = fromFile.Seek(offset, 0)
	if err != nil {
		t.Fatalf("Failed to seek in source file: %v", err)
	}

	// Читаем ожидаемые данные
	var expectedData []byte
	if limit > 0 {
		expectedData = make([]byte, limit)
		_, err = fromFile.Read(expectedData)
		if err != nil {
			t.Fatalf("Failed to read from source file: %v", err)
		}
	} else {
		expectedData, err = io.ReadAll(fromFile)
		if err != nil {
			t.Fatalf("Failed to read from source file: %v", err)
		}
	}

	return expectedData
}

func compareFileContent(t *testing.T, toPath string, expectedData []byte) {
	t.Helper() // Помечаем функцию как вспомогательную

	copiedData, err := os.ReadFile(toPath)
	if err != nil {
		t.Fatalf("Failed to read copied file: %v", err)
	}

	// Сравниваем данные
	if string(copiedData) != string(expectedData) {
		t.Errorf("Copied file content does not match expected content")
	}
}

func testCopyScenario(t *testing.T, fromPath, toPath string, offset, limit int64, wantErr bool) {
	t.Helper() // Помечаем функцию как вспомогательную

	err := Copy(fromPath, toPath, offset, limit)
	if (err != nil) != wantErr {
		t.Errorf("Copy() error = %v, wantErr %v", err, wantErr)
		return
	}

	// Если ошибка не ожидалась, проверяем содержимое файла
	if !wantErr {
		expectedData := readExpectedData(t, fromPath, offset, limit)
		compareFileContent(t, toPath, expectedData)
	}
}
