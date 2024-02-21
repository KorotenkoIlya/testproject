package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pervyi, vtoroi *int
var znaki = map[string]func() int{
	"+": func() int { return *pervyi + *vtoroi },
	"-": func() int { return *pervyi - *vtoroi },
	"*": func() int { return *pervyi * *vtoroi },
	"/": func() int { return *pervyi / *vtoroi },
}
var inputdata []string

const (
	Oshibka  = "Ошибка - строка не является математической операцией"
	Oshibka0 = "Ошибка - не удовлетворяет заданию(только одна операция)"
	Oshibka1 = "Ошибка - может использовать только одну систему исчисления"
	Oshibka2 = "Ошибка - в римской нет отрицательных чисел"
	Oshibka3 = "Ошибка - в римской нет числа равному 0"
	Oshibka4 = "Ошибка - не удовлетворяет заданию, принимаются только цифры от 1 до 10"
)

var fromromantoarabic = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var arabictoroman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

func arabictoromans(romanresult int) {
	var romanchislo string
	if romanresult == 0 {
		panic(Oshibka3)
	} else if romanresult < 0 {
		panic(Oshibka2)
	}
	for romanresult > 0 {
		for _, elem := range arabictoroman {
			for i := elem; i <= romanresult; {
				for index, value := range fromromantoarabic {
					if value == elem {
						romanchislo += index
						romanresult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanchislo)
}
func osn(s string) {
	var znak string
	var skolko uint
	chisla := make([]int, 0)
	romans := make([]string, 0)
	fromromans := make([]int, 0)
	for i := range znaki {
		for _, k := range s {
			if i == string(k) {
				znak += i
				inputdata = strings.Split(s, znak)
			}
		}
	}
	if len(znak) > 1 {
		panic(Oshibka0)
	} else if len(znak) < 1 {
		panic(Oshibka)
	}
	for _, elem := range inputdata {
		el, err := strconv.Atoi(elem)
		if err != nil {
			skolko++
			romans = append(romans, elem)
		} else {
			chisla = append(chisla, el)
		}
	}
	if skolko == 1 {
		panic(Oshibka1)
	} else if skolko == 0 {
		errcheck := chisla[0] > 0 && chisla[0] < 11 && chisla[1] > 0 && chisla[1] < 11
		if val, ok := znaki[znak]; ok && errcheck == true {
			pervyi, vtoroi = &chisla[0], &chisla[1]
			fmt.Println(val())
		} else {
			panic(Oshibka4)
		}
	} else if skolko == 2 {
		for _, elem := range romans {
			if val, ok := fromromantoarabic[elem]; ok && val > 0 && val < 11 {
				fromromans = append(fromromans, val)
			} else {
				panic(Oshibka4)
			}
		}
		if val, ok := znaki[znak]; ok {
			pervyi, vtoroi = &fromromans[0], &fromromans[1]
			arabictoromans(val())
		}
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		osn(strings.ToUpper(strings.TrimSpace(s)))
	}
}
