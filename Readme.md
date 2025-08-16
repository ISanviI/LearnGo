# How to run the go files of this repo

1. Go to the file you want to run and uncomment the line containing `func main() {`
2. Keep it commented for other files if all files are in the same directory.

# Go Modules vs Packages

## 1. Package

- **Smallest unit of code organization** in Go.
- A **package** = one directory containing `.go` files **with the same `package <name>`** declaration (where `name` is the path relative to the base path defined in `go.mod` using `module` keyword.)
- Conventionally directory name and package name should be same.
- The functions to be exported from the package should _start with a CAPITAL LETTER_
- Two types:
  - **Executable** → `package main` (must have `func main()`).
  - **Library** → any other name (meant to be imported).
- Example:  
  ~ utils/  
  strings.go <!-- package utils  
  math.go <!-- package utils

## 2. Module

- **Higher-level unit**: a versioned collection of packages.
- Defined by a `go.mod` file at the module root.
- `go.mod` specifies:
- **Module path** (base import path)
- Go version
- Dependencies and versions
- Example (`go.mod`):
  ```go
  module github.com/sanavi/myproject
  go 1.22
  ```

## 3. Relationship Between Modules and Packages

- A **module** can contain **multiple packages** in subdirectories.
- Import path = **`module path` + `/` + `relative folder path`**.
- Example:
  ```go
  import github.com/sanavi/myproject/utils
  import github.com/sanavi/myproject/services
  ```

## 4. Downloading and Using Packages

- You **cannot** download just one package from a module
- When you import/download a package using `go get`:

```go
import "github.com/sanavi/myproject/utils"
```

- Go:
  - Downloads the entire module containing that package (if not cached)
  - Stores it in `$GOPATH/pkg/mod`
  - Compiles only the imported package and its dependencies into your project's executable binary.
  - The executable binary of the project's dependencies is not stored locally in the `$GOPATH/pkg/mod` cache, but present in the project's executable on `go build` command
  - `go build` or `go run` will fetch and compile them automatically

# Go General

## `go.mod` vs `go.sum`

- A demo `go.mod` file is added as `go.mod.demo` in this repo for reference
- `go.sum`
  - Contains cryptographic checksums (hashes) of the exact versions of all dependencies (direct + indirect) used in your project.
  - This is auto generated on `go get`, `go mod tidy`, `go build` commands
  - Demo `go.sum` file:
  ```go
  github.com/gorilla/mux v1.8.0 h1:sm8O1C8H…
  github.com/gorilla/mux v1.8.0/go.mod h1:bnxGCyR…
  golang.org/x/net v0.10.0 h1:sda9HbdG…
  golang.org/x/net v0.10.0/go.mod h1:zC4j0sP…
  ```
  - Each line = module path + version + checksum.
  - Two entries per module (one for the module, one for its go.mod).

## ENV Variables

1. `$GOBIN` - Defines where go install puts binaries. Check using ` go env GOBIN` in terminal. Set it using the shell config (.bashrc or .zshrc) or exporting the variables using `export GOBIN=<path>`
2. `GOPATH/bin` - Fallback for $GOBIN. Check using `go env GOPATH` in terminal. (Default)
3. `$GOPATH/pkg/mod` - It is the module cache

- Keeps one copy of every version of every module fetched.
- Used by all your projects (saves bandwidth and disk).
- Ensures builds are reproducible — same version, same code, same checksum.

## `go` commands

1. `go install`

- Compiles, builds and installs binary (an executable file) globally in the local system in `$GOBIN` if set, else `$GOPATH/bin`
- Must have `package main`
  - Without arguments command executes on the current directory
  - With arguments as `go install <module-path>@<version>` for remote project installed as a CLI tool

2. `go get <remote_module_url>`

- Used to update library dependencies by modifying `go.mod` and `go.sum`
- Doesn't install binary/executable since Go1.16 version
- Downloads the module source into the local module cache `$GOPATH/pkg/mod`.
- Makes it available for compilation.

3. `go build` - Compiles the program and creates an executable file in the current directory but for packages, it doesn't create them, however it is useful for debugging and testing. (called using `./project_name`)
4. `go run` - Compiles and runs the program in one step, useful for quick testing
5. `go mod tidy` - Keeps your module dependencies clean and correct.

- Adds missing entries → Ensures go.mod/go.sum include everything your code actually imports.
- Removes unused entries → Cleans out modules not used anywhere in your project.

# Interesting Observation

> How `cp` command is superior than `Ctrl+C` and `Ctrl+V` for copying files in Unix-like systems:

- I generated the compiled code using `go build -o hello_world 1.basics.go`
- Can run the program by using `./hello_world` command on Unix-like systems.
- I did Ctrl+C and Ctrl+V of `hello_world` to a new file `binary`.
- Made it executable using `chmod +x binary`. But it dodn't run as expected.
- So I used `cp hello_world binary` to copy the file and it run as expected after making it executable.

# Interesting Fact about GO

> The Go programming language is self-hosted, meaning its compiler and runtime are primarily written in Go itself.

- Initially, the Go compiler and runtime were developed using C and C++.
- However, this changed with Go version 1.5, where the C code for the compiler was transitioned to Go, and the C++ code for the runtime was replaced with Go in version 1.4.
- This process is known as bootstrapping, where a language is used to compile itself.

# Differences

## Concurrency vs Parallelism

> In terms of OS vocabulary, concurrency is dealing with multiple processes at once whereas parallelism is executing a single process using multiple threads for high efficiency
> Watch Rob Pike - Concurrency vs Parallelism

1. Concurrency

- Multiple tasks are in progress at the same time conceptually.
- The CPU may switch between them quickly (time-slicing) but they don’t necessarily run at the same physical instant.
- It’s about dealing with many things at once (interleaving).

2. Parallelism

- Multiple tasks are executed literally at the same time.
- Requires multiple cores/CPUs.

## Synchronous vs Sequential

1. Synchronous

- An operation where tasks wait for each other to finish before moving on.
- Each step is completed in order — if one is blocked, the rest wait.

2. Sequential

- Tasks are executed one after another in a fixed order.
- All sequential programs are synchronous, but not all synchronous operations need to be sequential (you could wait for multiple things at once in sync).

# GO Proverbs

(When creating your own module, keep as less as exports possible as you would have to maintain those many exports. To keep the module stable do not update names of modules/functions often.)

> To know about GO Proverbs, watch **Rob Pike's - Gopherfest 2015**.

- Don't communicate by sharing memory, share memory by communicating.
- Concurrency is not parallelism.
- Channels orchestrate; mutexes serialize.
- The bigger the interface, the weaker the abstraction.
- Make the zero value useful.
- interface{} says nothing.
- Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite.
- A little copying is better than a little dependency.
- Syscall must always be guarded with build tags.
- Cgo must always be guarded with build tags.
- Cgo is not Go.
- With the unsafe package there are no guarantees.
- Clear is better than clever.
- Reflection is never clear.
- Errors are values.
- Don’t just check errors, handle them gracefully.
- Design the architecture, name the components, document the details.
- Documentation is for users.
- Don’t panic.
