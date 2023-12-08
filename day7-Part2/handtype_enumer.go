// Code generated by "enumer -type=HandType"; DO NOT EDIT.

package main

import (
	"fmt"
)

const _HandTypeName = "HighCardOnePairTwoPairThreeOfAKindFullHouseFourOfAKindFiveOfAKind"

var _HandTypeIndex = [...]uint8{0, 8, 15, 22, 34, 43, 54, 65}

func (i HandType) String() string {
	if i < 0 || i >= HandType(len(_HandTypeIndex)-1) {
		return fmt.Sprintf("HandType(%d)", i)
	}
	return _HandTypeName[_HandTypeIndex[i]:_HandTypeIndex[i+1]]
}

var _HandTypeValues = []HandType{0, 1, 2, 3, 4, 5, 6}

var _HandTypeNameToValueMap = map[string]HandType{
	_HandTypeName[0:8]:   0,
	_HandTypeName[8:15]:  1,
	_HandTypeName[15:22]: 2,
	_HandTypeName[22:34]: 3,
	_HandTypeName[34:43]: 4,
	_HandTypeName[43:54]: 5,
	_HandTypeName[54:65]: 6,
}

// HandTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func HandTypeString(s string) (HandType, error) {
	if val, ok := _HandTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to HandType values", s)
}

// HandTypeValues returns all values of the enum
func HandTypeValues() []HandType {
	return _HandTypeValues
}

// IsAHandType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i HandType) IsAHandType() bool {
	for _, v := range _HandTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
