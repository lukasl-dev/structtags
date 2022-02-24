package structtags

import "encoding/json"

// OptionSet is a set of options.
type OptionSet map[string]bool

// NewOptionSet returns a new OptionSet whereby the given opts are set.
func NewOptionSet(opts ...string) OptionSet {
	set := make(OptionSet)
	for _, opt := range opts {
		set.Enable(opt)
	}
	return set
}

// Enable marks the option as set.
func (set OptionSet) Enable(option string) {
	set[option] = true
}

// Contains returns true whether the option is set.
func (set OptionSet) Contains(option string) bool {
	return set[option]
}

// Slice transforms the set into a slice of strings.
func (set OptionSet) Slice() []string {
	sl := make([]string, 0, len(set))
	for option := range set {
		sl = append(sl, option)
	}
	return sl
}

// MarshalJSON marshals the set into a JSON array.
func (set OptionSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(set.Slice())
}

// UnmarshalJSON unmarshals into the set from a JSON array.
func (set OptionSet) UnmarshalJSON(data []byte) error {
	var sl []string
	if err := json.Unmarshal(data, &sl); err != nil {
		return err
	}
	for _, option := range sl {
		set.Enable(option)
	}
	return nil
}
