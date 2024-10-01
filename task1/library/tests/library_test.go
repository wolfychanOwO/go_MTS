package library

import (
	"fmt"
	"main/task1/library"
	"testing"
)

func TestLibraryAddAndFindBook(t *testing.T) {
	fmt.Println("Test Starting...")
	idCounter := 1
	idGenerator := func() int {
		id := idCounter
		idCounter++
		return id
	}
	lib := library.NewLibrary(idGenerator)

	books := []library.Book{
		{Title: "1984"},
		{Title: "Far Beyond the World"},
	}

	fmt.Println("Adding books...")
	for _, book := range books {
		lib.AddBook(book.Title)
	}

	fmt.Println("Getting book by title...")
	book, err := lib.GetBookByTitle("1984")
	if err != nil {
		t.Errorf("Expected book '1984' but got error: %v", err)
	}
	if book.Title != "1984" {
		t.Errorf("Expected book '1984' but got '%s'", book.Title)
	}

	book, err = lib.GetBookByTitle("Far Beyond the World")
	if err != nil {
		t.Errorf("Expected book 'Far Beyond the World' but got error: %v", err)
	}
	if book.Title != "Far Beyond the World" {
		t.Errorf("Expected book 'Far Beyond the World' but got '%s'", book.Title)
	}
}

func TestLibraryReplaceIdGenerator(t *testing.T) {
	idCounter := 1
	idGenerator := func() int {
		id := idCounter
		idCounter++
		return id
	}
	lib := library.NewLibrary(idGenerator)

	fmt.Println("Adding books with first generator...")
	lib.AddBook("1984")

	fmt.Println("Generator change...")
	idCounter = 1
	NewidGenerator := func() int {
		id := idCounter*2 + 1
		idCounter++
		return id
	}
	lib.SetIDGenerator(NewidGenerator)

	fmt.Println("Adding book with another generator...")
	newBookID := lib.AddBook("The Catcher in the Rye")

	fmt.Println("Getting book by id...")
	book, err := lib.GetBookById(newBookID)
	if err != nil {
		t.Errorf("Expected book 'The Catcher in the Rye' but got error: %v", err)
	}
	if book.Title != "The Catcher in the Rye" {
		t.Errorf("Expected book 'The Catcher in the Rye' but got '%s'", book.Title)
	}
	if book.ID != 3 {
		t.Errorf("Expected ID '3' but got '%d'", book.ID)
	}
}

func TestLibraryReplaceStorage(t *testing.T) {
	idCounter := 1
	idGenerator := func() int {
		id := idCounter
		idCounter++
		return id
	}
	lib := library.NewLibrary(idGenerator)

	books := []library.Book{
		{Title: "1984"},
		{Title: "Far Beyond the World"},
	}

	fmt.Println("Adding books with one storage...")
	for _, book := range books {
		lib.AddBook(book.Title)
	}

	lib.ReplaceStorage()

	books = []library.Book{
		{Title: "Moby Dick"},
		{Title: "Kakaya ti kNIGGA"},
	}

	fmt.Println("Adding books with another storage...")
	for _, book := range books {
		lib.AddBook(book.Title)
	}

	fmt.Println("Getting book by title...")
	book, err := lib.GetBookByTitle("Moby Dick")
	if err != nil {
		t.Errorf("Expected title 'Moby Dick' but got error: %v", err)
	}
	if book.Title != "Moby Dick" {
		t.Errorf("Expected title 'Moby Dick' but got '%s'", book.Title)
	}

	book, err = lib.GetBookByTitle("Kakaya ti kNIGGA")
	if err != nil {
		t.Errorf("Expected title 'Kakaya ti kNIGGA' but got error: %v", err)
	}
	if book.Title != "Kakaya ti kNIGGA" {
		t.Errorf("Expected title 'Kakaya ti kNIGGA' but got '%s'", book.Title)
	}
}
