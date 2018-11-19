package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

var priceList = template.Must(template.New("priceList").Parse(`
<html>
    <table>
	    <tr>
		    <th>item</th>
			<th>prime</th>
		</tr>
		{{range $key, $val := .}}
		<tr>
		    <td>{{$key}}</td>
			<td>{{$val}}</td>
		</tr>
		{{end}}
	</table>
</html>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	priceList.Execute(w, db)
	/*
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}*/
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.FormValue("item")
	price := req.FormValue("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	//before := db[item]
	db[item] = dollars(f)
	//fmt.Fprintf(w, "%s'price is updated %f to %f\n", item, before, f)
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
