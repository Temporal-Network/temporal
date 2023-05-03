package types

func (f Frequency) IsNil() bool {
	return f.OnceEvery.BigInt() == nil
}

func (f Frequency) IsBlock() bool {
	return f.SecondsOrBlocks == "block"
}

func (f Frequency) IsSecond() bool {
	return f.SecondsOrBlocks == "second"
}

func (f Frequency) IsSecondsOrBlocksValid() bool {
	return f.IsBlock() || f.IsSecond()
}

// TODO: setup unit tests for the above functions
