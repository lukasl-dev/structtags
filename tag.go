package structtags

// Tag represents an enhanced struct tag.
type Tag struct {
	// Value is the actual value of the tag. In most cases, this is used to
	// change the name of the annotated fields internally. For example, the
	// struct tag `json:"foo,omitempty` has the value "foo".
	Value string `json:"value,omitempty"`

	// Options is a slice of strings that are used as boolean options in the
	// struct tag. For example, the struct tag `json:"foo,omitempty"` has an
	// option labelled "omitempty".
	Options []string `json:"options,omitempty"`
}
