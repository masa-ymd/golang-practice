# 4 コンポジット型

# 4.1 配列
* C言語は配列を関数に渡した時、暗黙的にポインタが渡されるが、GOはコピーが渡される。ポインタを渡したい場合は明示的に書く

## 4.2 スライス
* cap()…スライスの容量を返す
* スライスの容量はたいてい基底配列上の開始点から、基底配列の「終わり」までの要素数
* len()を超えてもcap()を超えない範囲ならスライスの拡張が可能
* スライスは配列と異なり==で比較できない(nilとのみ比較可)

## 4.3 マップ
* キーと値のハッシュテーブル
* 値の取得、更新、削除が可能
* キーに関連する値の特定ははデータ量に関係なく一定回数で可能
* キーに格納する値は比較可能であればよい（NaNが入るものはよくない）
* mapの要素に対するアドレスは取得不可　要素の追加によりメモリの位置が移動する可能性があるため
* map同士の比較は不可のため、loopを使って、中身を１つずつ確認していく必要がある

## 4.4 構造体
* すべてのフィールドが比較可能であれば、== !=を使って比較可能
* 比較可能な構造体はmapのキーとしても使える
