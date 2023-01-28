package errorwrapper

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWrap(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		targetErr := errors.New("test error")
		Wrap(&targetErr, "wrapped")
		if targetErr == nil {
			t.Errorf("Wrap(): targetErr is nil")
		}

		if diff := cmp.Diff(targetErr.Error(), "wrapped: test error"); diff != "" {
			t.Errorf(diff)
		}
	})
	t.Run("success", func(t *testing.T) {
		targetErr := errors.New("test error")
		Wrap(&targetErr, "wrapped %v", 100)
		if targetErr == nil {
			t.Errorf("Wrap(): targetErr is nil")
		}
		if diff := cmp.Diff(targetErr.Error(), "wrapped 100: test error"); diff != "" {
			t.Errorf(diff)
		}
	})
}
