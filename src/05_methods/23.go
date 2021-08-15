/*
	Exercise: rot13Reader
	よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。

	例えば、 gzip.NewReader は、 io.Reader (gzipされたデータストリーム)を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開したデータストリーム)を実装しています。

	io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。

	rot13Reader 型は提供済みです。 この Read メソッドを実装することで io.Reader インタフェースを満たしてください。
*/
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13Reader *rot13Reader) Read(arr []byte) (int, error) {
	var len int

	if i, e := rot13Reader.r.Read(arr); e != io.EOF {
		len = i
		for j, v := range arr {
			if v < 'N' || (v >= 'a' && v < 'n') {
				arr[j] += 13
			} else {
				arr[j] -= 13
			}
		}
	} else {
		len = i
	}

	return len, io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
