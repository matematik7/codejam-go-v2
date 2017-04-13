// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package slice

import (
	"container/heap"
	"fmt"
	"sort"
)

// SliceInt type
type SliceInt []int

// Set sets all element to c
func (slice SliceInt) Set(c int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceInt) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceInt) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceInt) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceInt) Get(i int) int {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceInt) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceInt) Push(c interface{}) {
	*slice = append(*slice, c.(int))
}

// Pop element
func (slice *SliceInt) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceInt) Append(values ...int) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceInt) Prepend(values ...int) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceInt) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceInt struct for min heap
type MinHeapSliceInt struct {
	Slice *SliceInt
}

// MinHeap returns struct with min heap functionality based on SliceInt
func (slice *SliceInt) MinHeap() MinHeapSliceInt {
	hp := MinHeapSliceInt{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceInt) Min() int {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceInt) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceInt) Pop() int {
	return heap.Pop(hp.Slice).(int)
}

// Push v to heap
func (hp MinHeapSliceInt) Push(v int) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceInt) Remove(i int) int {
	return heap.Remove(hp.Slice, i).(int)
}

// MaxHeapSliceInt struct for max heap
type MaxHeapSliceInt struct {
	Slice *SliceInt
}

// MaxHeap returns struct with max heap functionality based on SliceInt
func (slice *SliceInt) MaxHeap() MaxHeapSliceInt {
	hp := MaxHeapSliceInt{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceInt) Max() int {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceInt) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceInt) Pop() int {
	return heap.Pop(HeapReverse(hp.Slice)).(int)
}

// Push v to heap
func (hp MaxHeapSliceInt) Push(v int) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceInt) Remove(i int) int {
	return heap.Remove(HeapReverse(hp.Slice), i).(int)
}

// SliceSliceInt type
type SliceSliceInt []SliceInt

// Set sets all element to c
func (slice SliceSliceInt) Set(c SliceInt) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceSliceInt) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceSliceInt) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceSliceInt) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceSliceInt) Get(i int) SliceInt {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceSliceInt) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceSliceInt) Push(c interface{}) {
	*slice = append(*slice, c.(SliceInt))
}

// Pop element
func (slice *SliceSliceInt) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceSliceInt) Append(values ...SliceInt) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceSliceInt) Prepend(values ...SliceInt) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceSliceInt) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceSliceInt struct for min heap
type MinHeapSliceSliceInt struct {
	Slice *SliceSliceInt
}

// MinHeap returns struct with min heap functionality based on SliceSliceInt
func (slice *SliceSliceInt) MinHeap() MinHeapSliceSliceInt {
	hp := MinHeapSliceSliceInt{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceSliceInt) Min() SliceInt {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceSliceInt) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceSliceInt) Pop() SliceInt {
	return heap.Pop(hp.Slice).(SliceInt)
}

// Push v to heap
func (hp MinHeapSliceSliceInt) Push(v SliceInt) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceSliceInt) Remove(i int) SliceInt {
	return heap.Remove(hp.Slice, i).(SliceInt)
}

// MaxHeapSliceSliceInt struct for max heap
type MaxHeapSliceSliceInt struct {
	Slice *SliceSliceInt
}

// MaxHeap returns struct with max heap functionality based on SliceSliceInt
func (slice *SliceSliceInt) MaxHeap() MaxHeapSliceSliceInt {
	hp := MaxHeapSliceSliceInt{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceSliceInt) Max() SliceInt {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceSliceInt) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceSliceInt) Pop() SliceInt {
	return heap.Pop(HeapReverse(hp.Slice)).(SliceInt)
}

// Push v to heap
func (hp MaxHeapSliceSliceInt) Push(v SliceInt) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceSliceInt) Remove(i int) SliceInt {
	return heap.Remove(HeapReverse(hp.Slice), i).(SliceInt)
}

// SliceFloat64 type
type SliceFloat64 []float64

// Set sets all element to c
func (slice SliceFloat64) Set(c float64) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceFloat64) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceFloat64) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceFloat64) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceFloat64) Get(i int) float64 {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceFloat64) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceFloat64) Push(c interface{}) {
	*slice = append(*slice, c.(float64))
}

// Pop element
func (slice *SliceFloat64) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceFloat64) Append(values ...float64) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceFloat64) Prepend(values ...float64) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceFloat64) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceFloat64 struct for min heap
type MinHeapSliceFloat64 struct {
	Slice *SliceFloat64
}

// MinHeap returns struct with min heap functionality based on SliceFloat64
func (slice *SliceFloat64) MinHeap() MinHeapSliceFloat64 {
	hp := MinHeapSliceFloat64{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceFloat64) Min() float64 {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceFloat64) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceFloat64) Pop() float64 {
	return heap.Pop(hp.Slice).(float64)
}

// Push v to heap
func (hp MinHeapSliceFloat64) Push(v float64) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceFloat64) Remove(i int) float64 {
	return heap.Remove(hp.Slice, i).(float64)
}

// MaxHeapSliceFloat64 struct for max heap
type MaxHeapSliceFloat64 struct {
	Slice *SliceFloat64
}

// MaxHeap returns struct with max heap functionality based on SliceFloat64
func (slice *SliceFloat64) MaxHeap() MaxHeapSliceFloat64 {
	hp := MaxHeapSliceFloat64{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceFloat64) Max() float64 {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceFloat64) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceFloat64) Pop() float64 {
	return heap.Pop(HeapReverse(hp.Slice)).(float64)
}

// Push v to heap
func (hp MaxHeapSliceFloat64) Push(v float64) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceFloat64) Remove(i int) float64 {
	return heap.Remove(HeapReverse(hp.Slice), i).(float64)
}

// SliceSliceFloat64 type
type SliceSliceFloat64 []SliceFloat64

// Set sets all element to c
func (slice SliceSliceFloat64) Set(c SliceFloat64) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceSliceFloat64) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceSliceFloat64) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceSliceFloat64) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceSliceFloat64) Get(i int) SliceFloat64 {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceSliceFloat64) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceSliceFloat64) Push(c interface{}) {
	*slice = append(*slice, c.(SliceFloat64))
}

// Pop element
func (slice *SliceSliceFloat64) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceSliceFloat64) Append(values ...SliceFloat64) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceSliceFloat64) Prepend(values ...SliceFloat64) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceSliceFloat64) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceSliceFloat64 struct for min heap
type MinHeapSliceSliceFloat64 struct {
	Slice *SliceSliceFloat64
}

// MinHeap returns struct with min heap functionality based on SliceSliceFloat64
func (slice *SliceSliceFloat64) MinHeap() MinHeapSliceSliceFloat64 {
	hp := MinHeapSliceSliceFloat64{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceSliceFloat64) Min() SliceFloat64 {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceSliceFloat64) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceSliceFloat64) Pop() SliceFloat64 {
	return heap.Pop(hp.Slice).(SliceFloat64)
}

// Push v to heap
func (hp MinHeapSliceSliceFloat64) Push(v SliceFloat64) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceSliceFloat64) Remove(i int) SliceFloat64 {
	return heap.Remove(hp.Slice, i).(SliceFloat64)
}

// MaxHeapSliceSliceFloat64 struct for max heap
type MaxHeapSliceSliceFloat64 struct {
	Slice *SliceSliceFloat64
}

// MaxHeap returns struct with max heap functionality based on SliceSliceFloat64
func (slice *SliceSliceFloat64) MaxHeap() MaxHeapSliceSliceFloat64 {
	hp := MaxHeapSliceSliceFloat64{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceSliceFloat64) Max() SliceFloat64 {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceSliceFloat64) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceSliceFloat64) Pop() SliceFloat64 {
	return heap.Pop(HeapReverse(hp.Slice)).(SliceFloat64)
}

// Push v to heap
func (hp MaxHeapSliceSliceFloat64) Push(v SliceFloat64) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceSliceFloat64) Remove(i int) SliceFloat64 {
	return heap.Remove(HeapReverse(hp.Slice), i).(SliceFloat64)
}

// SliceString type
type SliceString []string

// Set sets all element to c
func (slice SliceString) Set(c string) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceString) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceString) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceString) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceString) Get(i int) string {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceString) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceString) Push(c interface{}) {
	*slice = append(*slice, c.(string))
}

// Pop element
func (slice *SliceString) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceString) Append(values ...string) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceString) Prepend(values ...string) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceString) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceString struct for min heap
type MinHeapSliceString struct {
	Slice *SliceString
}

// MinHeap returns struct with min heap functionality based on SliceString
func (slice *SliceString) MinHeap() MinHeapSliceString {
	hp := MinHeapSliceString{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceString) Min() string {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceString) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceString) Pop() string {
	return heap.Pop(hp.Slice).(string)
}

// Push v to heap
func (hp MinHeapSliceString) Push(v string) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceString) Remove(i int) string {
	return heap.Remove(hp.Slice, i).(string)
}

// MaxHeapSliceString struct for max heap
type MaxHeapSliceString struct {
	Slice *SliceString
}

// MaxHeap returns struct with max heap functionality based on SliceString
func (slice *SliceString) MaxHeap() MaxHeapSliceString {
	hp := MaxHeapSliceString{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceString) Max() string {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceString) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceString) Pop() string {
	return heap.Pop(HeapReverse(hp.Slice)).(string)
}

// Push v to heap
func (hp MaxHeapSliceString) Push(v string) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceString) Remove(i int) string {
	return heap.Remove(HeapReverse(hp.Slice), i).(string)
}

// SliceSliceString type
type SliceSliceString []SliceString

// Set sets all element to c
func (slice SliceSliceString) Set(c SliceString) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceSliceString) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceSliceString) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceSliceString) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceSliceString) Get(i int) SliceString {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceSliceString) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceSliceString) Push(c interface{}) {
	*slice = append(*slice, c.(SliceString))
}

// Pop element
func (slice *SliceSliceString) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceSliceString) Append(values ...SliceString) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceSliceString) Prepend(values ...SliceString) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceSliceString) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceSliceString struct for min heap
type MinHeapSliceSliceString struct {
	Slice *SliceSliceString
}

// MinHeap returns struct with min heap functionality based on SliceSliceString
func (slice *SliceSliceString) MinHeap() MinHeapSliceSliceString {
	hp := MinHeapSliceSliceString{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceSliceString) Min() SliceString {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceSliceString) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceSliceString) Pop() SliceString {
	return heap.Pop(hp.Slice).(SliceString)
}

// Push v to heap
func (hp MinHeapSliceSliceString) Push(v SliceString) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceSliceString) Remove(i int) SliceString {
	return heap.Remove(hp.Slice, i).(SliceString)
}

// MaxHeapSliceSliceString struct for max heap
type MaxHeapSliceSliceString struct {
	Slice *SliceSliceString
}

// MaxHeap returns struct with max heap functionality based on SliceSliceString
func (slice *SliceSliceString) MaxHeap() MaxHeapSliceSliceString {
	hp := MaxHeapSliceSliceString{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceSliceString) Max() SliceString {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceSliceString) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceSliceString) Pop() SliceString {
	return heap.Pop(HeapReverse(hp.Slice)).(SliceString)
}

// Push v to heap
func (hp MaxHeapSliceSliceString) Push(v SliceString) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceSliceString) Remove(i int) SliceString {
	return heap.Remove(HeapReverse(hp.Slice), i).(SliceString)
}

// SliceByte type
type SliceByte []byte

// Set sets all element to c
func (slice SliceByte) Set(c byte) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceByte) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceByte) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceByte) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceByte) Get(i int) byte {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceByte) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceByte) Push(c interface{}) {
	*slice = append(*slice, c.(byte))
}

// Pop element
func (slice *SliceByte) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceByte) Append(values ...byte) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceByte) Prepend(values ...byte) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceByte) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceByte struct for min heap
type MinHeapSliceByte struct {
	Slice *SliceByte
}

// MinHeap returns struct with min heap functionality based on SliceByte
func (slice *SliceByte) MinHeap() MinHeapSliceByte {
	hp := MinHeapSliceByte{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceByte) Min() byte {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceByte) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceByte) Pop() byte {
	return heap.Pop(hp.Slice).(byte)
}

// Push v to heap
func (hp MinHeapSliceByte) Push(v byte) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceByte) Remove(i int) byte {
	return heap.Remove(hp.Slice, i).(byte)
}

// MaxHeapSliceByte struct for max heap
type MaxHeapSliceByte struct {
	Slice *SliceByte
}

// MaxHeap returns struct with max heap functionality based on SliceByte
func (slice *SliceByte) MaxHeap() MaxHeapSliceByte {
	hp := MaxHeapSliceByte{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceByte) Max() byte {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceByte) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceByte) Pop() byte {
	return heap.Pop(HeapReverse(hp.Slice)).(byte)
}

// Push v to heap
func (hp MaxHeapSliceByte) Push(v byte) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceByte) Remove(i int) byte {
	return heap.Remove(HeapReverse(hp.Slice), i).(byte)
}

// SliceSliceByte type
type SliceSliceByte []SliceByte

// Set sets all element to c
func (slice SliceSliceByte) Set(c SliceByte) {
	for i := 0; i < len(slice); i++ {
		slice[i] = c
	}
}

// SortAsc sort ascending
func (slice SliceSliceByte) SortAsc() {
	sort.Sort(slice)
}

// SortDesc sort descending
func (slice SliceSliceByte) SortDesc() {
	sort.Sort(sort.Reverse(slice))
}

// Len length
func (slice SliceSliceByte) Len() int {
	return len(slice)
}

// Get i-th element
func (slice SliceSliceByte) Get(i int) SliceByte {
	if i < 0 {
		i += len(slice)
	}

	return slice[i]
}

// Swap two elements
func (slice SliceSliceByte) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Push element
func (slice *SliceSliceByte) Push(c interface{}) {
	*slice = append(*slice, c.(SliceByte))
}

// Pop element
func (slice *SliceSliceByte) Pop() interface{} {
	element := slice.Get(-1)
	*slice = (*slice)[:len(*slice)-1]
	return element
}

// Append values
func (slice *SliceSliceByte) Append(values ...SliceByte) {
	*slice = append(*slice, values...)
}

// Prepend values
func (slice *SliceSliceByte) Prepend(values ...SliceByte) {
	*slice = append(values, (*slice)...)
}

// Print prints using separator
func (slice SliceSliceByte) Print(sep string) string {
	output := ""
	for _, c := range slice {
		if output != "" {
			output += sep
		}
		output += fmt.Sprintf("%v", c)
	}
	return output
}

// MinHeapSliceSliceByte struct for min heap
type MinHeapSliceSliceByte struct {
	Slice *SliceSliceByte
}

// MinHeap returns struct with min heap functionality based on SliceSliceByte
func (slice *SliceSliceByte) MinHeap() MinHeapSliceSliceByte {
	hp := MinHeapSliceSliceByte{
		Slice: slice,
	}
	heap.Init(hp.Slice)
	return hp
}

// Min returns min element
func (hp MinHeapSliceSliceByte) Min() SliceByte {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MinHeapSliceSliceByte) Fix(i int) {
	heap.Fix(hp.Slice, i)
}

// Pop removes the minimum element
func (hp MinHeapSliceSliceByte) Pop() SliceByte {
	return heap.Pop(hp.Slice).(SliceByte)
}

// Push v to heap
func (hp MinHeapSliceSliceByte) Push(v SliceByte) {
	heap.Push(hp.Slice, v)
}

// Remove i-th element
func (hp MinHeapSliceSliceByte) Remove(i int) SliceByte {
	return heap.Remove(hp.Slice, i).(SliceByte)
}

// MaxHeapSliceSliceByte struct for max heap
type MaxHeapSliceSliceByte struct {
	Slice *SliceSliceByte
}

// MaxHeap returns struct with max heap functionality based on SliceSliceByte
func (slice *SliceSliceByte) MaxHeap() MaxHeapSliceSliceByte {
	hp := MaxHeapSliceSliceByte{
		Slice: slice,
	}
	heap.Init(HeapReverse(hp.Slice))
	return hp
}

// Max returns max element
func (hp MaxHeapSliceSliceByte) Max() SliceByte {
	return (*hp.Slice)[0]
}

// Fix re-establishes heap ordering after i has changed value
func (hp MaxHeapSliceSliceByte) Fix(i int) {
	heap.Fix(HeapReverse(hp.Slice), i)
}

// Pop removes the minimum element
func (hp MaxHeapSliceSliceByte) Pop() SliceByte {
	return heap.Pop(HeapReverse(hp.Slice)).(SliceByte)
}

// Push v to heap
func (hp MaxHeapSliceSliceByte) Push(v SliceByte) {
	heap.Push(HeapReverse(hp.Slice), v)
}

// Remove i-th element
func (hp MaxHeapSliceSliceByte) Remove(i int) SliceByte {
	return heap.Remove(HeapReverse(hp.Slice), i).(SliceByte)
}
