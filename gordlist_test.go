package gordlist

import (
    "testing"
)
const STEP = 5000000
func TestGenerator(t *testing.T) {
    g := New("0123456789abcdefghijklmnopqrstuvwxyz")
    results := []string {
        "f9j80u7",
    }
    var i uint64 = 0
    for w := range g.GenerateFrom(6, 7, 35400000000) {
        i += 1
        if i % STEP == 0 {
            t.Log(i/STEP)
            if results[(i/STEP) - 1] != string(w) {
                t.Logf("expected %s instead of %s at %d\n", results[(i/STEP) - 1], w, i)
                t.Fail()
            }
            if i / STEP == uint64(len(results)) {
                break
            }
        }
    }
}

func TestPowUint32(t *testing.T) {
    if powUint32(38, 8) != 4347792138496 {
        t.Fail()
    }
}
