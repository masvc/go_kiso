package main

import (
	"fmt"
	"reflect"
)

// mainから始まり、ここが始まりになるのでエントリーポイントという、ここに他の関数を作成した場合は入れる
func main() {
	fmt.Println("Good Morning")
	fmt.Println("Good Afternoon")
	fmt.Println("Good Evening")
	var_test()
	type_test()
	math()
}

// var_test関数では以下の内容を練習しています：
// - 変数の基本的な宣言方法
// - :=を使った省略形の変数宣言
// - 複数の変数を一度に宣言する方法
// - reflect.TypeOfを使用した変数の型確認
func var_test() {
	var number int
	number = 1
	fmt.Println(number)

	num := 2
	fmt.Println(num)

	num2 := 3
	fmt.Println(num2)

	var (
		number1 int
		number2 int
	)

	number1 = 1
	number2 = 2
	fmt.Println(number1)
	fmt.Println(reflect.TypeOf(number1))
	fmt.Println(number2)
	fmt.Println(reflect.TypeOf(number2))
}

// type_test関数では以下の内容を練習しています：
// - 固定長配列の宣言と使用方法
// - 要素数を自動で数える配列の宣言 [...]
// - 多次元配列の使い方
func type_test() {
	// 固定長配列の宣言と値の変更
	hairetsu := [3]string{"Python", "Ruby", "Go"}
	hairetsu[1] = "JavaScript"

	fmt.Println(hairetsu[0])
	fmt.Println(hairetsu[1])
	fmt.Println(hairetsu[2])

	// 要素数を自動で数える配列
	yousomugen := [...]string{"HTML", "CSS", "PHP"}

	fmt.Println(yousomugen[0])
	fmt.Println(yousomugen[1])
	fmt.Println(yousomugen[2])

	// 2x2の2次元配列
	tajigenhairetsu := [2][2]string{
		{"Windows", "Mac"},
		{"Linux", "Android"},
	}

	fmt.Println(tajigenhairetsu[0][0])
	fmt.Println(tajigenhairetsu[0][1])
	fmt.Println(tajigenhairetsu[1][0])
	fmt.Println(tajigenhairetsu[1][1])
}

func math() {
	x := 10
	y := 2

	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y || x > y)
	fmt.Println(x == y && x > y)

	// インクリメントとデクリメント
	// インクリメントは1を加算すること
	// デクリメントは1を減算すること

	x++
	fmt.Println(x)
	x--
	fmt.Println(x)

	// 論理演算子
	// 論理和（OR）: ||
	// 論理積（AND）: &&
	// 否定（NOT）: !

	fmt.Println(true || false)
	fmt.Println(true && false)
	fmt.Println(!true)

	// 条件分岐
	// if文
	// if文は条件式がtrueの場合に実行される
	// if文は条件式がfalseの場合に実行されない

	if x > y {
		fmt.Println("xはyより大きい")
	} else if x < y {
		fmt.Println("xはyより小さい")
	} else {
		fmt.Println("xはyと等しい")
	}

	// switch文
	// switch文は条件式がtrueの場合に実行される
	// switch文は条件式がfalseの場合に実行されない

	switch x {
	case 1:
		fmt.Println("xは1")
	case 2:
		fmt.Println("xは2")
	}
}
