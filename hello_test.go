package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("saying Hello to people in Thai", func(t *testing.T) {
		got := Hello("Noon", "thai")
		want := "สวัสดี, Noon"

		if got != want {
			t.Errorf("got %s want %s ", got, want)
		}
	})

	t.Run("say 'Hello World' when an empty string is supplied in Thai", func(t *testing.T) {
		got := Hello("", "thai")
		want := "สวัสดี, World"

		if got != want {
			t.Errorf("got %s want %s ", got, want)
		}
	})

}
