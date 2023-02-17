package replication

import "github.com/lemon-mint/super-system/replication/protocol"

type MessageBus interface {
	SendMessageToClient(clientID uint64, message *protocol.Message)
	SendMessageToReplica(replicaID uint64, message *protocol.Message)
	BroadcastMessageToReplicas(message *protocol.Message)
}
