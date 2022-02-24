package structtags

import (
	"errors"
	"fmt"
	"strings"
)

type parser struct {
	// s is the raw string to parse. It is mutated during parsing.
	s string

	// tags is a map of tag names to parsed tags.
	tags map[string]Tag

	// name is the name of the tag being parsed.
	name string

	// value is the value of the tag being parsed.
	value string

	// opts is a slice of options of the tag being parsed.
	opts []string
}

// Parse parses the raw struct tag s into a map of tags identified by their
// names.
func Parse(s string) (map[string]Tag, error) {
	return newParser(s).parse()
}

// newParser returns a new parser that parses s.
func newParser(s string) *parser {
	return &parser{s: s, tags: make(map[string]Tag)}
}

// parse parses p's raw string into a map of tags.
func (p *parser) parse() (map[string]Tag, error) {
	for fn := p.start; fn != nil; {
		var err error
		fn, err = fn()
		if err != nil {
			return nil, err
		}
	}
	return p.tags, nil
}

type parse func() (parse, error)

func (p *parser) start() (parse, error) {
	p.name = ""
	p.value = ""
	p.opts = nil
	return p.parseName, nil
}

func (p *parser) parseName() (parse, error) {
	p.s = strings.TrimSpace(p.s)
	if p.s == "" {
		return nil, errors.New("structtags: raw string must not be empty")
	}

	pos := strings.IndexRune(p.s, ':')
	if pos == -1 {
		return nil, errors.New("structtags: no or invalid tag name (missing colon)")
	}

	p.name = p.s[:pos]
	p.s = p.s[pos+1:]

	return p.parseValue, nil
}

func (p *parser) parseValue() (parse, error) {
	if p.s == "" {
		return nil, fmt.Errorf("structtags: no value found for tag %q", p.name)
	}

	openQuote := strings.IndexRune(p.s, '"')
	if openQuote == -1 {
		return nil, fmt.Errorf("structtags: no opening quote found for tag %q", p.name)
	}
	p.s = p.s[openQuote+1:]

	closeQuote := strings.IndexRune(p.s, '"')
	if closeQuote == -1 {
		return nil, fmt.Errorf("structtags: no closing quote found for tag %q", p.name)
	}

	p.value = p.s[:closeQuote]
	p.s = p.s[closeQuote+1:]

	return p.parseOptions, nil
}

func (p *parser) parseOptions() (parse, error) {
	p.opts = strings.Split(p.value, ",")
	p.value = p.opts[0]
	p.opts = p.opts[1:]
	return p.end, nil
}

func (p *parser) end() (parse, error) {
	p.tags[p.name] = Tag{Value: p.value, Options: p.opts}

	if p.s == "" {
		return nil, nil
	}
	return p.start, nil
}
