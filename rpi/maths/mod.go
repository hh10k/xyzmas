package maths

func Mod(a int, b int) int {
	return ((a % b) + b) % b
}
