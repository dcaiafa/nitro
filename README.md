
# TODO
- Backtick string literal
- Dedicated type for Enumerator, to get rid of all this code accepting Closure
    assuming that it is an enumerator.
- Expose MachineFromContext, and remove context.Context from all machine
    wrappers (e.g. Next). Look at lib/util.go AsReader where it has to store
    ctx.
- Figure out `return;` The ; is currently needed when returning no values.
  Special syntax? `return nothing` 

# Open Questions
- `match(str, re)`: is the order confusing?
- `match`: is muti-return semantics confusing?


