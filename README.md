# Yet Another Slices Utility Package (for Go)

[![Go Report Card](https://goreportcard.com/badge/github.com/J4RV/yasup)](https://goreportcard.com/report/github.com/J4RV/yasup) [![GoDoc](https://godoc.org/github.com/J4RV/yasup?status.svg)](https://godoc.org/github.com/J4RV/yasup)

### What is it

An auto-generated library that contains helper functions for primitive type slices and a **code generator tool** that you can use for your custom types.

Current helper functions include:

- **Insert** an element at a given index
- **Delete** an element at a given index
- **Contains** an element
- **Index** of an element in the slice, or -1 if not found
- **LastIndex** of an element in the slice, or -1 if not found
- **Count** the amount of elements in the slice that match a given one
- **Push** an element to the end of the slice
- **Pop** an element from the end of the slice
- **FrontPush**: pushes an element to the start of the slice
- **FrontPop**: pops an element from the start of the slice
- **Replace**: replaces elements that match a given value with another given value N times
- **ReplaceAll**: replaces all elements that match a given value with another given value
- An **Equals** function for comparing slices
- **Shuffle** functions: One that uses math/rand and one that uses crypto/rand

### Examples of usage

YASUP is simple to use!

##### Using the library for slices of primitive types

See [examples/main.go](https://github.com/J4RV/yasup/blob/master/examples/main.go). It is runnable code that shows how to use many of the YASUP functions.

You can run the following command to see what that code prints:

```
go get github.com/J4RV/yasup
go run github.com/J4RV/yasup/examples
```

##### Using the generator tool for custom types

For an example about creating YASUP helper functions for non primitive or custom types, see [examples/genericStructExample.bat](https://github.com/J4RV/yasup/blob/master/examples/genericStructExample.bat).
