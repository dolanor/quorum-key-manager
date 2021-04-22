package interceptor

import (
	"fmt"
	"testing"

	"github.com/ConsenSysQuorum/quorum-key-manager/src/store/accounts"
	mockaccounts "github.com/ConsenSysQuorum/quorum-key-manager/src/store/accounts/mock"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
)

func TestEthSign(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	i, stores, _ := newInterceptor(ctrl)
	accountsStore := mockaccounts.NewMockStore(ctrl)

	tests := []*testHandlerCase{
		{
			desc:    "Signature",
			handler: i.EthSign(),
			prepare: func() {
				expectedFrom := ethcommon.HexToAddress("0x78e6e236592597c09d5c137c2af40aecd42d12a2")
				stores.EXPECT().GetAccountStoreByAddr(gomock.Any(), expectedFrom).Return(accountsStore, nil)
				accountsStore.EXPECT().Sign(gomock.Any(), expectedFrom, ethcommon.FromHex("0x2eadbe1f")).Return(ethcommon.FromHex("0xa6122e27"), nil)
			},
			reqBody:          []byte(`{"jsonrpc":"2.0","method":"test","params":["0x78e6e236592597c09d5c137c2af40aecd42d12a2", "0x2eadbe1f"]}`),
			expectedRespBody: []byte(`{"jsonrpc":"","result":"0xa6122e27","error":null,"id":null}`),
		},
		{
			desc:    "Account not found",
			handler: i.EthSign(),
			prepare: func() {
				expectedFrom := ethcommon.HexToAddress("0x78e6e236592597c09d5c137c2af40aecd42d12a2")
				stores.EXPECT().GetAccountStoreByAddr(gomock.Any(), expectedFrom).Return(nil, accounts.ErrorNotfound)
			},
			reqBody:          []byte(`{"jsonrpc":"2.0","method":"test","params":["0x78e6e236592597c09d5c137c2af40aecd42d12a2", "0x2eadbe1f"]}`),
			expectedRespBody: []byte(`{"jsonrpc":"","result":null,"error":{"code":-32000,"message":"account not found","data":null},"id":null}`),
		},
		{
			desc:    "Error signing",
			handler: i.EthSign(),
			prepare: func() {
				expectedFrom := ethcommon.HexToAddress("0x78e6e236592597c09d5c137c2af40aecd42d12a2")
				stores.EXPECT().GetAccountStoreByAddr(gomock.Any(), expectedFrom).Return(accountsStore, nil)
				accountsStore.EXPECT().Sign(gomock.Any(), expectedFrom, ethcommon.FromHex("0x2eadbe1f")).Return(nil, fmt.Errorf("error signing"))
			},
			reqBody:          []byte(`{"jsonrpc":"2.0","method":"test","params":["0x78e6e236592597c09d5c137c2af40aecd42d12a2", "0x2eadbe1f"]}`),
			expectedRespBody: []byte(`{"jsonrpc":"","result":null,"error":{"code":-32000,"message":"error signing","data":null},"id":null}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assertHandlerScenario(t, tt)
		})
	}
}
