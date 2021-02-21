package blockchain

import "math/big"

/*
- Take the data from the block
- Create a counter (nonce) which starts at 0
- Create the hash of the data plus the counter
- Check the hash to see if it meets a set of requirements
	- Requirements:
		The first few bytes must contains 0s
*/

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}