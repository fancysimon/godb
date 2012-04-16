package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1,1) != 2 {
		t.Error("1+1!=2")
	}
}
