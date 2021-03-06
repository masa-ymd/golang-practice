# 第3章

* goの4つの型…基本型、合成型、参照型、インターフェース型
* 基本型…数値、文字列、ブーリアン
* 合成型…配列、構造体
* 参照型…ポインタ、スライス、マップ、関数、チャネル
* インターフェース型

# 3.1 整数

* 符号付き int8, int16, int32, int64
* 符号なし uint8, uint16, uint32, uint64
* コンパイラが適切なサイズを選択する int, uint (32,64bit <- コンパイラにより選択が異なる)
  -> int32とintが実際のサイズが同じであっても変換が必要
* int(n)の範囲は-2^(n-1)〜2^(n-1)-1 uint(n)は0から2^n-1, 例 int8は-128から128, uint8は0から255
* 低レベルプログラミング用の型uintptrもある
* runeはint32のシノニム、byteはint8のシノニム

# 3.2 浮動小数点数

* float32, float64がある
* float32は誤差が発生しやすいので、通常はfloat64を利用する
* NaNは0で割るなど数学的に疑わしい演算の結果を表すが、NaNとの比較は必ずfalseになるので注意

# 3.3 複素数

* complex64とcomplex128があり、実部と虚部はfloat32, float64
* complex関数で複素数を実部、虚部から生成 real, imag関数で実部、虚部成分を取り出せる

# 3.4 ブーリアン

* 左オペランドの値ですでに答えが決定していれば、右オペランドは評価されない
* s != "" && s[0] == "" でs==""なら右の評価はしないで終了する

# 3.5 文字列

* 文字列の値は不変のため、s[0] = 'L'のような構文は不可

## 3.5.1 文字列リテラル

* リテラルとは、数値や文字列を直接に記述した定数のことで、 変数の対義語であり、変更されないことを前提とした値である。
* マルチバイト文字列について
* asciiが0-127までマルチバイトは128-255を複数桁組み合わせて用いる
* https://qiita.com/mpyw/items/a8dba1b80fe68523b8eb
* http://d.hatena.ne.jp/snaka72/20100710/SUMMARY_ABOUT_JAPANESE_CHARACTER_CODE
* UTFとunicodeの違いについて
* https://qiita.com/kasei-san/items/fd0adedc398480cf4ce8
* unicodeを保持するのに適したデータ型はint32->runeというシノニムをもつ
* UTF-8は長さ可変で効率が良い
* unicodeは16進数、unicodeエスケープを使った16ビット形式、32ビット形式で表現できる
* https://seiai.ed.jp/sys/text/csd/mcodes/ucodeindex.html

## 3.5.4 文字列とバイトスライス
* s := "abc" -> b := []byte(s)でsのバイトのコピーを保持する新たなバイト配列が割り当てられ、その配列全体を参照するスライスが生成される
* bytes.Buffer…[]byte をラップして Read(), Write() などを付けたもの

# 3.6 定数
* 定数はコンパイル時に評価が行われる
* 定数を用いた計算式、lenなどの組み込み関数の結果も定数となる

## 3.6.2 型付けなし定数
* 定数は特定の型に結び付けない形で宣言できる
* 型付けなしブーリアン、ルーン、浮動小数点、複素数、文字列、整数がある
* 型付けなし定数は基本型よりも遥かに高い数値精度で表現可能
* 定数宣言時に明確な型を指定された場合、変数に代入された場合は型付きに変換される
