/*
	Exercise: Images
	前に解いた、 画像ジェネレーター　を覚えていますか？ 今回は、データのスライスの代わりに image.Image インタフェースの実装を返すようにしてみましょう。

	自分の Image 型を定義し、 インタフェースを満たすのに必要なメソッド を実装し、 pic.ShowImage を呼び出してみてください。

	Bounds は、 image.Rect(0, 0, w, h) のようにして image.Rectangle を返すようにします。

	ColorModel は、 color.RGBAModel を返すようにします。

	At は、ひとつの色を返します。 生成する画像の色の値 v を color.RGBA{v, v, 255, 255} を利用して返すようにします。
*/
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rectangle{
		image.Point{0, 0},
		image.Point{255, 255},
	}
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{
		uint8(x * y), uint8(x * y), 255, 255,
	}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
