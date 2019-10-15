# stack
--
    import "go.ajitem.com/stack"

A stack is a fundamental data structure which is extensively used. This package
provides a basic implementation of a stack in Go. A stack works using Last In
First Out (LIFO) method.

## Usage

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
func (s *Stack) Peek() int
```
Return the item on the top of the stack

#### func (*Stack) Pop

```go
func (s *Stack) Pop() int
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
