package policy

// Kind of policy
type Kind uint

const (
	// Blank returns a blank policy with nothing allowed or permitted.
	Blank Kind = iota
	// Strict returns an empty policy, which will effectively strip all HTML elements and their attributes from a document.
	// (source: https://pkg.go.dev/github.com/microcosm-cc/bluemonday#StrictPolicy)
	Strict
	// UGC returns a policy aimed at user generated content that is a result of HTML WYSIWYG tools and Markdown conversions.
	// (source: https://pkg.go.dev/github.com/microcosm-cc/bluemonday#UGCPolicy)
	UGC
)
