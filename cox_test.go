package cox

import (
	"testing"

	"dw1.io/cox/policy"
	"github.com/stretchr/testify/assert"
)

type ExampleData struct {
	Field1 string
	Field2 []string
	Field3 []byte
	Field4 int64
	Field5 interface{}
}

func TestClean(t *testing.T) {
	i := ExampleData{
		Field1: "Foo <b>Bar</b>",
		Field2: []string{
			"Hello,",
			`world<script>alert("world")</script>!`,
		},
		Field3: []byte("<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>"),
		Field4: 911,
		Field5: nil,
	}

	e := ExampleData{
		Field1: "Foo <b>Bar</b>",
		Field2: []string{"Hello,", "world!"},
		Field3: []byte("<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>"),
		Field4: 911,
		Field5: nil,
	}

	assert.Exactly(t, e, Clean[ExampleData](i, policy.UGC))
}

func TestCleanWithStrict(t *testing.T) {
	i := ExampleData{
		Field1: "Foo <b>Bar</b>",
		Field2: []string{
			"Hello,",
			`world<script>alert("world")</script>!`,
		},
		Field3: []byte("<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>"),
		Field4: 911,
		Field5: nil,
	}

	e := ExampleData{
		Field1: "Foo Bar",
		Field2: []string{"Hello,", "world!"},
		Field3: []byte("Lorem ipsum."),
		Field4: 911,
		Field5: nil,
	}

	assert.Exactly(t, e, Clean[ExampleData](i, policy.Strict))
}

func TestCleanPtr(t *testing.T) {
	i := &ExampleData{
		Field1: "Foo <b>Bar</b>",
		Field2: []string{
			"Hello,",
			`world<script>alert("world")</script>!`,
		},
		Field3: []byte("<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>"),
		Field4: 911,
		Field5: nil,
	}

	e := &ExampleData{
		Field1: "Foo <b>Bar</b>",
		Field2: []string{"Hello,", "world!"},
		Field3: []byte("<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>"),
		Field4: 911,
		Field5: nil,
	}

	assert.Exactly(t, e, CleanPtr[ExampleData](i, policy.UGC))
}

func TestCleanPtrWithStrict(t *testing.T) {
	i := &ExampleData{
		Field1: "Foo <b>Bar</b>",
		Field2: []string{
			"Hello,",
			`world<script>alert("world")</script>!`,
		},
		Field3: []byte("<blockquote><h1>Lorem</h1> <p>ipsum.</p></blockquote>"),
		Field4: 911,
		Field5: nil,
	}

	e := &ExampleData{
		Field1: "Foo Bar",
		Field2: []string{"Hello,", "world!"},
		Field3: []byte("Lorem ipsum."),
		Field4: 911,
		Field5: nil,
	}

	assert.Exactly(t, e, CleanPtr[ExampleData](i, policy.Strict))
}
