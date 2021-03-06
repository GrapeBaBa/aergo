/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */
package rpc

import (
	"context"
	"encoding/hex"
	"fmt"
	"reflect"
	"testing"

	"github.com/aergoio/aergo/internal/enc"
	"github.com/aergoio/aergo/p2p"

	"github.com/aergoio/aergo-actor/actor"

	"github.com/aergoio/aergo/message"
	"github.com/aergoio/aergo/message/mocks"
	"github.com/aergoio/aergo/pkg/component"
	"github.com/aergoio/aergo/types"
	"github.com/mr-tron/base58/base58"
	"github.com/stretchr/testify/mock"
)

func TestAergoRPCService_dummys(t *testing.T) {
	fmt.Println("dummyBlockHash")
	fmt.Printf("HEX : %s \n", hex.EncodeToString(dummyBlockHash))
	fmt.Printf("B64 : %s \n", enc.ToString(dummyBlockHash))
	fmt.Printf("B58 : %s \n", base58.Encode(dummyBlockHash))
	fmt.Println()
	fmt.Println("dummyTx")
	fmt.Printf("HEX : %s \n", hex.EncodeToString(dummyTxHash))
	fmt.Printf("B64 : %s \n", enc.ToString(dummyTxHash))
	fmt.Printf("B58 : %s \n", base58.Encode(dummyTxHash))
	fmt.Println()

	fmt.Println("Address1")
	fmt.Printf("HEX : %s \n", hex.EncodeToString(dummyWalletAddress))
	fmt.Printf("B64 : %s \n", enc.ToString(dummyWalletAddress))
	fmt.Printf("B58 : %s \n", base58.Encode(dummyWalletAddress))
	fmt.Println()

	fmt.Println("Address2")
	fmt.Printf("HEX : %s \n", hex.EncodeToString(dummyWalletAddress2))
	fmt.Printf("B64 : %s \n", enc.ToString(dummyWalletAddress2))
	fmt.Printf("B58 : %s \n", base58.Encode(dummyWalletAddress2))
	fmt.Println()

}

var dummyBlockHash, _ = hex.DecodeString("4f461d85e869ade8a0544f8313987c33a9c06534e50c4ad941498299579bd7ac")
var dummyBlockHeight uint32 = 100215
var dummyTxHash, _ = hex.DecodeString("218bdab4e87fb332b55eb89854ef553f9e3d440c81fff4161b672adede1261ee")

// base64 encoding of dummyTxHash is ""
var dummyWalletAddress, _ = base58.Decode("1Ee8uhLFXzkSRRU1orBpgXFAPpVi64aSYo")
var dummyWalletAddress2, _ = base58.Decode("16Uiu2HAkwgfFvViH6j2QpQYKtGKKdveEKZvU2T5mRkqFLTZKU4Vp")
var dummyPayload = []byte("OPreturn I am groooot")

var hubStub *component.ComponentHub
var mockCtx context.Context
var mockMsgHelper *mocks.Helper
var mockActorHelper *MockActorService

func init() {
	hubStub = &component.ComponentHub{}
	mockMsgHelper = &mocks.Helper{}
	mockActorHelper = &MockActorService{}

	mockCtx = &Context{}
}
func TestAergoRPCService_GetTX(t *testing.T) {
	dummyTxBody := types.TxBody{Account: dummyWalletAddress, Amount: 4332, Recipient: dummyWalletAddress2, Payload: dummyPayload}
	sampleTx := &types.Tx{Hash: dummyTxHash, Body: &dummyTxBody}
	mockActorHelper.On("CallRequestDefaultTimeout", message.MemPoolSvc, mock.Anything).Return(message.MemPoolGetRsp{}, nil)
	mockMsgHelper.On("ExtractTxFromResponse", mock.AnythingOfType("message.MemPoolGetRsp")).Return(sampleTx, nil)
	type fields struct {
		hub         *component.ComponentHub
		actorHelper p2p.ActorService
		msgHelper   message.Helper
	}
	type args struct {
		ctx context.Context
		in  *types.SingleBytes
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Tx
		wantErr bool
	}{
		{args: args{ctx: mockCtx, in: &types.SingleBytes{Value: append(dummyTxHash, 'b', 'd')}}, fields: fields{hubStub, mockActorHelper, mockMsgHelper},
			want: &types.Tx{Hash: dummyTxHash, Body: &dummyTxBody}, wantErr: false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rpc := &AergoRPCService{
				hub: tt.fields.hub, actorHelper: mockActorHelper, msgHelper: mockMsgHelper,
			}
			got, err := rpc.GetTX(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("AergoRPCService.GetTX() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AergoRPCService.GetTX() = %v, want %v", got, tt.want)
			}
		})
	}
}

type FutureStub struct {
	actor.Future
	dumbResult interface{}
}

func (fs *FutureStub) Result() interface{} {
	return fs.dumbResult
}

func NewFutureStub(result interface{}) FutureStub {
	return FutureStub{dumbResult: result}
}
