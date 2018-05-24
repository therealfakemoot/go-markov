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

// Nodes is a list of probability entries
type Nodes []*Node

// Get iterates through a Nodes sequence and finds a matching probability entry.
func (n *Nodes) Get(nk NodeKey) *Node {
	for _, n := range *n {
		if n.HashKey() == nk {
			return n
		}
	}
	return nil
}

// Chain does stuff
type Chain struct {
	Mode  int
	Chain map[NodeKey]Nodes
}

// Count returns the number of keys in the markov chain
func (c *Chain) Count() int {
	return len(c.Chain)
}

// Update adds new keys and updates existing node counts based on a given input string
func (c *Chain) Update(s string) {

	words := Normalize(s)

	for _, w := range words {
		k := NodeKey(w)
		nodes, ok := c.Chain[k]

		if ok {
			n := nodes.Get(k)
			n.Count++
		} else {
			n := Node{Key: []string{w}, Count: 1}
			c.Chain[k] = Nodes{&n}
		}

	}
}

// Speak produces a sentence at random from the chain.
func (c *Chain) Speak(count int) string {
	var res string
	return res
}

// Respond produces a sentence given a prompt.
func (c *Chain) Respond(prompt string, count int) string {
	var res string
	return res
}

// ImportFile does stuff
func (c *Chain) ImportFile(f *os.File) {

}

// NewChain returns a new Markov Chain
func NewChain(mode int) *Chain {
	c := Chain{Mode: mode, Chain: make(map[NodeKey]Nodes)}
	return &c
}
