package maintain

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/parameters"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type TestKeepers struct {
	MaintainersKeeper helpers.AuxiliaryKeeper
}

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := baseHelpers.NewMapper(key.Prototype, mappable.Prototype).Initialize(storeKey)
	paramsKeeper := params.NewKeeper(
		Codec,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	Parameters := parameters.Prototype().Initialize(paramsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	Error := commitMultiStore.LoadLatestVersion()
	require.Nil(t, Error)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	keepers := TestKeepers{
		MaintainersKeeper: keeperPrototype().Initialize(Mapper, Parameters, []interface{}{}).(helpers.AuxiliaryKeeper),
	}

	return context, keepers
}

func Test_Auxiliary_Keeper_Help(t *testing.T) {

	context, keepers := CreateTestInput(t)

	classificationID := base.NewID("classificationID")
	identityID := base.NewID("identityID")

	maintainedModules := base.NewMutables(base.NewProperties())
	maintainerID := key.NewMaintainerID(classificationID, identityID)
	keepers.MaintainersKeeper.(auxiliaryKeeper).mapper.NewCollection(context).Add(mappable.NewMaintainer(maintainerID, base.NewMutables(base.NewProperties()), false, false, false))

	t.Run("PositiveCase", func(t *testing.T) {
		want := newAuxiliaryResponse(nil)
		if got := keepers.MaintainersKeeper.Help(context, NewAuxiliaryRequest(classificationID, identityID, maintainedModules)); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Maintainer not present", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.EntityNotFound)
		if got := keepers.MaintainersKeeper.Help(context, NewAuxiliaryRequest(base.NewID("classificationID1"), base.NewID("identityID1"), base.NewMutables(base.NewProperties()))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

	t.Run("NegativeCase-Maintainer Unauthorized", func(t *testing.T) {
		t.Parallel()
		want := newAuxiliaryResponse(errors.NotAuthorized)
		if got := keepers.MaintainersKeeper.Help(context, NewAuxiliaryRequest(classificationID, identityID, base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID"), base.NewFact(base.NewStringData("Data"))))))); !reflect.DeepEqual(got, want) {
			t.Errorf("Transact() = %v, want %v", got, want)
		}
	})

}