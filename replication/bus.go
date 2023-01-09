package replication

import "v8.run/go/supersystem/replication/vsrproto"

type MessageBus interface {
	SendMessageToClient(clientID uint64, message *vsrproto.Message)
	SendMessageToReplica(replicaID uint64, message *vsrproto.Message)
	BroadcastMessageToReplicas(message *vsrproto.Message)
}
