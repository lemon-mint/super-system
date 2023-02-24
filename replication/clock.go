package replication

type Clock struct {
	TickCounter uint64
}

func (c *Clock) Tick() uint64 {
	c.TickCounter++
	return c.TickCounter
}
