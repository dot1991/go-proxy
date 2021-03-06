package client

import "proxy/network"

type PongPacket struct {
	Serial int32
	Time   int32
}

func (pp *PongPacket) Read(p *network.GamePacket) {
	pp.Serial = p.ReadInt32()
	pp.Time = p.ReadInt32()
}

func (pp PongPacket) Write(p *network.GamePacket) {
	p.WriteInt32(pp.Serial)
	p.WriteInt32(pp.Time)
}
