# rueban[鲁班]

## A toolkit for slicing map and other data processing


### Array[slice]

#### Chunk

```go
array.chunk([]string{"a", "b", "c", "d", "e"}, 2)
// ==> [["a", "b"],["c", "d"]["e"]]
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
array.Drop([]string{"a", "b", "c", "d", "e"}, 2,3)
// ==> ["a", "b", "e"]
```


#### Fill

```go
array.Fill([]string{"a", "b", "c", "d", "e"},"d", 1,5)
// ==> ["e","d","d","d","d"]
```


#### Find

```go
array.Find([]string{"a", "b", "c", "d", "e"}, func(v string) bool {
	return v != "c"
})

// ==> ["a", "b", "d", "e"], true, [0, 1, 3, 4]
```

#### FindFirst

```go
array.Find([]string{"a", "b", "c", "d", "e"}, func(v string) bool {
	return v != "c"
})

// ==> "a", true, 0
```

#### Head

```go
array.Head([]string{"a", "b", "c", "d", "e"})

// ==> "a"
```


#### IndexOf

```go
array.IndexOf([]string{"a", "b", "c", "d", "e"}, "a", func(i string) string { return i }, 3)
// ==> -1
array.IndexOf([]string{"a", "b", "c", "d", "e"}, "b", func(i string) string { return i }, 0)
// ==> 1
```

#### Initial

```go
array.Initial([]string{"a", "b", "c", "d", "e"})
// ==> ["a","b","c","d"]
```

#### Intersection

```go
type TestIntersectionStruct struct {
    Name string
    Age  int
}


var TArr = []TestIntersectionStruct{
    TestIntersectionStruct{
        Name: "qm",
        Age:  10,
    },
    TestIntersectionStruct{
        Name: "qm",
        Age:  12,
    },
    TestIntersectionStruct{
        Name: "qm",
        Age:  18,
    },
}

var TArr2 = []TestIntersectionStruct{
    TestIntersectionStruct{
        Name: "qm",
        Age:  10,
    },
    TestIntersectionStruct{
        Name: "qm",
        Age:  12,
    },
    TestIntersectionStruct{
        Name: "qm",
        Age:  13,
    },
}

var TArr3 = []TestIntersectionStruct{
    TestIntersectionStruct{
        Name: "qm",
        Age:  10,
    },
    TestIntersectionStruct{
        Name: "qm",
        Age:  11,
    },
    TestIntersectionStruct{
        Name: "qm",
        Age:  18,
    },
}

array.Intersection(
	func(item TestIntersectionStruct) int {
            return item.Age
        }, TArr, TArr2, TArr3)
// ==> [{qm 10}]

array.Intersection(
	func(i string) string {
            return i
        },[]string{"a", "b", "c"}, []string{"b", "c", "d"}, []string{"b", "c", "f"})
// ==> ["b","c"]
```

#### Uniq

```go
array.Uniq(
	func(i string) string {
            return i
        },[]string{"a", "b", "c","b","c"})

// ==> ["a", "b", "c"]
```

#### Nth

```go
    Nth([]int{1,2,3,4,5}, 10)  ==> 5
    Nth([]int{1,2,3,4,5}, 2)   ==> 3
    Nth([]int{1,2,3,4,5}, -3)   ==> 3
```

#### Pull

```go
    Pull(func(i int) int { return i }, []int{1, 2, 2, 2, 23, 3, 12, 4, 5}, 2) ==> [1 23 3 12 4 5]
    Pull(func(i TestStruct) string { return i.Name }, TestArr2, "name2") ==> [{name1} {name3} {name4} {name5}]
```

#### Remove

```go
    Remove(func(i TestStruct) bool {
		return i.Name == "name2"
	}, TestArr2) ==> [{name1} {name3} {name4} {name5}]

    Remove(func(i int) bool {
        return i < 3
    }, []int{1, 2, 3, 4, 5}) ==> [3 4 5]
```

#### Filter

```go
    Filter(func(i TestStruct) bool {
		return i.Name == "name2"
	}, TestArr2) ==> [{name2}]

    Filter(func(i int) bool {
        return i < 3
    }, []int{1, 2, 3, 4, 5}) ==> [1 2]
```