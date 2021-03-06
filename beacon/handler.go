package beacon

import "github.com/phoreproject/synapse/beacon/primitives"

// HandleNewBlocks handles any incoming blocks.
func (b *Blockchain) HandleNewBlocks(blocks chan primitives.Block) error {
	for {
		block := <-blocks
		err := b.ProcessBlock(&block)
		if err != nil {
			return err
		}
	}
}
