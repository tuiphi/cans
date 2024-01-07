package main

import "github.com/tuiphy/cans/list"

var _ list.Item = (*Item)(nil)

type Item struct {
	title, description string
}

func (i Item) FilterValue() string {
	return i.title
}

func (i Item) Title() string {
	return i.FilterValue()
}

func (i Item) Description() string {
	return i.description
}
