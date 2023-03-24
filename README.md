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
    - [x] FromMap
    - [x] ToMap
    - [x] Add - Modifies original DS, does not return a new one
    - [x] AddAll - Modifies original DS, does not return a new one
    - [x] Remove - Modifies original DS, does not return a new one
    - [x] RemoveAll - Modifies original DS, does not return a new one 
    - [x] Contains
    - [x] Keys
    - [x] Values


- [ ] HashableMap
    - [ ] FromHashableMap
    - [ ] ToHashableMap
    - [ ] Add - Modifies original DS, does not return a new one
    - [ ] AddAll - Modifies original DS, does not return a new one
    - [ ] Remove - Modifies original DS, does not return a new one
    - [ ] RemoveAll - Modifies original DS, does not return a new one
    - [ ] Contains
    - [ ] Keys
    - [ ] Values

- [ ] Set
    - [x] FromSet
    - [x] ToSet
    - [ ] Add
    - [ ] Contains
    - [ ] Difference
    - [ ] Intersection
    - [ ] Union
    - [ ] Remove
    - [ ] IsSubsetOf
    - [ ] IsSupersetOf
    - [ ] Size

- [ ] HashableSet
    - [ ] FromHashableSet
    - [ ] ToHashableSet
    - [ ] Add
    - [ ] Contains
    - [ ] Difference
    - [ ] Intersection
    - [ ] Union
    - [ ] Remove
    - [ ] IsSubsetOf
    - [ ] IsSupersetOf
    - [ ] Size

- [ ] Slices/Arrays 
    - [ ] Chunk
    - [ ] Compact
    - [ ] Contains
    - [ ] ContainsWith
    - [ ] Drop
    - [ ] DropWhile
    - [ ] DropRight
    - [ ] DropRightWhile
    - [x] Filter
    - [-] ForEach
    - [ ] Head
    - [-] Map
    - [ ] Range
    - [ ] RangeRight
    - [-] Reduce
    - [ ] Tail
    - [ ] Take
    - [ ] TakeWhile
    - [ ] TakeRight
    - [ ] TakeRightWhile
    - [ ] Zip
    - [ ] Unzip
    - [ ] ZipWithIndex
    - [ ] Difference
    - [ ] DifferenceWith
    - [ ] Intersection
    - [ ] IntersectionWith
    - [ ] Union
    - [ ] UnionWith
    - [ ] FlatMap
    - [ ] Flatten

- [ ] Linked List
    - [ ] Append
    - [ ] Insert
    - [ ] Remove
    - [ ] RemoveAt
    - [ ] Contains
    - [ ] Get
    - [ ] Size
    - [ ] ToSlice
    - [ ] FromSlice