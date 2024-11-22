package library

import (
	"main/task1/storage"
)

type Library struct {
	storage storage.Storage
	nextID  func() int
}

func NewLibrary(storage storage.Storage, idGenerator func() int) *Library {
	return &Library{
		storage: storage,
		nextID:  idGenerator,
	}
}

func (lib *Library) AddBook(title string) int {
	id := lib.nextID()
	book := storage.Book{
		ID:    id,
		Title: title,
	}
	lib.storage.AddBook(book)
	return id
}

func (lib *Library) GetBookByID(id int) (storage.Book, error) {
	return lib.storage.GetBookByID(id)
}

func (lib *Library) GetBookByTitle(title string) (storage.Book, error) {
	return lib.storage.GetBookByTitle(title)
}

func (lib *Library) SetIDGenerator(newGenerator func() int) {
	lib.nextID = newGenerator
}

func (lib *Library) ReplaceStorage() {
	lib.storage.ReplaceStorage()
}
