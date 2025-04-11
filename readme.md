# Go Workspace Example

This repository provides a basic example demonstrating how to set up and use a [Go workspace](https://go.dev/doc/tutorial/workspaces). Go workspaces allow you to work on multiple Go modules simultaneously without needing to edit `go.mod` files for each module or rely on `replace` directives, making local development across related modules much easier.

This example includes two modules: `module1` and `module2`.

## Prerequisites

* **Go 1.18 or later**: Workspaces were introduced in Go 1.18. You can check your version with `go version`.

## Directory Structure

The expected directory structure for this example setup is:

/workspace              # Your top-level project directory
├── go.work             # Defines the workspace and included modules
├── module1/
│   ├── go.mod          # Defines module1
│   └── ...             # Source files for module1 (e.g., main.go, packages)
└── module2/
├── go.mod          # Defines module2
└── ...             # Source files for module2 (e.g., packages)

*(For a concrete example, imagine `module1` contains a `main.go` and `module2` contains a utility package used by `module1`)*

## Setup Instructions

Follow these steps to create and configure the Go workspace.

1.  **Create Project Structure and Modules:**
    If you haven't already, create the directories and initialize the Go modules. Replace `example.com/module1` and `example.com/module2` with your actual module paths.

    ```bash
    # Create the main workspace directory and navigate into it
    mkdir workspace
    cd workspace

    # Create module1 and initialize its go.mod
    mkdir module1
    cd module1
    # Add some Go code here (e.g., a simple main.go)
    # touch main.go
    go mod init [example.com/module1](https://example.com/module1)
    cd ..

    # Create module2 and initialize its go.mod
    mkdir module2
    cd module2
    # Add some Go code here (e.g., a utility package)
    # mkdir util
    # touch util/util.go
    go mod init [example.com/module2](https://example.com/module2)
    cd ..
    ```

2.  **Initialize the Go Workspace:**
    From the root `workspace` directory, initialize the workspace, telling Go which modules it should contain:

    ```bash
    # Ensure you are in the 'workspace' directory
    pwd # Should output /path/to/your/workspace

    # Initialize the workspace with the modules
    go work init ./module1 ./module2
    ```
    This command creates the `go.work` file in the current directory (`workspace`).

3.  **Review the `go.work` File:**
    The generated `go.work` file defines the Go version and lists the modules included in the workspace. It will look similar to this:

    ```go
    // Example go.work content
    go 1.23 // Your Go version might differ

    use (
        ./module1
        ./module2
    )
    ```

4.  **Add More Modules (Optional):**
    If you create another module (e.g., `module3`) later and want to include it in the workspace, navigate to the workspace root directory and use the `go work use` command:

    ```bash
    # Assuming you are in the 'workspace' directory
    # mkdir module3
    # (cd module3 && go mod init [example.com/module3](https://example.com/module3))

    go work use ./module3
    ```
    This command appends `./module3` to the `use` directive in your `go.work` file.

## Usage

The primary benefit of a workspace is that Go commands run within the workspace (like `go build`, `go run`, `go test`, `go list`, `go vet`) operate in the context of all modules defined in `go.work`.

* **Dependency Resolution**: Go commands will resolve dependencies using the modules specified in `go.work` first. If `module1` imports a package from `module2`, Go will use the local `module2` code found via `./module2` in `go.work`, rather than trying to download it based on `module1`'s `go.mod`.
* **Running Code**: You can build or run code from any module *while inside the workspace directory* or any subdirectory within it.

    ```bash
    # Example: Run the main package of module1
    # Assumes module1 has a main package and its path is [example.com/module1](https://example.com/module1)

    # From the workspace root
    cd /path/to/workspace
    go run [example.com/module1](https://example.com/module1)

    # Or, from within module1's directory
    cd /path/to/workspace/module1
    go run .
    ```

## Useful `go work` Commands

* `go work init [mod-dirs...]`: Creates a `go.work` file including the specified module directories.
* `go work use [-r] [mod-dirs...]`: Adds module directories to the `go.work` file. The `-r` flag recursively searches directories for modules.
* `go work edit`: Provides a command-line interface for editing `go.work` (e.g., adding/removing `use` directives or `replace` directives specific to the workspace).
* `go work sync`: Syncs dependencies from workspace modules' `go.mod` files into the `go.work` file (less commonly needed for basic local development).
