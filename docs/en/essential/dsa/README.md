# Datatructure & Algorithms (DSA)

## Content

- [Datatructure \& Algorithms (DSA)](#datatructure--algorithms-dsa)
  - [Content](#content)
  - [Complexity](#complexity)
    - [Time Complexity - Big O](#time-complexity---big-o)
    - [Space Complexity - Big Theta](#space-complexity---big-theta)
  - [Datastructure](#datastructure)
    - [Array](#array)
    - [Linked List](#linked-list)
    - [Hash Table](#hash-table)
      - [Concept](#concept)
      - [Use Cases](#use-cases)
      - [Operations](#operations)
      - [Props/Cons Trade-off](#propscons-trade-off)
    - [Hash Map](#hash-map)
      - [Concept](#concept-1)
      - [Use Cases](#use-cases-1)
      - [Operations](#operations-1)
      - [Props/Cons Trade-off](#propscons-trade-off-1)
    - [Stack](#stack)
      - [Concept](#concept-2)
      - [Use Cases](#use-cases-2)
      - [Operations](#operations-2)
      - [Props/Cons Trade-off](#propscons-trade-off-2)
    - [Queue](#queue)
      - [Concept](#concept-3)
      - [Use Cases](#use-cases-3)
      - [Operations](#operations-3)
      - [Props/Cons Trade-off](#propscons-trade-off-3)
    - [Tree](#tree)
      - [Concept](#concept-4)
      - [Use Cases](#use-cases-4)
      - [Operations](#operations-4)
      - [Props/Cons Trade-off](#propscons-trade-off-4)
    - [Graph](#graph)
      - [Concept](#concept-5)
      - [Use Cases](#use-cases-5)
      - [Operations](#operations-5)
      - [Props/Cons Trade-off](#propscons-trade-off-5)
    - [Trie](#trie)
      - [Concept](#concept-6)
      - [Use Cases](#use-cases-6)
      - [Operations](#operations-6)
      - [Props/Cons Trade-off](#propscons-trade-off-6)
  - [Algorithms](#algorithms)
    - [Search Algorithms](#search-algorithms)
      - [Linear Search](#linear-search)
      - [Binary Search](#binary-search)
    - [Sorting Algorithms](#sorting-algorithms)
      - [Selection Sort](#selection-sort)
      - [Insertion Sort](#insertion-sort)
      - [Bubble Sort](#bubble-sort)
      - [Quick Sort](#quick-sort)
      - [Merge Sort](#merge-sort)
      - [Heap Sort](#heap-sort)
      - [Tim Sort](#tim-sort)

## Complexity

### Time Complexity - Big O

Time Complexity, or Big O, is a measure of the rate of growth of an algorithm time spend when the input size increases.

Some common time complexity:

- Constant time complexity: O(1) -> *The input size does not affect the time it takes to execute the algorithm.*
- Linear time complexity: O(n) -> *If you add n items to your input, the algorithm will also increase n units of time*
- Parabolic time complexity: O(n^2) -> *If you add n items to your input, the algorithm will also increase n^2 units of time*
- Quadratic time complexity: O(n^3) -> *If you add n items to your input, the algorithm will also increase n^3 units of time*
- Exponential time complexity: O(2^n) -> *If you add n items to your input, the algorithm will also increase 2^n units of time*
- Logarithmic time complexity: O(log n) -> *If you add n items to your input, the algorithm will also increase log n units of time*
- Factorial time complexity: O(n!) -> *If you add n items to your input, the algorithm will also increase n! units of time*
- Polynomial time complexity: O(n^k) -> *If you add n items to your input, the algorithm will also increase n^k units of time*
- O (n log n) time complexity: O(n log n) -> *If you add n items to your input, the algorithm will also increase n log n units of time*

> Log n mean the log2 of n.

**Time Complexity - Efficiency - From More Efficient to Less Efficient**

O (1) < O (log n) < O (n) < O (n log n) < O (n^2) < O (n^3) < O (2^n) < O (n!)

### Space Complexity - Big Theta

Space Complexity, or Big Theta, is a measure of the amount of memory that an algorithm uses to store the input and output of the algorithm.

## Datastructure

### Array

**Concept:**

An array is a fundamental data structure that stores a fixed-size, sequential collection of elements of the same data type. Each element in an array has a unique index (position), starting from 0, that allows direct access to its value. Think of it as a numbered list where all the elements are stored contiguously in memory.

**Use Cases:**

* **Storing homogeneous data sets**: Arrays are ideal for storing collections of similar data types, like lists of numbers, student grades, or product prices.
* **Implementing other data structures**: Arrays are the building blocks for more complex data structures like stacks, queues, and some types of trees.
* **Random access**: Since elements are stored contiguously, accessing any element by its index is very efficient. This makes them suitable for scenarios where you need to frequently retrieve specific elements from the collection.

**Operations:**

* **Access:** Retrieving the value of an element at a specific index (constant time complexity - O(1) on average).
* **Modification:** Updating the value of an element at a specific index (constant time complexity).
* **Insertion:** Adding a new element at a specific index (shifting elements can be expensive, O(n) in worst case).
* **Deletion:** Removing an element at a specific index (shifting elements can be expensive, O(n) in worst case).
* **Search:** Linear search for a specific element (linear time complexity - O(n) in worst case).

**Pros & Cons (Trade-offs):**

**Pros:**

* **Simple and efficient for random access:** Retrieving elements by index is very fast.
* **Memory efficient:** Elements are stored contiguously in memory, minimizing wasted space.
* **Fixed size:** Knowing the size upfront allows for efficient memory allocation.

**Cons:**

* **Fixed size:**  The size of an array cannot be changed after creation.
* **Inefficient insertions/deletions:** Inserting or deleting elements in the middle can be slow because it requires shifting other elements.
* **Limited data type:** All elements must be of the same data type.

**Trade-offs:**

Arrays offer efficient random access but lack flexibility for insertions and deletions in the middle. Consider alternatives like linked lists if frequent insertions/deletions are crucial. However, linked lists have slower random access compared to arrays.

**How array stores data in memory:**

![How array stores data in memory](images/dsa/array.png)

**Examples (Golang):**

```go
package main

import "fmt"

func main() {
  // Initialize an array with 5 elements
  arr := [5]int{1,2,3,4,5}

  fmt.Printf("Array[%d]: %v\n", 0, arr[0]) // Output: Array[0]: 1

  for i := 0; i < len(arr); i++ {
    fmt.Printf("Array[%d]: %v\n", i, arr[i])
    // Will output: Array[0]: 1, Array[1]: 2, Array[2]: 3, Array[3]: 4, Array[4]: 5
  }

  // Modify an element
  arr[0] = 10
  fmt.Printf("Array[%d]: %v\n", 0, arr[0]) // Output: Array[0]: 10
}
```

### Linked List

**Concept:**

A linked list is a linear data structure where elements, called nodes, are not stored contiguously in memory. Each node contains two parts:

1. **Data:** The actual information stored in the node (e.g., a number, a string, an object).
2. **Next pointer:** A reference (link) that points to the next node in the sequence. This pointer tells the program where to find the next element in the list.  

Imagine a train with cars linked together. Each car (node) holds cargo (data) and has a hitch (next pointer) that connects it to the next car in the line. Unlike arrays, linked lists don't rely on physical memory location for order, but on these pointers to establish the sequence.

**Use Cases:**

* **Dynamic data sets:** Linked lists are ideal for collections where the size is unknown beforehand or changes frequently.  Since elements don't need to be shifted to accommodate insertions or deletions, they excel in these scenarios.
* **Implementing other data structures:** Linked lists serve as building blocks for more complex data structures like stacks, queues, and some graph implementations.

**Operations:**

* **Access:** Unlike arrays with direct index-based retrieval, accessing elements in a linked list requires traversing the list, starting from the head (the first node) and following the next pointers until the target element is found. This has a linear time complexity of O(n) in the worst case.
* **Modification:** Updating the data within a node itself is a simple operation.
* **Insertion:**  Inserting a new node at a specific position is generally faster than arrays, especially for insertions at the head or tail. The time complexity can be constant (O(1)) for head/tail insertions.
* **Deletion:** Removing a node is also efficient compared to arrays for non-head/tail deletions. However, finding the node to delete requires traversal, leading to a worst-case time complexity of O(n). 
* **Search:** Similar to access, searching for a specific element involves traversing the list, resulting in a linear time complexity of O(n) in the worst case.

**Pros & Cons (Trade-offs):**

**Pros:**

* **Dynamic size:** The size of a linked list can grow or shrink as needed, making it adaptable to changing data requirements.
* **Efficient insertions/deletions:** Adding or removing elements is generally faster than arrays, especially for insertions/deletions in the middle of the list.
* **No wasted space:** Memory is allocated only for the nodes themselves, reducing wasted space compared to arrays with unused elements.

**Cons:**

* **Slower random access:**  Since access relies on traversing the list, finding a specific element by its position (like you can with an index in an array) is slower.
* **More complex memory management:** Using pointers introduces overhead for managing memory references.

**Trade-offs:**

Linked lists shine when insertions/deletions are frequent operations, and the order of elements is more important than random access speed. However, if fast retrieval by index is crucial and the size is fixed, arrays are a better choice.

**How Linked List stores data in memory**

![How Linked List stores data in memory](images/dsa/linkedlist.png)

**Example (Golang):**

```go
package main

import "fmt"

type Node[T any] struct {
  data T
  next *Node[T]
}

type LinkedList[T any] struct {
  head *Node[T]
  tail *Node[T]
  size int
}

func NewLinkedList[T any]() *LinkedList[T] {
  return &LinkedList[T]{}
}

// Add a new node to the end of the linked list
//
// Time Complexity: O(1)
func (l *LinkedList[T]) Add(data T) {
  newNode := &Node[T]{data: data}
  if l.head == nil {
    l.head = newNode
    l.tail = newNode
  }

  l.tail.next = newNode
  l.tail = newNode

  l.size++
}

// Add a new node at a specific index
//
// Time Complexity: O(n)
func (l *LinkedList[T]) AddAt(index int, data T) {
  newNode := &Node[T]{data: data}

  if index == 0 {
    newNode.next = l.head
    l.head = newNode
  } else {
    current := l.head
    for i := 0; i < index-1; i++ {
      current = current.next
    }
    newNode.next = current.next
    current.next = newNode
  }

  l.size++
}

// Remove a node at a specific index
//
// Time Complexity: O(n)
func (l *LinkedList[T]) RemoveAt(index int) {
  if index == 0 {
    l.head = l.head.next
  } else {
    current := l.head
    for i := 0; i < index-1; i++ {
      current = current.next
    }
    current.next = current.next.next
  }

  l.size--
}

// Get the data at a specific index
//
// Time Complexity: O(n)
func (l *LinkedList[T]) Get(index int) T {
  current := l.head
  for i := 0; i < index; i++ {
    current = current.next
  }
  return current.data
}

// Implement the Stringer interface to print the linked list
func (l *LinkedList[T]) String() string {
  str := ""
  current := l.head
  for current != nil {
    str += fmt.Sprintf("%v ", current.data)
    current = current.next
  }
  return str
}

func main() {
  list := NewLinkedList[int]()
  list.Add(1)
  list.Add(2)
  list.Add(3)
  list.AddAt(1, 4)
  list.RemoveAt(2)
  fmt.Println(list.Get(1))
  // Output: 4
  fmt.Printf("List: %v\n", list)
  // Output: 1 4 3
}
```

### Hash Table

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Hash Map

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Stack

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Queue

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Tree

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Graph

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Trie

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

## Algorithms

### Search Algorithms

#### Linear Search

- Time Complexity - Big O: O(n)

#### Binary Search

- Time Complexity - Big O: O(log n)

### Sorting Algorithms

#### Selection Sort

- Time Complexity - Big O: O(n^2)

#### Insertion Sort

- Time Complexity - Big O: O(n^2)

#### Bubble Sort

- Time Complexity - Big O: O(n^2)

#### Quick Sort

- Time Complexity - Big O: O(n log n)

#### Merge Sort

- Time Complexity - Big O: O(n log n)

#### Heap Sort

- Time Complexity - Big O: O(n log n)

#### Tim Sort

- Time Complexity - Big O: O(n log n)