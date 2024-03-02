package testhelpers

import (
	"github.com/google/go-cmp/cmp"
)

func FilterPathLast(last []string, opt cmp.Option) cmp.Option { //nolint:ireturn
	return cmp.FilterPath(
		func(path cmp.Path) bool {
			for i := range last {
				j := len(path) - len(last) + i
				if j < 0 {
					return false
				}
				if last[i] != path[j].String() {
					return false
				}
			}
			return true
		},
		opt,
	)
}
