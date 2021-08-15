/*
	Exercise: Loops and Functions
	関数とループを使った簡単な練習として、平方根の計算を実装してみましょう: 数値 x が与えられたときに z² が最も x に近い数値 z を求めたいと思います。

	コンピュータは通常ループを使って x の平方根を計算します。 いくつかの z を推測することから始めて、z² がどれほど x に近づいているかに応じて z を調整できます。

	z -= (z*z - x) / (2*z)
	実際の平方根に近い答えになるまでこの調整を繰り返すことによって、推測はより良いものなります。

	これを func Sqrt に実装してください。 何が入力されても z の適切な開始推測値は 1 です。 まず計算を 10 回繰り返してそれぞれの z を表示します。 x (1, 2, 3, ...) のさまざまな値に対する答えがどれほど近似し、 推測が速くなるかを確認してください。

	Hint: 浮動小数点の変数を初期化して宣言するには、型でキャストするか、浮動小数点を使ってみてください:

	z := 1.0
	z := float64(1)
	次に値が変化しなくなった (もしくはごくわずかな変化しかしなくなった) 場合にループを停止させます。 それが 10 回よりも多いか少ないかを確認してください。 x や x/2 のように他の初期推測の値を z に与えてみてください。 あなたの関数の結果は標準ライブラリの math.Sqrt にどれくらい近づきましたか？

	(メモ: アルゴリズムの詳細について興味がある人のために説明すると、 上の z² − xという式は、z² が最終的な期待値 x からどのくらい離れているかを表しています。 除算の 2z は z² の導関数で、z² の変化の大きさに応じて z の調整値を変化させます。 この一般的なアプローチはニュートン法と呼ばれています。 多くの関数で有効ですが、平方根で特に有効です。)
*/
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	cnt := 0
	var temp, z = 0.0, 1.0
	for i := 0; i < 10; i++ {
		cnt++
		temp = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v: %v\n", cnt, z)
		if math.Abs(temp-z) < 0.0000000001 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}