package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type romanAndArabicNumeral struct {
	roman  string
	arabic int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var operator string
	var num1 int
	var num2 int
	var isRoman bool
	var operation func(num1, num2 int) int

	for exit := false; exit != true; {
		fmt.Printf("Добро пожаловать в калькулятор! \n" +
			"Введите математическое выражение в котором два операнда и один оператор (+, -, /, *)\n" +
			"Или же введите команду (exit) для выхода:\n")

		inputText, _ := reader.ReadString('\n')
		if strings.TrimSpace(inputText) == "exit" {
			exit = true
			break
		}

		num1, num2, operator, isRoman = GetOperandsAndOperator(inputText)

		switch operator {
		case "+":
			operation = func(num1, num2 int) int { return num1 + num2 }
		case "-":
			operation = func(num1, num2 int) int { return num1 - num2 }
		case "/":
			operation = func(num1, num2 int) int { return num1 / num2 }
		case "*":
			operation = func(num1, num2 int) int { return num1 * num2 }
		default:
			panic("Недопустимый оператор!")
		}

		action(num1, num2, operation, isRoman)
		fmt.Println("Для очистки консоли и продолжения работы нажмите ввод...")
		fmt.Scanf("\n")
		CallClear()
	}
}

func action(n1 int, n2 int, operation func(int, int) int, isRoman bool) {
	var result string
	value := operation(n1, n2)

	if isRoman && value < 1 {
		panic("В римской системе не существует отрицательных чисел и нуля!")
	} else if isRoman {
		result = ConvertArabicToRoman(value)
	} else {
		result = strconv.Itoa(value)
	}
	fmt.Println(result)
}

func GetOperandsAndOperator(expression string) (firstNum, secondNum int, operator string, isRoman bool) {
	var separateExpression = strings.Split(expression, " ")

	if len(separateExpression) == 3 {
		var isFirstNumRoman bool
		var isSecondNumRoman bool
		firstOperand := strings.TrimSpace(separateExpression[0])
		secondOperand := strings.TrimSpace(separateExpression[2])

		firstNum, isFirstNumRoman = TryConvertToNumber(firstOperand)
		secondNum, isSecondNumRoman = TryConvertToNumber(secondOperand)
		operator = separateExpression[1]

		if isFirstNumRoman != isSecondNumRoman {
			panic("Запрещено проводить операции над числами из разных систем!")
		} else {
			isRoman = isFirstNumRoman
		}
	} else {
		panic("Неккоректный ввод!")
	}
	return
}

func TryConvertToNumber(operand string) (number int, isRoman bool) {
	var value, ok = strconv.Atoi(operand)
	if ok != nil {
		number = ConvertRomanToArabic(operand)
		isRoman = true
	} else if value > 0 && value <= 10 {
		number = value
		isRoman = false
	} else {
		panic("Недопустимое значение операнда/-ов")
	}
	return
}

func ConvertArabicToRoman(arabicNum int) (romanNum string) {
	numerals := [9]romanAndArabicNumeral{
		{
			roman: "C", arabic: 100,
		},
		{
			roman: "XC", arabic: 90,
		},
		{
			roman: "L", arabic: 50,
		},
		{
			roman: "XL", arabic: 40,
		},
		{
			roman: "X", arabic: 10,
		},
		{
			roman: "IX", arabic: 9,
		},
		{
			roman: "V", arabic: 5,
		},
		{
			roman: "IV", arabic: 4,
		},
		{
			roman: "I", arabic: 1,
		},
	}

	for arabicNum > 0 {
		for i := 0; i < len(numerals); i++ {
			if arabicNum >= numerals[i].arabic {
				arabicNum -= numerals[i].arabic
				romanNum += numerals[i].roman
				break
			}
		}
	}
	return romanNum
}

func ConvertRomanToArabic(romanNum string) (arabicNum int) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	value, ok := romanNumerals[romanNum]
	if ok {
		arabicNum = value
	} else {
		panic("Недопустимое число!")
	}

	return
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Ваша платформа не поддерживается! Я не могу очистить экран терминала :(")
	}
}
