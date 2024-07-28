package sqrt_test

import (
	"example/sqrt" // ${module_name}/${package_name}でインポート
	"fmt"
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	testCases := []struct {
		input    float64
		expected float64
		epsilon  float64
	}{
		// テストケース一覧
		{4, 2, 1e-8},
		{2, math.Sqrt2, 1e-8},
		{0, 0, 1e-8},
		{100, 10, 1e-8},
		{-1, math.NaN(), 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Sqrt(%v)", tc.input), func(t *testing.T) {
			result := sqrt.Sqrt(tc.input) // 平方根を計算
			if tc.input < 0 {             // 引数が負の場合、NaNを返すか
				if !math.IsNaN(result) {
					t.Errorf("Sqrt(%v) = %v, want NaN", tc.input, result)
				}
			} else if math.Abs(result-tc.expected) > tc.epsilon { // 引数が正の場合、誤差が1e-8以下になっているか
				t.Errorf("Sqrt(%v) = %v, want %v (within %v)", tc.input, result, tc.expected, tc.epsilon)
			}
		})
	}
}
