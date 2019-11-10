package main

func AppendInt(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + len(y)
	if zLen <= cap(x) {
		z = x[:zLen]
	} else {
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
