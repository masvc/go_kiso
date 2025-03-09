//構造体についての説明と練習をここには書く

package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

// 戻り値の例を示す関数
// 【解説】
// 1. func getAverage(math, english float64) の部分は引数の定義
//   - math, english は引数の名前
//   - float64 は引数の型（小数を扱える数値型）
//
// 2. 最後の float64 は「この関数が返す値の型」を示している
// 3. return で実際の値を返す（ここでは計算結果を返している）
func getAverage(math, english float64) float64 {
	return (math + english) / 2
}

// 複数の戻り値を返す関数
// 【解説】
// 1. (string, float64) は「この関数が2つの値を返す」ことを示している
//   - 1つ目の戻り値はstring型（文字列）
//   - 2つ目の戻り値はfloat64型（小数）
//
// 2. return s.name, average のように、カンマで区切って複数の値を返せる
// 3. 受け取る側は name, avg := getScoreInfo(score) のように2つの変数で受け取る
func getScoreInfo(s Student) (string, float64) {
	average := (s.math + s.english) / 2
	return s.name, average
}

func kouzoutai() {
	var score Student
	score.name = "John"
	score.math = 80
	score.english = 90
	fmt.Println(score)

	// 平均点を計算
	// 【解説】getAverageは1つの値（float64型）を返すので、1つの変数で受け取る
	average := getAverage(score.math, score.english)
	fmt.Printf("%sさんの平均点は%.1f点です\n", score.name, average)

	// 複数の戻り値を受け取る
	// 【解説】getScoreInfoは2つの値を返すので、2つの変数で受け取る
	// もし平均点だけが必要な場合は、 _, avg := getScoreInfo(score) と書くこともできる
	// _ は「この戻り値は使わない」という意味
	name, avg := getScoreInfo(score)
	fmt.Printf("名前: %s, 平均点: %.1f\n", name, avg)

	// 以下でもOK
	// s := Student{name: "John", math: 80, english: 90}
	// fmt.Println(s)
}
