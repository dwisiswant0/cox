package cox

import (
	"reflect"

	"github.com/dwisiswant0/cox/policy"
	"github.com/microcosm-cc/bluemonday"
)

func deepClean(v reflect.Value, p *bluemonday.Policy) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}

	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		// Currently string & bytes are supported
		switch f.Kind() {
		case reflect.String:
			f.SetString(p.Sanitize(f.String()))
		case reflect.Slice:
			switch f.Type().Elem().Kind() {
			case reflect.Uint8: // []byte handler
				f.SetBytes(p.SanitizeBytes(f.Bytes()))
			case reflect.String: // []string handler
				for j := 0; j < f.Len(); j++ {
					s := f.Index(j)
					s.SetString(p.Sanitize(s.String()))
				}
			}
		case reflect.Ptr:
			deepClean(f, p)
		}
	}

	return v
}

func setPolicy(i policy.Kind, a []*bluemonday.Policy) *bluemonday.Policy {
	var n *bluemonday.Policy

	switch i {
	case policy.Blank:
		n = bluemonday.NewPolicy()
	case policy.Strict:
		n = bluemonday.StrictPolicy()
	case policy.UGC:
		n = bluemonday.UGCPolicy()
	}

	if len(a) > 0 {
		a = append([]*bluemonday.Policy{n}, a...)

		for _, p := range a {
			n = p
		}
	}

	return n
}
