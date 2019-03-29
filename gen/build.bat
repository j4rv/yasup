::  bool
::  string
go run . -type bool       --test -singleVal false                -multipleVals="true, false, true, false, true, false, true, false, true, false, true, false"
go run . -type string     --test -singleVal "\"foobar\""         -multipleVals="\"0\", \"1\", \"2\", \"3\", \"4\", \"5\", \"6\", \"7\", \"lorem\", \"ipsum\""

::  int  int8  int16  int32  int64
go run . -type int        --test -singleVal -2147483648          -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type int8       --test -singleVal -128                 -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type int16      --test -singleVal -32768               -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type int32      --test -singleVal -2147483648          -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type int64      --test -singleVal -9223372036854775807 -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"

::  uint uint8 uint16 uint32 uint64 uintptr
go run . -type uint       --test -singleVal 4294967295           -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type uint8      --test -singleVal 255                  -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type uint16     --test -singleVal 65535                -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type uint32     --test -singleVal 4294967295           -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type uint64     --test -singleVal 18446744073709551615 -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"

::  byte // alias for uint8
::  rune // alias for int32
::       // represents a Unicode code point
go run . -type byte       --test -singleVal 255                  -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type rune       --test -singleVal 2147483647           -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"

::  float32 float64
go run . -type float32    --test -singleVal -2147483648          -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"
go run . -type float64    --test -singleVal -9223372036854775807 -multipleVals="0, 1, 2, 3, 4, 5, 6, 7, 8, 9"

::  complex64 complex128
go run . -type complex64  --test -singleVal -5.64+15.82i         -multipleVals="0+0i, 1+2i, -2+7.5i, 3+42.1i, 4-74.6i, -5+4i, 6-88i, 7-0i, 8+100i, 9+99i"
go run . -type complex128 --test -singleVal -52.6084+155.80287i  -multipleVals="0+0i, 1+2i, -2+7.5i, 3+42.1i, 4-74.6i, -5+4i, 6-88i, 7-0i, 8+100i, 9+99i"

cd ..
gofmt -s -w .
go test .