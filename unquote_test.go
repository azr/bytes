package bytes_test

import (
	"testing"

	"github.com/azr/bytes"
)

func TestUnquoteToString(t *testing.T) {
	b, ok := bytes.UnquoteToString([]byte(`"hey\n"`))
	bs := `hey
`
	if !ok {
		t.Fatal("UnquoteToString")
	}
	if b != bs {
		t.Fatalf("UnqoteToString failed : %+q != %+q", b, bs)
	}
}
