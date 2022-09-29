package main

import "testing"

func TestEx1(t *testing.T) {
/* 	file := `<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>`
 */
	got := 1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}