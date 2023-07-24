package atoi_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

var (
	testcasesb = [][]byte{[]byte("0"), []byte("1"), []byte("123"), []byte("123456"), []byte("123456789"), []byte("2147483647")}
	testcases  = []string{"0", "1", "123", "123456", "123456789", "2147483647"}
	want       = []int{0, 1, 123, 123456, 123456789, 2147483647}
)

func b2s(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

func s2b(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func atoisimple(s string) int {
	ret := 0
	sgn := 1

	if s[0] == '-' {
		sgn = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, ch := range s2b(s) {
		if ch < '0' || ch > '9' {
			return sgn * ret
		}
		ret = ret*10 + int(ch) - '0'
	}

	return sgn * ret
}

func fastatoiadd(s string) int {
	res := 0

	if s[0] == '-' {
		s = s[1:]
		for _, ch := range s2b(s) {
			d := ch - '0'
			if d > 9 {
				return res
			}
			res = res*10 - int(d)
		}
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, ch := range s2b(s) {
		d := ch - '0'
		if d > 9 {
			return res
		}
		res = res*10 + int(d)
	}

	return res
}

func Test_Atoi(t *testing.T) {
	for i, tt := range testcases {
		t.Run(fmt.Sprintf("Atoi_%s", tt), func(t *testing.T) {
			if got, err := strconv.Atoi(tt); got != want[i] || err != nil {
				t.Errorf("atoi() = %v, want %v - %v", got, want[i], err)
			}
		})

		t.Run(fmt.Sprintf("atoisimple_%s", tt), func(t *testing.T) {
			if got := atoisimple(tt); got != want[i] {
				t.Errorf("atoi() = %v, want %v", got, want[i])
			}
		})

		t.Run(fmt.Sprintf("fastatoiadd_%s", tt), func(t *testing.T) {
			if got := fastatoiadd(tt); got != want[i] {
				t.Errorf("atoi() = %v, want %v", got, want[i])
			}
		})
	}
}

func BenchmarkAtoi(b *testing.B) {
	for j, tt := range testcases {
		b.Run(fmt.Sprintf("Atoi_%s", tt), func(bb *testing.B) {
			for i := 0; i < bb.N; i++ {
				if got, err := strconv.Atoi(tt); got != want[j] || err != nil {
					_ = got
					b.Errorf("atoi() = %v, want %v - %v", got, want[j], err)
				}
			}
		})

		b.Run(fmt.Sprintf("atoisimple_%s", tt), func(bb *testing.B) {
			for i := 0; i < bb.N; i++ {
				if got := atoisimple(tt); got != want[j] {
					_ = got
					b.Errorf("atoi() = %v, want %v", got, want[j])
				}
			}
		})

		b.Run(fmt.Sprintf("fastatoiadd_%s", tt), func(bb *testing.B) {
			for i := 0; i < bb.N; i++ {
				got := fastatoiadd(tt)
				if got != want[j] {
					_ = got
					b.Errorf("atoi() = %v, want %v", got, want[j])
				}
				_ = got
			}
		})
	}
}
