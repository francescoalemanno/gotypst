package gotypst_test

import (
	"strings"
	"testing"

	"github.com/francescoalemanno/gotypst"
)

func TestPDF(t *testing.T) {
	bts, err := gotypst.PDF([]byte("= hello"))

	if err != nil || len(bts) == 0 {
		t.Errorf("%v", err)
	}
}

func TestVersion(t *testing.T) {
	out, err := gotypst.Version()
	if err != nil || !strings.HasPrefix(out, "typst") {
		t.Errorf("%v ; %v", err, out)
	}
}
