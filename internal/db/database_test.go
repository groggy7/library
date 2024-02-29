package db

import "testing"

func TestConnection(t *testing.T) {
	_, err := NewDB()
	if err != nil {
		t.Fatal(err)
	}
}
