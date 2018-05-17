package wgtconv

import (
	"fmt"
)

type Pound float64
type Kilogram float64

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func PToK(p Pound) Kilogram       { return Kilogram(p * 0.453592) }
func KToP(k Kilogram) Pound       { return Pound(k * 2.20462) }
