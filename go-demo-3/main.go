package main

import "fmt"

type bookmarkType = map[string]string

func main() {
	bookmarks := bookmarkType{}
	fmt.Println("Приложение для закладок")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			printBookmarks(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Scan(&variant)
	return variant
}

func printBookmarks(bookmarks bookmarkType) {
	if len(bookmarks) == 0 {
		fmt.Println("Пока нет закладок")
	}
	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}

}

func addBookmark(bookmarks bookmarkType) bookmarkType {
	var newBookmarkKey string
	var newBookmarkVal string
	fmt.Println("Введите ключ")
	fmt.Scan(&newBookmarkKey)
	fmt.Println("Введите значение")
	fmt.Scan(&newBookmarkVal)
	bookmarks[newBookmarkKey] = newBookmarkVal
	return bookmarks
}

func deleteBookmark(bookmarks bookmarkType) bookmarkType {
	var bookmarkKeyToDelete string
	fmt.Println("Введите ключ на удаление")
	fmt.Scan(&bookmarkKeyToDelete)
	delete(bookmarks, bookmarkKeyToDelete)
	return bookmarks
}
