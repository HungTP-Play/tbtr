# Datatructure & Algorithms (DSA)

## Content

- [Datatructure \& Algorithms (DSA)](#datatructure--algorithms-dsa)
  - [Content](#content)
  - [Complexity](#complexity)
    - [Time Complexity - Big O](#time-complexity---big-o)
    - [Space Complexity - Big Theta](#space-complexity---big-theta)
  - [Datastructure](#datastructure)
    - [Array](#array)
    - [List](#list)
      - [Concept](#concept)
      - [Use Cases](#use-cases)
      - [Operations](#operations)
      - [Props/Cons Trade-off](#propscons-trade-off)
    - [Linked List](#linked-list)
      - [Concept](#concept-1)
      - [Use Cases](#use-cases-1)
      - [Operations](#operations-1)
      - [Props/Cons Trade-off](#propscons-trade-off-1)
    - [Hash Table](#hash-table)
      - [Concept](#concept-2)
      - [Use Cases](#use-cases-2)
      - [Operations](#operations-2)
      - [Props/Cons Trade-off](#propscons-trade-off-2)
    - [Hash Map](#hash-map)
      - [Concept](#concept-3)
      - [Use Cases](#use-cases-3)
      - [Operations](#operations-3)
      - [Props/Cons Trade-off](#propscons-trade-off-3)
    - [Stack](#stack)
      - [Concept](#concept-4)
      - [Use Cases](#use-cases-4)
      - [Operations](#operations-4)
      - [Props/Cons Trade-off](#propscons-trade-off-4)
    - [Queue](#queue)
      - [Concept](#concept-5)
      - [Use Cases](#use-cases-5)
      - [Operations](#operations-5)
      - [Props/Cons Trade-off](#propscons-trade-off-5)
    - [Tree](#tree)
      - [Concept](#concept-6)
      - [Use Cases](#use-cases-6)
      - [Operations](#operations-6)
      - [Props/Cons Trade-off](#propscons-trade-off-6)
    - [Graph](#graph)
      - [Concept](#concept-7)
      - [Use Cases](#use-cases-7)
      - [Operations](#operations-7)
      - [Props/Cons Trade-off](#propscons-trade-off-7)
    - [Trie](#trie)
      - [Concept](#concept-8)
      - [Use Cases](#use-cases-8)
      - [Operations](#operations-8)
      - [Props/Cons Trade-off](#propscons-trade-off-8)
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

### List

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

### Linked List

#### Concept

#### Use Cases

#### Operations

#### Props/Cons Trade-off

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