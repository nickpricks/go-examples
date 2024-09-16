
# Go Workspace Example

  

This repository demonstrates how to set up a Go workspace with multiple modules. The workspace includes `module1` and `module2`, each with its own packages.

  

## Directory Structure

  

/workspace ├── go.work ├── module1 │ ├── go.mod │ ├── mod1Pkg1 │ └── mod1Pkg2 └── module2 ├── go.mod └── mod2Pkg1

  

## Setting Up the Workspace

  

1.  **Initialize the Workspace**:

Navigate to the workspace directory and initialize the workspace with the modules.

  

```sh```

```cd /path/to/your/workspace```

```go work init ./module1 ./module2```

  

2.  **Add More Modules: If you need to add more modules later, use the go work use command**:

```go work use ./module3```

  
  

3. **go.work File**

The go.work file manages the workspace and includes the paths to the modules.

  
```
go 1.23.1

use ./module1

use ./module2
```