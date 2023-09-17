package main

import "fmt"

// //
// **Toy Gene Example 1**

// >gene

// TCCCTGGGCTCTTTTAGTGGACGGAGACCCAGCTGTCAGTTTGTTGTAATAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACACCACTTACTCCATGGAGGGATTAGATGATTCACGGTAGGCTTGGGCAG

// >oligo1

// TCCCTGGGCTCTTTTAGTGGACGGAGACCCAGCTGTCAGTTTGTTGTAATAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

// >oligo2

// CTGCCCAAGCCTACCGTGAATCATCTAATCCCTCCATGGAGTAAGTGGTGTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT

func main() {
	//naively assume any site >10 nucleotides will result in binding regardless of affinity
	//assume only pcr
	//check for complements in every oligo
	//compute all combinations of complements down order of the tree
	//check oligos against database of dangerous dna strands
	
 }

 func oligoComb(oligo string) []string {
	for i := 0; i < len(oligo); i++ {
		
	}
	return nil
 }