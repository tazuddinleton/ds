# Data Structures in Go
An implementation of data structures in golang and solutions to common problems.

### To run tests
`go test -v -count=1 ./...`
### To run a specific test
`go test -v -count=1 -run <name of the test> <path to package>`

- [Array](array/array.go) | [Tests](array/array_test.go)
- [Stack](stack/Readme.md)
  - [Array based stack](stack/arraybased/stack.go) | [Tests](stack/arraybased/stack_test.go)
  - [Linked list based stack](stack/linkedlistbased/stack_ll.go) | [Tests](stack/linkedlistbased/stack_ll_test.go)
- Queue
  - [Ring buffer](queue/ringbuffer_based/queue.go) | [Tests](queue/ringbuffer_based/queue_test.go)
- Linked List
  - [Singly](linkedlist/singlyoo/singlyoo.go) | [Tests](linkedlist/singlyoo/singlyoo.go)
  - [Doubly](linkedlist/doubly/doubly.go) | [Tests](linkedlist/doubly/doubly_test.go)
  - [Singly Circular](linkedlist/singlyoocircular/singlycircular.go) | [Tests](linkedlist/singlyoocircular/singlycircular_test.go)
