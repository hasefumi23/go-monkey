# go-monkey

```go
let counter = fn(x) {
    if (x > 100) {
        return true;
    } else {
        let foobar = 9999;
        counter(x + 1);
    }
}
counter(0);
```

```go
parser_test.go:318: expected="((a * ([1, 2, 3, 4][(b * c)])) * d)", got="((a * ([1, 2, 3, 4][(b * c)]) * d)"
parser_test.go:318: expected="add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))", got="add((a * (b[2]), (b[1], (2 * ([1, 2][1]))"
```

```go
{"name": "Monkey", "age": 0, "type": "Language", "status": "awesome"}
let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
people[0]["name"]
people[0]["age"] + people[1]["age"]
```
