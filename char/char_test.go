package char

import (
	"testing"
)

func TestFwSlash(t *testing.T) {
	if FwSlash(`\hello\world`) != "/hello/world" {
		t.Error(`expected FwSlash("\hello\world") equal to  "/hello/world"`)
	}
}

func TestBwSlash(t *testing.T) {
	if BwSlash(`/hello/world`) != `\hello\world` {
		t.Error(`expected BwSlash("\hello\world") equal to  "\hello\world"`)
	}
}
