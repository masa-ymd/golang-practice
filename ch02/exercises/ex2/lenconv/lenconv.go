package lenconv

import (
	"fmt"
)

type Meter float64
type Feet float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func MToF(m Meter) Feet        { return Feet(m * 3.28084) }
func FToM(f Feet) Meter        { return Meter(f * 0.3048) }