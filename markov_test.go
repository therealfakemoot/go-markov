package main

import "testing"

func TestMarkovImport(t *testing.T) {

	t.Run("Simple Sentence", func(t *testing.T) {

		c := NewChain(1)
		s := "June jumps up and down."
		c.Update(s)

		t.Run("WordCount", func(t *testing.T) {

			if c.Count() != 5 {
				t.Fail()
			}
		})
	})

	t.Run("Duplicate Words", func(t *testing.T) {

		c := NewChain(1)
		s := "The dog is near the bed."
		c.Update(s)

		t.Run("WordCount", func(t *testing.T) {

			if c.Count() != 5 {
				t.Fail()
			}
		})
	})
}
