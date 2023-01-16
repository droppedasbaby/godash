# godash

Modeled after the [lodash](https://lodash.com/) library, godash is a collection of useful functions for working with
slices/arrays, maps, sets, hashable sets, optionals and more. It is meant to help you write more concise and readable
code in a functional way. So, in any case, the original data structure passed in is never edited. A copy is always
returned.

It is a work in progress, I will be adding more functions as I need them/as I have time.

Feel free to contribute. You can either contribute by adding functions, by adding tests or by adding documentation.

## To-Do:

Only functions that are themselves written and are unit tested are marked as done.

- [x] Conditions

- [x] Types

- [ ] Utils
    - [ ] Optional
    - [ ] Panics
    - [ ] ResultWithError
    - [ ] Testing

- [ ] Maps
    - [ ] Add
    - [ ] Remove
    - [ ] Contains
    - [ ] Keys
    - [ ] Values
    - [ ] ToPairs
    - [ ] FromPairs


- [ ] HashableMap
    - [ ] Add
    - [ ] Remove
    - [ ] Contains
    - [ ] Keys
    - [ ] Values
    - [ ] ToPairs
    - [ ] FromPairs

- [ ] Set
    - [ ] FromSet
    - [ ] ToSet
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

- [ ] Collections
    - [ ] Chunk
    - [ ] Compact
    - [ ] Contains
    - [ ] ContainsWith
    - [ ] Drop
    - [ ] DropWhile
    - [ ] DropRight
    - [ ] DropRightWhile
    - [ ] Filter
    - [ ] ForEach
    - [ ] Head
    - [ ] Map
    - [ ] Range
    - [ ] RangeRight
    - [ ] Reduce
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