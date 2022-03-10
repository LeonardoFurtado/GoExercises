package main

import "testing"

func TestName(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("Testing Origem", func(t *testing.T) {
		got := CoordinatePoint(0.0, 0.0)
		want := "Origem"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing Q4", func(t *testing.T) {
		got := CoordinatePoint(4.5, -2.2)
		want := "Q4"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing Q3", func(t *testing.T) {
		got := CoordinatePoint(-1.0, -2.0)
		want := "Q3"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing Q2", func(t *testing.T) {
		got := CoordinatePoint(-1.0, 2.0)
		want := "Q2"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing Q1", func(t *testing.T) {
		got := CoordinatePoint(1.0, 2.0)
		want := "Q1"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing Eixo X", func(t *testing.T) {
		got := CoordinatePoint(0.1, 0.0)
		want := "Eixo X"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing Eixo Y", func(t *testing.T) {
		got := CoordinatePoint(0.0, 1.0)
		want := "Eixo Y"

		assertCorrectMessage(t, got, want)
	})
}
