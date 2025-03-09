package main

import "fmt"

//for文の練習

func for_test() {
	// for文は繰り返し処理を行うための文
	// 繰り返し処理は、同じ処理を繰り返し行う場合に使う
	// 例：配列の要素を順番に処理する場合など

	// ※注意：無限ループを避けるために、必ず終了条件を設定し、
	// インクリメント/デクリメントの処理を忘れないようにしましょう

	// 繰り返し処理の基本形
	for index := 0; index <= 4; index++ {

		// break文は繰り返し処理を中断するための文
		// 例：for文を途中で抜ける場合など
		if index == 3 {
			continue
		}
		fmt.Println(index)
	}

	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i, v := range arr {
		sum += v
		fmt.Println(i, v, sum)
	}
	fmt.Println(sum)

}

func for_training() {
	// 練習問題：1から100までの数字を足し合わせて、その合計を表示するプログラムを作成してください。
	// ただし、3の倍数と5の倍数は加算しないようにしてください。

	sum := 0
	for i := 1; i <= 10; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
