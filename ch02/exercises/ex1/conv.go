package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c - AbsoluteZero) }
func KToC(k Kelvin) Celsius     { return Celsius(k + Kelvin(float64(AbsoluteZero))) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin(FToC(f) - AbsoluteZero) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit(KToC(k)*9/5 + 32) }
