/*
	Buffered Channels
	チャネルは、 バッファ ( buffer )として使えます。 バッファを持つチャネルを初期化するには、 make の２つ目の引数にバッファの長さを与えます:

	ch := make(chan int, 100)
	バッファが詰まった時は、チャネルへの送信をブロックします。 バッファが空の時には、チャネルの受信をブロックします。

	バッファが詰まるようにサンプルコードを変更し、何が起きるのかを見てみてください。
*/
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// ch := make(chan int, 1)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
