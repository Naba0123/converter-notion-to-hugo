package logger

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/zap"
)

func TestInit(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := Init()
		if err != nil {
			t.Errorf("Init() unexpected error: %v", err)
		}

		l, err := zap.NewDevelopment()
		want := &Client{
			logger: l,
		}
		if err != nil {
			t.Fatalf("cannot prepare zap")
		}

		opts := []cmp.Option{
			cmp.AllowUnexported(Client{}),
			cmpopts.IgnoreUnexported(zap.Logger{}),
		}
		if diff := cmp.Diff(got, want, opts...); diff != "" {
			t.Errorf(diff)
		}
	})
}

func TestLogger(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		c, err := Init()
		if err != nil {
			t.Fatalf("cannot prepare because err %v", err)
		}

		_, err = c.Logger()
		if err != nil {
			t.Errorf("Logger() unexpected error = %v", err)
		}
	})
}
