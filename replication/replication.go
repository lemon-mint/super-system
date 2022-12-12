package replication

// SuperSystem Uses Viewstamped Replication based protocol to replicate data across multiple nodes
//
// Terminology:
// View Number: A monotonically increasing number that is incremented every time a new view is installed
// Commit Number: A monotonically increasing number that is incremented every time a new commit is made
// Snapshot: A snapshot of the state of the system at a particular commit number
// Log: A log of all the operations that have been performed on the system since the last snapshot
// Quorum: A set of nodes that can be used to make a decision
// View Change: A process of changing the current view to a new view
// Reconfiguration: A process of changing cluster membership
//
// # Implementation Detail
//
// ## High Throughput Mode
//
// - Leader Batch Processes Requests received in 50ms.
//
// ## Low Latency Mode
//
// - Leader Processes Requests in Realtime.
//
// ## Replication Group Sharding
//
// SuperSystem Use Replication Group Sharding to Provide Scalability
//
// - Every Replication Group has a unique ID (Replication Group ID)
// - All Packets sent to a Replication Group are tagged with the Replication Group ID
//
// # Normal Operation
//
// 1. A client sends a <Propose> to the primary
//
// Propose [
//	ClientID: The ID of the client.
//	SequenceNumber: The client's sequence number.
//	Operation: The operation to be performed.
// ]
//
// 2. When the primary node receives a <Propose>,
// IF SequenceNumber is not greater than the primary's sequence number
// THEN the primary node sends a <Reject> to the client
//
// Reject [
//	ClientID: The ID of the client.
//	SequenceNumber: The client's sequence number.
//  Reason: The reason for the rejection.
// ]
//
// ELSE
// 3. The primary node sets the primary's sequence number to SequenceNumber.
// 4. OpNumber = OpNumber + 1
// 5. n = OpNumber
// 6. The primary node adds the <Propose> to the primary's log.
// 7. The primary node sends a <Prepare> to all the other nodes in the cluster.
//
// Prepare [
//  ViewNumber: The view number.
//  OpNumber: The operation number.
//  Operation: The operation to be performed.
//  CommitNumber: The commit number.
// ]
//
// 8. When the primary node receives a <PrepareOK> from a quorum of nodes
