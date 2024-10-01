package library

import (
	"errors"
	"sync"
)

type Book struct {
	ID    int
	Title string
}

type Library struct {
	books    []Book
	bookIdx  map[int]Book
	titleIdx map[string]int
	nextID   func() int
	mu       sync.Mutex
}

func NewLibrary(idGenerator func() int) *Library {
	return &Library{
		books:    []Book{},
		bookIdx:  make(map[int]Book),
		titleIdx: make(map[string]int),
		nextID:   idGenerator,
	}
}

func (lib *Library) AddBook(title string) int {
	id := lib.nextID()
	book := Book{
		ID:    id,
		Title: title,
	}

	lib.mu.Lock()
	defer func() { lib.mu.Unlock() }()
	lib.books = append(lib.books, book)
	lib.bookIdx[id] = book
	lib.titleIdx[title] = id

	return id
}

func (lib *Library) GetBookById(id int) (Book, error) {
	book, ex := lib.bookIdx[id]
	if !ex {
		return Book{}, errors.New("book not found")
	}

	return book, nil
}

func (lib *Library) GetBookByTitle(title string) (Book, error) {
	id, ex := lib.titleIdx[title]
	if !ex {
		return Book{}, errors.New("book not found")
	}
	book, err := lib.GetBookById(id)
	if err != nil {
		return Book{}, err
	}

	return book, nil
}

func (lib *Library) SetIDGenerator(newGenerator func() int) {
	lib.mu.Lock()
	defer func() { lib.mu.Unlock() }()
	lib.nextID = newGenerator
}

func (lib *Library) ReplaceStorage() {
	lib.bookIdx = make(map[int]Book)
	lib.titleIdx = make(map[string]int)
	lib.books = []Book{}
}
