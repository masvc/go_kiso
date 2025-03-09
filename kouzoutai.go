//構造体についての説明と練習をここには書く

package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

func kouzoutai() {
	var score Student
	score.name = "John"
	score.math = 80
	score.english = 90
	fmt.Println(score)

	// 以下でもOK
	// s := Student{name: "John", math: 80, english: 90}
	// fmt.Println(s)
}
