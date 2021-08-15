/*
	Exercise: Readers
	ASCII文字 'A' の無限ストリームを出力する Reader 型を実装してください。
*/
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (myreader MyReader) Read(arr []byte) (int, error) {
	for i := range arr {
		arr[i] = 'A'
	}
	return len(arr), nil
}

func main() {
	reader.Validate(MyReader{})
}
