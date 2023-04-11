package types

// ValidateBasic is used for validating the packet
func (p DustpacketPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p DustpacketPacketData) GetBytes() ([]byte, error) {
	var modulePacket GatePacketData

	modulePacket.Packet = &GatePacketData_DustpacketPacket{&p}

	return modulePacket.Marshal()
}
