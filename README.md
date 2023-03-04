# SuperSystem

Distributed System Tools ( replication, membership, gossip )

## Requirements

- gobe
- stringer

```shell
go install github.com/lemon-mint/gobe@latest
go install golang.org/x/tools/cmd/stringer@latest
``` 

## Design Rules

- Prefer consistency and correctness over performance.
- Allow hardware failures.
- Quickly evict failed nodes.
