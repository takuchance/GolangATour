/*
	Exercise: Fibonacci closure
	関数を用いた面白い例を見てみましょう。

	fibonacci (フィボナッチ)関数を実装しましょう。この関数は、連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返します。
*/
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var arr []int
	return func() int {
		switch len(arr) {
		case 0:
			arr = append(arr, 0)
		case 1:
			arr = append(arr, 1)
		default:
			arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
		}
		return arr[len(arr)-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
