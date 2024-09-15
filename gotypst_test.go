package gotypst_test

import (
	"testing"

	"github.com/francescoalemanno/gotypst"
)

func TestPDF(t *testing.T) {
	bts, err := gotypst.PDF([]byte("= hello"))

	if bts[0] != 37 || err != nil {
		t.Errorf("NOT WORKING")
	}
}
