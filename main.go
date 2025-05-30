package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	
	file, err := os.Open("data_prog_contest_problem_2.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Файл пуст")
		os.Exit(1)
	}
	
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Ошибка чтения числа n:", err)
		os.Exit(1)
	}
	
	numbers := make([]int, n)
	scanner.Scan()
	str := scanner.Text()
	fields := strings.Fields(str)
	
	for i := 0; i < n; i++ {
		numbers[i], _ = strconv.Atoi(fields[i])
	}
	
	res := Foo(numbers)
	if res == -1 {
		fmt.Println("NONE")
	} else {
		fmt.Printf("%d\n", res)
	}
}

func Foo(numbers []int) int {
	if len(numbers) < 26 {
		return -1
	}
	alf := make([]bool, 27) 
	for _, num := range numbers {
		if num >= 1 && num <= 26 {
			alf[num] = true
		}
	}
	for i := 1; i <= 26; i++ {
		if !alf[i] {
			return -1
		}
	}
	minLen := len(numbers)
	count := make([]int, 27)
	uniqueCount := 0
	left := 0
	for right := 0; right < len(numbers); right++ {
		num := numbers[right]
		if num >= 1 && num <= 26 {
			if count[num] == 0 {
				uniqueCount++
			}
			count[num]++
			for uniqueCount == 26 {
				if right-left+1 < minLen {
					minLen = right - left + 1
				}
				leftNum := numbers[left]
				if leftNum >= 1 && leftNum <= 26 {
					count[leftNum]--
					if count[leftNum] == 0 {
						uniqueCount--
					}
				}
				left++
			}
		}
	}
	return minLen
}