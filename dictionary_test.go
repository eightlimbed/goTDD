package dictionary

import "testing"

func TestSearch(t *testing.T) {

	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("lol")
		assertError(t, err, ErrNotFound)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("update a word", func(t *testing.T) {
		word, definition := "edamame", "a green bean"
		dictionary := Dictionary{word: definition}
		newDefinition := "a soy bean"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("try to update a word that is already defined", func(t *testing.T) {
		dictionary := Dictionary{}
		word, definition := "edamame", "a green bean"
		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestAdd(t *testing.T) {

	dictionary := Dictionary{}
	t.Run("add a word", func(t *testing.T) {
		word, definition := "edamame", "a green bean"
		err := dictionary.Add(word, definition)
		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {
		word, definition := "edamame", "a green bean"
		err := dictionary.Add(word, definition)
		err = dictionary.Add(word, "lol")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestDelete(t *testing.T) {

	t.Run("delete a word that exists", func(t *testing.T) {
		word, definition := "edamame", "a green bean"
		dictionary := Dictionary{word: definition}
		err := dictionary.Delete(word)
		assertError(t, err, nil)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("try to delete a word that does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("lol")
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
