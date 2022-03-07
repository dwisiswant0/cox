# cox

[![GoDoc](https://godoc.org/dw1.io/cox?status.svg)](https://pkg.go.dev/dw1.io/cox)

Cox is [bluemonday](https://github.com/microcosm-cc/bluemonday)-wrapper to perform a deep-clean and/or sanitization of <i>(nested-)</i>interfaces from HTML to prevent XSS payloads. It'll sanitize all fields in the structure, supports fields with _(slice of, and/or just)_ `string` and `byte` types.

## Install

<table>
	<td><b>NOTE:</b> <a href="https://go.dev/dl/">Go1.18+</a> compiler should be installed & configured.</td>
</table>

It's fairly simple!

```console
go get dw1.io/cox
```

## Usage

You can import `cox` using a basic statement:

```golang
import (
	"dw1.io/cox"
	"dw1.io/cox/policy"
)
```

### Examples

```golang
t := T{/* ... */}
t = cox.Clean[T](t, policy.Strict) // Sanitizing with strict policy, returning to its type
// For pointer, use cox.CleanPtr method.
```

> Kind of policy: `Blank`, `UGC` and `Strict`. See [policy](https://pkg.go.dev/dw1.io/cox/policy).

#### Additional policies

If you want additional policies, you can add some of them as variadic arguments at the end.

```golang
p := bluemonday.NewPolicy()

t := T{/* ... */}
t = cox.Clean[T](t, policy.Blank, p.AllowRelativeURLs(true), p.AllowElements("br", "div", "hr", "p", "span"))
```

> See [bluemonday documentation index](https://pkg.go.dev/github.com/microcosm-cc/bluemonday#pkg-index) as a reference for any methods that support policy returns.

### Workaround

The following is an example of how this library is implemented & works:

```golang
type Info struct {
	Fname string
	Lname string
	Phone int64
	Notes []string
	Story []byte
}

func main() {
	i := &Info{
		Fname: "Foo",
		Lname: "<b>Bar</b>",
		Phone: 911,
		Notes: []string{
			"Hello,",
			`world<script>alert("world")</script>!`,
		},
		Story: []byte(`<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>`),
	}
	i = cox.CleanPtr[Info](i, policy.Strict)

	fmt.Printf("%+v\n", i)

	// Output:
	// &{Fname:Foo Lname:Bar Phone:911 Notes:[Hello, world!] Story:[76 111 114 101 109 32 105 112 115 117 109 46]}
}
```

## Why this name?

Mbuh, cok!

> F\*ck! Dunno.

## Limitations

Nested types only work for pointer, not struct.

## License

**cox** is distributed under MIT. See `LICENSE`.