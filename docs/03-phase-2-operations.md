# Phase 2 — Stack Operations

## Overview

The goal of Phase 2 was to implement all mandatory
push-swap stack operations and create a reusable
instruction execution system.

---

# Operations Implemented

## Swap Operations

- sa
- sb
- ss

## Push Operations

- pa
- pb

## Rotate Operations

- ra
- rb
- rr

## Reverse Rotate Operations

- rra
- rrb
- rrr

---

# Architecture Decision

Operations were designed to optionally record instructions.

Example:

```go
func Sa(a *Stack, ops *[]string)
```

This allows the same operations to work for:

- push-swap (record operations)
- checker (silent execution)

This prevents code duplication.

---

# Push Operation Logic

Example:

Before pb:

Stack A:

3
2
1

Stack B:

(empty)

After pb:

Stack A:

2
1

Stack B:

3

---

# Rotate Operation Logic

Example:

Before ra:

1
2
3

After ra:

2
3
1

---

# Reverse Rotate Logic

Example:

Before rra:

1
2
3

After rra:

3
1
2

---

# Instruction Executor

An instruction dispatcher was implemented using
a switch statement.

Example:

```go
Execute("sa", &a, &b)
```

This provides a centralized execution engine for:

- checker
- testing
- future optimizations

---

# Challenges

## Double Recording Operations

Combined operations such as:

- ss
- rr
- rrr

initially caused duplicate recordings.

Example:

sa
sb
ss

This was solved by executing internal operations with:

```go
nil
```

instead of passing the operation recorder.

---

# Result

At the end of Phase 2:

- all stack operations were implemented
- operations safely handle empty stacks
- instruction execution works
- operations can record instructions
- reusable architecture was established

The project is now ready for:

- checker implementation
- stdin instruction processing
- operation validation