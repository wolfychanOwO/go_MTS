package storage

import (
	"errors"
)

type Book struct {
	ID    int
	Title string
}

type Storage interface {
	AddBook(book Book)
	GetBookByID(id int) (Book, error)
	GetBookByTitle(title string) (Book, error)
	ReplaceStorage()
}

type MapStorage struct {
	bookIdx  map[int]Book
	titleIdx map[string]int
}

func NewMapStorage() *MapStorage {
	return &MapStorage{
		bookIdx:  make(map[int]Book),
		titleIdx: make(map[string]int),
	}
}

func (s *MapStorage) AddBook(book Book) {
	s.bookIdx[book.ID] = book
	s.titleIdx[book.Title] = book.ID
}

func (s *MapStorage) GetBookByID(id int) (Book, error) {
	book, exists := s.bookIdx[id]
	if !exists {
		return Book{}, errors.New("book not found")
	}
	return book, nil
}

func (s *MapStorage) GetBookByTitle(title string) (Book, error) {
	id, exists := s.titleIdx[title]
	if !exists {
		return Book{}, errors.New("book not found")
	}
	return s.GetBookByID(id)
}

func (s *MapStorage) ReplaceStorage() {
	s.bookIdx = make(map[int]Book)
	s.titleIdx = make(map[string]int)
}

type SliceStorage struct {
	books []Book
}

func NewSliceStorage() *SliceStorage {
	return &SliceStorage{
		books: []Book{},
	}
}

func (s *SliceStorage) AddBook(book Book) {
	s.books = append(s.books, book)
}

func (s *SliceStorage) GetBookByID(id int) (Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, errors.New("book not found")
}

func (s *SliceStorage) GetBookByTitle(title string) (Book, error) {
	for _, book := range s.books {
		if book.Title == title {
			return book, nil
		}
	}
	return Book{}, errors.New("book not found")
}

func (s *SliceStorage) ReplaceStorage() {
	s.books = []Book{}
}
