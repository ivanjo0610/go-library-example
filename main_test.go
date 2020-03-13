package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_add(t *testing.T) {
	books := make(map[string]Book)

	cases := []struct {
		desc     string
		exp      map[string]Book
		commands []string
	}{
		{
			desc: "Add new book with good params (book name without space)",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
			},
			commands: []string{"add", "code1", "book"},
		},
		{
			desc: "Add new book with bad params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
			},
			commands: []string{"add", "code1"},
		},
		{
			desc: "Add new book with good params (book name with space)",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"add", "code2", "book", "2"},
		},
		{
			desc: "Add new book but code already exist",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"add", "code2", "book", "3"},
		},
	}

	for _, c := range cases {
		books = cmdAdd(books, c.commands)
		if !reflect.DeepEqual(books, c.exp) {
			t.Errorf("add was incorrect, got: %v, want: %v", books, c.exp)
		}
	}

}

func Test_rent(t *testing.T) {
	books := map[string]Book{
		"code1": Book{Name: "book", Status: 0},
		"code2": Book{Name: "book 2", Status: 0},
	}

	cases := []struct {
		desc     string
		exp      map[string]Book
		commands []string
	}{
		{
			desc: "Rent book with good params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"rent", "code1"},
		},
		{
			desc: "Rent book with bad params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"rent"},
		},
		{
			desc: "Rent book that already rented",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"rent", "code1"},
		},
	}

	for _, c := range cases {
		books = cmdRent(books, c.commands)
		if !reflect.DeepEqual(books, c.exp) {
			t.Log(c.desc)
			t.Errorf("rent was incorrect, got: %v, want: %v", books, c.exp)
		}
	}

}

func Test_return(t *testing.T) {
	books := map[string]Book{
		"code1": Book{Name: "book", Status: 1},
		"code2": Book{Name: "book 2", Status: 1},
	}

	cases := []struct {
		desc     string
		exp      map[string]Book
		commands []string
	}{
		{
			desc: "Return book with good params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"return", "code1"},
		},
		{
			desc: "Return book with bad params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"return"},
		},
		{
			desc: "Return book that already returned",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"return", "code1"},
		},
	}

	for _, c := range cases {
		books = cmdReturn(books, c.commands)
		if !reflect.DeepEqual(books, c.exp) {
			t.Log(c.desc)
			t.Errorf("return was incorrect, got: %v, want: %v", books, c.exp)
		}
	}

}
