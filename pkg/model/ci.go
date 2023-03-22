package model

import (
	"fmt"
	"strings"
)

type Ci struct {
	Title      string   `json:"rhythmic"`
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
}

type CiJi []Ci

func (c *CiJi) Get(poemIndex int) string {
	poemIndex = poemIndex % c.len()
	poem := (*c)[poemIndex]
	if poem.Title == "" {
		poem.Title = "无题"
	}
	if poem.Author == "" {
		poem.Author = "佚名"
	}
	return fmt.Sprintf("《%s》\n\t——%s\n%s", poem.Title, poem.Author, strings.Join(poem.Paragraphs, "\n"))
}

func (c *CiJi) len() int {
	return len(*c)
}
