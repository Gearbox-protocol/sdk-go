package priceFetcher

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/ethereum/go-ethereum/common"
)

func NewGearboxOraclev3(addr common.Address, version core.VersionType, client core.ClientI) GearboxOracleI {
	po := &GearboxOracle{
		Address:     addr,
		TokenToFeed: map[string]common.Address{},
		Node: &pkg.Node{
			Client: client,
		},
		topics:  []common.Hash{core.Topic("SetPriceFeed(address,address,uint32,bool,bool)")},
		version: version,
	}
	return po
}
