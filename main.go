package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

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

func ReadTextFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)
}

func FindMinMaxMedian(args []int) (int, int, float64) {
	args = BubleSort(args)
	var sum float64 = 0
	for _, v := range args {
		sum += float64(v)
	}
	return args[0], args[len(args)-1], sum / float64(len(args))
}

func BubleSort(iniSlice []int) []int {
	for i := 1; i < len(iniSlice); i++ {
		if iniSlice[i] < iniSlice[0] {
			iniSlice = swapToFirst(iniSlice, i)
			continue
		}
		for j := 1; j < i; j++ {
			if iniSlice[i] < iniSlice[j] {
				iniSlice = swap(iniSlice, j, i)
				break
			}
		}
	}
	return iniSlice
}

func swap(aSlice []int, i, j int) []int {
	newSlice := []int{}
	newSlice = append(newSlice, aSlice[:i]...)
	newSlice = append(newSlice, aSlice[j])
	newSlice = append(newSlice, aSlice[i])
	newSlice = append(newSlice, aSlice[i+1:j]...)
	newSlice = append(newSlice, aSlice[j+1:]...)

	return newSlice
}
func swapToFirst(aSlice []int, i int) []int {
	newSlice := []int{}
	newSlice = append(newSlice, aSlice[i])
	newSlice = append(newSlice, aSlice[0:i]...)
	newSlice = append(newSlice, aSlice[i+1:]...)
	return newSlice
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
func contains(s []int, element int) bool {
	for _, v := range s {
		if v == element {
			return true
		}
	}

	return false
}
func main() {
	aSlice := []int{2, 3, 6, 15, 5, 17, 34}
	fmt.Println(PrimeNumberDetector(aSlice))
}
