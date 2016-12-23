// binarysearch_test
package binarysearch_test

import (
	"testing"

	"github.com/cheekybits/is"
)

func TestRandom(t *testing.T) {
	is := is.New(t)
	a, b := "aa2", "aa1"
	is.True(a < b)
}
