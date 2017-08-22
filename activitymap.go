package activitymap

import "fmt"

/*Tuple tuple type for storing two values*/
type Tuple struct {
	a int
	b int
}

/*NewTuple constructor for tuple type*/
func NewTuple(x, y int) Tuple {
	return Tuple{
		a: x,
		b: y,
	}
}

/*ActivityMap type, has a single member acitivity which represents a hour of*/
/*time, each bit indicating if that minute was indeed active                */
type ActivityMap struct {
	acitivity uint64
}

/*New constructor for activity map*/
func New(mask uint64) *ActivityMap {
	return &ActivityMap{
		acitivity: mask,
	}
}

/*GetActivty read only accessor for activity map*/
func (am *ActivityMap) GetActivty() uint64 {
	return am.acitivity
}

/*AddMinute adds a entry to the activity map*/
func (am *ActivityMap) AddMinute(min int) {
	validateMinute(min)
	mask := uint64(1) << uint64(min-1)
	am.acitivity = am.acitivity | mask
}

/*WasOn returns whether or not a minute is active*/
func (am *ActivityMap) WasOn(min int) bool {
	validateMinute(min)
	mask := uint64(1) << uint64(min-1)
	if am.acitivity&mask > 0 {
		return true
	}
	return false
}

/*GetRanges returns all the contiguous ranges of 1's in the mask*/
func (am *ActivityMap) GetRanges() []Tuple {
	activeRanges := []Tuple{}
	prevIndex := 0
	currentIndex := 0
	isOn := false
	for i := 1; i <= 60; i++ {
		currentIndex = i
		if am.WasOn(i) {
			if !isOn {
				prevIndex = currentIndex
			}
			isOn = true
		} else {
			if isOn {
				isOn = false
				activeRanges = append(activeRanges, NewTuple(prevIndex, currentIndex-1))
			}
			prevIndex = currentIndex
		}
	}
	if isOn == true {
		activeRanges = append(activeRanges, NewTuple(prevIndex, currentIndex))
	}
	return activeRanges
}

func validateMinute(min int) {
	if min > 60 || min <= 0 {
		panic(fmt.Sprintf("minute out of range, was %v", min))
	}
}
