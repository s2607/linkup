package main

import (
)

type response struct {
	key int
	value string
	q question
}

func (r response) Validate(c []criterion) bool {
	//this checks the user's responses to questions it depends on.
	//the dependancy can be checked by q.something
	return true
}

