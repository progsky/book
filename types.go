package main

import "time"

// IndexPage - основная страница, доступная на адресе "/"
// Здесь отображается список книг, которые хранятся в БД
type IndexPage struct {
	AllBooks []Book
}

// BookPage - это страница описания книги
// Находится по адресу "/book.html"
type BookPage struct {
	TargetBook Book
}

//Структура объекта Book
type Book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}

//Функция для форматирования даты публикации
func (b Book) PublicationDateStr() string {
	return b.PublicationDate.Format("2006-01-02")
}

//Страница ошибки
type ErrorPage struct {
	ErrorMsg string
}
