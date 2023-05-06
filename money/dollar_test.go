package money

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDollar_times(t *testing.T) {
	t.Run("失敗", func(t *testing.T) {
		five := Dollar{5}
		got := five.Times(3)
		want := 12
		if d := cmp.Diff(want, got); d != "" {
			t.Errorf("+got, -want \n%s", d)
		}
	})
}
