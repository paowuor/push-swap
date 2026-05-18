# Phase 1 — Foundation & Project Setup

## Overview

The goal of Phase 1 was to establish the foundational architecture
of the push-swap project and implement reliable input parsing,
validation, and error handling.

At this stage, no sorting algorithms or stack operations were implemented.
The focus was entirely on building a stable and maintainable base
for future development phases.

---

# Objectives

The main objectives of this phase were:

- Initialize the Go project
- Design the project architecture
- Create executable entry points
- Implement stack data structures
- Parse command-line arguments
- Validate input data
- Detect duplicate values
- Handle integer overflow
- Detect already sorted input
- Implement standardized error handling

---

# Project Structure

The project structure was designed to follow good Go engineering practices.

```text
paowuor-push-swap/
│
├── cmd/
│   ├── checker/
│   │   └── main.go
│   │
│   └── push-swap/
│       └── main.go
│
├── internal/
│   ├── models/
│   │   └── stack.go
│   │
│   ├── parser/
│   │   ├── parser.go
│   │   └── validation.go
│   │
│   └── utils/
│       ├── errors.go
│       └── helpers.go
│
├── tests/
├── docs/
├── go.mod
├── README.md
└── Makefile
```

---

# Architecture Decisions

## Why `cmd/` Was Used

Initially, the project used directories named:

```text
push-swap/
checker/
```

However, this caused a naming conflict during compilation because
Go attempted to generate executables with the same names.

For example:

```bash
go build -o push-swap ./push-swap
```

could not generate the executable because a directory named
`push-swap` already existed.

The project structure was therefore refactored into:

```text
cmd/push-swap/
cmd/checker/
```

This follows standard Go project conventions and clearly separates:

- executable entry points
- reusable internal packages

This structure is more scalable and maintainable.

---

# Stack Model

## File

```text
internal/models/stack.go
```

## Implementation

```go
type Stack struct {
    Values []int
}
```

## Design Decision

Slices were chosen because they:

- dynamically resize automatically
- simplify stack manipulation
- provide efficient indexing
- integrate naturally with Go operations

Using slices also keeps the implementation lightweight and easy to debug.

---

# Argument Parsing

## File

```text
internal/parser/parser.go
```

## Responsibilities

The parser was designed to:

- support multiple input formats
- convert strings into integers
- detect invalid integers
- detect overflow
- return validated integer slices

---

# Supported Input Formats

The parser supports both:

```bash
./push-swap "1 2 3"
```

and:

```bash
./push-swap 1 2 3
```

This was achieved using:

```go
strings.Join(args, " ")
strings.Fields(joined)
```

---

# Why `strings.Fields()` Was Used

`strings.Fields()` automatically handles:

- multiple spaces
- tabs
- irregular formatting

Example:

```text
"1     2      3"
```

This improves parser reliability.

---

# Integer Conversion

## Why `strconv.ParseInt()` Was Used

Instead of using:

```go
strconv.Atoi()
```

the project uses:

```go
strconv.ParseInt()
```

because it provides safer overflow handling.

This allows explicit validation against:

```go
math.MinInt32
math.MaxInt32
```

which ensures compatibility with project requirements.

---

# Duplicate Detection

## File

```text
internal/parser/validation.go
```

## Implementation Strategy

Duplicate detection was implemented using a hash map.

Example logic:

```go
seen := make(map[int]bool)
```

## Why a Hash Map Was Chosen

Using a hash map reduces complexity from:

| Method | Complexity |
|--------|------------|
| Nested loops | O(n²) |
| Hash map | O(n) |

This significantly improves efficiency.

---

# Error Handling

## File

```text
internal/utils/errors.go
```

## Responsibilities

A centralized error utility was created to standardize all project errors.

Example:

```go
func PrintError() {
    fmt.Fprintln(os.Stderr, "Error")
}
```

---

# Invalid Cases Handled

The parser correctly detects:

- invalid integers
- duplicates
- integer overflow
- malformed input

Examples:

```bash
./push-swap 1 two 3
```

```bash
./push-swap 1 2 2
```

```bash
./push-swap 999999999999999999
```

All produce:

```text
Error
```

---

# Sorted Detection

## File

```text
internal/utils/helpers.go
```

## Purpose

The helper function:

```go
IsSorted()
```

checks whether input is already sorted.

Example:

```bash
./push-swap 1 2 3
```

produces no output, as required by the project specification.

---

# Executable Entry Points

## push-swap

File:

```text
cmd/push-swap/main.go
```

Responsibilities:

- parse arguments
- validate input
- detect sorted input

---

## checker

File:

```text
cmd/checker/main.go
```

Responsibilities:

- parse arguments
- validate input

Checker logic itself will be implemented in later phases.

---

# Build System

## Makefile

A Makefile was created to simplify:

- compilation
- cleaning binaries
- rebuilding
- testing

Example commands:

```bash
make build
make clean
make re
```

---

# Testing Performed

The following cases were tested successfully.

---

## Valid Input

```bash
./push-swap 3 2 1
```

Expected:

```text
(no output)
```

---

## Duplicate Detection

```bash
./push-swap 1 2 2
```

Expected:

```text
Error
```

---

## Invalid Integer

```bash
./push-swap 1 two 3
```

Expected:

```text
Error
```

---

## Overflow

```bash
./push-swap 999999999999999999999
```

Expected:

```text
Error
```

---

## Already Sorted

```bash
./push-swap 1 2 3
```

Expected:

```text
(no output)
```

---

# Challenges Encountered

## Executable Naming Conflict

One major issue encountered was a naming conflict between:

- source directories
- generated executables

Example:

```text
push-swap/
```

directory conflicted with:

```text
push-swap
```

binary executable.

This issue was solved by restructuring the project into:

```text
cmd/push-swap/
cmd/checker/
```

which follows standard Go project conventions.

---

# Results

At the end of Phase 1:

- project architecture was completed
- both executables compile successfully
- argument parsing works correctly
- invalid input is detected
- duplicates are handled safely
- overflow detection works
- sorted input detection works
- foundational utilities are stable

The project is now ready for:

- stack operations
- instruction execution
- checker implementation

which will be developed during Phase 2.

---

# Next Phase

Phase 2 will focus on implementing:

- push operations
- swap operations
- rotate operations
- reverse rotate operations
- instruction execution system

These operations will become the core mechanics of the project.