# go-where
Provides methods for extracting subsets which satisfy some condition.

## Usage

TODO

```go
package main

import(
  "github.com/shawnohare/go-where"
)

type Circle struct {
  Radius float64
}

// Circles represents a collection of circle objects.
type Circles struct []Circle

// implement the where.Interface interface
func (cs Circles) Len() int {
  return len(cs)
} 
```
