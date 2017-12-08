package main

import (
	"flag"
	"fmt"
	"os"
    "bufio"
    "strconv"
	"golang-practice/ch02/tempconv"
)

type Kilogram float64
type Pound float64

func (kg Kilogram) String() string {
    return fmt.Sprintf("%fkg", kg)
}
func (p Pound) String() string {
    return fmt.Sprintf("%flbs", p)
}

func main() {
    var t, w *float64
    var rt, rw float64 = 0.0, 0.0
    t, w = &rt, &rw
	// 引数の数が1以上
	if len(os.Args) > 1 {
		t = flag.Float64("t", 0.0, "temperature")
		w = flag.Float64("w", 0.0, "weight")
		flag.Parse()
	} else {
        fmt.Print("type? (t|w):")
        stdin := bufio.NewScanner(os.Stdin)
        stdin.Scan()
        tp := stdin.Text()
        fmt.Print("num? (float):")
        stdin.Scan()
        n, _ := strconv.ParseFloat(stdin.Text(),  64)
        if tp == "t" {
           t = &n
        } else if tp == "w" {
           w = &n
        }
    }

    if t != nil && *t != 0 {
        f := tempconv.Fahrenheit(*t)
        c := tempconv.Celsius(*t)
        fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
    } else if w != nil && *w != 0 {
        k := Kilogram(*w)
        p := Pound(*w)
        fmt.Printf("%s = %s, %s = %s\n", k, KToP(k), p, PToK(p))
    }
}

func KToP(kg Kilogram) Pound {
    return Pound(kg * 0.454)
}

func PToK(p Pound) Kilogram {
    return Kilogram(p / 0.454)
}
