package gordlist

import (
    "math"
    "fmt"
)
var Debug bool
type Generator struct {
    charset []byte
    charsetLen uint64
    wordCount uint64
}

func New(charset string) *Generator {
    return &Generator{charset: parseCharset(charset)}
}

func (m *Generator) WordCount() uint64 {
    return m.wordCount
}

func (m *Generator) GenerateFrom(min, max, start uint64) (chan []byte) {
    if min < 1 || max < min {
        panic("Invalid parameters")
    }
    m.charsetLen = uint64(len(m.charset))
    out := make(chan []byte, 1000)
    go m.generateRangeWords(min, max, start, out)
    return out
}

func (m *Generator) Generate(min, max uint64) (chan []byte) {
    return m.GenerateFrom(min, max, 0)
}

func (m *Generator) generateRangeWords(min, max, start uint64, out chan []byte) {
    for i := min; i<= max; i++ {
        m.generateWords(i, start, out)
    }
    close(out)
}

func (m *Generator) generateWords(wordLen, start uint64, out chan []byte) {
    var i, j, pos, posNo, wordNo uint64
    wordCount := powUint64(m.charsetLen, wordLen)
    relativeStart := uint64(math.Max(0, float64(start) - float64(m.wordCount)))
    m.wordCount += uint64(math.Min(float64(relativeStart), float64(wordCount)))
    if Debug {
        fmt.Printf("Current Len: %-3d Count: %-14d Start: %-14d | Global Count: %-14d Start: %-14d\n", wordLen, wordCount, relativeStart, m.wordCount, start)
    }
    for i = relativeStart; i < wordCount; i++ {
        word := make([]byte, wordLen)
        wordNo = i
        for j = 0; j < wordLen ; j++ {
            if j == wordLen-1 {
                word[j] = m.charset[wordNo % m.charsetLen]
            } else {
                posNo = powUint64(m.charsetLen, wordLen-j-1)
                pos = wordNo/posNo
                wordNo -= (pos * posNo)
                word[j] = m.charset[pos]
            }
        }
        out <- word
        m.wordCount += 1
    }
}

func powUint64(x, n uint64) uint64 {
    if n == 1 {
        return x
    } else if (n % 2) == 0 {
        return powUint64(x*x, n/2)
    } else {
        return x * powUint64(x*x, (n-1)/2)
    }
}

func parseCharset(charset string) []byte {
    return []byte(charset)
}

