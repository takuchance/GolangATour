/*
Forever
ループ条件を省略すれば、無限ループ( infinite loop )になりますので、無限ループをコンパクトに表現できます。
*/
package main

import "time"

func main() {
	for {
		println("now is", time.Now().String())
	}
}
