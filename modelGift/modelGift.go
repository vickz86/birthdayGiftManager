// package that define gift structure and method
package modelgift

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	dataLoader "github.com/vickz86/GoFunctions/utiFile"
)

// define gift structure
type Gift struct {
	// Unique Index for identifying each gift
	Index int
	// Name of the gift
	Name string
	// Price of the gift
	Price int
	// Recommended age for the gift
	Age int
	// Targeted sex: 0 for men, 1 for women, 2 for both
	Sex int
	// URL of the gift
	Url string
}

//define slice of Gift
type SliceGift []Gift

//create a counter that get the last index when loading the file!

// Create a new gift based on input and return it
func CreateGift(name, url string,index, price, age, sex int) Gift {
	newGift := Gift{
		Index: index,
		Name:  name,
		Url:   url,
		Price: price,
		Age:   age,
		Sex:   sex,
	}

	return newGift
}


// return a slice of string from filePath , each element of the slice will define a gift struct
func SliceStringGift(filePath string) ([]string, error) {
	sliceGift, err := dataLoader.SliceStringFromFile(filePath)
	if err != nil {
		fmt.Println("ERROR loading gift file!")
	}

	//check there is the same amount of elememt in each string
	for nb, el := range sliceGift {
		//dont check the first element
		if nb == 0 {
			continue
		}

		if strings.Count(el, ";") != strings.Count(sliceGift[nb-1], ";") {
			return nil, errors.New("error , not the same amount of element in each string of gift")

		}

	}

	return sliceGift, nil
}


// GiftFromString creates a Gift struct from a semicolon-separated string
func GiftFromString(theStringGift string) (Gift, error) {
	// Split the input string by ";"
	allElements := strings.Split(theStringGift, ";")

	// Ensure we have the expected number of elements
	if len(allElements) != 6 {
		return Gift{}, errors.New("invalid format: expected 6 elements")
	}

	// Parse each element and handle any conversion errors
	theIndex, err := strconv.Atoi(allElements[0])
	if err != nil {
		return Gift{}, errors.New("error parsing index: not an integer")
	}

	name := allElements[1]

	thePrice, err := strconv.Atoi(allElements[2])
	if err != nil {
		return Gift{}, errors.New("error parsing price: not an integer")
	}

	theAge, err := strconv.Atoi(allElements[3])
	if err != nil {
		return Gift{}, errors.New("error parsing age: not an integer")
	}

	theSex, err := strconv.Atoi(allElements[4])
	if err != nil {
		return Gift{}, errors.New("error parsing sex: not an integer")
	}

	theUrl := allElements[5]

	// Create and return the Gift struct
	return Gift{
		Index: theIndex,
		Name:  name,
		Price: thePrice,
		Age:   theAge,
		Sex:   theSex,
		Url:   theUrl,
	}, nil
}


// create a sliceOfGift type from a slice of string AND return last index of the file
func CreateSliceGifts(sliceString []string) (SliceGift, int, error) {
	// Create an empty SliceGift and index
	var allGifts SliceGift
	var lastIndex int

	// Check if sliceString is empty
	if len(sliceString) == 0 {
		return nil, 0, errors.New("input slice is empty")
	}

	// Iterate over each string in the slice
	for nb, el := range sliceString {
		// Attempt to create a Gift from the string
		newGift, err := GiftFromString(el)
		if err != nil {
			// Handle the error, e.g., by printing a message and continuing
			fmt.Printf("Error creating gift from string at index %d: %v\n", nb, err)
			continue // Skip this entry if it has an error
		}

		// Add the new Gift to allGifts
		allGifts = append(allGifts, newGift)

		// Update lastIndex to the current gift's index if it's the last entry
		if nb == len(sliceString)-1 {
			lastIndex = newGift.Index
		}
	}

	// Return an error if no valid gifts were loaded
	if len(allGifts) == 0 {
		return nil, 0, errors.New("no valid gifts found in input")
	}

	return allGifts, lastIndex, nil
}

//create a SliceGift type from a pathFile
func LoadGift(filePath string) (SliceGift, int, error){

}
//DELETE INDEX IN THE SLICE
//ADD ELEMENT IN THE SLICE
