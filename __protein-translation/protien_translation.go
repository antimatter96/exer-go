package protein

import (
	"errors"
)

var codToPro = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrInvalidBase error = errors.New("Invalid Base")
var STOP error = errors.New("STOP")

// FromCodon returns the corresponding protien to a codon
func FromCodon(codon string) (string, error) {
	pro, present := codToPro[codon]
	if !present {
		return "", ErrInvalidBase
	}
	if pro == "STOP" {
		return "", STOP
	}
	return pro, nil
}

// FromRNA returns an array of protiens
func FromRNA(rna string) ([]string, error) {
	var result []string
	for i := 0; i < len(rna)-2; i += 3 {
		codon := rna[i : i+3]
		pro, present := codToPro[codon]
		if !present {
			return result, ErrInvalidBase
		}
		if pro == "STOP" {
			return result, nil
		}
		result = append(result, pro)
	}
	return result, nil
}
