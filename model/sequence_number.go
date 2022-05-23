package model

import "fmt"

type SequenceNumber int

func (number SequenceNumber) String() string {
	return fmt.Sprintf("%03d/%03d/%03d", int(number/1000000), int(number/1000)-(int(number/1000000)*1000), number%1000)
}
