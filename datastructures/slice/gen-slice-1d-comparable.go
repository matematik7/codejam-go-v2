// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slice

func (slice SliceInt) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice SliceFloat64) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice SliceString) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice SliceByte) Less(i, j int) bool {
	return slice[i] < slice[j]
}
