package structtags

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse_EmptyString(t *testing.T) {
	r := require.New(t)

	tags, err := Parse("")
	r.Nil(tags)
	r.Error(err)
}

func TestParse_InvalidTag(t *testing.T) {
	r := require.New(t)

	tags, err := Parse("invalid")
	r.Nil(tags)
	r.Error(err)
}

func TestParse_ValidTag(t *testing.T) {
	r := require.New(t)

	tags, err := Parse(`json:"value"`)
	r.NoError(err)
	r.Len(tags, 1)
	r.Contains(tags, "json")

	tag := tags["json"]
	r.Equal("value", tag.Value)
	r.Empty(tag.Options)
}

func TestParse_ValidTagWithOption(t *testing.T) {
	r := require.New(t)

	tags, err := Parse(`json:"value,omitempty"`)
	r.NoError(err)
	r.Len(tags, 1)
	r.Contains(tags, "json")

	tag := tags["json"]
	r.Equal("value", tag.Value)

	r.Len(tag.Options, 1)
	r.Equal("omitempty", tag.Options.Slice()[0])
}

func TestParse_ValidTagWithOptions(t *testing.T) {
	r := require.New(t)

	tags, err := Parse(`json:"value,omitempty,foo"`)
	r.NoError(err)
	r.Len(tags, 1)
	r.Contains(tags, "json")

	tag := tags["json"]
	r.Equal("value", tag.Value)

	sl := tag.Options.Slice()
	r.Len(sl, 2)
	r.Equal("omitempty", sl[0])
	r.Equal("foo", sl[1])
}
