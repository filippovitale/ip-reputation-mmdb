package fileutil

import (
	"testing"
)

func TestFromIPBeginAndEndToIPNet(t *testing.T) {
	if "a.html" != CreateOutputFilename("a.css", "html") {
		t.Error("Error on the extension substitution.")
	}

	if "a/b/cdefg.html" != CreateOutputFilename("a/b/cdefg.css", "html") {
		t.Error("Error on the extension substitution.")
	}
}
