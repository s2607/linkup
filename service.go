package main

import (
)

type service struct {
	key int
	criteria map [question] []criterion
	name string
	description string
	url string
}

func (s *service) check(rid  int) bool {
	return true
}

