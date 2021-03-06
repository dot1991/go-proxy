package client

import "proxy/network"

type PlayerHitPacket struct {
	BulletID byte
	ObjectID int32
}

func (ph *PlayerHitPacket) Read(p *network.GamePacket) {
	ph.BulletID = p.ReadByte()
	ph.ObjectID = p.ReadInt32()
}

func (ph PlayerHitPacket) Write(p *network.GamePacket) {
	p.WriteByte(ph.BulletID)
	p.WriteInt32(ph.ObjectID)
}
