package main

import (
)

type responder struct {
	key int
	responses []response
	fname string
	lname string
	dob string//TODO: time
	zip string
	suggestions []service


}

func (r responder) update_suggestions() error {
	return nil
}

