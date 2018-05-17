package tempconv

import (
	"fmt"
	//"testing"
)

func ExampleChange() {
	fmt.Println(CToF(AbsoluteZero))
	fmt.Println(CToF(FreezingC))
	fmt.Println(CToF(BiolingC))
	fmt.Println(FToC(32))
	fmt.Println(CToK(0))
	fmt.Println(KToC(0))
	fmt.Println(FToK(32))
	fmt.Println(KToF(273.15))

	// output:
	// -459.66999999999996°F
	// 32°F
	// 212°F
	// 0°C
	// 273.15°K
	// -273.15°C
	// 273.15°K
	// 32°F
}
