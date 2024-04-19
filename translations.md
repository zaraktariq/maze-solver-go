#### The code for the struct in Go:
```go
type Point struct {
    x int
    y int
}
```
#### Is equivalent to this code in Python:
```python
class Point:
def __init__(self, x, y):
    self.x = x
    self.y = y
```
------
- := is used for short variable declaration and initialization.
- = is used for variable assignment.
------
#### This is how one prints in Go
```go
fmt.Print(cell, " ")
```

#### In comparison to our familiar python
```go
print(cell, end=" ")
```
----

In Go, the `make()` function is used to create slices, maps, and channels, which are all reference types.
1. #### Slices
    - Dynamic data structures for managing sequences of elements.
    - Can grow or shrink in size dynamically.
    - Created using the make() function or by slicing existing arrays or other slices.
2. #### Maps:
    - Unordered collections of key-value pairs.
    - Keys must be unique and can be of any comparable type.
    - Created using the make() function or a composite literal (map[KeyType]ValueType{}).
3. #### Channels:
    - Facilitate communication and synchronization between goroutines.
    - Typed, allowing only values of a specific type.
    - Created using the make() function with the chan keyword followed by the value type.
---