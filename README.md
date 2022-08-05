# rueban[鲁班]

## A toolkit for slicing map and other data processing


### Array[slice]

#### Chunk

```go
array.chunk(['a', 'b', 'c', 'd', 'e'], 2)
// ==> [['a', 'b'],['c', 'd']['e']]
```

#### Diffrernce

```go
Diffrernce([]int{1,2,3,4,5}, []int{2, 3},
func(i int) int {
	return i
})

// ==> []int{1,4,5}
```


#### Drop

```go
array.Drop(['a', 'b', 'c', 'd', 'e'], 2,3)
// ==> ['a', 'b', 'e']
```


#### Fill

```go
array.Fill(['a', 'b', 'c', 'd', 'e'],'d', 1,5)
// ==> ['e','d','d','d','d']
```


#### Find

```go
array.Find(['a', 'b', 'c', 'd', 'e'], func(v string) bool {
	return v != "c"
})

// ==> ['a', 'b', 'd', 'e'], true, [0, 1, 3, 4]
```

#### FindFirst

```go
array.Find(['a', 'b', 'c', 'd', 'e'], func(v string) bool {
	return v != "c"
})

// ==> 'a', true, 0
```
