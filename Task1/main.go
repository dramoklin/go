package main

import "fmt"

var appleOne float64 = 5.99

var peareOne int = 7

var myMoney int = 23

func main() {
	firstQuestion(peareOne, appleOne)
	secondQuestion(peareOne, myMoney)
	thirdQuestion(appleOne, myMoney)
	fourthQuestion(peareOne, appleOne)

}

func firstQuestion(peareOne int, appleOne float64) {
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Нужно столько денег для груш", peareOne*8)
	fmt.Println("Нужно столько денег для яблок", appleOne*9)

}

func secondQuestion(peareOne int, myMoney int) {
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Println("Мы можем купить такое кол-во груш", myMoney/peareOne)

}

func thirdQuestion(appleOne float64, myMoney int) {
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Println("Мы можем купить такое кол-во яблок", float64(myMoney)/float64(appleOne))

}

func fourthQuestion(peareOne int, appleOne float64) {
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Println("Стоимость двух груш ", peareOne+peareOne)
	fmt.Println("Стоимость двух яблок ", appleOne+appleOne)
	fmt.Println("Мы не можем купить 2 груши и 2 яблока")

}
