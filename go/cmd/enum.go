package cmd

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"
)

// EnumValue is a flag.Value that can be used to restrict the
// values of a flag to a list of allowed values.
type EnumValue struct {
	Enum     []string
	Default  string
	selected string
}

func (e *EnumValue) Get() any {
	if e.selected == "" {
		return e.Default
	}

	return e.selected
}

func (e *EnumValue) Set(value string) error {
	for _, enum := range e.Enum {
		if enum == value {
			e.selected = value
			return nil
		}
	}

	return fmt.Errorf("allowed values are %s", strings.Join(e.Enum, ", ")) //nolint: err113
}

func (e *EnumValue) String() string {
	if s, ok := e.Get().(string); ok {
		return s
	}

	return ""
}

func GetEnumValue(cmd *cli.Command, flagName string) string {
	g := GetGeneric[EnumValue](cmd, flagName)
	if g == nil {
		return ""
	}

	if s, ok := g.Get().(string); ok {
		return s
	}

	return ""
}

func GetGeneric[T any](cmd *cli.Command, name string) *T {
	g := cmd.Generic(name)
	if g == nil {
		return nil
	}

	if v, ok := any(g).(*T); ok {
		return v
	}

	return nil
}
