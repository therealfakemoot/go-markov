package main

import (
	"os"
	"regexp"
	"strings"
)

// Normalize takes an input string and returns a slice of whitespace spearated words in all lowercase with all punctuation removed.
func Normalize(s string) []string {
	// This regex is neat. \p{L} means "any letter in any language". \p{Z} means "any whitespace character in any unicode language". I'm using these so the markov engine can be 100% unicode friendly and language agnostic.
	reg, _ := regexp.Compile("[^\\p{L}\\p{Z}]+")
	words := strings.Split(strings.ToLower(reg.ReplaceAllString(s, "")), " ")
	return words

}

// Key represents a n-tuple containing a word sequence
type Key []string

// HashKey creates a string value suitable for use as a map key.
func (k Key) HashKey() NodeKey {
	return NodeKey(strings.Join(k, ","))
}

// NodeKey does stuff
type NodeKey string

// Node contains the probability of a given Key occuring after another.
type Node struct {
	Key   Key
	Count int
}

// HashKey creates a string value suitable for use as a map key.
func (n Node) HashKey() NodeKey {
	return n.Key.HashKey()
}

// Chain does stuff
type Chain struct {
	chain map[NodeKey][]Node
}

// ImportFile does stuff
func (c *Chain) ImportFile(f *os.File) {

}
