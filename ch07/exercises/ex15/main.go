package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/masa-ymd/golang-practice/ch07/exercises/ex15/eval"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("function?:")
	stdin.Scan()
	text := stdin.Text()
	expr, err := eval.Parse(text)
	if err != nil {
		log.Fatal("err")
	}
	v := eval.Env{}
	r := regexp.MustCompile(`[a-zA-Z]`)
	for _, k := range []byte(expr.String()) {
		k := string(k)
		if r.MatchString(k) {
			fmt.Printf("%s?:", k)
			stdin.Scan()
			val := stdin.Text()
			f, _ := strconv.ParseFloat(val, 64)
			v[eval.Var(k)] = f
		}
	}
	fmt.Println(expr.Eval(v))
}
