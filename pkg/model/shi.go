package model

import (
	"fmt"
	"strings"
)

type Shi struct {
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Note       []string `json:"note,omitempty"`
}

type ShiJi []Shi

func (s *ShiJi) Get(poemIndex int) string {
	poemIndex = poemIndex % s.len()
	poem := (*s)[poemIndex]
	if poem.Title == "" {
		poem.Title = "无题"
	}
	if poem.Author == "" {
		poem.Author = "佚名"
	}
	return fmt.Sprintf("《%s》\n\t——%s\n%s", poem.Title, poem.Author, strings.Join(poem.Paragraphs, "\n"))
}

func (s *ShiJi) len() int {
	return len(*s)
}
