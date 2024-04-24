# Datatructure & Algorithms (DSA)

## Content

- [Datatructure \& Algorithms (DSA)](#datatructure--algorithms-dsa)
  - [Content](#content)
  - [Complexity](#complexity)
    - [Time Complexity - Big O](#time-complexity---big-o)
    - [Space Complexity - Big Theta](#space-complexity---big-theta)
  - [Datastructure](#datastructure)
    - [Array](#array)
      - [Concept](#concept)
      - [Use Cases](#use-cases)
      - [Operations](#operations)
      - [Props/Cons Trade-off](#propscons-trade-off)
    - [List](#list)
      - [Concept](#concept-1)
      - [Use Cases](#use-cases-1)
      - [Operations](#operations-1)
      - [Props/Cons Trade-off](#propscons-trade-off-1)
    - [Linked List](#linked-list)
      - [Concept](#concept-2)
      - [Use Cases](#use-cases-2)
      - [Operations](#operations-2)
      - [Props/Cons Trade-off](#propscons-trade-off-2)
    - [Hash Table](#hash-table)
      - [Concept](#concept-3)
      - [Use Cases](#use-cases-3)
      - [Operations](#operations-3)
      - [Props/Cons Trade-off](#propscons-trade-off-3)
    - [Hash Map](#hash-map)
      - [Concept](#concept-4)
      - [Use Cases](#use-cases-4)
      - [Operations](#operations-4)
      - [Props/Cons Trade-off](#propscons-trade-off-4)
    - [Stack](#stack)
      - [Concept](#concept-5)
      - [Use Cases](#use-cases-5)
      - [Operations](#operations-5)
      - [Props/Cons Trade-off](#propscons-trade-off-5)
    - [Queue](#queue)
      - [Concept](#concept-6)
      - [Use Cases](#use-cases-6)
      - [Operations](#operations-6)
      - [Props/Cons Trade-off](#propscons-trade-off-6)
    - [Tree](#tree)
      - [Concept](#concept-7)
      - [Use Cases](#use-cases-7)
      - [Operations](#operations-7)
      - [Props/Cons Trade-off](#propscons-trade-off-7)
    - [Graph](#graph)
      - [Concept](#concept-8)
      - [Use Cases](#use-cases-8)
      - [Operations](#operations-8)
      - [Props/Cons Trade-off](#propscons-trade-off-8)
    - [Trie](#trie)
      - [Concept](#concept-9)
      - [Use Cases](#use-cases-9)
      - [Operations](#operations-9)
      - [Props/Cons Trade-off](#propscons-trade-off-9)
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

#### Concept

Array is a **fixed size** data structure that stores values in a **contiguous block of memory**. Once the array is created, you can access and modify the values by referring to an index number.

Array items will be indexed **starting from 0**. The annotation A[0] represents the first element of the array, A[1] represents the second element, and so on.

When you create an array, the address of the first element is returned.

In many programming languages, arrays are implemented as pointers. So it will **pass-by-reference** by default.

**The elements of the array are required to have the same data type.**

#### Use Cases

- To store a collection of data when the size is known (or guess) and all items have the same data type.

#### Operations

- Access an element: O(1) - Both read and write operations are O(1).

#### Props/Cons Trade-off

- Props:
  - Access operation has constant time complexity
- Cons:
  - Fixed size
  - Require to have the same data type

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