package main

import (
	"fmt"

	"dw1.io/cox"
	"dw1.io/cox/policy"
)

// Info ...
type Info struct {
	Fname string
	Lname string
	Phone int64
	Notes []string
	Story []byte
}

func main() {
	i := Info{
		Fname: "Foo",
		Lname: "<b>Bar</b>",
		Phone: 911,
		Notes: []string{
			"Hello,",
			`world<script>alert("world")</script>!`,
		},
		Story: []byte(`<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>`),
	}
	i = cox.Clean[Info](i, policy.Strict)

	fmt.Printf("Info: %+v\n", i)
	fmt.Printf("Info.Story: %+v\n", string(i.Story))

	// Output:
	// Info: {Fname:Foo Lname:Bar Phone:911 Notes:[Hello, world!] Story:[76 111 114 101 109 32 105 112 115 117 109 46]}
	// Info.Story: Lorem ipsum.
}
