package gordlist

type Generator struct {
    charset []byte
    charsetLen uint32
}

func New(charset string) *Generator {
    return &Generator{charset: parseCharset(charset)}
}

func (m *Generator) Generate(min, max uint32) (chan []byte) {
    if min < 1 || max < min {
        panic("Invalid parameters")
    }
    m.charsetLen = uint32(len(m.charset))
    out := make(chan []byte)
    go m.generateRangeWords(min, max, out)
    return out
}

func (m *Generator) generateRangeWords(min, max uint32, out chan []byte) {
    var i uint32
    for i = min; i<= max; i++ {
        m.generateWords(i, out)
    }
    close(out)
}

func (m *Generator) generateWords(wordLen uint32, out chan []byte) {
    var i, j, posNo, pos uint32
    wordCount := powUint32(m.charsetLen, wordLen)
    for i = 0; i < wordCount; i++ {
        word := make([]byte, wordLen)
        wordNo := i
        for j = 0; j < wordLen ; j++ {
            if j == wordLen-1 {
                word[j] = m.charset[wordNo % m.charsetLen]
            } else {
                posNo = powUint32(m.charsetLen, wordLen-j-1)
                pos = wordNo/posNo
                wordNo -= (pos * posNo)
                word[j] = m.charset[pos]
            }
        }
        out <- word
    }
}

func powUint32(x, n uint32) uint32 {
    if n == 1 {
        return x
    } else if (n % 2) == 0 {
        return powUint32(x*x, n/2)
    } else {
        return x * powUint32(x*x, (n-1)/2)
    }
}

func parseCharset(charset string) []byte {
    return []byte(charset)
}

