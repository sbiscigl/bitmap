package activitymap

import "testing"

func TestShouldCreateNone(t *testing.T) {
	a := New(0)
	if r := a.GetActivty(); r != 0 {
		t.Errorf("should have created none got %v", r)
	}
}

func TestShouldSome(t *testing.T) {
	a := New(1)
	if r := a.GetActivty(); r != 1 {
		t.Errorf("should have created 1 got %v", r)
	}
}

func TestShouldAddMinuteFirstMinute(t *testing.T) {
	a := New(0)
	a.AddMinute(1)
	if r := a.GetActivty(); r != 1 {
		t.Errorf("should have added 1 got %v", r)
	}
}

func TestShouldAddMinuteLastMinute(t *testing.T) {
	a := New(0)
	a.AddMinute(60)
	if r := a.GetActivty() != 576460752303423488; r {
		t.Errorf("should have added 60 minutes (576460752303423488 decimal) got decimal %v", r)
	}
}

func TestShouldAddMiddleMinute(t *testing.T) {
	a := New(0)
	a.AddMinute(30)
	if r := a.GetActivty() != 536870912; r {
		t.Errorf("should have added 30 minutes (536870912 decimal) got decimal %v", r)
	}
}

func TestShouldPanicWhenMinuteBad(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	a := New(0)
	a.AddMinute(61)
}

func TestShouldVerifyFirstMinute(t *testing.T) {
	a := New(1)
	if r := a.WasOn(1) != true; r {
		t.Errorf("should have found it instead: %v", r)
	}
}

func TestShouldVerifyLastMinute(t *testing.T) {
	a := New(0)
	a.AddMinute(60)
	if r := a.WasOn(60) != true; r {
		t.Errorf("should have found it instead: %v", r)
	}
}

func TestShouldVerifyMiddleMinute(t *testing.T) {
	a := New(0)
	a.AddMinute(30)
	if r := a.WasOn(30) != true; r {
		t.Errorf("should have found it instead: %v", r)
	}
}

func TestShouldGetCorrectRanges(t *testing.T) {
	//11100011100
	a := New(1820)
	ranges := a.GetRanges()
	expected := []Tuple{NewTuple(3, 5), NewTuple(9, 11)}
	if len(ranges) != len(expected) {
		t.Errorf("length should have been %v, was %v", len(expected), len(ranges))
	}
	for i := range ranges {
		if ranges[i] != expected[i] {
			t.Errorf("element should have been %v, was %v", expected[i], ranges[i])
		}
	}
}

func TestShouldGetCorrectRangesFirstBit(t *testing.T) {
	//111
	a := New(7)
	ranges := a.GetRanges()
	expected := []Tuple{NewTuple(1, 3)}
	if len(ranges) != len(expected) {
		t.Errorf("length should have been %v, was %v", len(expected), len(ranges))
	}
	for i := range ranges {
		if ranges[i] != expected[i] {
			t.Errorf("element should have been %v, was %v", expected[i], ranges[i])
		}
	}
}

func TestShouldGetCorrectRangesJustFirstBit(t *testing.T) {
	//1
	a := New(1)
	ranges := a.GetRanges()
	expected := []Tuple{NewTuple(1, 1)}
	if len(ranges) != len(expected) {
		t.Errorf("length should have been %v, was %v", len(expected), len(ranges))
	}
	for i := range ranges {
		if ranges[i] != expected[i] {
			t.Errorf("element should have been %v, was %v", expected[i], ranges[i])
		}
	}
}

func TestShouldGetCorrectRangesLastBit(t *testing.T) {
	a := New(0)
	a.AddMinute(60)
	a.AddMinute(59)
	a.AddMinute(58)
	ranges := a.GetRanges()
	expected := []Tuple{NewTuple(58, 60)}
	if len(ranges) != len(expected) {
		t.Errorf("length should have been %v, was %v", len(expected), len(ranges))
	}
	for i := range ranges {
		if ranges[i] != expected[i] {
			t.Errorf("element should have been %v, was %v", expected[i], ranges[i])
		}
	}
}

func TestShouldGetCorrectRangesJustLastBit(t *testing.T) {
	a := New(0)
	a.AddMinute(60)
	ranges := a.GetRanges()
	expected := []Tuple{NewTuple(60, 60)}
	if len(ranges) != len(expected) {
		t.Errorf("length should have been %v, was %v", len(expected), len(ranges))
	}
	for i := range ranges {
		if ranges[i] != expected[i] {
			t.Errorf("element should have been %v, was %v", expected[i], ranges[i])
		}
	}
}
