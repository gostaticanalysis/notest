package notest

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "notest",
	Doc:  Doc,
	Run:  run,
}

const Doc = "notest checks either the package has test files"

// Match returns either err represents error of notest or not.
func Match(err error) bool {
	_, isErrNoTests := err.(*errNoTests)
	return isErrNoTests
}

type errNoTests struct {
	pass *analysis.Pass
}

func (err *errNoTests) Error() string {
	return fmt.Sprintf("%s has not test files", err.pass.Pkg.Path())
}

func run(pass *analysis.Pass) (interface{}, error) {
	// ignore main and _test package
	switch pkgname := pass.Pkg.Name(); {
	case pkgname == "main", strings.HasSuffix(pkgname, "_test"):
		return nil, nil
	}

	for i := range pass.Files {
		pos := pass.Files[i].Pos()
		dir := filepath.Dir(pass.Fset.File(pos).Name())

		fis, err := ioutil.ReadDir(dir)
		if err != nil {
			return nil, err
		}

		for _, fi := range fis {
			if !fi.IsDir() && strings.HasSuffix(fi.Name(), "_test.go") {
				return nil, nil
			}
		}
	}
	return nil, &errNoTests{pass: pass}
}
