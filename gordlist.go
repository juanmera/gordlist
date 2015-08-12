package gordlist

import "math"

type Generator struct {
    charset []byte
    charsetLen uint32
    wordCount uint64
}

func New(charset string) *Generator {
    return &Generator{charset: parseCharset(charset)}
}

func (m *Generator) WordCount() uint64 {
    return m.wordCount
}

func (m *Generator) GenerateFrom(min, max uint32, start uint64) (chan []byte) {
    if min < 1 || max < min {
        panic("Invalid parameters")
    }
    m.charsetLen = uint32(len(m.charset))
    out := make(chan []byte)
    go m.generateRangeWords(min, max, start, out)
    return out
}

func (m *Generator) Generate(min, max uint32) (chan []byte) {
    return m.GenerateFrom(min, max, 0)
}

func (m *Generator) generateRangeWords(min, max uint32, start uint64, out chan []byte) {
    var i uint32
    for i = min; i<= max; i++ {
        m.generateWords(i, start, out)
    }
    close(out)
}

func (m *Generator) generateWords(wordLen uint32, start uint64, out chan []byte) {
    var i, pos, posNo, wordNo uint64
    var j  uint32
    wordCount := powUint32(m.charsetLen, wordLen)
    relativeStart := uint64(math.Max(0, float64(start) - float64(m.wordCount)))
    for i = relativeStart; i < wordCount; i++ {
        word := make([]byte, wordLen)
        wordNo = i
        for j = 0; j < wordLen ; j++ {
            if j == wordLen-1 {
                word[j] = m.charset[wordNo % uint64(m.charsetLen)]
            } else {
                posNo = powUint32(m.charsetLen, wordLen-j-1)
                pos = wordNo/posNo
                wordNo -= (pos * posNo)
                word[j] = m.charset[pos]
            }
        }
        out <- word
        m.wordCount += 1
    }
    if relativeStart >= wordCount {
        m.wordCount += wordCount
    }
}

func powUint32(x, n uint32) uint64 {
    if n == 1 {
        return uint64(x)
    } else if (n % 2) == 0 {
        return powUint32(x*x, n/2)
    } else {
        return uint64(x) * powUint32(x*x, (n-1)/2)
    }
}

func parseCharset(charset string) []byte {
    return []byte(charset)
}

