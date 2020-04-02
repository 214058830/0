package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Casablanca", Year: 1943, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Casablanca", Year: 1944, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
}

func main() {
	// fmt.Println(movies)
	data, err := json.Marshal(movies) // 返回字节slice格式的json
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	data2, err := json.MarshalIndent(movies, "", "	") // 返回格式化字节slice格式的json
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)

	// Json > 结构体
	var title []struct {
		Year int `json:"released"`
	}
	// 只需Year时只定义即可，其余解析内容将丢弃
	if err := json.Unmarshal(data, &title); err != nil {
		log.Fatalf("JSON Unmarshaling failed: %s", err)
	}
	fmt.Println(title)
}
