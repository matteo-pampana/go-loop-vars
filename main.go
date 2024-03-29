package main

import "fmt"

type Item struct {
	ID   string
	Name string
}

type ListItem struct {
	ReferenceID *string
}

func main() {
	m := map[string]Item{
		"1": {"1", "Item 1"},
		"2": {"2", "Item 2"},
		"3": {"3", "Item 3"},
		"4": {"4", "Item 4"},
		"5": {"5", "Item 5"},
		"6": {"6", "Item 6"},
	}

	var refs []ListItem
	for _, v := range m {
		refs = append(refs, ListItem{ReferenceID: &v.ID})
		fmt.Println("Item: ", v.ID)
	}

	for _, ref := range refs {
		fmt.Println("Referenced item with id:", *ref.ReferenceID)
	}

}
