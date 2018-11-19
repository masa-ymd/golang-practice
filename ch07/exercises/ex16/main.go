package main

import (
	//"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/masa-ymd/golang-practice/ch07/exercises/ex15/eval"
)

var expr eval.Expr

type data struct {
	Values []string
}

var d data

func formula(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t := template.Must(template.ParseFiles("formula.tmpl"))
		t.Execute(w, nil)
	}
}

func calc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		expr, _ = eval.Parse(r.FormValue("formula"))
		t := template.Must(template.ParseFiles("calc.tmpl"))
		d.Values = []string{}
		r := regexp.MustCompile(`[a-zA-Z]`)
		for _, k := range []byte(expr.String()) {
			k := string(k)
			if r.MatchString(k) {
				d.Values = append(d.Values, k)
			}
		}
		t.Execute(w, d)
	}
}

func result(w http.ResponseWriter, r *http.Request) {
	env := eval.Env{}
	if r.Method == http.MethodPost {
		for _, v := range d.Values {
			f, _ := strconv.ParseFloat(r.FormValue(v), 64)
			env[eval.Var(v)] = f
		}
		w.Write([]byte(fmt.Sprintf("%f", expr.Eval(env))))
	}
}

func main() {
	http.HandleFunc("/", formula)
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/result", result)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
