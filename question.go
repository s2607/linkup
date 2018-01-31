package main

import (
)

type question struct {
	key int
	prompt string
	qtype int
}

func (q *question) New_response ()  *response {
	return nil
}
