# godash

Modeled after the [lodash](https://lodash.com/) library, godash is a collection of useful functions for working with
slices/arrays, maps, sets, hashable sets, optionals and more. It is meant to help you write more concise and readable
code in a functional way.

Work in progress, I will be adding more functions as I need them/as I have time.
Feel free to contribute. You can either contribute by adding functions, by adding tests or by adding documentation.

## Install

```shell
go get github.com/GrewalAS/godash
```

## Completed:

Only functions that are themselves written and are unit tested are marked as done unless they are one-liners. Function
that do not have tests are marked with a `-` instead of a `x`.

- [x] Conditions
    - [x] ComparableEquality
    - [x] NumberLessThan
    - [x] HashableEquality
    - [x] AnyEquality

- [x] Types
    - [x] Number
    - [x] Pair
    - [x] Hashable
    - [x] Iterable

- [x] Utils
    - [x] Optional
    - [x] Panics
    - [x] ResultWithError
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