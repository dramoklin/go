package main

import "fmt"

const appleOne float64 = 5.99

const peareOne int = 7
const myMoney int = 23

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
	fmt.Println("Мы можем купить такое кол-во яблок", myMoney/int(appleOne))

}

func fourthQuestion(peareOne int, appleOne float64) {
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Println("Мы не можем купить такое кол-во груш и яблок", peareOne+peareOne+int(appleOne)+int(appleOne))

}
