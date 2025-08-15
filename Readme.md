# Go Modules vs Packages — Complete Guide

## 1. Package

- **Smallest unit of code organization** in Go.
- A **package** = one directory containing `.go` files **with the same `package <name>`** declaration (where `name` is the path relative to the base path defined in `go.mod` using `module` keyword.)
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

- You **cannot** download just one package from a module.
- When you import a package:

```go
import "github.com/sanavi/myproject/utils"
```

- Go:
  - Downloads the entire module containing that package (if not cached).
  - Stores it in $GOPATH/pkg/mod.
  - Compiles only the imported package and its dependencies into your binary.

## 5. `go install` vs `import`

- For libraries (like utils):

  - Just import them in your code.
  - `go build` or `go run` will fetch and compile them automatically.

- For executables (CLI tools):
  - Use `go install <module-path>@<version>` to build and place the binary locally.

# `go` commands

1. `go install` without arguments compiles the current directory project and creates an executable file globally in the local system. (called using `project_name`)
2. `go build` - Compiles the program and creates an executable file in the current directory but for packages, it doesn't create them, however it is useful for debugging and testing. (called using `./project_name`)
3. `go run` - Compiles and runs the program in one step, useful for quick testing

# How to run the go files

1. Go to the file you want to run and uncomment the line containing `func main() {`
2. Keep it commented for other files if all files are in the same directory.

# Interesting Observation

> How `cp` command is superior than `Ctrl+C` and `Ctrl+V` for copying files in Unix-like systems:

- I generated the compiled code using `go build -o hello_world 1.basics.go`
- Can run the program by using `./hello_world` command on Unix-like systems.
- I did Ctrl+C and Ctrl+V of `hello_world` to a new file `binary`.
- Made it executable using `chmod +x binary`. But it dodn't run as expected.
- So I used `cp hello_world binary` to copy the file and it run as expected after making it executable.
