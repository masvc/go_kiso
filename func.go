package main

import "fmt"

//【Goの関数について】
// 関数とは、特定の処理をまとめた再利用可能なコードブロックです。
// Goの関数は以下の特徴があります：
// 1. func キーワードで始まります
// 2. 引数（パラメータ）を受け取ることができます
// 3. 戻り値を返すことができます（複数の戻り値も可能）
// 4. パッケージ内であれば、別ファイルからでも呼び出せます

// func_test は関数の基本的な使い方を示すための関数です。
// 引数として2つの整数を受け取り、その積を出力します。
//
// 【解説】
// 1. func func_test(x int, y int) の部分が関数の定義です
//   - func_test は関数名
//   - (x int, y int) は引数リスト。各引数に型を指定します
//
// 2. この関数は戻り値を持ちません（void関数）
//
// 3. 関数の中身（{}の中）が実際の処理内容です
//
// 使用例：
// func_test(5, 3) // 出力: 15
func func_test(x int, y int) {
	fmt.Println(x * y)
}

// func_return は戻り値のある関数の例です。
// 引数として2つの整数を受け取り、その積を返します。
//
// 【解説】
// 1. func func_return(x int, y int) int の最後のintが戻り値の型です
//
// 2. return文で値を返します
//   - return文は関数の実行を終了し、指定した値を呼び出し元に返します
//
// 3. 戻り値は変数に代入して使うことができます
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
// 【解説】
// 1. (int, int) は2つの戻り値があることを示します
//   - カッコの中に複数の型を書くことで、複数の値を返せます
//
// 2. return文も複数の値をカンマで区切って書きます
//
// 3. 呼び出し側では := で複数の変数に同時に代入できます
//   - 使わない戻り値がある場合は _ で無視できます
//     例：sum, _ := func_return2(5, 3) // 積の値は無視
//
// 使用例：
// sum, product := func_return2(5, 3) // sum = 8, product = 15
func func_return2(x int, y int) (int, int) {
	fmt.Printf("和: %d, 積: %d\n", x+y, x*y) // 計算結果を表示
	return x + y, x * y
}
