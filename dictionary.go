package dictionary

// Dictionary is a map of words and their definitions
type Dictionary map[string]string

// DictErr is a light wrapper around an error created for clatification
type DictErr string

// Error is a DictionaryErr error
func (e DictErr) Error() string {
	return string(e)
}

var (
	// ErrNotFound occurs when a user looks up a word that is not defined
	ErrNotFound = DictErr("that word is not defined")
	// ErrWordExists occurs when a user tries to add a word that is already defined
	ErrWordExists = DictErr("that word is already defined")
	// ErrWordDoesNotExist occurs when a user tries to update a word that is already defined
	ErrWordDoesNotExist = DictErr("that word does not exist")
)

// Search searches for a word in a Dictionary and returns its defintion, or ErrNotFound
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Update adds a word to a Dictionary, or overwrites it if it exists
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}
	d[word] = definition
	return nil
}

// Add adds a word to a Dictionary if it isn't already defined
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		d[word] = definition
		return nil
	}
	return ErrWordExists

}

// Delete deletes a word from a Dictionary. If the word doesn't exist, ErrWordDoesNotExist is returned
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		return ErrWordDoesNotExist
	}
	delete(d, word)
	return nil
}
