package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// mainから始まり、ここが始まりになるのでエントリーポイントという、ここに他の関数を作成した場合は入れる
// プログラムの実行順序：
// 1. 基本的な出力を確認
// 2. 変数の使い方を学習
// 3. 配列の使い方を学習
// 4. 計算の仕方を学習
// 5. 条件分岐を学習
func main() {
	fmt.Println("=== プログラム開始 ===")
	fmt.Println("基本的な出力を確認します：")
	fmt.Println("Good Morning")
	fmt.Println("Good Afternoon")
	fmt.Println("Good Evening")

	fmt.Println("\n=== 変数の練習 ===")
	var_test()

	fmt.Println("\n=== 配列の練習 ===")
	type_test()

	fmt.Println("\n=== 計算の練習 ===")
	math()

	fmt.Println("\n=== 条件分岐の練習 ===")
	if_test()

	fmt.Println("\n=== 乱数の練習 ===")
	random_test()
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

// math関数では以下の内容を練習しています：
// - 基本的な四則演算（足し算、引き算、掛け算、割り算）
// - 比較演算子（等しい、等しくない、大きい、小さい、など）
// - 論理演算子（AND、OR、NOT）
func math() {
	// 変数の初期値を設定
	x := 10 // xに10を代入
	y := 2  // yに2を代入

	// 基本的な四則演算
	fmt.Println(x + y) // 足し算
	fmt.Println(x - y) // 引き算
	fmt.Println(x * y) // 掛け算
	fmt.Println(x / y) // 割り算
	fmt.Println(x % y) // 余り（割り算の余りを求める）

	// 比較演算子の使い方
	fmt.Println(x == y)          // 等しいかどうか
	fmt.Println(x != y)          // 等しくないかどうか
	fmt.Println(x > y)           // より大きいかどうか
	fmt.Println(x < y)           // より小さいかどうか
	fmt.Println(x >= y)          // 以上かどうか
	fmt.Println(x <= y)          // 以下かどうか
	fmt.Println(x == y || x > y) // 等しいまたは大きいかどうか（OR）
	fmt.Println(x == y && x > y) // 等しくかつ大きいかどうか（AND）

	// インクリメントとデクリメント
	// インクリメントは1を加算すること
	// デクリメントは1を減算すること
	// よく使うのはfor文のカウンター変数の増減

	x++ // xの値を1増やす
	fmt.Println(x)
	x-- // xの値を1減らす
	fmt.Println(x)

	// 論理演算子
	// 論理和（OR）: || 　→ どちらかがtrueならtrue
	// 論理積（AND）: && 　→ 両方がtrueならtrue
	// 否定（NOT）: !　→ trueならfalse、falseならtrue

	fmt.Println(true || false) // どちらかがtrue
	fmt.Println(true && false) // 両方がtrue
	fmt.Println(!true)         // 反対にする

}

func if_test() {
	fmt.Println("テストの合否判定プログラム")
	fmt.Println("------------------------")

	// 条件分岐
	// if文は条件に応じて処理を分けたい時に使う
	// 例：テストの点数で合格/不合格を判定する場合など

	// 複数のテストケースを試してみる
	scores := []int{80, 59, 100, 40}

	for _, score := range scores {
		fmt.Printf("\n点数: %d点の判定\n", score)

		// if文による判定
		fmt.Println("【if文での判定】")
		if score >= 80 {
			fmt.Println("優秀です！")
		} else if score >= 60 {
			fmt.Println("合格です")
		} else {
			fmt.Println("不合格です。もう少し頑張りましょう。")
		}

		// switch文による判定
		// if文の特殊なバージョン
		// 値が一致するケースの処理を実行する
		fmt.Println("\n【switch文での判定】")
		switch {
		case score >= 80:
			fmt.Println("Aランク")
		case score >= 60:
			fmt.Println("Bランク")
		case score < 60:
			fmt.Println("Cランク")
		}
	}

	fmt.Println("\n曜日判定プログラム")
	fmt.Println("----------------")

	// 曜日の例
	weekday := 3 // 1:月曜日 ～ 7:日曜日

	switch weekday {
	case 1:
		fmt.Println("月曜日：今週も頑張りましょう！")
	case 2:
		fmt.Println("火曜日：調子出てきました")
	case 3:
		fmt.Println("水曜日：週の真ん中です")
	case 4:
		fmt.Println("木曜日：もう少しです")
	case 5:
		fmt.Println("金曜日：週末が近いです")
	case 6:
		fmt.Println("土曜日：よい週末を！")
	case 7:
		fmt.Println("日曜日：ゆっくり休みましょう")
	default:
		fmt.Println("正しい曜日番号を入力してください（1-7）")
	}
}

// random_test関数では以下の内容を練習しています：
// - 乱数の基本的な生成方法
// - シード値の設定
// - 範囲を指定した乱数生成
// - 乱数を使ったミニゲーム
func random_test() {
	fmt.Println("乱数生成プログラム")
	fmt.Println("----------------")

	// 乱数のシード値を設定
	// 新しい乱数生成器を作成し、現在時刻をシードとして使用
	// これにより、毎回異なる乱数列が生成される
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 1. 基本的な乱数生成
	fmt.Println("\n【基本的な乱数】")
	for i := 0; i < 5; i++ {
		num := r.Int()
		fmt.Printf("生成された乱数 %d: %d\n", i+1, num)
	}

	// 2. 範囲を指定した乱数（サイコロ）
	fmt.Println("\n【サイコロを振る】")
	for i := 0; i < 3; i++ {
		dice := r.Intn(6) + 1 // 1から6までの乱数
		fmt.Printf("%d回目の結果: %d\n", i+1, dice)
	}

	// 3. おみくじプログラム
	fmt.Println("\n【おみくじを引く】")
	fortunes := []string{"大吉", "中吉", "小吉", "吉", "末吉", "凶"}
	fortune := fortunes[r.Intn(len(fortunes))]
	fmt.Printf("今日の運勢: %s\n", fortune)

	// 4. 数当てゲーム
	fmt.Println("\n【数当てゲーム】")
	target := r.Intn(10) + 1 // 1から10までの乱数
	fmt.Println("1から10までの数字を生成しました")
	fmt.Printf("答え（確認用）: %d\n", target)

	// 5. じゃんけんゲーム
	fmt.Println("\n【じゃんけんゲーム】")
	hands := []string{"グー", "チョキ", "パー"}
	computerHand := hands[r.Intn(len(hands))]
	fmt.Printf("コンピュータの手: %s\n", computerHand)
	fmt.Println("あなたの手を決めてじゃんけんしましょう！")

	// 6. ランダムな色生成
	fmt.Println("\n【ランダムな色生成】")
	red := r.Intn(256) // 0-255のRGB値
	green := r.Intn(256)
	blue := r.Intn(256)
	fmt.Printf("生成された色のRGB値: (%d, %d, %d)\n", red, green, blue)
}
