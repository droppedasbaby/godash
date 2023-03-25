# godash

Modeled after the [lodash](https://lodash.com/) library, godash is a collection of useful functions for working with
slices/arrays, maps, sets, hashable sets, optionals and more. It is meant to help you write more concise and readable
code in a functional way.

Work in progress, I will be adding more functions as I need them/as I have time.
Feel free to contribute. You can either contribute by adding functions, by adding tests or by adding documentation.

## To-Do:

Only functions that are themselves written and are unit tested are marked as done unless they are one-liners. Function 
that do not have tests are marked with a `-` instead of a `x`.

- [x] Conditions

- [x] Types

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

- [-] HashableMap
    - [-] Add - Modifies original DS, does not return a new one
    - [-] AddAll - Modifies original DS, does not return a new one
    - [-] Contains
    - [-] FromHashableMap
    - [-] Keys
    - [-] Remove - Modifies original DS, does not return a new one
    - [-] RemoveAll - Modifies original DS, does not return a new one
    - [-] ToHashableMap
    - [-] Values

- [x] Set
    - [x] Add - Modifies original DS, does not return a new one
    - [x] AddAll - Modifies original DS, does not return a new one
    - [x] Contains
    - [x] Difference
    - [x] FromSet
    - [x] Get
    - [x] Intersection
    - [x] IsSubsetOf
    - [x] IsSupersetOf
    - [x] Remove - Modifies original DS, does not return a new one
    - [x] RemoveAll - Modifies original DS, does not return a new one
    - [x] ToSet
    - [x] Union

- [ ] HashableSet
    - [ ] Add - Modifies original DS, does not return a new one
    - [ ] AddAll - Modifies original DS, does not return a new one
    - [ ] Contains
    - [ ] Difference
    - [ ] FromHashableSet
    - [ ] Get
    - [ ] Intersection
    - [ ] IsSubsetOf
    - [ ] IsSupersetOf
    - [ ] Remove - Modifies original DS, does not return a new one
    - [ ] RemoveAll - Modifies original DS, does not return a new one
    - [ ] Size
    - [ ] ToHashableSet
    - [ ] Union

- [ ] Slices/Arrays 
    - [ ] Chunk
    - [ ] Compact
    - [ ] Contains
    - [ ] ContainsWith
    - [ ] Difference
    - [ ] DifferenceWith
    - [ ] Drop
    - [ ] DropRight
    - [ ] DropRightWhile
    - [ ] DropWhile
    - [x] Filter
    - [ ] FlatMap
    - [ ] Flatten
    - [x] ForEach
    - [ ] Head
    - [ ] Intersection
    - [ ] IntersectionWith
    - [x] Map
    - [ ] Range
    - [ ] RangeRight
    - [x] Reduce
    - [ ] Tail
    - [ ] Take
    - [ ] TakeRight
    - [ ] TakeRightWhile
    - [ ] TakeWhile
    - [ ] Union
    - [ ] UnionWith
    - [ ] Unzip
    - [ ] Zip
    - [ ] ZipWithIndex

- [ ] Linked List
    - [ ] Append
    - [ ] Contains
    - [ ] FromSlice
    - [ ] Get
    - [ ] Insert
    - [ ] Remove
    - [ ] RemoveAt
    - [ ] Size
    - [ ] ToSlice