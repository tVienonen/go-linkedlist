# LinkedList data structure for Go

Simple structure that supports serializing to JSON and back. I made this just to practice Go.

# Usage:

```
package main

import "github.com/tVienonen/go-linkedlist/linkedlist"

func main() {
    // Create
    list := linkedlist.NewLinkedList("some data", 123123, []int{3,4,5})
    // Insert
    list.Insert("asd")
    // Remove
    list.Remove(list.Head())
    // Iterate
    i := list.GetIterator()
    for item := i.Next(); item != nil; item = i.Next() {
        // do stuff with item
        
    }
}
```
