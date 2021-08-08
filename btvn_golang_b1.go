package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

//Viết các API cộng trừ nhân chia
func MathOperator(operator string, operand1, operand2 int) (int, error) {
	switch operator {
	case "+":
		return operand1 + operand2, nil
	case "-":
		return operand1 - operand2, nil
	case "*":
		return operand1 * operand2, nil
	case "/":
		if operand2 != 0 {
			return operand1 / operand2, nil
		} else {
			err := fmt.Errorf("Can not divide for 0")
			return 0, err
		}
	}
	return 0, nil
}

//Đọc file text
func ReadTextFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}

//số lớn nhất, nhỏ nhất, trung bình
func FindMinMaxMedian(args []int) (int, int, float64) {
	args = InsertionSort(args)
	var sum float64 = 0
	for _, v := range args {
		sum += float64(v)
	}
	return args[0], args[len(args)-1], sum / float64(len(args))
}

func BubleSort(iniSlice []int) []int {
	for i := len(iniSlice) - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if iniSlice[j] < iniSlice[j-1] {
				iniSlice = swap(iniSlice, j)
				fmt.Println(iniSlice)
			}
		}
	}
	return iniSlice
}
func swap(iniSlice []int, j int) []int {
	newSlice := make([]int, 0, len(iniSlice))
	newSlice = append(newSlice, iniSlice[:j-1]...)
	newSlice = append(newSlice, iniSlice[j], iniSlice[j-1])
	newSlice = append(newSlice, iniSlice[j+1:]...)
	return newSlice
}
func InsertionSort(iniSlice []int) []int {
	for i := 1; i < len(iniSlice); i++ {
		if iniSlice[i] < iniSlice[0] {
			iniSlice = insertToFirst(iniSlice, i)
			continue
		}
		for j := 1; j < i; j++ {
			if iniSlice[i] < iniSlice[j] {
				iniSlice = normalInsert(iniSlice, j, i)
				break
			}
		}
	}
	return iniSlice
}

func normalInsert(aSlice []int, i, j int) []int {
	newSlice := []int{}
	newSlice = append(newSlice, aSlice[:i]...)
	newSlice = append(newSlice, aSlice[j])
	newSlice = append(newSlice, aSlice[i])
	newSlice = append(newSlice, aSlice[i+1:j]...)
	newSlice = append(newSlice, aSlice[j+1:]...)

	return newSlice
}
func insertToFirst(aSlice []int, i int) []int {
	newSlice := []int{}
	newSlice = append(newSlice, aSlice[i])
	newSlice = append(newSlice, aSlice[0:i]...)
	newSlice = append(newSlice, aSlice[i+1:]...)
	return newSlice
}

//Kiểm tra xem các giá trị ở file đầu vào có là số nguyên tố ko
func PrimeNumberDetector(aSlice []int) []int {
	_, max, _ := FindMinMaxMedian(aSlice)
	primerList := PrimeNumberGenerator(max)
	outputSlice := make([]int, 0)
	for _, v := range aSlice {
		if contains(primerList, v) {
			outputSlice = append(outputSlice, v)
		}
	}
	return outputSlice
}
func PrimeNumberGenerator(n int) []int {
	var list = make([]int, 0, 1000)
	list = append(list, 2)
mainlopp:
	for i := 3; i <= n; i++ {
		for _, v := range list {
			if i%v == 0 {
				continue mainlopp
			}
		}
		list = append(list, i)
	}
	return list
}

//Viết hàm kiểm tra xem với 1 số  giá trị cho trước thì có tồn tại trong file hay
func contains(s []int, element int) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}

	return false
}

//In ra file
func Writefile(input []int) {
	inputConverted := make([]string, 0, len(input))
	for _, v := range input {
		inputConverted = append(inputConverted, strconv.Itoa(v))
	}
	stringToWrite := strings.Join(inputConverted, ", ")
	ioutil.WriteFile("outputSlice.txt", []byte(stringToWrite), 0644)

}
func main() {
	aSlice := []int{6, 1, 4, 15, 5, 71, 34}
	fmt.Println(aSlice)
	fmt.Println(BubleSort(aSlice))
	Writefile(BubleSort(aSlice))
}
