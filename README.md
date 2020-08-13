# go-monkey

## For eval test

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
expected="((a * ([1, 2, 3, 4][(b * c)])) * d)", got="((a * ([1, 2, 3, 4][(b * c)]) * d)"
expected="add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))", got="add((a * (b[2]), (b[1], (2 * ([1, 2][1]))"
```

```go
{"name": "Monkey", "age": 0, "type": "Language", "status": "awesome"}
let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
people[0]["name"]
people[0]["age"] + people[1]["age"]
```

```go
let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); };
unless(10 > 5, puts("not greater"), puts("greater"));
```
