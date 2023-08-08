package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanArray = []string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
	"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
	"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
	"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
	"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
	"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func main() {
	fmt.Print("Введите пример в арабской или римской СС:  ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	operatorIndex := strings.IndexAny(text, "+-*/")
	if operatorIndex == -1 {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	}
	sign := string(text[operatorIndex])
	checkSign(sign)
	array := strings.Split(text, sign)
	if len(array) != 2 {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if array[1] == "" {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	}
	array[0] = strings.ReplaceAll(array[0], " ", "")
	array[1] = strings.ReplaceAll(array[1], " ", "")

	if (itRomanOrNot(array[0]) && itRomanOrNot(array[1])) == true {
		number1 := convertToArabian(array[0])
		number2 := convertToArabian(array[1])
		finalRes := arabianCounter(number1, number2, sign)
		if finalRes < 0 {
			panic("Вывод ошибки, так как в римской системе нет отрицательных чисел1")
		} else {
			fmt.Println(finalRes)
		}
	} else if (itRomanOrNot(array[0]) && itRomanOrNot(array[1])) == false {
		fmt.Println(arabianCounter(array[0], array[1], sign))
	} else {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
}

func arabianCounter(number1 string, number2 string, sign string) int {
	var result int
	num1, err := strconv.Atoi(number1)
	if err != nil {
		panic(err)
	}
	num2, err1 := strconv.Atoi(number2)
	if err1 != nil {
		panic(err)
	}
	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		panic("На вход принимаются числа от 1 до 10 включительно")
	}
	switch sign {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}

func checkSign(sign string) {
	if sign != "+" && sign != "-" && sign != "*" && sign != "/" {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
}

func checkExpression(expression []string) {
	if len(expression) > 3 {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}
}

func itRomanOrNot(number string) bool {
	for i := 0; i < len(romanArray); i++ {
		if number == romanArray[i] {
			return true
		}
	}
	return false
}

func convertToArabian(number string) string {

	for i := 0; i < len(romanArray); i++ {
		if number == romanArray[i] {
			number = strconv.Itoa(i)
			return number
		}
	}
	return ""
}
