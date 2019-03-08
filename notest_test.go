package notest_test

import (
	"testing"

	"github.com/gostaticanalysis/notest"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []struct {
		dir       string
		expectErr bool
	}{
		{"a", true},
		{"b", false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.dir, func(t *testing.T) {
			var errors []error
			testingT := testutil.Filter(t, func(format string, args ...interface{}) bool {
				err, isErr := args[len(args)-1].(error)
				if !isErr || !notest.Match(err) {
					return true
				}
				errors = append(errors, err)
				return false
			})
			analysistest.Run(testingT, testdata, notest.Analyzer, tt.dir)
			switch {
			case tt.expectErr && len(errors) == 0:
				t.Error("expected error did not occur")
			case !tt.expectErr && len(errors) != 0:
				t.Errorf("unexpected error: %v", errors)
			}
		})
	}
}
