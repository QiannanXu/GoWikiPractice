package main

import (
	"io/ioutil"
	"fmt"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	page := Page{Title: "Testpage", Body: []byte("This is the test page.")}
	page.save()
	loadedPage, _ := loadPage("Testpage")
	fmt.Println(string(loadedPage.Body))
}
