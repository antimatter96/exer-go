package dna

import (
	"errors"
	"fmt"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	return 0, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := make(Histogram)
	h['A'] = 0
	h['C'] = 0
	h['T'] = 0
	h['G'] = 0
	for i := 0; i < len(d); i++ {
		if d[i] == 'A' || d[i] == 'C' || d[i] == 'T' || d[i] == 'G' {
			h[d[i]]++
		} else {
			return nil, errors.New("Unexpected nucleotide")
		}
	}
	fmt.Println(h)
	return h, nil
}
