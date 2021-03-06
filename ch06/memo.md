# 6. メソッド

## 6.1 メソッド宣言
* func (h Hoge) Func(f Fuga)... のhはメソッドのレシーバ(レシーバパラメータ)
* メソッドを呼び出すときのh.Funcのhはレシーバ引数
* 式h.Funcはセレクタと呼ばれる　セレクタは構造体のフィールドを選択するときにも使われるので、フィールド名と同名のメソッドは定義不可
* パッケージの関数と型に指定したメソッドに衝突はない
* 異なる型には同名のメソッドを宣言可能
* 基底型がポインタかインターフェースでなければどのような方でもメソッドは宣言可能

## 6.2 ポインタレシーバを持つメソッド
* 関数の呼び出しは個々の引数値のコピーを行うため、関数が変数を更新する必要がある場合、引数が大きくコピーを避けたい場合はポインタ型にメソッドを結びつける
* レシーバとメソッドの関係1: レシーバ引数とレシーバパラメータが同じ型
* レシーバとメソッドの関係2: レシーバ引数が型Hの変数hで、レシーバパラメータがHのポインタ型の時は、h.Func()で暗黙的に&h.Func()になる
* レシーバとメソッドの関係3: レシーバ引数が型Hのポインタ変数hで、レシーバパラメータがH型の時は、h.Func()で暗黙的に(*h).Func()になる

## 6.3 構造体埋め込みによる型の合成
* 埋め込みの関係はis aではなくhas a