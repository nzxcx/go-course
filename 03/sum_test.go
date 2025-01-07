package sum

import "testing"

func TestSum(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"positive numbers", 1, 3, 4},
		{"negative numbers", -1, -3, -4},
		{"mixed numbers", -5, 3, -2},
		{"zeros", 0, 0, 0},
		{"large numbers", 1000000, 2000000, 3000000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Sum(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
