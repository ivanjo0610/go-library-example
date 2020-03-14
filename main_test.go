package main

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func Test_list(t *testing.T) {
	cases := []struct {
		desc     string
		expErr   error
		books    map[string]Book
		commands []string
	}{
		{
			desc:     "List book with no books",
			expErr:   errors.New("This library currently doesn't have any books"),
			books:    map[string]Book{},
			commands: []string{"list"},
		},
		{
			desc:   "List book with bad params",
			expErr: errors.New("Illegal parameter"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"list", "code1"},
		},
	}

	for _, c := range cases {
		err := cmdList(c.books, c.commands)
		if c.expErr != err { //check error
			if c.expErr != nil && err != nil { //check if error content is same
				if c.expErr.Error() != err.Error() {
					t.Log(c.desc)
					t.Errorf("list error was incorrect, got: %v, want: %v", err, c.expErr)
				}
			} else {
				t.Log(c.desc)
				t.Errorf("list error was incorrect, got: %v, want: %v", err, c.expErr)
			}
		}
	}
}

func Test_get(t *testing.T) {
	cases := []struct {
		desc     string
		exp      string
		expErr   error
		books    map[string]Book
		commands []string
	}{
		{
			desc:   "Get book with bad params",
			exp:    "",
			expErr: errors.New("Illegal parameter"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"get"},
		},
		{
			desc:   "Get book with bad code parameter",
			exp:    "",
			expErr: errors.New("Book not found"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"get", "code4"},
		},
	}

	for _, c := range cases {
		name, err := cmdGet(c.books, c.commands)
		if name != c.exp {
			t.Log(c.desc)
			t.Errorf("get was incorrect, got: %v, want: %v", err, c.expErr)
		}
		if c.expErr != err { //check error
			if c.expErr != nil && err != nil { //check if error content is same
				if c.expErr.Error() != err.Error() {
					t.Log(c.desc)
					t.Errorf("get error was incorrect, got: %v, want: %v", err, c.expErr)
				}
			} else {
				t.Log(c.desc)
				t.Errorf("get error was incorrect, got: %v, want: %v", err, c.expErr)
			}
		}
	}
}
func Test_add(t *testing.T) {
	cases := []struct {
		desc     string
		exp      map[string]Book
		expErr   error
		books    map[string]Book
		commands []string
	}{
		{
			desc: "Add new book with good params (book name without space)",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
			},
			expErr:   nil,
			books:    map[string]Book{},
			commands: []string{"add", "code1", "book"},
		},
		{
			desc:     "Add new book with bad params",
			exp:      map[string]Book{},
			expErr:   errors.New("Illegal parameter"),
			books:    map[string]Book{},
			commands: []string{"add", "code1"},
		},
		{
			desc: "Add new book with good params (book name with space)",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			expErr: nil,
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
			},
			commands: []string{"add", "code2", "book", "2"},
		},
		{
			desc: "Add new book but code already exist",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			expErr: errors.New("Failed to add book! Book code already exist"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"add", "code2", "book", "3"},
		},
	}

	var err error
	for _, c := range cases {
		c.books, err = cmdAdd(c.books, c.commands)
		if !reflect.DeepEqual(c.books, c.exp) { //check books items
			t.Log(c.desc)
			t.Errorf("add books was incorrect, got: %v, want: %v", c.books, c.exp)
		}
		if c.expErr != err { //check error
			if c.expErr != nil && err != nil { //check if error content is same
				if c.expErr.Error() != err.Error() {
					t.Log(c.desc)
					t.Errorf("add error was incorrect, got: %v, want: %v", err, c.expErr)
				}
			} else {
				t.Log(c.desc)
				t.Errorf("add error was incorrect, got: %v, want: %v", err, c.expErr)
			}
		}
	}

}

func Test_rent(t *testing.T) {
	cases := []struct {
		desc     string
		exp      map[string]Book
		expErr   error
		books    map[string]Book
		commands []string
	}{
		{
			desc: "Rent book with good params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 0},
			},
			expErr: nil,
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"rent", "code1"},
		},
		{
			desc: "Rent book with bad params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			expErr: errors.New("Illegal parameter"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"rent"},
		},
		{
			desc: "Rent book that already rented",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			expErr: errors.New("Failed to rent book! book 2 is already rented"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"rent", "code2"},
		},
		{
			desc: "Rent book with bad code parameter",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			expErr: errors.New("Book not found"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"rent", "code3"},
		},
	}

	var err error
	for _, c := range cases {
		c.books, err = cmdRent(c.books, c.commands)
		if !reflect.DeepEqual(c.books, c.exp) {
			t.Log(c.desc)
			t.Errorf("rent was incorrect, got: %v, want: %v", c.books, c.exp)
		}
		if c.expErr != err { //check error
			if c.expErr != nil && err != nil { //check if error content is same
				if c.expErr.Error() != err.Error() {
					t.Log(c.desc)
					t.Errorf("rent error was incorrect, got: %v, want: %v", err, c.expErr)
				}
			} else {
				t.Log(c.desc)
				t.Errorf("rent error was incorrect, got: %v, want: %v", err, c.expErr)
			}
		}
	}

}

func Test_return(t *testing.T) {
	cases := []struct {
		desc     string
		exp      map[string]Book
		expErr   error
		books    map[string]Book
		commands []string
	}{
		{
			desc: "Return book with good params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			expErr: nil,
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"return", "code1"},
		},
		{
			desc: "Return book with bad params",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 1},
			},
			expErr: errors.New("Illegal parameter"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
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
			expErr: errors.New("Failed to return book! book is not rented"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"return", "code1"},
		},
		{
			desc: "Return book with bad code parameter",
			exp: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 1},
			},
			expErr: errors.New("Book not found"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"return", "code4"},
		},
	}

	var err error
	for _, c := range cases {
		c.books, err = cmdReturn(c.books, c.commands)
		if !reflect.DeepEqual(c.books, c.exp) {
			t.Log(c.desc)
			t.Errorf("return was incorrect, got: %v, want: %v", c.books, c.exp)
		}
		if c.expErr != err { //check error
			if c.expErr != nil && err != nil { //check if error content is same
				if c.expErr.Error() != err.Error() {
					t.Log(c.desc)
					t.Errorf("return error was incorrect, got: %v, want: %v", err, c.expErr)
				}
			} else {
				t.Log(c.desc)
				t.Errorf("return error was incorrect, got: %v, want: %v", err, c.expErr)
			}
		}
	}
}

func Test_rented(t *testing.T) {
	cases := []struct {
		desc     string
		expErr   error
		books    map[string]Book
		commands []string
	}{
		{
			desc:   "List rented book with no books rented",
			expErr: errors.New("Currently no books is rented"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 0},
				"code2": Book{Name: "book 2", Status: 0},
			},
			commands: []string{"rented"},
		},
		{
			desc:   "List rented book with bad params",
			expErr: errors.New("Illegal parameter"),
			books: map[string]Book{
				"code1": Book{Name: "book", Status: 1},
				"code2": Book{Name: "book 2", Status: 1},
			},
			commands: []string{"rented", "code1"},
		},
	}

	for _, c := range cases {
		err := cmdRented(c.books, c.commands)
		if c.expErr != err { //check error
			if c.expErr != nil && err != nil { //check if error content is same
				if c.expErr.Error() != err.Error() {
					t.Log(c.desc)
					t.Errorf("list error was incorrect, got: %v, want: %v", err, c.expErr)
				}
			} else {
				t.Log(c.desc)
				t.Errorf("list error was incorrect, got: %v, want: %v", err, c.expErr)
			}
		}
	}
}
