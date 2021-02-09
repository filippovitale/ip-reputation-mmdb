package mmdbutil

import (
	"testing"
)

func TestFromIPBeginAndEndToIPNet(t *testing.T) {
	if "a.mmdb" != CreateFilename("a.css") {
		t.Error("Error on the extension substitution.")
	}

	if "a/b/cdefg.mmdb" != CreateFilename("a/b/cdefg.css") {
		t.Error("Error on the extension substitution.")
	}
}
