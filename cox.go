package cox

import (
	"reflect"

	"dw1.io/cox/policy"
	"github.com/microcosm-cc/bluemonday"
)

// Clean implements deep-clean and/or sanitization of interface.
// 
// `i` is your interface.
// `p` as kind of policy. See [policy](https://pkg.go.dev/dw1.io/cox/policy).
// `a` for additional policies. See [bluemonday documentation index](https://pkg.go.dev/github.com/microcosm-cc/bluemonday#pkg-index) as a reference for any methods that support policy returns.
func Clean[T any](i T, p policy.Kind, a ...*bluemonday.Policy) T {
	deepClean(reflect.ValueOf(&i), setPolicy(p, a))

	return i
}

// CleanPtr implements Clean but for pointer.
func CleanPtr[T any](i *T, p policy.Kind, a ...*bluemonday.Policy) *T {
	deepClean(reflect.ValueOf(i), setPolicy(p, a))

	return i
}

// CleanSlice implements Clean but for slice type.
func CleanSlice[T any](i []T, p policy.Kind, a ...*bluemonday.Policy) []T {
	for _, m := range i {
		deepClean(reflect.ValueOf(&m), setPolicy(p, a))
	}

	return i
}

// CleanSlicePtr implements Clean but for slice of pointer type.
func CleanSlicePtr[T any](i []*T, p policy.Kind, a ...*bluemonday.Policy) []*T {
	for _, m := range i {
		deepClean(reflect.ValueOf(m), setPolicy(p, a))
	}

	return i
}
