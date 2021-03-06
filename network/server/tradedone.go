package server

import "proxy/network"

type TradeDonePacket struct {
	ResultCode int32
	Message    string
}

func (t *TradeDonePacket) Read(p *network.GamePacket) {
	t.ResultCode = p.ReadInt32()
	t.Message = p.ReadString()
}

func (t TradeDonePacket) Write(p *network.GamePacket) {
	p.WriteInt32(t.ResultCode)
	p.WriteString(t.Message)
}
