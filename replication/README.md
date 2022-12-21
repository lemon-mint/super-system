SuperSystem Uses Viewstamped Replication based protocol to replicate data across multiple nodes

# Terminology

- View Number: A monotonically increasing number that is incremented every time a new view is installed
- Commit Number: A monotonically increasing number that is incremented every time a new commit is made
- Snapshot: A snapshot of the state of the system at a particular commit number
- Log: A log of all the operations that have been performed on the system since the last snapshot
- Quorum: A set of nodes that can be used to make a decision
- View Change: A process of changing the current view to a new view
- Reconfiguration: A process of changing cluster membership

# Implementation Detail

## Consistency Models

SuperSystem can operate with precise clocks (like Atomic Clocks) or without them.

without precise clocks, the Consistency level is limited.

### Linearizable Mode

When Precise Clocks are available (like Atomic Clocks), Linearizable Consistency is available.

When Replication Group Sharding is disabled, SuperSystem uses a single replication group for all data, in this case, Linearizable Consistency is available even if Precise Clocks are not available.

When Linearizable Mode is enabled, SuperSystem waits for clock boundaries to ensure that all operations are ordered in the same order on all nodes.

## High Throughput Mode

- Leader Batch Processes Requests received in 50ms.

## Low Latency Mode

- Leader Processes Requests in Realtime.

## Replication Group Sharding

SuperSystem Use Replication Group Sharding to Provide Scalability
- Every Replication Group has a unique ID (Replication Group ID)
- All Packets sent to a Replication Group are tagged with the Replication Group ID

# Normal Operation

## On Propose

1. A client sends a Propose to the primary
```
Propose [
	ClientID: The ID of the client.
    SequenceNumber: The client's sequence number.
    Operation: The operation to be performed.
]
```

2. When the primary node receives a Propose,
IF SequenceNumber is not greater than the primary's sequence number
THEN the primary node sends a Reject to the client
```
Reject [
    ClientID: The ID of the client.
    SequenceNumber: The client's sequence number.
    Reason: The reason for the rejection.
]
```
ELSE

3. The primary node sets the primary sequence number to SequenceNumber.

4. OperationNumber++

5. n = OperationNumber

6. The primary node adds the Propose to the primary's log.

7. The primary node sends a Prepare to all the other nodes in the cluster.

```
Prepare [
 ViewNumber: The view number.
 OperationNumber: The operation number.
 Operation: The operation to be performed.
 CommitNumber: The commit number.
]
```

8. When the primary node receives a PrepareOK from a quorum of nodes, it commits the operation.

9. CommitNumber = n

## On Prepare

1. When a node receives a Prepare, it checks if the Prepare is valid.

2. If Prepare.ViewNumber != ViewNumber, the node ignores the Prepare.

3. If Prepare.OperationNumber < CommitNumber, the node ignores the Prepare.
