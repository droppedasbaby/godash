# godash

**Note: This project has been sun-setted. It remains available for reference and
educational purposes. The goal was to gain a better understanding of basic Go
concepts through practical implementation.**

Modeled after the [lodash](https://lodash.com/) library in Javascript/Typescript,
godash is a collection of useful functions for working with slices/arrays, maps,
sets, hashable sets, optionals, and more. It is meant to help you write more
concise and readable code in a functional way.

Although the project is no longer actively developed, feel free to explore its
codebase for learning purposes or to build upon it.

## Install

```shell
go get github.com/GrewalAS/godash
```

## Completed

- [x] Conditions

  - [x] ComparableEquality
  - [x] NumberEquality
  - [x] NumberLessThan
  - [x] NumberLessThanOrEqual
  - [x] NumberGreaterThan
  - [x] NumberGreaterThanOrEqual
  - [x] HashableEquality
  - [x] AnyEquality

- [x] Types

  - [x] Number
  - [x] Pair
  - [x] Hashable
  - [x] Iterable

- [x] Utils

  - [x] Option
  - [x] Panics
  - [x] Result
  - [x] Testing

- [x] Maps

  - [x] Add - Modifies original DS, does not return a new one
  - [x] AddAll - Modifies original DS, does not return a new one
  - [x] Contains
  - [x] FromMap
  - [x] Keys
  - [x] Remove - Modifies original DS, does not return a new one
  - [x] RemoveAll - Modifies original DS, does not return a new one
  - [x] ToMap
  - [x] Values

- [x] HashableMap

  - [x] Add - Modifies original DS, does not return a new one
  - [x] AddAll - Modifies original DS, does not return a new one
  - [x] Contains
  - [x] FromHashableMap
  - [x] Keys
  - [x] Remove - Modifies original DS, does not return a new one
  - [x] RemoveAll - Modifies original DS, does not return a new one
  - [x] ToHashableMap
  - [x] Values

- [x] Set

  - [x] Add - Modifies original DS, does not return a new one
  - [x] AddAll - Modifies original DS, does not return a new one
  - [x] Contains
  - [x] Difference
  - [x] FromSet
  - [x] Intersection
  - [x] IsSubsetOf
  - [x] IsSupersetOf
  - [x] Remove - Modifies original DS, does not return a new one
  - [x] RemoveAll - Modifies original DS, does not return a new one
  - [x] ToSet
  - [x] Union

- [x] HashableSet

  - [x] Add - Modifies original DS, does not return a new one
  - [x] AddAll - Modifies original DS, does not return a new one
  - [x] Contains
  - [x] Difference
  - [x] FromHashableSet
  - [x] Intersection
  - [x] IsSubsetOf
  - [x] IsSupersetOf
  - [x] Remove - Modifies original DS, does not return a new one
  - [x] RemoveAll - Modifies original DS, does not return a new one
  - [x] Size
  - [x] ToHashableSet
  - [x] Union

- [x] Slices/Arrays

  - [x] Chunk
  - [x] Compact
  - [x] Contains
  - [x] ContainsHashable
  - [x] ContainsWith
  - [x] Difference
  - [x] DifferenceHashable
  - [x] DifferenceWith
  - [x] Filter
  - [x] Flatten
  - [x] ForEach
  - [x] Head // One liner, no tests
  - [x] Intersection
  - [x] IntersectionHashable
  - [x] IntersectionWith
  - [x] Map
  - [x] Tail // One liner, no tests
  - [x] Union
  - [x] UnionHashable
  - [x] UnionWith
  - [x] Unzip
  - [x] Zip
  - [x] ZipWithIndex
  - [x] Sample
  - [x] SampleWithSeed
  - [x] SampleN
  - [x] SampleNWithSeed
  - [x] Shuffle
  - [x] ShuffleWithSeed

- [x] Linked List

  - [x] AddFirst
  - [x] AddLast
  - [x] Clear
  - [x] Contains
  - [x] FromSlice
  - [x] Get
  - [x] Insert
  - [x] Remove
  - [x] ToSlice

- [x] Search

  - [x] LinearSearch
  - [x] BinarySearch
  - [x] Bisect
    - [x] BisectLeftWith
    - [x] BisectWith
    - [x] BisectRightWith

- [x] Queue // Incredibly simple, no tests

  - [x] Enqueue
  - [x] Dequeue
  - [x] Peek
  - [x] Size

- [x] Stack // Incredibly simple, no tests

  - [x] Push
  - [x] Pop
  - [x] Peek
  - [x] Size

- [x] Strings/strutil
  - [x] CamelCase
  - [x] SnakeCase
  - [x] KebabCase
  - [x] PadLeft
  - [x] PadRight
  - [x] Pad
  - [x] Truncate
