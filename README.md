# stack

    import "go.ajitem.com/stack"

A stack is a fundamental data structure which is extensively used. This package
provides a basic implementation of a stack in Go. A stack works using Last In
First Out (LIFO) method.

## Usage

```go
var IsEmpty = errors.New("stack is empty")
```

#### type Stack

```go
type Stack struct {
}
```


#### func  New

```go
func New() *Stack
```
Create a new stack

#### func (*Stack) Peek

```go
func (s *Stack) Peek() (int, error)
```
Return the item on the top of the stack

#### func (*Stack) Pop

```go
func (s *Stack) Pop() (int, error)
```
Pop (delete) the item on the top of stack and return it

#### func (*Stack) Push

```go
func (s *Stack) Push(number int)
```
Push (add) a new value to the top of the stack

#### func (*Stack) Size

```go
func (s *Stack) Size() int
```
Returns the size (number of elements) in the stack

## Example

```go
package main

import (
    "fmt"
    "go.ajitem.com/stack"
)

func main() {
    s := stack.New()

    s.Push(1)
    s.Push(2)
    s.Push(3)
    s.Push(4)
    
    fmt.Println("size:", s.Size())
    
    number, _ := s.Pop()
    fmt.Println("number:", number)
    number, _ = s.Pop()
    fmt.Println("number:", number)
    number, _ = s.Pop()
    fmt.Println("number:", number)
    
    fmt.Println("size:", s.Size())
    number, _ = s.Peek()
    fmt.Println("number:", number)
    
    // Output:
    // size: 4
    // number: 4
    // number: 3
    // number: 2
    // size: 1
    // number: 1
}
```