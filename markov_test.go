package main

import "testing"

func TestMarkovImport(t *testing.T) {

	t.Run("SimpleSentence", func(t *testing.T) {

		c := NewChain()
		s := "June jumps up and down."
		c.Update(s)

		t.Run("WordCount", func(t *testing.T) {

			if c.Count() != 5 {

				t.Fail()
			}
		})
	})
}
