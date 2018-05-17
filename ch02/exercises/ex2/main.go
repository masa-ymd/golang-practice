package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/masa-ymd/golang-practice/ch02/exercises/ex1"
	"github.com/masa-ymd/golang-practice/ch02/exercises/ex2/lenconv"
	"github.com/masa-ymd/golang-practice/ch02/exercises/ex2/wgtconv"
)

var (
	tmp = flag.Float64("t", 0.0, "temparature")
	lgt = flag.Float64("l", 0.0, "length")
	wgt = flag.Float64("w", 0.0, "weight")
)

func main() {
	flag.Parse()

	if flag.NFlag() > 0 {
		flag.Visit(func(f *flag.Flag) {
			if f.Name == "t" {
				f := tempconv.Fahrenheit(*tmp)
				c := tempconv.Celsius(*tmp)
				fmt.Printf("%s = %s, %s = %s\n",
					f, tempconv.FToC(f), c, tempconv.CToF(c))
			}
			if f.Name == "l" {
				m := lenconv.Meter(*lgt)
				ft := lenconv.Feet(*lgt)
				fmt.Printf("%s = %s, %s = %s\n",
					m, lenconv.MToF(m), ft, lenconv.FToM(ft))
			}
			if f.Name == "w" {
				p := wgtconv.Pound(*wgt)
				k := wgtconv.Kilogram(*wgt)
				fmt.Printf("%s = %s, %s = %s\n",
					p, wgtconv.PToK(p), k, wgtconv.KToP(k))
			}
		})
	} else {
		stdin := bufio.NewScanner(os.Stdin)
		fmt.Print("Please Input Type [t/l/w]:")
		stdin.Scan()
		t := stdin.Text()
		fmt.Print("Please Input Value:")
		stdin.Scan()
		v, err := strconv.ParseFloat(stdin.Text(), 64)
		if err != nil {
			fmt.Errorf("%v", err)
			os.Exit(1)
		}
		switch t {
		case "t":
			f := tempconv.Fahrenheit(v)
			c := tempconv.Celsius(v)
			fmt.Printf("%s = %s, %s = %s\n",
				f, tempconv.FToC(f), c, tempconv.CToF(c))
		case "l":
			m := lenconv.Meter(v)
			ft := lenconv.Feet(v)
			fmt.Printf("%s = %s, %s = %s\n",
				m, lenconv.MToF(m), ft, lenconv.FToM(ft))
		case "w":
			p := wgtconv.Pound(v)
			k := wgtconv.Kilogram(v)
			fmt.Printf("%s = %s, %s = %s\n",
				p, wgtconv.PToK(p), k, wgtconv.KToP(k))
		default:
			fmt.Printf("Unknown Type (%s)", t)
		}
	}
}
