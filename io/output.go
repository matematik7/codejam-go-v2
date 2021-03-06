package io

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
	"unicode"

	"gonum.org/v1/plot/plotter"
)

type Output struct {
	finalOutput io.Writer
	buffer      *bytes.Buffer

	output io.Writer

	caseN int
	input *Input

	periodicPrint       chan struct{}
	previousPeriodicInt int
	periodicCount       int
	prevPeriodicCount   int

	points plotter.XYs

	startTime time.Time
	timers    map[string]*Timer
}

func newOutput(w io.Writer, buffered bool) *Output {
	output := &Output{
		periodicPrint: make(chan struct{}, 10),
		timers:        make(map[string]*Timer),
	}
	if buffered {
		output.buffer = &bytes.Buffer{}
		output.output = output.buffer
		output.finalOutput = w
	} else {
		output.output = w
	}
	return output
}

func (o *Output) resetPeriodic() {
	for {
		select {
		case <-o.periodicPrint:
		default:
			return
		}
	}
}

func (o *Output) triggerPeriodic() {
	o.resetPeriodic()
	o.periodicPrint <- struct{}{}
}

func (o *Output) Print(a ...interface{}) {
	fmt.Fprint(o.output, a...)
}

func (o *Output) PrintFloat64(a float64) {
	fmt.Fprintf(o.output, "%f", a)
}

func (o *Output) Println(a ...interface{}) {
	fmt.Fprintln(o.output, a...)
}

func (o *Output) Printf(format string, a ...interface{}) {
	fmt.Fprintf(o.output, format, a...)
}

func (o *Output) Fatal(a ...interface{}) {
	o.DebugCase()
	log.Fatalln(a...)
}

func (o *Output) Fatalf(format string, a ...interface{}) {
	o.DebugCase()
	log.Fatalf(format, a...)
}

func (o *Output) Debug(a ...interface{}) {
	o.DebugCase()
	log.Println(a...)
}

func (o *Output) Debugf(format string, a ...interface{}) {
	o.DebugCase()
	log.Printf(format, a...)
}

func (o *Output) DebugCase() {
	outStr := ""
	if o.buffer != nil {
		outStr = string(o.buffer.Bytes())
	}
	log.Printf("Case #%d, input: %v, output: %q\n", o.caseN, o.input.currentCase(), outStr)
}

func (o *Output) Periodic(a ...interface{}) {
	select {
	case <-o.periodicPrint:
		log.Println(a...)
	default:
	}
}

func (o *Output) PeriodicInt(a int) {
	select {
	case <-o.periodicPrint:
		log.Println(a, "Rate =", a-o.previousPeriodicInt)
		o.previousPeriodicInt = a
	default:
	}
}

func (o *Output) PeriodicCount() {
	o.periodicCount++
	select {
	case <-o.periodicPrint:
		c := o.periodicCount
		log.Println(c, "Rate =", c-o.prevPeriodicCount)
		o.prevPeriodicCount = c
	default:
	}
}

func (o *Output) Periodicf(format string, a ...interface{}) {
	select {
	case <-o.periodicPrint:
		log.Printf(format, a...)
	default:
	}
}

func (o *Output) Point(x, y float64) {
	o.points = append(o.points, struct{ X, Y float64 }{
		X: x,
		Y: y,
	})
}

func (o *Output) PointInt(x, y int) {
	o.Point(float64(x), float64(y))
}

func (o *Output) init(input *Input, caseN int) {
	o.input = input
	o.caseN = caseN
}

func (o *Output) flush() {
	if o.buffer != nil {
		if o.buffer.Len() <= 0 {
			o.Fatal("No output")
		}
		fmt.Fprintf(o.finalOutput, "Case #%d:", o.caseN)
		if !unicode.In(rune(o.buffer.Bytes()[0]), unicode.White_Space) {
			o.finalOutput.Write([]byte{' '})
		}
		o.finalOutput.Write(o.buffer.Bytes())
		if o.buffer.Bytes()[o.buffer.Len()-1] != '\n' {
			o.buffer.Write([]byte{'\n'})
		}
		o.buffer.Reset()
	}
}

func (o *Output) assertOutput(fatal []bool, a ...interface{}) {
	output := o.Debug
	if len(fatal) > 0 && fatal[0] {
		output = o.Fatal
	}
	output(a...)
}

func (o *Output) assertOutputf(fatal []bool, format string, a ...interface{}) {
	outputf := o.Debugf
	if len(fatal) > 0 && fatal[0] {
		outputf = o.Fatalf
	}
	outputf(format, a...)
}

func (o *Output) AssertByteCount(a byte, count int, fatal ...bool) {
	for _, b := range o.buffer.Bytes() {
		if a == b {
			count--
		}
	}
	if count != 0 {
		o.assertOutputf(fatal, "AssertByteCount: byte: %q difference: %d", string(a), count*-1)
	}
}

func (o *Output) AssertCount(count int, fatal ...bool) {
	if o.buffer.Len() != count {
		o.assertOutput(fatal, "AssertCount: ", count)
	}
}

func (o *Output) AssertEqual(data string, fatal ...bool) {
	if strings.TrimSpace(string(o.buffer.Bytes())) != strings.TrimSpace(data) {
		o.assertOutputf(fatal, "Output should be: %q", data)
	}
}

func (o *Output) AssertIntEqual(a, b int, fatal ...bool) {
	if a != b {
		o.assertOutputf(fatal, "Ints %d and %d not equal", a, b)
	}
}

func (o *Output) AssertTrue(a bool, fatal ...bool) {
	if !a {
		o.assertOutput(fatal, "Not true")
	}
}

func (o *Output) AssertNoError(e error, fatal ...bool) {
	if e != nil {
		o.assertOutputf(fatal, "Error: %v", e)
	}
}

func (o *Output) TimerStart(key string) {
	if o.timers[key] == nil {
		o.timers[key] = &Timer{}
	}
	o.timers[key].Start()
}

func (o *Output) TimerStop(key string) {
	if o.timers[key] == nil {
		o.timers[key] = &Timer{}
	}
	o.timers[key].Stop(o.startTime)
}
