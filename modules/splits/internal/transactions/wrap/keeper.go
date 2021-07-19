/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package wrap

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
)

type transactionKeeper struct {
	mapper          helpers.Mapper
	parameters      helpers.Parameters
	bankKeeper      bank.Keeper
	verifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	if Error := transactionKeeper.bankKeeper.SendCoinsFromAccountToModule(context, message.From, module.Name, message.Coins); Error != nil {
		return newTransactionResponse(Error)
	}

	for _, coin := range message.Coins {
		if _, Error := utilities.AddSplits(transactionKeeper.mapper.NewCollection(context), message.FromID, test_types.NewID(coin.Denom), sdkTypes.NewDecFromInt(coin.Amount)); Error != nil {
			return newTransactionResponse(Error)
		}
	}

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bank.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			default:
				panic(errors.UninitializedUsage)
			}
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
