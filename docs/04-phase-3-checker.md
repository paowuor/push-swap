# Phase 3 — Checker Implementation

## Overview

The goal of Phase 3 was to implement the checker executable.

The checker validates whether a sequence of push-swap
instructions correctly sorts a stack.

---

# Responsibilities

The checker program was designed to:

- parse and validate arguments
- read instructions from stdin
- validate instructions
- execute operations
- verify final stack state
- output OK or KO

---

# Architecture

The checker was divided into multiple responsibilities:

| Component | Responsibility |
|-----------|----------------|
| reader.go | stdin reading |
| validator.go | instruction validation |
| execute.go | operation execution |
| helpers.go | sorted verification |

This separation improves maintainability and debugging.

---

# Instruction Reading

Instructions are read using:

```go
bufio.Scanner
```

Each instruction is stored inside a slice.

Example:

```text
sa
pb
rra
```

becomes:

```go
[]string{"sa", "pb", "rra"}
```

---

# Instruction Validation

A validation system was implemented to reject invalid instructions.

Allowed instructions:

- sa
- sb
- ss
- pa
- pb
- ra
- rb
- rr
- rra
- rrb
- rrr

Invalid instructions produce:

```text
Error
```

---

# Execution Engine

The checker uses the centralized execution dispatcher:

```go
Execute(op, &a, &b)
```

This prevents duplicated logic between:

- checker
- push-swap
- testing

---

# Final Verification

After all instructions are executed, the checker verifies:

- stack A is sorted
- stack B is empty

If true:

```text
OK
```

Otherwise:

```text
KO
```

---

# Example Execution

Initial stack:

```text
3 2 1
```

Instructions:

```text
sa
rra
```

Execution:

After sa:

```text
2 3 1
```

After rra:

```text
1 2 3
```

Final result:

```text
OK
```

---

# Edge Cases Tested

## Empty Input

```bash
./checker
```

Expected:

```text
(no output)
```

---

## Invalid Instruction

```bash
echo "invalid" | ./checker 3 2 1
```

Expected:

```text
Error
```

---

## Duplicate Values

```bash
./checker 1 2 2
```

Expected:

```text
Error
```

---

# Challenges

## Instruction Validation

One challenge was ensuring invalid instructions were rejected
before execution.

This was solved using a centralized validation map.

---

## Shared Logic

Another challenge was preventing duplicated operation logic
between push-swap and checker.

This was solved by implementing a shared execution engine.

---

# Results

At the end of Phase 3:

- checker fully validates instructions
- stdin reading works correctly
- stack operations execute properly
- invalid instructions are handled safely
- final stack verification works correctly

The project is now ready for:

- sorting algorithms
- small stack optimizations
- radix sort implementation
