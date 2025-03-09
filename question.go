package main

import "fmt"

// 【練習問題：基本的な計算関数】
// この関数は2つの整数を受け取り、その和を計算して表示します
//
// 【解説】
// 1. func cal(x int, y int) で2つの整数型引数を受け取る関数を定義
// 2. calsum := x + y で引数の和を計算
// 3. fmt.Println(calsum) で計算結果を表示
//
// 【発展課題のヒント】
// - 引数の型を float64 に変更すると小数点の計算もできます
// - return文を追加すると計算結果を返せます
// - 別の演算（引き算、掛け算、割り算）も試してみましょう
func cal(x int, y int) {
	calsum := x + y
	fmt.Println(calsum)
}
