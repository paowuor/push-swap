# Phase 4 — Sorting Algorithms

## Overview

The goal of Phase 4 was to implement the core sorting logic
for the push-swap program.

At this stage, the project evolved from a stack operation engine
into a fully functional sorting program capable of generating
valid instruction sequences for multiple input sizes.

This phase introduced:

- small stack sorting
- coordinate compression
- radix sort
- algorithm selection
- operation optimization strategies

---

# Objectives

The main objectives of this phase were:

- implement optimized sorting for small inputs
- implement scalable sorting for large inputs
- minimize operation count
- normalize values for radix sort
- dynamically select sorting strategies
- integrate sorting algorithms with stack operations

---

# Algorithm Selection Strategy

Different algorithms were implemented depending on input size.

| Input Size | Algorithm |
|------------|------------|
| 2 | sort_two |
| 3 | sort_three |
| 4–5 | sort_five |
| >5 | radix sort |

This hybrid strategy improves efficiency because:

- small datasets benefit from hardcoded optimal moves
- large datasets benefit from scalable algorithms

---

# Why Radix Sort Was Chosen

Radix sort was selected because it performs consistently
well under push-swap constraints and avoids the complexity
of comparison-based algorithms within stack-only operations.

Advantages:

- predictable performance
- efficient with stack operations
- scalable for large datasets
- simple implementation using binary representation

Complexity:

| Algorithm | Complexity |
|------------|------------|
| Radix Sort | O(n × bits) |

For push-swap, radix sort is one of the most reliable
approaches for achieving acceptable move counts.

---

# Coordinate Compression

## Problem

Radix sort works best with non-negative sequential values.

However, push-swap inputs may contain:

- negative numbers
- large values
- irregular ranges

Example:

```text
-500 1000 42
```

Using radix sort directly on such values is inefficient.

---

# Solution

Coordinate compression (normalization) was implemented.

Example:

Original:

```text
[-5, 100, 2]
```

Sorted copy:

```text
[-5, 2, 100]
```

Assigned indexes:

```text
-5  -> 0
2   -> 1
100 -> 2
```

Normalized result:

```text
[0, 2, 1]
```

---

# Benefits Of Normalization

Normalization provides:

- sequential integer ranges
- reduced radix complexity
- simplified bit operations
- predictable sorting behavior

---

# sort_two Implementation

## File

```text
internal/algorithms/sort_two.go
```

## Strategy

Sorting two numbers requires only one condition.

Example:

Before:

```text
2 1
```

Operation:

```text
sa
```

After:

```text
1 2
```

Complexity:

```text
O(1)
```

---

# sort_three Implementation

## File

```text
internal/algorithms/sort_three.go
```

## Strategy

The algorithm handles all six permutations of three numbers.

Possible states:

| Input | Operations |
|-------|-------------|
| 1 2 3 | none |
| 1 3 2 | sa + ra |
| 2 1 3 | sa |
| 2 3 1 | rra |
| 3 1 2 | ra |
| 3 2 1 | sa + rra |

This implementation guarantees optimal or near-optimal moves.

---

# Example

Input:

```text
3 2 1
```

Operations:

```text
sa
rra
```

Result:

```text
1 2 3
```

---

# sort_five Implementation

## File

```text
internal/algorithms/sort_five.go
```

## Strategy

The algorithm sorts 4–5 numbers using the following process:

1. locate the minimum value
2. move minimum to top
3. push minimum to stack B
4. repeat until only 3 remain
5. sort remaining 3 numbers
6. restore values from stack B

---

# Example

Input:

```text
5 1 4 2 3
```

Possible operations:

```text
ra
pb
rra
pb
sa
ra
pa
pa
```

Final result:

```text
1 2 3 4 5
```

---

# Why This Strategy Works

Sorting 3 numbers optimally is easy.

Reducing larger small sets into a sortable 3-element subset
significantly simplifies logic.

This is a common push-swap optimization technique.

---

# Radix Sort Implementation

## File

```text
internal/algorithms/radix.go
```

## Core Concept

Radix sort processes numbers bit by bit.

Example binary representation:

```text
0 = 000
1 = 001
2 = 010
3 = 011
```

The algorithm processes:

```text
bit 0
bit 1
bit 2
```

---

# Radix Sorting Process

For each bit:

- numbers with bit = 0 are pushed to stack B
- numbers with bit = 1 are rotated inside stack A
- all numbers are restored back to A

This repeats until all bits are processed.

---

# Example

Normalized input:

```text
[2, 0, 1]
```

Binary:

```text
2 = 10
0 = 00
1 = 01
```

Processing bit 0:

```text
pb
pb
ra
pa
pa
```

The stack gradually becomes sorted.

---

# Bit Calculation

The number of required radix passes depends on:

```text
maximum number of bits
```

Example:

```text
100 -> 1100100
```

Requires:

```text
7 bits
```

Therefore:

```text
7 radix passes
```

---

# Operation Recording

All algorithms use shared stack operations that automatically
record executed instructions.

Example:

```go
stack.Pb(a, b, ops)
```

This allows:

- instruction generation
- checker validation
- reusable execution logic

without duplicated code.

---

# Algorithm Dispatcher

The push-swap executable dynamically selects algorithms.

Example logic:

```go
if size == 2
```

```go
else if size == 3
```

```go
else if size <= 5
```

```go
else radix sort
```

This ensures the most efficient algorithm is chosen.

---

# Testing Performed

## sort_two

Command:

```bash
./push-swap 2 1
```

Expected:

```text
sa
```

---

## sort_three

Command:

```bash
./push-swap 3 2 1
```

Expected:

```text
sa
rra
```

---

## checker Integration

Command:

```bash
ARG="3 2 1"; ./push-swap $ARG | ./checker $ARG
```

Expected:

```text
OK
```

---

## Five Numbers

Command:

```bash
ARG="5 1 4 2 3"; ./push-swap $ARG | ./checker $ARG
```

Expected:

```text
OK
```

---

## Large Input

Command:

```bash
ARG=$(shuf -i 1-100 -n 100)
./push-swap $ARG | ./checker $ARG
```

Expected:

```text
OK
```

---

# Performance Benchmarks

## 100 Numbers

Command:

```bash
ARG=$(shuf -i 1-100 -n 100)
./push-swap $ARG | wc -l
```

Target:

```text
< 700 operations
```

---

## 500 Numbers

Command:

```bash
ARG=$(shuf -i 1-500 -n 500)
./push-swap $ARG | wc -l
```

Target:

```text
< 5500 operations
```

---

# Challenges Encountered

## Coordinate Compression

One challenge was adapting radix sort to arbitrary integers.

This was solved using normalization, which converts values
into sequential indexes before sorting.

---

## Small Stack Optimization

Radix sort is inefficient for very small datasets.

This was solved by implementing specialized algorithms for:

- 2 numbers
- 3 numbers
- 5 numbers

which significantly reduces operation counts.

---

## Operation Explosion

Without optimization, unnecessary rotations and pushes
greatly increase instruction count.

Specialized small sorting algorithms were implemented
to reduce this issue.

---

# Results

At the end of Phase 4:

- push-swap became fully functional
- small stack optimization works correctly
- radix sort handles large datasets
- coordinate compression works correctly
- checker integration succeeds
- operation generation is stable
- sorting validation passes

The project is now capable of sorting:

- small datasets efficiently
- large datasets reliably

using only stack operations.

---

# Next Phase

Phase 5 will focus on:

- reducing operation counts
- combined operations optimization
- redundant instruction removal
- benchmark improvements
- automated testing
- README completion
- final project polishing
- performance tuning
- evaluation preparation