# golang_study

Goの勉強メモ。
- https://docs.google.com/presentation/d/1RVx8oeIMAWxbB7ZP2IcgZXnbZokjCmTUca-AbIpORGk/edit?slide=id.g4f417182ce_0_0#slide=id.g4f417182ce_0_0
- https://go.dev/tour/
を参照。

## Goの誕生

2009年11月に最初のバージョンがオープンソースで公開。
毎年8月と2月にリリースされる。

## Goの特徴

- 強力でシンプルな言語設計と文法
- 並行プログラミング
- 豊富な標準ライブラリ群
- 周辺ツールの充実
- シングルバイナリ・クロスコンパイル

## 文法

### モジュールから変数をexportする時

`math.Pi`のように、大文字から始める。

### 型

```go
var i int = 1
var i int （→ zero valueが初期値に代入される）
```

特有の型として、文字を示す`rune`型（UTF）がある。

型定義が後ろについている理由は、可読性のため。

キャストも可能。
```go
var i int = 1
var str string = string(i)
```

関数の中では、暗黙の型定義もできる。

```go
i := 1
j := "Hello"
```

constで定数化する。

```go
const i = 1
```

### for文

初期条件、終了条件は省略可能。

```go
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

```

↑ これのセミコロンは省略でき、while文にできる。

```go
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```

loop条件の記載がないfor文は、無限ループになる。

```go
package main

func main() {
	for {
	}
}
```

if文の冒頭で、変数を用意することができる。
この変数は、elseの中でも使える。

```go
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
```

Newton法の実装例
```go
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := 1e-15
	maxIter := 10
	next := x
	
	for i := 0; i < maxIter; i++ {
		next = z - (z*z - x) / (2*z)
		fmt.Println(next)
		
		if math.Abs(z - next) < eps {
			return next
		}
		
		z = next
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

```

### Switch文

```go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
}
```

switch文の条件式を書かなければ、if-then-elseの代わりとして使える。
```go
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```

### defer

deferで指定した処理は、親の関数がreturnするまで実行されなくなる。
LIFO方式。

### ポインタ

&変数で、ポインタを取得。
*ポインタ変数で、値を取得。

- スライス
- マップ
- チャネル

ではポインタを利用する必要がない。

### 構造体

フィールドの集まり。
ポインタ経由でアクセスもできる。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

```

## テクニック

空の構造体はゼロ容量のインスタンスになる。
keyだけが重要なmapとかで利用したりする。

...to be continued