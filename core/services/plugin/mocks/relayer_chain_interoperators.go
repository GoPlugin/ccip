package mocks

import (
	"context"
	"slices"

	services2 "github.com/goplugin/pluginv3.0/v2/core/services"
	"github.com/goplugin/pluginv3.0/v2/core/services/plugin"

	"github.com/goplugin/pluginv3.0/v2/core/chains/legacyevm"

	"github.com/goplugin/plugin-common/pkg/loop"

	"github.com/goplugin/plugin-common/pkg/types"
)

// FakeRelayerChainInteroperators is a fake plugin.RelayerChainInteroperators.
// This exists because mockery generation doesn't understand how to produce an alias instead of the underlying type (which is not exported in this case).
type FakeRelayerChainInteroperators struct {
	Relayers  []loop.Relayer
	EVMChains legacyevm.LegacyChainContainer
	Nodes     []types.NodeStatus
	NodesErr  error
}

func (f *FakeRelayerChainInteroperators) LegacyEVMChains() legacyevm.LegacyChainContainer {
	return f.EVMChains
}

func (f *FakeRelayerChainInteroperators) NodeStatuses(ctx context.Context, offset, limit int, relayIDs ...types.RelayID) (nodes []types.NodeStatus, count int, err error) {
	return slices.Clone(f.Nodes), len(f.Nodes), f.NodesErr
}

func (f *FakeRelayerChainInteroperators) Services() []services2.ServiceCtx {
	panic("unimplemented")
}

func (f *FakeRelayerChainInteroperators) List(filter plugin.FilterFn) plugin.RelayerChainInteroperators {
	panic("unimplemented")
}

func (f *FakeRelayerChainInteroperators) Get(id types.RelayID) (loop.Relayer, error) {
	panic("unimplemented")
}

func (f *FakeRelayerChainInteroperators) GetIDToRelayerMap() (map[types.RelayID]loop.Relayer, error) {
	panic("unimplemented")
}

func (f *FakeRelayerChainInteroperators) Slice() []loop.Relayer {
	return f.Relayers
}

func (f *FakeRelayerChainInteroperators) LegacyCosmosChains() plugin.LegacyCosmosContainer {
	panic("unimplemented")
}

func (f *FakeRelayerChainInteroperators) ChainStatus(ctx context.Context, id types.RelayID) (types.ChainStatus, error) {
	panic("unimplemented")
}

func (f *FakeRelayerChainInteroperators) ChainStatuses(ctx context.Context, offset, limit int) ([]types.ChainStatus, int, error) {
	panic("unimplemented")
}
