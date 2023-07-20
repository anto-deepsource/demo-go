package main

func HexLiteral() bool {
	x := 0xFff
	y := 0xFFF
	z := 0xfff

	_ = "AKIAIOSFODNN73943434"

	return (x == y) && (y == z) || false
}
