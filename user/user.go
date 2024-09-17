package user

import (
	"fmt"
	"net/http"
)

type TODO struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func PerformGetRequest () {
	const url = "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)
	if (err != nil) {
		panic(err);
	}

	defer response.Body.Close()

	if (response.StatusCode != http.StatusOK) {
		fmt.Println("failed To fetch data:", response.Status)
	}

	
}