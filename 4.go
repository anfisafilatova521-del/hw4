package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Функция для обработки -f, -s и -i вместе
func getCompareString(line string, numFields int, numChars int, ignoreCase bool) string {
	s := line

	// -f: пропустить первые N полей
	if numFields > 0 {
		fields := strings.Fields(s)
		if numFields >= len(fields) {
			s = ""
		} else {
			s = strings.Join(fields[numFields:], " ")
		}
	}

	// -s: пропустить первые N символов
	if numChars > 0 {
		if numChars >= len(s) {
			s = ""
		} else {
			s = s[numChars:]
		}
	}

	// -i: игнорировать регистр
	if ignoreCase {
		s = strings.ToLower(s)
	}

	return s
}

func main() {
	// Объявляем флаги
	var cFlag = flag.Bool("c", false, "подсчитать количество")
	var dFlag = flag.Bool("d", false, "вывести повторяющиеся")
	var uFlag = flag.Bool("u", false, "вывести уникальные")
	var iFlag = flag.Bool("i", false, "игнорировать регистр")
	var fFlag = flag.Int("f", 0, "пропустить первые N полей")
	var sFlag = flag.Int("s", 0, "пропустить первые N символов")

	// Парсим флаги
	flag.Parse()

	// Проверка: нельзя использовать -c, -d, -u вместе
	if (*cFlag && *dFlag) || (*cFlag && *uFlag) || (*dFlag && *uFlag) {
		fmt.Fprintln(os.Stderr, "uniq: нельзя использовать -c, -d, -u вместе")
		os.Exit(1)
	}

	// Получаем аргументы (файлы)
	args := flag.Args()

	// Открываем входной файл или используем stdin
	var inputFile *os.File
	var err error
	if len(args) > 0 {
		inputFile, err = os.Open(args[0])
		if err != nil {
			log.Fatalf("Не удалось открыть файл: %v", err)
		}
		defer inputFile.Close()
	} else {
		inputFile = os.Stdin
	}

	// Открываем выходной файл или используем stdout
	var outputFile *os.File
	if len(args) > 1 {
		outputFile, err = os.Create(args[1])
		if err != nil {
			log.Fatalf("Не удалось создать файл: %v", err)
		}
		defer outputFile.Close()
	} else {
		outputFile = os.Stdout
	}

	// Читаем все строки
	scanner := bufio.NewScanner(inputFile)
	var mas []string
	for scanner.Scan() {
		mas = append(mas, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}

	// Если нет строк — выходим
	if len(mas) == 0 {
		return
	}

	// Обработка в зависимости от флагов
	if *cFlag {
		// для -c
		k := 1
		for i := 0; i < len(mas)-1; i++ {
			key1 := getCompareString(mas[i], *fFlag, *sFlag, *iFlag)
			key2 := getCompareString(mas[i+1], *fFlag, *sFlag, *iFlag)
			if key1 == key2 {
				k++
			} else {
				fmt.Fprintf(outputFile, "%d %s\n", k, mas[i])
				k = 1
			}
		}
		fmt.Fprintf(outputFile, "%d %s\n", k, mas[len(mas)-1])
	} else if *dFlag {
		// для -d
		k := 1
		for i := 0; i < len(mas)-1; i++ {
			key1 := getCompareString(mas[i], *fFlag, *sFlag, *iFlag)
			key2 := getCompareString(mas[i+1], *fFlag, *sFlag, *iFlag)
			if key1 == key2 {
				k++
			} else {
				if k >= 2 {
					fmt.Fprintln(outputFile, mas[i])
				}
				k = 1
			}
		}
		if k >= 2 {
			fmt.Fprintln(outputFile, mas[len(mas)-1])
		}
	} else if *uFlag {
		// для -u
		k := 1
		for i := 0; i < len(mas)-1; i++ {
			key1 := getCompareString(mas[i], *fFlag, *sFlag, *iFlag)
			key2 := getCompareString(mas[i+1], *fFlag, *sFlag, *iFlag)
			if key1 == key2 {
				k++
			} else {
				if k == 1 {
					fmt.Fprintln(outputFile, mas[i])
				}
				k = 1
			}
		}
		if k == 1 {
			fmt.Fprintln(outputFile, mas[len(mas)-1])
		}
	} else {
		// без флагов — просто вывод уникальных строк
		k := 1
		for i := 0; i < len(mas)-1; i++ {
			key1 := getCompareString(mas[i], *fFlag, *sFlag, *iFlag)
			key2 := getCompareString(mas[i+1], *fFlag, *sFlag, *iFlag)
			if key1 == key2 {
				k++
			} else {
				fmt.Fprintln(outputFile, mas[i])
				k = 1
			}
		}
		fmt.Fprintln(outputFile, mas[len(mas)-1])
	}
}