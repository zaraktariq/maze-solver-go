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

- `nil` is used in Go to represent the absence of a value for certain types.
- `null` is not used in Go, more common in the other languages we are familiar with
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
    - Maps are akin to Python's dictionaries.
3. #### Channels:
    - Facilitate communication and synchronization between goroutines.
    - Typed, allowing only values of a specific type.
    - Created using the make() function with the chan keyword followed by the value type.
---

Other notes.
1. #### Some file handling notes:
    - Go provides a comprehensive os package for file handling.
    - The `bufio.Scanner` type is used to read files line by line efficiently.
    - The `defer` statement is used to ensure that the file is closed after its use.

2. #### Looping Constructs:
    - Go has a for loop, similar to Python, but it doesn't have a foreach construct.
    - Go's range keyword is used to iterate over slices and maps.

3. #### Conditional Statements:
    - Go uses if-else constructs similar to Python.
    - Boolean operators like && (and) and || (or) are used for conditions.
