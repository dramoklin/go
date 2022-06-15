package main

import "fmt"

var appleOne float64 = 5.99

var peareOne float64 = 7

var myMoney float64 = 23

func main() {
	firstQuestion(peareOne, appleOne)
	secondQuestion(peareOne, myMoney)
	thirdQuestion(appleOne, myMoney)
	fourthQuestion(peareOne, appleOne)

}

func firstQuestion(peareOne float64, appleOne float64) {
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Нужно столько денег для груш", peareOne*8)
	fmt.Println("Нужно столько денег для яблок", appleOne*9)

}

func secondQuestion(peareOne float64, myMoney float64) {
	fmt.Println("Скільки груш ми можемо купити?")
	fmt.Println("Мы можем купить такое кол-во груш", myMoney/peareOne)

}

func thirdQuestion(appleOne float64, myMoney float64) {
	fmt.Println("Скільки яблук ми можемо купити?")
	fmt.Println("Мы можем купить такое кол-во яблок", float64(myMoney)/float64(appleOne))

}

func fourthQuestion(peareOne float64, appleOne float64) {
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	sumPeare := peareOne + peareOne
	sumApple := appleOne + appleOne
	sumAll := sumApple + sumPeare
	if sumAll > myMoney {
		fmt.Println("Мы не можем себе это позволить")
	} else {
		fmt.Println("Мы  можем себе это позволить")
	}

}
