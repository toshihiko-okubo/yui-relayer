package ethereum

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/hyperledger-labs/yui-relayer/core"
)

// RegisterInterfaces register the module interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*core.ChainConfigI)(nil),
		&ChainConfig{},
	)
}

func makeEncodingConfig() params.EncodingConfig {
	encodingConfig := core.MakeEncodingConfig()
	RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}