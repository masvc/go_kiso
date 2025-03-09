//【メソッドとは？】
// Goにはクラス概念がありませんが、構造体に関連付けた関数（メソッド）を定義できます。
// メソッドを使うと、データ（構造体）と、そのデータを処理する関数を紐付けることができます。

package main

import "fmt"

// StudentScore構造体の定義
// 【解説】
// この構造体は生徒の名前と科目の点数を保持します
type StudentScore struct {
	name          string
	math, english float64
}

// average メソッドの定義
// 【解説】
// 1. func (s StudentScore) average() の部分がメソッドの定義です
//   - (s StudentScore) はレシーバと呼ばれ、このメソッドが StudentScore 構造体に属することを示します
//   - s は構造体のインスタンスを参照する変数名（自由に名付けられます）
//
// 2. メソッドは構造体のインスタンスに対して呼び出します（例：s.average()）
//
// 3. メソッド内では s. を使って構造体のフィールドにアクセスできます
func (s StudentScore) average() {
	fmt.Println(s.name, (s.math+s.english)/2)
}

// method関数（メインの処理）
// 【解説】
// 1. 構造体のインスタンスを作成し、初期値を設定
// 2. 作成したインスタンスに対してメソッドを呼び出し
func method() {
	// 構造体の初期化と同時にフィールドに値を設定
	s := StudentScore{name: "John", math: 80, english: 90}

	// メソッドの呼び出し
	// 通常の関数と違い、構造体のインスタンス（この場合はs）に対して . でメソッドを呼び出します
	s.average()
}
