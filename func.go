package main

import "fmt"

// func_test は関数の基本的な使い方を示すための関数です。
// 引数として2つの整数を受け取り、その積を出力します。
//
// 注意点：
// - この関数はmainパッケージに属しています
// - main.goから呼び出すことができます
// - 同じパッケージ内の関数は、別ファイルでも直接呼び出し可能です
//
// 引数：
// - x int: 1つ目の整数
// - y int: 2つ目の整数
//
// 使用例：
// func_test(5, 3) // 出力: 15
func func_test(x int, y int) {
	fmt.Println(x * y)
}

// func_return は戻り値のある関数の例です。
// 引数として2つの整数を受け取り、その積を返します。
//
// 引数：
// - x int: 1つ目の整数
// - y int: 2つ目の整数
//
// 戻り値：
// - int: xとyの積
//
// 使用例：
// result := func_return(5, 3) // result = 15
func func_return(x int, y int) int {
	fmt.Println(x * y) // 計算結果を表示
	return x * y
}

// func_return2 は複数の戻り値がある関数の例です。
// 引数として2つの整数を受け取り、その和と積を返します。
//
// 引数：
// - x int: 1つ目の整数
// - y int: 2つ目の整数
//
// 戻り値：
// - int: xとyの和
// - int: xとyの積
//
// 使用例：
// sum, product := func_return2(5, 3) // sum = 8, product = 15
func func_return2(x int, y int) (int, int) {
	fmt.Printf("和: %d, 積: %d\n", x+y, x*y) // 計算結果を表示
	return x + y, x * y
}
