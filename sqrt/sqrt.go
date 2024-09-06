package sqrt

// 平方根を返す関数
func Sqrt(x float64) float64 {

	// 負の数の場合はNaNを返す
	if x < 0 {
		return 0
	}

	// ニュートン法で近似値計算
	z := 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z

}
