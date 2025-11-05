package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCompareString(line string, numFields int, numChars int, ignoreCase bool) string {
	s := line

	if numFields > 0 {
		fields := strings.Fields(s)
		if numFields >= len(fields) {
			s = ""
		} else {
			s = strings.Join(fields[numFields:], " ")
		}
	}

	if numChars > 0 {
		if numChars >= len(s) {
			s = ""
		} else {
			s = s[numChars:]
		}
	}

	if ignoreCase {
		s = strings.ToLower(s)
	}

	return s
}

func process(mas []string, mode string, fFlag, sFlag int, iFlag bool, outputFile *os.File) {
	if len(mas) == 0 {
		return
	}

	k := 1
	for i := 0; i < len(mas)-1; i++ {
		curKey := getCompareString(mas[i], fFlag, sFlag, iFlag)
		nextKey := getCompareString(mas[i+1], fFlag, sFlag, iFlag)

		if curKey == nextKey {
			k++
		} else {
			if mode == "c" {
				fmt.Fprintf(outputFile, "%d %s\n", k, mas[i])
			} else if mode == "d" {
				if k >= 2 {
					fmt.Fprintln(outputFile, mas[i])
				}
			} else if mode == "u" {
				if k == 1 {
					fmt.Fprintln(outputFile, mas[i])
				}
			} else {
				fmt.Fprintln(outputFile, mas[i])
			}
			k = 1
		}
	}

	 
	if mode == "c" {
		fmt.Fprintf(outputFile, "%d %s\n", k, mas[len(mas)-1])
	} else if mode == "d" {
		if k >= 2 {
			fmt.Fprintln(outputFile, mas[len(mas)-1])
		}
	} else if mode == "u" {
		if k == 1 {
			fmt.Fprintln(outputFile, mas[len(mas)-1])
		}
	} else {
		fmt.Fprintln(outputFile, mas[len(mas)-1])
	}
}

func main() {
	var cFlag = flag.Bool("c", false, "подсчитать количество")
	var dFlag = flag.Bool("d", false, "вывести повторяющиеся")
	var uFlag = flag.Bool("u", false, "вывести уникальные")
	var iFlag = flag.Bool("i", false, "игнорировать регистр")
	var fFlag = flag.Int("f", 0, "пропустить первые N полей")
	var sFlag = flag.Int("s", 0, "пропустить первые N символов")

	flag.Parse()

	if (*cFlag && *dFlag) || (*cFlag && *uFlag) || (*dFlag && *uFlag) {
		fmt.Fprintln(os.Stderr, "uniq: нельзя использовать -c, -d, -u вместе")
		os.Exit(1)
	}

	args := flag.Args()

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

	scanner := bufio.NewScanner(inputFile)
	var mas []string
	for scanner.Scan() {
		mas = append(mas, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}

	mode := ""
	if *cFlag {
		mode = "c"
	} else if *dFlag {
		mode = "d"
	} else if *uFlag {
		mode = "u"
	}

	process(mas, mode, *fFlag, *sFlag, *iFlag, outputFile)
}