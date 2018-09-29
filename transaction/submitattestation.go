package transaction

import (
	"io"

	"github.com/phoreproject/synapse/serialization"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// SubmitAttestationTransaction submits a signed attestation to the
// beacon chain.
type SubmitAttestationTransaction struct {
	AggregateSignature []byte
	AttesterBitField   []byte
	Attestation
}

// Deserialize reads a submitattestationtransaction from the reader.
func (sat SubmitAttestationTransaction) Deserialize(r io.Reader) error {

}

func (sat SubmitAttestationTransaction) Serialize() []byte {

}

// Attestation is a signed attestation of a shard block.
type Attestation struct {
	Slot                uint64
	ShardID             uint64
	JustifiedSlot       uint64
	JustifiedBlockHash  *chainhash.Hash
	ShardBlockHash      *chainhash.Hash
	ObliqueParentHashes []*chainhash.Hash
}

// Deserialize reads an attestation from the provided reader.
func (a Attestation) Deserialize(r io.Reader) error {
	slot, err := serialization.ReadUint64(r)
	if err != nil {
		return err
	}

	shardID, err := serialization.ReadUint64(r)
	if err != nil {
		return err
	}

	justifiedSlot, err := serialization.ReadUint64(r)
	if err != nil {
		return err
	}

	justifiedBlockHash, err := serialization.ReadHash(r)
	if err != nil {
		return err
	}

	shardBlockHash, err := serialization.ReadHash(r)
	if err != nil {
		return err
	}

	obliqueParentHashes, err := serialization.ReadHashArray(r)
	if err != nil {
		return err
	}

	a.Slot = slot
	a.ShardID = shardID
	a.ShardBlockHash = shardBlockHash
	a.JustifiedSlot = justifiedSlot
	a.JustifiedBlockHash = justifiedBlockHash
	a.ObliqueParentHashes = obliqueParentHashes
	return nil
}

func (a Attestation) Serialize() []byte {

}
