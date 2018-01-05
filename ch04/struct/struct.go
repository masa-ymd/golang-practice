package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	// 構造体のそれぞれのフィールドへは.表記でアクセス可能
	dilbert.Salary += 5000
	// フィールドのアドレスを取得し、ポインタ経由でアクセスも可能
	position := &dilbert.Position
	*position = "hoge"

	fmt.Printf("%v\n", dilbert)

	// .表記は構造体のポインタでも利用可能
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "test"
	// 上と同じ内容　*は省略可
	(*employeeOfTheMonth).Position += "test2"

	fmt.Printf("%v\n", dilbert)

	// 構造体のポインタを返す関数があったとき、返り値のフィールドにアクセスする際、以下のように.が使える
	fmt.Println(EmployeeByID(1).ID)
}

func EmployeeByID(id int) *Employee {
	e := &Employee{}
	// 構造体を初期化し、アドレスを得る方法は他にも以下のようなものがある
	// pp := new(Employee)
	// *pp = Employee{}
	e.ID = id
	e.Position = "fuga"
	return e
}
