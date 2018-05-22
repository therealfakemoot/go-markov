package main

import (
	"os"
	"strings"
)

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
