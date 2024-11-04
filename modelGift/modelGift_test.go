package modelgift

import (
	"fmt"
	"testing"
)

// func TestCreateGift(t *testing.T) {

// 	newGift := CreateGift("book","www.book.com",25,5,2)

// }


func TestSliceStringGift(t *testing.T) {
	theSliceString,err := SliceStringGift("gift.txt")
	if err!=nil{
		t.Error(err)
	}

	fmt.Println(theSliceString)
	
}