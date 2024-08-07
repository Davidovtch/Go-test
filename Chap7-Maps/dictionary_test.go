package main

import "testing"

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dic Dic, word, definition string) {
	t.Helper()

	got, err := dic.Search(word)

	if err != nil {
		t.Fatal("should find addded word:", err)
	}

	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dic := Dic{"test": "this is a test"}

	t.Run("know value", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknow value", func(t *testing.T) {
		_, err := dic.Search("unknow")

		if err == nil {
			t.Fatal("Expected an error")
		}

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dic{}
		word := "test"
		definition := "this is a test"

		err := dic.Add(word, definition)

		assertErrors(t, err, nil)
		assertDefinition(t, dic, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"

		dic := Dic{word: definition}
		err := dic.Add(word, "new test")

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dic, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dic := Dic{word: definition}
		newDefinition := "new def"

		err := dic.Update(word, newDefinition)

		assertErrors(t, err, nil)
		assertDefinition(t, dic, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dic := Dic{}

		err := dic.Update(word, definition)

		assertErrors(t, err, ErrWordDontExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dic{word: "test def"}

	dic.Delete(word)
	_, err := dic.Search(word)

	assertErrors(t, err, ErrNotFound)
}
