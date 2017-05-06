// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slice

// NewSliceSliceInt creates slice length n
func NewSliceSliceInt(n, m int) SliceSliceInt {
	newSlice := make([]SliceInt, n)
	for i := range newSlice {
		newSlice[i] = NewSliceInt(m)
	}
	return newSlice
}

// Copy makes a new independent copy of slice
func (slice SliceSliceInt) Copy() SliceSliceInt {
	newSlice := make([]SliceInt, len(slice))
	for i := range newSlice {
		newSlice[i] = slice[i].Copy()
	}
	return newSlice
}

// String is for print
func (slice SliceSliceInt) String() string {
	return slice.Print("\n")
}

// SpiralIterator returns []Coordinate in spiral order
func (slice SliceSliceInt) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(slice)*len(slice[0]))
	return spiralTopRight(data, 0, 0, len(slice)-1, len(slice[0])-1)
}

// NewSliceSliceFloat64 creates slice length n
func NewSliceSliceFloat64(n, m int) SliceSliceFloat64 {
	newSlice := make([]SliceFloat64, n)
	for i := range newSlice {
		newSlice[i] = NewSliceFloat64(m)
	}
	return newSlice
}

// Copy makes a new independent copy of slice
func (slice SliceSliceFloat64) Copy() SliceSliceFloat64 {
	newSlice := make([]SliceFloat64, len(slice))
	for i := range newSlice {
		newSlice[i] = slice[i].Copy()
	}
	return newSlice
}

// String is for print
func (slice SliceSliceFloat64) String() string {
	return slice.Print("\n")
}

// SpiralIterator returns []Coordinate in spiral order
func (slice SliceSliceFloat64) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(slice)*len(slice[0]))
	return spiralTopRight(data, 0, 0, len(slice)-1, len(slice[0])-1)
}

// NewSliceSliceString creates slice length n
func NewSliceSliceString(n, m int) SliceSliceString {
	newSlice := make([]SliceString, n)
	for i := range newSlice {
		newSlice[i] = NewSliceString(m)
	}
	return newSlice
}

// Copy makes a new independent copy of slice
func (slice SliceSliceString) Copy() SliceSliceString {
	newSlice := make([]SliceString, len(slice))
	for i := range newSlice {
		newSlice[i] = slice[i].Copy()
	}
	return newSlice
}

// String is for print
func (slice SliceSliceString) String() string {
	return slice.Print("\n")
}

// SpiralIterator returns []Coordinate in spiral order
func (slice SliceSliceString) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(slice)*len(slice[0]))
	return spiralTopRight(data, 0, 0, len(slice)-1, len(slice[0])-1)
}

// NewSliceSliceByte creates slice length n
func NewSliceSliceByte(n, m int) SliceSliceByte {
	newSlice := make([]SliceByte, n)
	for i := range newSlice {
		newSlice[i] = NewSliceByte(m)
	}
	return newSlice
}

// Copy makes a new independent copy of slice
func (slice SliceSliceByte) Copy() SliceSliceByte {
	newSlice := make([]SliceByte, len(slice))
	for i := range newSlice {
		newSlice[i] = slice[i].Copy()
	}
	return newSlice
}

// String is for print
func (slice SliceSliceByte) String() string {
	return slice.Print("\n")
}

// SpiralIterator returns []Coordinate in spiral order
func (slice SliceSliceByte) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(slice)*len(slice[0]))
	return spiralTopRight(data, 0, 0, len(slice)-1, len(slice[0])-1)
}

// NewSliceSliceBool creates slice length n
func NewSliceSliceBool(n, m int) SliceSliceBool {
	newSlice := make([]SliceBool, n)
	for i := range newSlice {
		newSlice[i] = NewSliceBool(m)
	}
	return newSlice
}

// Copy makes a new independent copy of slice
func (slice SliceSliceBool) Copy() SliceSliceBool {
	newSlice := make([]SliceBool, len(slice))
	for i := range newSlice {
		newSlice[i] = slice[i].Copy()
	}
	return newSlice
}

// String is for print
func (slice SliceSliceBool) String() string {
	return slice.Print("\n")
}

// SpiralIterator returns []Coordinate in spiral order
func (slice SliceSliceBool) SpiralIterator() []Coordinate {
	data := make([]Coordinate, len(slice)*len(slice[0]))
	return spiralTopRight(data, 0, 0, len(slice)-1, len(slice[0])-1)
}
