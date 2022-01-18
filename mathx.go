package gomathx

import "math"

//最大公约数
func GCD(x, y int64) int64 {
	x = int64(math.Abs(float64(x)))
	y = int64(math.Abs(float64(y)))

	var tmp int64
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

//最小公倍数
func LCM(x, y int64) int64 {
	return (x * y) / GCD(x, y)
}
