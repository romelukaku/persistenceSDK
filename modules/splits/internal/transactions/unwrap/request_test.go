/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package unwrap

import (
	"encoding/json"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"testing"

	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	xprtErrors "github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/schema"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	baseHelpers "github.com/persistenceOne/persistenceSDK/schema/helpers/base"
	"github.com/stretchr/testify/require"
)

func Test_Unwrap_Request(t *testing.T) {
	var Codec = codec.NewLegacyAmino()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterLegacyAminoCodec(Codec)
	cryptoCodec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vestingTypes.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{flags.FromID, flags.OwnableID, flags.Value})
	cliContext := client.Context{}.WithLegacyAmino(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	testBaseReq := test_types.BaseReq{From: fromAddress, ChainId: "test", Fees: sdkTypes.NewCoins()}
	testTransactionRequest := newTransactionRequest(testBaseReq, "fromID", "ownableID", "2")

	require.Equal(t, TransactionRequest{BaseReq: testBaseReq, FromID: "fromID", OwnableID: "ownableID", Value: "2"}, testTransactionRequest)
	require.Equal(t, nil, testTransactionRequest.Validate())

	requestFromCLI, Error := TransactionRequest{}.FromCLI(cliCommand, cliContext)
	require.Equal(t, nil, Error)
	require.Equal(t, TransactionRequest{BaseReq: test_types.BaseReq{From: cliContext.GetFromAddress().String(), ChainId: cliContext.ChainID, Simulate: cliContext.Simulate}, FromID: "", OwnableID: "", Value: ""}, requestFromCLI)

	jsonMessage, _ := json.Marshal(testTransactionRequest)
	transactionRequestUnmarshalled, Error := TransactionRequest{}.FromJSON(jsonMessage)
	require.Equal(t, nil, Error)
	require.Equal(t, testTransactionRequest, transactionRequestUnmarshalled)

	randomUnmarshall, Error := TransactionRequest{}.FromJSON([]byte{})
	require.Equal(t, nil, randomUnmarshall)
	require.NotNil(t, Error)

	require.Equal(t, testBaseReq, testTransactionRequest.GetBaseReq())

	msg, Error := testTransactionRequest.MakeMsg()
	require.Equal(t, newMessage(fromAccAddress, test_types.NewID("fromID"), test_types.NewID("ownableID"), sdkTypes.NewInt(2)), msg)
	require.Nil(t, Error)

	msg2, Error := newTransactionRequest(test_types.BaseReq{From: "randomFromAddress", ChainId: "test"}, "fromID", "ownableID", "2").MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	msg2, Error = newTransactionRequest(test_types.BaseReq{From: fromAddress, ChainId: "test"}, "fromID", "ownableID", "2.5").MakeMsg()
	require.Equal(t, xprtErrors.InvalidRequest, Error)
	require.Nil(t, msg2)

	msg2, Error = newTransactionRequest(testBaseReq, "fromID", "ownableID", "randomString").MakeMsg()
	require.NotNil(t, Error)
	require.Nil(t, msg2)

	require.Equal(t, TransactionRequest{}, requestPrototype())
	require.NotPanics(t, func() {
		requestPrototype().RegisterCodec(codec.NewLegacyAmino())
	})
}
