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
    - [Stack](#stack)
      - [Concept](#concept)
      - [Use Cases](#use-cases)
      - [Operations](#operations)
      - [Props/Cons Trade-off](#propscons-trade-off)
    - [Queue](#queue)
      - [Concept](#concept-1)
      - [Use Cases](#use-cases-1)
      - [Operations](#operations-1)
      - [Props/Cons Trade-off](#propscons-trade-off-1)
    - [Tree](#tree)
      - [Concept](#concept-2)
      - [Use Cases](#use-cases-2)
      - [Operations](#operations-2)
      - [Props/Cons Trade-off](#propscons-trade-off-2)
    - [Graph](#graph)
      - [Concept](#concept-3)
      - [Use Cases](#use-cases-3)
      - [Operations](#operations-3)
      - [Props/Cons Trade-off](#propscons-trade-off-3)
    - [Trie](#trie)
      - [Concept](#concept-4)
      - [Use Cases](#use-cases-4)
      - [Operations](#operations-4)
      - [Props/Cons Trade-off](#propscons-trade-off-4)
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

**[Example (Golang)]**

```go
// example/go/dsa/array.go

package dsa

import "fmt"

type Array[T any] struct {
	array []T
}

func NewArray[T any](size int) *Array[T] {
	return &Array[T]{
		array: make([]T, size),
	}
}

func (a *Array[T]) Get(index int) T {
	return a.array[index]
}

func (a *Array[T]) Set(index int, val T) {
	a.array[index] = val
}

func (a *Array[T]) Size() int {
	return len(a.array)
}

func (a *Array[T]) InsertAt(index int, val T) {
	for i := a.Size() - 1; i > index; i-- {
		a.array[i] = a.array[i-1]
	}

	a.array[index] = val
}

func (a *Array[T]) DeleteAt(index int) {
	for i := index; i < a.Size()-1; i++ {
		a.array[i] = a.array[i+1]
	}
}

func (a *Array[T]) String() string {
	str := "["
	for _, i := range a.array {
		str += fmt.Sprintf("%v ", i)
	}
	str += "]"
	return str
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
// example/go/dsa/linked_list.go
package dsa

import "fmt"

type Item[T any] struct {
	Data T
	Next *Item[T]
}

type LinkedList[T any] struct {
	head *Item[T]
	tail *Item[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add an item to the tail of the list
func (l *LinkedList[T]) Add(data T) {
	newItem := &Item[T]{
		Data: data,
		Next: nil,
	}

	if l.head == nil {
		l.head = newItem
		l.tail = newItem
	}

	l.tail.Next = newItem
	l.tail = newItem

	l.size++
}

func (l *LinkedList[T]) AddAt(index int, data T) {
	newItem := &Item[T]{
		Data: data,
		Next: nil,
	}

	if index >= l.size {
		l.tail.Next = newItem
		l.tail = newItem
	}

	currentItem := l.head

	for i := 1; i < index; i++ {
		currentItem = currentItem.Next
	}

	newItem.Next = currentItem.Next
	currentItem.Next = newItem

	l.size++
}

// Remove the last item
func (l *LinkedList[T]) Remove() {
	currentItem := l.head
	for i := 1; i < l.size-1; i++ {
		currentItem = currentItem.Next
	}

	currentItem.Next = nil
	l.size--
}

func (l *LinkedList[T]) RemoveAt(index int) {
	currentItem := l.head
	for i := 1; i < index; i++ {
		currentItem = currentItem.Next
	}

	currentItem.Next = currentItem.Next.Next
	l.size--
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) String() string {
	str := ""
	currentItem := l.head
	for currentItem.Next != nil {
		str += fmt.Sprintf("%v -> ", currentItem.Data)
	}
	str += "nil"
	return str
}

```

### Hash Table

**Concept:**

A hash table is a data structure that implements an associative array, also known as a dictionary or map. It stores key-value pairs and allows for fast retrieval of values based on their keys. Unlike arrays or linked lists that rely on order or position, hash tables use a hashing function to map keys to unique indexes within a fixed-size array (called a hash table). This mapping allows for near-constant time (average case) lookups for elements. 

Imagine a library catalog with books categorized by Dewey Decimal System codes (keys). The codes act like unique identifiers for finding specific books (values). A well-designed hash table is like a very efficient librarian who can quickly locate a book (value) based on its Dewey Decimal code (key).

Hash table often uses LinkedList as a backend data structure to store key-value pairs.

Typical Hash Table data structure look like

![Typical Hash Tabl data structure](https://res.cloudinary.com/dmsb4anlx/image/upload/v1713975104/hungtp/aiyunkegh78qdf3eolip.png)

**Use Cases:**

* **Fast key-value lookups:** Hash tables excel when you need to retrieve data based on a unique identifier (key). This makes them ideal for:
    * Caching frequently accessed data
    * Implementing symbol tables in compilers
    * Representing user accounts in applications (username as key, user data as value)
    * Network routing (IP addresses as keys, routes as values)
    * Index in database

**Operations:**

* **Insertion:** Adding a new key-value pair to the hash table.
* **Deletion:** Removing a key-value pair from the hash table.
* **Search:** Finding the value associated with a specific key. (Average case constant time - O(1))
* **Update:** Modifying the value associated with an existing key.

**Pros & Cons (Trade-offs):**

**Pros:**

* **Fast average-case search:**  Retrieving elements by key is very efficient (O(1) on average) due to the hashing function.
* **Efficient for key-value lookups:** Ideal for scenarios where you need to find data based on unique identifiers.
* **Dynamic size:** Although the hash table itself has a fixed size, it can be resized to accommodate more elements while maintaining efficiency.

**Cons:**

* **Potential collisions:**  Hash functions might map different keys to the same index (collision). Collision resolution techniques are needed to handle these situations, which can add some overhead.
* **Slower worst-case search:**  In the worst case (e.g., very bad hash function or overloaded hash table), lookups might become slower.
* **Wasted space:**  Due to collisions and resizing, some space in the hash table might be unused.

**Trade-offs:**

Hash tables offer exceptional speed for key-value lookups but require careful consideration of **hash function selection and collision resolution techniques**. If the order of elements is not important, and fast retrieval based on unique keys is essential, hash tables are a powerful choice.

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