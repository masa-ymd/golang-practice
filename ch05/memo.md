# 5 関数
## 5.1 関数宣言
* 本体のない関数宣言はGo以外の言語で実装されいるもの。その宣言は関数のシグネチャ（引数と結果の型）を定義している
## 5.3 複数戻り値
* Goのガベージコレクタはファイルやネットワーク接続などの使われていないOSの資源を開放することは想定していないので自分で実施する必要あり
* 関数の結果に名前を付けることで、空リターンを行うことができるが、コードがわかりづらくなるため、利用箇所は限定すべき
## 5.4 エラー
* 関数f(x)を呼び出す側は、呼び出し側のみが持っている情報をエラーに付与する責任がある
* すべての関数呼び出しの後に、エラーを考慮する処理を書く。書かない場合は意図を明確にドキュメントに残す
## 5.8 遅延関数呼び出し
* deferは関数終了まで呼び出されないので、loop内のファイルクローズ処理などをdeferで書くとファイル記述子が枯渇する可能性がある
## 5.9 パニック
* パニックは起こるはずのない状況が発生した場合に起こす
* パニックを起こしたプログラムはすべての遅延関数を呼び出した後、ログメッセージを表示してクラッシュする
* Goランタイムが検査してくれる状態（パニック）を自前のパニックで確認することはない
## 5.10 リカバー
* recover関数はパニックになっている場合にパニックの現状を終了させてパニック値を返す
* 異常な状態になった場合の後処理をするときなどに用いる
* 安全性が判断できないため、他のパッケージなど自分で管理していない関数のパニックからはリカバーすべきではない