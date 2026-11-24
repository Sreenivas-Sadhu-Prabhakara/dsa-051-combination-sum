# Combination Sum

> Striver SDE Sheet 051 · Day 9 Recursion — backtracking where staying on the same index encodes unlimited reuse; sortedness gives the cutoff. Go, stdlib only, tested.

## Problem

Find all unique combinations of candidate numbers summing to a target, where each candidate may be used any number of times. `[2,3,6,7]` with target 7 gives `[2,2,3]` and `[7]`.

## Approach

Backtrack with a start index and remaining target. The whole semantics of "unlimited reuse without duplicate combinations" lives in one character: recursing with `i` (not `i+1`) lets the same candidate repeat, while never looking *left* of `start` ensures each combination is generated in non-decreasing order exactly once. Sorting the candidates first turns "this candidate exceeds the remainder" into a loop break that prunes the whole tail, not just one branch.

**Complexity:** exponential in the worst case (output-bound), O(target/min) recursion depth.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
