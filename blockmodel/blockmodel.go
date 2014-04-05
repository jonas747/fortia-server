package blockmodel

import ()

type Vector3f struct {
	x, y, z float64
}

type BlockModel struct {
	Pos Vector3f
	// Animations
	// Physics bounds
	// blocks
}

// Returns a BlockModel from a messages.BlockModel protobuf source
func FromProto(source []byte) (*BlockModel, error) {
	return nil, nil
}

// Serializes the BlockModel into messages.BlockModel proto
func (b *BlockModel) Serialize() ([]byte, error) {
	return make([]error, nil)
}
