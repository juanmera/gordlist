package gordlist

import (
    "testing"
)

func TestGenerator(t *testing.T) {
    g := New("abcd0123")
    i := 0
    for _ = range g.Generate(1, 5) {
        i += 1
    }
    if i != 37448 {
        t.Fail()
    }
}
