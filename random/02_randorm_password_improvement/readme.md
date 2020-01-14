## Read Me

You have to save it into a file named `random_test.go` and run it with

```sh
go test -bench . -benchmem
```

## I. Improvements

### 1. Genesis (Runes)

As a reminder, the original, general solution we're improving is this:

````go
func init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
``

### 2. Bytes
If the characters to choose from and assemble the random string contains only the uppercase and lowercase letters of the English alphabet, we can work with bytes only because the English alphabet letters map to bytes 1-to-1 in the UTF-8 encoding (which is how Go stores strings).

So instead of:

```go
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
``

we can use:

```go
var letters = []bytes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
````

Or even better:

```go
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
```
