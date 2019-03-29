## 'slices' makes working with basic type slices easier and more clear

Want to shuffle an int array/slice or compare two float64 arrays/slices? Just do:

`slices.IntShuffle(&myInts)`  
`eq := slices.Float64Equals(myFloats, yourFloats)`

Supports all basic types, and you can add your own types by using the script in the 'factory' directory.  
You will just need to pass some values to the script and it will generate your slices methods by executing a template.

For example, executing the following command in the factory folder will generate the `int` .go files that exist in the root folder of this repository:  
`go run . -type int --test -singleVal -2147483648 -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"`  
The singleVal and multipleVals flag values will be used to auto-generate the tests.
