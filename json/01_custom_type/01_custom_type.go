// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// RecordStatus represents the record status, which is int16 underneath.
type RecordStatus int16

// Record status const
const (
	Delete   RecordStatus = -1
	Active   RecordStatus = 1
	Inactive RecordStatus = 0
)

// MarshalJSON returns *r as the JSON encoding of r.
func (r RecordStatus) MarshalJSON() ([]byte, error) {
	switch int16(r) {
	case -1:
		return json.Marshal("Delete")
	case 1:
		return json.Marshal("Active")
	default:
		return json.Marshal("Inactive")
	}
}

// UnmarshalJSON sets *r to a copy of data
func (r *RecordStatus) UnmarshalJSON(data []byte) error {
	if r == nil {
		return errors.New("RecordStatus: UnmarshalJSON on nil pointer")
	}

	// If data is string
	if data[0] == '"' {
		var status string
		if err := json.Unmarshal(data, &status); err != nil {
			return err
		}

		switch status {
		case "Active":
			*r = RecordStatus(1)
		case "Inactive":
			*r = RecordStatus(0)
		default:
			*r = RecordStatus(-1)
		}
	} else {
		*r = RecordStatus(-1)
	}

	return nil
}

type Foo struct {
	Value RecordStatus
}

func main() {
	var f1 Foo
	var f2 Foo

	f1 = Foo{Value: RecordStatus(-1)}
	data, err := json.Marshal(f1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Marshal result = %s\n", data)

	if err := json.Unmarshal([]byte(`{ "value": "Delete" }`), &f2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(f2.Value)

	var int1 int32
	var int2 int16
	int1 = 123
	if v, ok := int16(int1); ok {
		fmt.Println(v)
	}

}
