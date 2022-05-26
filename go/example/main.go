package main

import "fmt"

type Blog struct {
	Title   string
	Content string
	Author  string
	Remark  string
}
type B_interface interface {
	Insert(data map[string]interface{}) error
}

func (b *Blog) Insert(data map[string]interface{}) error {
	blog := b
	blog.Title = data["title"].(string)
	blog.Content = data["content"].(string)
	blog.Author = data["author"].(string)
	blog.Remark = data["remark"].(string)
	return nil
}
func main() {
	fmt.Printf("---------------------interface Start--------------------\n")
	var b B_interface = &Blog{}
	data := map[string]interface{}{
		"title":   "你好",
		"author":  "Ink",
		"content": "内容",
		"remark":  "备注",
	}
	b.Insert(data)
	fmt.Printf("%+v\n", b)
	fmt.Printf("\n---------------------interface End----------------------")
}
