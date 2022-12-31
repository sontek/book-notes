# Notes for book Learning Go
## Chapter 1
- Recommended tools:
    - go fmt
    - goimports
    - golint
    - go vet
    - [golangci-link](https://github.com/golangci/golangci-lint) runs many tools at the same time.

- Recommend reading: https://go.dev/doc/effective_go
- Recommend reading: https://github.com/golang/go/wiki/CodeReviewComments

## Chapter 2
- Integers can be represented with underscores for better readability
    - 1_234_567
    - Can use `var foo = ...` to assign variables or `foo := ...`

## Chapter 3
- Go considers the size of an array part of the type, so you should rarely use arrays
    - [3]int is a different type than [4]int and they can't be type casted.
- Go has type "slice" that is a better thing to use than arrays
    - built in `append(<slice>, <number>, <number2>, <number3>)` to add items to it, `len` to check the length.
    - Since slices will have a dynamic capacity allocated you can check that with `cap`
    - Can define length and capacity of a slice if you need to with `make`:
        - `x := make([]int, 5)`
    - If you Slice a slice like `x[:2]` it is referencing the same memory is the parent slice.
        - You can use `copy` to get around this.
- Maps will return the empty value of the type even if the key doesn't exist, so you can create a map
  and immediately start using keys in it, even if they don't exist.  Like `foo["bar"]++` would succeed
  even if you never added `bar` into the map.
    - If you need to know if the key actually exists you can use the "comma ok idiom" which looks like,
      `v, ok = m["world"]`
- How hashmaps work: https://www.youtube.com/watch?v=Tl7mi9QmLns&themeRefresh=1
- Go has structs, not classes.  Structs do not support inheritance. 
- You can declare "anonymous structs" which means just not giving a name and using it at that time.

## Chapter 4
- Shadowing variable is easy if you use `:=` so its important to using `shadow` for linting and finding it

## Chapter 5
- Go is call by value
    - caveat is that map and slices are implemented via pointers, so you *can* modify them from another func.
- `defer` is a way to run clean up like closing file handles / sql connections 

## Chapter 6
- For pointers, `&` is the address operator. `&x` is the address to the variable x.
- `*` is the indirection operator, also called dereferencing.  This turns an address into a value.
- You cannot derefence a `nil` pointer.   You must check if its `nil`.
- Slices are are struct that has a pointer in it.  So when its passed and copied the capacity and length
  aren't modified even though the value from the pointer can be.

## Chapter 7
- There isn't any inheritance but there is "embedded composition" which will drop a child object's methods
  onto the parent.
- Can use `interface{}` or generics for generic methods
- Great tool for dependency injection: https://github.com/google/wire

## Chapter 8
- There are no exceptions in go, you return the error as a the last return value
- When using custom errors, never define a variable to be of the type of your custom error. Either explicitly return nil when no error occurs or define the variable to be of type error.
- errors.Is and errors.As solve error chain handling.   You can also use `errors.Unwrap`.

## Chapter 9
- Code in go becomes a module when it has a go.mod file in the root.  Don't manage this manually, use `go mod init`.
- Go uses capitalization to determine if a package-level identifier is visible outside of the package. So if you start
  something with a capital letter as its name it'll be exported and can be imported by other modules.
- Relative paths aren't used, we give the modules and packages unique names.
- If two packages have the same name you can give one of them a different name:
  ```
  import (
    crand "crypto/rand"
    "math/rand"
  )
  ```
- Can view docs of a package with `go doc <package>`
- If you have internal code you want to share between your own packages you can name it a special name `internal` and
  it will not be exported to any other packages that don't share a parent with the internal package.
- go.mod and go.sum should be committed to git
- List versions available of a package with `go list -m -versions github.com/gin-gonic/gin`
- Can use `go get github.com/gin-gonic/gin@v1.8.2` to pin to a specific version
  - `go get -u <repo>` will upgrade to the latest
  - `go get -u patch` will upgrade to the most recent patch version of the major version you are pinned to
- `go mod tiny` to remove unused dependencies from go.sum

## Chapter 10
TODO: Skipped this chapter, go routines are worth their own learning.

- `select` is a switch statement that randomly selects which one it resolves first.  Great for channels and async
   coding.  It prevents deadlocks by checking if any of its cases can proceed.
- Pass variables into go routines if they are changing (like in a for loop). https://go.dev/doc/faq#closures_and_goroutines

## Chapter 11
Nothing useful, just some standard library discussions.  They can be done in their own learning.

## Chapter 12
- There is a `context.Context` interface that can be used to pass context around as `ctx` 
- You can use `context.WithCancel` if you want a cancellable context that is used in coroutines.
- You can use `context.WithTimeout` if you want to cancel a long running process.

## Chapter 13
- `t.Error` and `t.Fatal` can be used to fail tests
- `TestingMain` func runs once before tests are ran and can be used for setup.
- `t.Cleanup` can be used for tearing down things like files / db connections
- `testdata` is a reserved folder for test data in golang.
- Put your tests in a different package like `<package_name>_test` to test the public API.
- `go-cmp` is a third party module thats great for diffing structs. https://github.com/google/go-cmp
- Table tests are a way to test many variations on the same function
  - `t.Run` is called multiple times in a test for this to work.
- Go has built in benchmarking which are triggered with any functions that start with `Benchmark` and take `testing.B`
  - Every Go benchmark must have a loop that iterates from 0 to b.N.
  - It supports `b.Run` as well for multiple benchmarks with different iputs.
- https://github.com/stretchr/testify and https://github.com/golang/mock are good mock/stub libraries.
- `httptest` can be used to mock APIs.
- You can have a magic comment `// +build integration` that will only build a file if you pass
  `go test -tags integration` otherwise it'll be ignored.   This is a nice way to prevent certain tests from being
  included.

## Resources
- Book Code Examples: https://github.com/learning-go-book
- VSCode with Go: https://www.youtube.com/watch?v=1MXIGYrMk80
- Journey of Go's Garbage Collector: https://go.dev/blog/ismmkeynote
- Don't just check errors, handle them: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
- How do you structure Go Apps: https://www.youtube.com/watch?v=oL6JBUk6tj0
- Go Modules v2 and beyond: https://go.dev/blog/v2-go-modules
- The scheduler saga: https://www.youtube.com/watch?v=YHRO5WQGh0k
- Profiling go: https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/