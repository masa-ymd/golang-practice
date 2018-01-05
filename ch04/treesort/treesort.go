package main

import "fmt"

// 再帰構造をもつ構造体
type tree struct {
	value       int
	left, right *tree
}

// intのスライスを与えると、tree構造体を利用してソートする関数
func Sort(values []int) []int {
	var root *tree
	for _, v := range values {
		// 再帰的にvalueの大小関係を表す木構造を作る関数
		root = add(root, v)
	}
	// スライスと木構造を渡すと、木構造の自分自身と左右の枝を大きさに従って再帰的にappendしていく関数
	// 最初なのでスライスの長さは0
	return appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		// まずは左（小さいほう）を再帰的にvalueｓに追加する
		values = appendValues(values, t.left)
		// 自分自身のvalue
		values = append(values, t.value)
		// 最後に右（大きいほう）を再帰的にvaluesに追加する
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	// 自分自身より小さいものは左に、大きいものは右の枝に配置する
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	values := []int{1, 4, 2, 5, 3, 9, 10}
	sortedValues := Sort(values)
	for _, v := range sortedValues {
		fmt.Println(v)
	}
}
