package main

import (
	"fmt"
)

func main() {
	bookmarks := map[int]string{}
	realizeMenu(bookmarks)
}

func menu() int {
	fmt.Println("Меню")
	fmt.Println("1. Посмотреть закладку")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	var userAnswer int
	fmt.Scan(&userAnswer)
	return userAnswer
}

func seeBookmarks(bookmarks map[int]string) {
	if len(bookmarks) == 0 {
		fmt.Println("Список закладок пуст")
		return
	}
	for k, v := range bookmarks {
		fmt.Printf("Ключ: %v, Значение: %v\n", k, v)
	}
}

func addBookmarks(bookmarks map[int]string) {
	var bokKey int
	fmt.Println("Введите ключ: ")
	fmt.Scan(&bokKey)
	var bokVal string
	fmt.Println("Введите значение: ")
	fmt.Scan(&bokVal)
	bookmarks[bokKey] = bokVal
}

func delBookmarks(bookmarks map[int]string) {
	var delKey int
	fmt.Println("Ведите ключ для удаления закладки: ")
	fmt.Scan(&delKey)
	delete(bookmarks, delKey)
}

func realizeMenu(bookmarks map[int]string) {
	for {
		fmt.Println("Закладки")
		answer := menu()
		switch answer {
		case 1:
			seeBookmarks(bookmarks)
		case 2:
			addBookmarks(bookmarks)
		case 3:
			delBookmarks(bookmarks)
		case 4:
			return
		}
	}
}
