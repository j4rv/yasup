[![Go Report Card](https://goreportcard.com/badge/github.com/J4RV/yasup)](https://goreportcard.com/report/github.com/J4RV/yasup)

# Yet Another Slices Utility Package (for Go)

### What is it

An auto-generated library that contains helper functions for primitive types and a **code generator tool** that you can use for your custom types.

Current utility functions include:

- **Insert** an element at a given index
- **Delete** an element at a given index
- **Push** an element
- **Pop** an element
- An **Equals** function for comparing slices
- **Shuffle** functions: One that uses math/rand and another one that uses crypto/rand.

### Example code

Simple to use, see the **/example/main.go**. It is runnable code that shows how to use many of the YASUP functions.

Or you can just run the following command to see what that code does:

```
go get github.com/J4RV/yasup
go run github.com/J4RV/yasup/example
```

To create YASUP functions for non primitive or custom types, see **/example/genericStructExample.go**
