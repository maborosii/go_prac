package benchtest

import (
	"fmt"
	"strconv"
	"testing"
)

// Sprintf 方式将 int 类型转换成 string 类型
func BenchmarkSprintf(b *testing.B) {
	num := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", num)
	}
}

// Itoa 方式将 int 类型转换成 string类型
func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}

//  通过 FormatInt 方式将 int 类型转换成 string 类型
func BenchmarkFormatInt(b *testing.B) {
	num := 10
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(num), num)
	}
}
