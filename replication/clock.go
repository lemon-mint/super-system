package replication

import "time"

type Clock struct {
	TickCounter uint64

	HLC_WALL uint64
	HLC_VIRT uint64
}

func (c *Clock) Tick() uint64 {
	c.TickCounter++

	now := uint64(time.Now().UnixNano())
	if now > c.HLC_WALL {
		c.HLC_WALL = now
		c.HLC_VIRT = 0
	} else {
		c.HLC_VIRT++
	}

	return c.TickCounter
}

func (c *Clock) NowTicks() uint64 {
	return c.TickCounter
}

func (c *Clock) NowHLC() (uint64, uint64) {
	return c.HLC_WALL, c.HLC_VIRT
}
