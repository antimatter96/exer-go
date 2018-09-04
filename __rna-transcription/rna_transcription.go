package strand

import "strings"

// var m = map[byte]byte{
// 	'G': 'C',
// 	'C': 'G',
// 	'T': 'A',
// 	'A': 'U',
// }

func ToRNA(dna string) string {
	var b strings.Builder
	b.Grow(len(dna))
	for i := 0; i < len(dna); i++ {
		switch dna[i] {
		case 'G':
			b.WriteByte('C')
		case 'C':
			b.WriteByte('G')
		case 'T':
			b.WriteByte('A')
		case 'A':
			b.WriteByte('U')
		}
	}
	return b.String()
}
