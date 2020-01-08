# Matrix is a Go matrix manipulation package

With this package, you can manipulate easiest matrix.
Every cell of the matrix is an interface{} so you can put what you want in it.

Here is an example :

```go
package main

import "matrix"

func main() {
    m := matrix.NewMatrix(5, 6, 0)  // fill the matrix with 0 integer
    m.Set(1, 1, "X")                // put a "X" in position 1/1
    m.SetCol(2, 9)                  // put 9 in every cells of the column 2
    m.GetRow(1)                     // return [0, "X", 9, 0, 0, 0]
}
```

You can read the file [matrix_test](matrix_test.go) for more examples. 

## Test

Launch `go test -cover`