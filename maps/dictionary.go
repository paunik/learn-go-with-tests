package main

import "fmt"

// Dictionary type represents map used for dictionary
type Dictionary map[string]string

const (
	// ErrNotFound no word found
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists word already exists
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordDoesNotExist word cant be updated as there is no word found
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr is error type for dictionary errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search method finds word if there and if not returns ErrNotFound
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add method adds word to dictionary or returns error if word already in dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update method updated word or returns error if word not found
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	fmt.Println(err)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete method deletes string from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
