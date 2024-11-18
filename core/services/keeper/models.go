package keeper

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/goplugin/pluginv3.0/v2/core/chains/evm/types"
	"github.com/goplugin/pluginv3.0/v2/core/chains/evm/utils"
	"github.com/goplugin/pluginv3.0/v2/core/chains/evm/utils/big"
	"github.com/goplugin/pluginv3.0/v2/core/null"
)

type KeeperIndexMap map[types.EIP55Address]int32

type Registry struct {
	ID                int64
	BlockCountPerTurn int32
	CheckGas          uint32
	ContractAddress   types.EIP55Address
	FromAddress       types.EIP55Address
	JobID             int32
	KeeperIndex       int32
	NumKeepers        int32
	KeeperIndexMap    KeeperIndexMap
}

type UpkeepRegistration struct {
	ID                  int32
	CheckData           []byte
	ExecuteGas          uint32
	LastRunBlockHeight  int64
	RegistryID          int64
	Registry            Registry
	UpkeepID            *big.Big
	LastKeeperIndex     null.Int64
	PositioningConstant int32
}

func (k *KeeperIndexMap) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &k)
		return err
	case string:
		err := json.Unmarshal([]byte(v), &k)
		return err
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func (k *KeeperIndexMap) Value() (driver.Value, error) {
	return json.Marshal(&k)
}

func (upkeep UpkeepRegistration) PrettyID() string {
	return NewUpkeepIdentifier(upkeep.UpkeepID).String()
}

func NewUpkeepIdentifier(i *big.Big) *UpkeepIdentifier {
	val := UpkeepIdentifier(*i)
	return &val
}

type UpkeepIdentifier big.Big

// String produces a hex encoded value, zero padded, prefixed with UpkeepPrefix
func (ui UpkeepIdentifier) String() string {
	val := big.Big(ui)
	result, err := utils.Uint256ToBytes(val.ToInt())
	if err != nil {
		panic(errors.Wrap(err, "invariant, invalid upkeepID"))
	}
	return fmt.Sprintf("%s%s", UpkeepPrefix, hex.EncodeToString(result))
}