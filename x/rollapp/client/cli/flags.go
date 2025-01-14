package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagInitSequencer   = "init-sequencer"
	FlagGenesisChecksum = "genesis-checksum"
	FlagAlias           = "alias"
	FlagMetadata        = "metadata"
	FlagByAlias         = "by-alias"
)

// FlagSetCreateRollapp returns flags for creating rollapps.
func FlagSetCreateRollapp() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	return fs
}

// FlagSetUpdateRollapp returns flags for updating rollapps.
func FlagSetUpdateRollapp() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	fs.String(FlagInitSequencer, "", "The address of the sequencer that will be used to initialize the rollapp")
	fs.String(FlagGenesisChecksum, "", "The checksum of the genesis file of the rollapp")
	fs.String(FlagAlias, "", "The alias of the rollapp")
	fs.String(FlagMetadata, "", "The metadata of the rollapp")

	return fs
}
