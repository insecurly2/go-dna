package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("dnafile")
	check(err)
	fmt.Println(rna2ammino(dna2rna(string(dat))))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func rna2ammino(r string) string {
	// u
	r = strings.ReplaceAll(r, "uuu", "phe")
	r = strings.ReplaceAll(r, "uuc", "phe")
	r = strings.ReplaceAll(r, "uua", "leu")
	r = strings.ReplaceAll(r, "uug", "leu")

	r = strings.ReplaceAll(r, "ucu", "ser")
	r = strings.ReplaceAll(r, "ucc", "ser")
	r = strings.ReplaceAll(r, "uca", "ser")
	r = strings.ReplaceAll(r, "ucg", "ser")

	r = strings.ReplaceAll(r, "uau", "tyr")
	r = strings.ReplaceAll(r, "uac", "tyr")
	r = strings.ReplaceAll(r, "uaa", "stp")
	r = strings.ReplaceAll(r, "uag", "stp")

	r = strings.ReplaceAll(r, "ugu", "cys")
	r = strings.ReplaceAll(r, "ugc", "cys")
	r = strings.ReplaceAll(r, "uga", "stp")
	r = strings.ReplaceAll(r, "ugg", "trp")

	// c
	r = strings.ReplaceAll(r, "cuu", "leu")
	r = strings.ReplaceAll(r, "cuc", "leu")
	r = strings.ReplaceAll(r, "cua", "leu")
	r = strings.ReplaceAll(r, "cug", "leu")

	r = strings.ReplaceAll(r, "ccu", "pro")
	r = strings.ReplaceAll(r, "ccc", "pro")
	r = strings.ReplaceAll(r, "cca", "pro")
	r = strings.ReplaceAll(r, "ccg", "pro")

	r = strings.ReplaceAll(r, "cau", "his")
	r = strings.ReplaceAll(r, "cac", "his")
	r = strings.ReplaceAll(r, "caa", "gln")
	r = strings.ReplaceAll(r, "cag", "gln")

	r = strings.ReplaceAll(r, "cgu", "arg")
	r = strings.ReplaceAll(r, "cgc", "arg")
	r = strings.ReplaceAll(r, "cga", "arg")
	r = strings.ReplaceAll(r, "cgg", "arg")

	// a
	r = strings.ReplaceAll(r, "auu", "ile")
	r = strings.ReplaceAll(r, "auc", "ile")
	r = strings.ReplaceAll(r, "aua", "ile")
	r = strings.ReplaceAll(r, "aug", "met")

	r = strings.ReplaceAll(r, "acu", "thr")
	r = strings.ReplaceAll(r, "acc", "thr")
	r = strings.ReplaceAll(r, "aca", "thr")
	r = strings.ReplaceAll(r, "acg", "thr")

	r = strings.ReplaceAll(r, "aau", "asn")
	r = strings.ReplaceAll(r, "aac", "asn")
	r = strings.ReplaceAll(r, "aaa", "lys")
	r = strings.ReplaceAll(r, "aag", "lys")

	r = strings.ReplaceAll(r, "agu", "ser")
	r = strings.ReplaceAll(r, "agc", "ser")
	r = strings.ReplaceAll(r, "aga", "arg")
	r = strings.ReplaceAll(r, "agg", "arg")

	// g
	r = strings.ReplaceAll(r, "guu", "val")
	r = strings.ReplaceAll(r, "guc", "val")
	r = strings.ReplaceAll(r, "gua", "val")
	r = strings.ReplaceAll(r, "gug", "val")

	r = strings.ReplaceAll(r, "gcu", "ala")
	r = strings.ReplaceAll(r, "gcc", "ala")
	r = strings.ReplaceAll(r, "gca", "ala")
	r = strings.ReplaceAll(r, "gcg", "ala")

	r = strings.ReplaceAll(r, "gau", "asp")
	r = strings.ReplaceAll(r, "gac", "asp")
	r = strings.ReplaceAll(r, "gaa", "glu")
	r = strings.ReplaceAll(r, "gag", "glu")

	r = strings.ReplaceAll(r, "ggu", "gly")
	r = strings.ReplaceAll(r, "ggc", "gly")
	r = strings.ReplaceAll(r, "gga", "gly")
	r = strings.ReplaceAll(r, "ggg", "gly")

	return r
}

func dna2rna(r string) string {
	r = strings.ReplaceAll(r, "a", "U")
	r = strings.ReplaceAll(r, "t", "A")
	r = strings.ReplaceAll(r, "g", "C")
	r = strings.ReplaceAll(r, "c", "G")

	r = strings.ReplaceAll(r, "A", "a")
	r = strings.ReplaceAll(r, "U", "u")
	r = strings.ReplaceAll(r, "G", "g")
	r = strings.ReplaceAll(r, "C", "c")

	return r
}
