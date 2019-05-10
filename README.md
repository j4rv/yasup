# Yet Another Slices Utility Package (for Go)

[![Go Report Card](https://goreportcard.com/badge/github.com/J4RV/yasup)](https://goreportcard.com/report/github.com/J4RV/yasup)

### What is it

An auto-generated library that contains helper functions for primitive type slices and a **code generator tool** that you can use for your custom types.

Current helper functions include:

- **Contains** an element
- **Insert** an element at a given index
- **Delete** an element at a given index
- **Push** an element to the end of the slice
- **Pop** an element from the end of the slice
- **FrontPush**: pushes an element to the start of the slice
- **FrontPop**: pops an element from the start of the slice
- An **Equals** function for comparing slices
- **Shuffle** functions: One that uses math/rand and one that uses crypto/rand

### Example code

Simple to use, see the **/example/main.go**. It is runnable code that shows how to use many of the YASUP functions.

Or you can just run the following command to see what that code does:

```
go get github.com/J4RV/yasup
go run github.com/J4RV/yasup/example
```

To create YASUP functions for non primitive or custom types, see **/example/genericStructExample.go**
