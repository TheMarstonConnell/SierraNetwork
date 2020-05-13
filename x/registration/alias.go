// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/enigmampc/EnigmaBlockchain/x/compute/internal/types
// ALIASGEN: github.com/enigmampc/EnigmaBlockchain/x/compute/internal/keeper
package registration

import (
	"github.com/enigmampc/EnigmaBlockchain/x/registration/internal/keeper"
	"github.com/enigmampc/EnigmaBlockchain/x/registration/internal/types"
)

const (
	ModuleName           = types.ModuleName
	StoreKey             = types.StoreKey
	TStoreKey            = types.TStoreKey
	QuerierRoute         = types.QuerierRoute
	RouterKey            = types.RouterKey
	QueryEncryptedSeed   = keeper.QueryEncryptedSeed
	SecretNodeSeedConfig = types.SecretNodeSeedConfig
	SecretNodeCfgFolder  = types.SecretNodeCfgFolder
	EncryptedKeyLength   = types.EncryptedKeyLength
	PublicKeyLength      = types.PublicKeyLength
)

var (
	// functions aliases
	RegisterCodec   = types.RegisterCodec
	ValidateGenesis = types.ValidateGenesis
	InitGenesis     = keeper.InitGenesis
	ExportGenesis   = keeper.ExportGenesis
	NewKeeper       = keeper.NewKeeper
	NewQuerier      = keeper.NewQuerier

	IsHexString = keeper.IsHexString
	//MakeTestCodec             = keeper.MakeTestCodec
	//CreateTestInput           = keeper.CreateTestInput

	// variable aliases
	ModuleCdc               = types.ModuleCdc
	DefaultCodespace        = types.DefaultCodespace
	ErrAuthenticateFailed   = types.ErrAuthenticateFailed
	ErrSeedInitFailed       = types.ErrSeedInitFailed
	RegistrationStorePrefix = types.RegistrationStorePrefix
)

type (
	MsgRaAuthenticate = types.RaAuthenticate
	GenesisState      = types.GenesisState
	Keeper            = keeper.Keeper
	SeedConfig        = types.SeedConfig
)
