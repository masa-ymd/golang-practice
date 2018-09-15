package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	SPACE = iota
	SYMBOL
	MARK
	DIGIT
	PRINT
	PUNCT
	LETTER
	NUMBER
	CONTROL
	GRAPHIC
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsSpace(r):
			counts[SPACE]++
		case unicode.IsSymbol(r):
			counts[SYMBOL]++
		case unicode.IsMark(r):
			counts[MARK]++
		case unicode.IsDigit(r):
			counts[DIGIT]++
		case unicode.IsPrint(r):
			counts[PRINT]++
		case unicode.IsPunct(r):
			counts[PUNCT]++
		case unicode.IsLetter(r):
			counts[LETTER]++
		case unicode.IsNumber(r):
			counts[NUMBER]++
		case unicode.IsControl(r):
			counts[CONTROL]++
		case unicode.IsGraphic(r):
			counts[GRAPHIC]++
		}
		utflen[n]++
	}
	fmt.Printf("unicode type\tcounts\n")
	for c, n := range counts {
		switch {
		case c == SPACE:
			fmt.Printf("SPACE\t%d\n", n)
		case c == SYMBOL:
			fmt.Printf("SYMBOL\t%d\n", n)
		case c == MARK:
			fmt.Printf("MARK\t%d\n", n)
		case c == DIGIT:
			fmt.Printf("DIGIT\t%d\n", n)
		case c == PRINT:
			fmt.Printf("PRINT\t%d\n", n)
		case c == PUNCT:
			fmt.Printf("PUNCT\t%d\n", n)
		case c == LETTER:
			fmt.Printf("LETTER\t%d\n", n)
		case c == NUMBER:
			fmt.Printf("NUMBER\t%d\n", n)
		case c == CONTROL:
			fmt.Printf("CONTROL\t%d\n", n)
		case c == GRAPHIC:
			fmt.Printf("GRAPHIC\t%d\n", n)
		}
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
