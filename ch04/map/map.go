package main

import "fmt"

func main() {
	// map作成の方法1
	ages := map[string]int{
		"alice": 31,
		"bob":   34,
	}
	// 方法2
	ages2 := make(map[string]int)
	ages2["alice"] = 31
	ages2["bob"] = 34

	fmt.Println(ages["alice"])
	ages["bob"]++
	fmt.Println(ages["bob"])
	// 存在しないキーを使った場合は、valueのゼロ値が返される
	fmt.Println(ages["taro"])
	// mapの繰り返し順序は定義されていないため、stringのスライスにいれてからsortするなどの対応をする必要がある
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// nilのmapを作成する方法
	var ages3 map[string]int
	// ハッシュテーブルを指していない状態のため、nilのマップに値を保存するとパニックになる
	//　新しく作ったりすでにあるハッシュテーブルを代入すれば保存できるようになる
	ages3 = make(map[string]int)
	ages3["taro"]++
	fmt.Printf("%v\n", ages3)

	// 要素の有無を確認する方法
	age, ok := ages3["taro"]
	fmt.Printf("%d, %t\n", age, ok)

}
