package gotypst_test

import (
	"testing"

	"github.com/francescoalemanno/gotypst"
)

func TestPDF(t *testing.T) {
	bts, err := gotypst.PDF([]byte("= hello"))

	if err != nil || len(bts) == 0 {
		t.Errorf("%v", err)
	}
}
