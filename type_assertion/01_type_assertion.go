package main

import "fmt"

// RecordStatus represents the record status, which is int16 underneath.
type RecordStatus int16

// Record status const
const (
	Delete   RecordStatus = -1
	Active   RecordStatus = 1
	Inactive RecordStatus = 0
)

func main() {
	t := Delete
	v := bool(t)

	fmt.Println(v)
}
