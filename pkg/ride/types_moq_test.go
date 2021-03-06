// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ride

import (
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/proto"
	"github.com/wavesplatform/gowaves/pkg/types"
	"sync"
)

// Ensure, that MockSmartState does implement types.SmartState.
// If this is not the case, regenerate this file with moq.
var _ types.SmartState = &MockSmartState{}

// MockSmartState is a mock implementation of types.SmartState.
//
//     func TestSomethingThatUsesSmartState(t *testing.T) {
//
//         // make and configure a mocked types.SmartState
//         mockedSmartState := &MockSmartState{
//             AddingBlockHeightFunc: func() (uint64, error) {
// 	               panic("mock out the AddingBlockHeight method")
//             },
//             BlockVRFFunc: func(blockHeader *proto.BlockHeader, height uint64) ([]byte, error) {
// 	               panic("mock out the BlockVRF method")
//             },
//             EstimatorVersionFunc: func() (int, error) {
// 	               panic("mock out the EstimatorVersion method")
//             },
//             IsNotFoundFunc: func(err error) bool {
// 	               panic("mock out the IsNotFound method")
//             },
//             NewestAccountBalanceFunc: func(account proto.Recipient, asset []byte) (uint64, error) {
// 	               panic("mock out the NewestAccountBalance method")
//             },
//             NewestAddrByAliasFunc: func(alias proto.Alias) (proto.Address, error) {
// 	               panic("mock out the NewestAddrByAlias method")
//             },
//             NewestAssetInfoFunc: func(assetID crypto.Digest) (*proto.AssetInfo, error) {
// 	               panic("mock out the NewestAssetInfo method")
//             },
//             NewestAssetIsSponsoredFunc: func(assetID crypto.Digest) (bool, error) {
// 	               panic("mock out the NewestAssetIsSponsored method")
//             },
//             NewestFullAssetInfoFunc: func(assetID crypto.Digest) (*proto.FullAssetInfo, error) {
// 	               panic("mock out the NewestFullAssetInfo method")
//             },
//             NewestFullWavesBalanceFunc: func(account proto.Recipient) (*proto.FullWavesBalance, error) {
// 	               panic("mock out the NewestFullWavesBalance method")
//             },
//             NewestHeaderByHeightFunc: func(height uint64) (*proto.BlockHeader, error) {
// 	               panic("mock out the NewestHeaderByHeight method")
//             },
//             NewestTransactionByIDFunc: func(in1 []byte) (proto.Transaction, error) {
// 	               panic("mock out the NewestTransactionByID method")
//             },
//             NewestTransactionHeightByIDFunc: func(in1 []byte) (uint64, error) {
// 	               panic("mock out the NewestTransactionHeightByID method")
//             },
//             RetrieveNewestBinaryEntryFunc: func(account proto.Recipient, key string) (*proto.BinaryDataEntry, error) {
// 	               panic("mock out the RetrieveNewestBinaryEntry method")
//             },
//             RetrieveNewestBooleanEntryFunc: func(account proto.Recipient, key string) (*proto.BooleanDataEntry, error) {
// 	               panic("mock out the RetrieveNewestBooleanEntry method")
//             },
//             RetrieveNewestIntegerEntryFunc: func(account proto.Recipient, key string) (*proto.IntegerDataEntry, error) {
// 	               panic("mock out the RetrieveNewestIntegerEntry method")
//             },
//             RetrieveNewestStringEntryFunc: func(account proto.Recipient, key string) (*proto.StringDataEntry, error) {
// 	               panic("mock out the RetrieveNewestStringEntry method")
//             },
//         }
//
//         // use mockedSmartState in code that requires types.SmartState
//         // and then make assertions.
//
//     }
type MockSmartState struct {
	// AddingBlockHeightFunc mocks the AddingBlockHeight method.
	AddingBlockHeightFunc func() (uint64, error)

	// BlockVRFFunc mocks the BlockVRF method.
	BlockVRFFunc func(blockHeader *proto.BlockHeader, height uint64) ([]byte, error)

	// EstimatorVersionFunc mocks the EstimatorVersion method.
	EstimatorVersionFunc func() (int, error)

	// IsNotFoundFunc mocks the IsNotFound method.
	IsNotFoundFunc func(err error) bool

	// NewestAccountBalanceFunc mocks the NewestAccountBalance method.
	NewestAccountBalanceFunc func(account proto.Recipient, asset []byte) (uint64, error)

	// NewestAddrByAliasFunc mocks the NewestAddrByAlias method.
	NewestAddrByAliasFunc func(alias proto.Alias) (proto.Address, error)

	// NewestAssetInfoFunc mocks the NewestAssetInfo method.
	NewestAssetInfoFunc func(assetID crypto.Digest) (*proto.AssetInfo, error)

	// NewestAssetIsSponsoredFunc mocks the NewestAssetIsSponsored method.
	NewestAssetIsSponsoredFunc func(assetID crypto.Digest) (bool, error)

	// NewestFullAssetInfoFunc mocks the NewestFullAssetInfo method.
	NewestFullAssetInfoFunc func(assetID crypto.Digest) (*proto.FullAssetInfo, error)

	// NewestFullWavesBalanceFunc mocks the NewestFullWavesBalance method.
	NewestFullWavesBalanceFunc func(account proto.Recipient) (*proto.FullWavesBalance, error)

	// NewestHeaderByHeightFunc mocks the NewestHeaderByHeight method.
	NewestHeaderByHeightFunc func(height uint64) (*proto.BlockHeader, error)

	// NewestTransactionByIDFunc mocks the NewestTransactionByID method.
	NewestTransactionByIDFunc func(in1 []byte) (proto.Transaction, error)

	// NewestTransactionHeightByIDFunc mocks the NewestTransactionHeightByID method.
	NewestTransactionHeightByIDFunc func(in1 []byte) (uint64, error)

	// RetrieveNewestBinaryEntryFunc mocks the RetrieveNewestBinaryEntry method.
	RetrieveNewestBinaryEntryFunc func(account proto.Recipient, key string) (*proto.BinaryDataEntry, error)

	// RetrieveNewestBooleanEntryFunc mocks the RetrieveNewestBooleanEntry method.
	RetrieveNewestBooleanEntryFunc func(account proto.Recipient, key string) (*proto.BooleanDataEntry, error)

	// RetrieveNewestIntegerEntryFunc mocks the RetrieveNewestIntegerEntry method.
	RetrieveNewestIntegerEntryFunc func(account proto.Recipient, key string) (*proto.IntegerDataEntry, error)

	// RetrieveNewestStringEntryFunc mocks the RetrieveNewestStringEntry method.
	RetrieveNewestStringEntryFunc func(account proto.Recipient, key string) (*proto.StringDataEntry, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddingBlockHeight holds details about calls to the AddingBlockHeight method.
		AddingBlockHeight []struct {
		}
		// BlockVRF holds details about calls to the BlockVRF method.
		BlockVRF []struct {
			// BlockHeader is the blockHeader argument value.
			BlockHeader *proto.BlockHeader
			// Height is the height argument value.
			Height uint64
		}
		// EstimatorVersion holds details about calls to the EstimatorVersion method.
		EstimatorVersion []struct {
		}
		// IsNotFound holds details about calls to the IsNotFound method.
		IsNotFound []struct {
			// Err is the err argument value.
			Err error
		}
		// NewestAccountBalance holds details about calls to the NewestAccountBalance method.
		NewestAccountBalance []struct {
			// Account is the account argument value.
			Account proto.Recipient
			// Asset is the asset argument value.
			Asset []byte
		}
		// NewestAddrByAlias holds details about calls to the NewestAddrByAlias method.
		NewestAddrByAlias []struct {
			// Alias is the alias argument value.
			Alias proto.Alias
		}
		// NewestAssetInfo holds details about calls to the NewestAssetInfo method.
		NewestAssetInfo []struct {
			// AssetID is the assetID argument value.
			AssetID crypto.Digest
		}
		// NewestAssetIsSponsored holds details about calls to the NewestAssetIsSponsored method.
		NewestAssetIsSponsored []struct {
			// AssetID is the assetID argument value.
			AssetID crypto.Digest
		}
		// NewestFullAssetInfo holds details about calls to the NewestFullAssetInfo method.
		NewestFullAssetInfo []struct {
			// AssetID is the assetID argument value.
			AssetID crypto.Digest
		}
		// NewestFullWavesBalance holds details about calls to the NewestFullWavesBalance method.
		NewestFullWavesBalance []struct {
			// Account is the account argument value.
			Account proto.Recipient
		}
		// NewestHeaderByHeight holds details about calls to the NewestHeaderByHeight method.
		NewestHeaderByHeight []struct {
			// Height is the height argument value.
			Height uint64
		}
		// NewestTransactionByID holds details about calls to the NewestTransactionByID method.
		NewestTransactionByID []struct {
			// In1 is the in1 argument value.
			In1 []byte
		}
		// NewestTransactionHeightByID holds details about calls to the NewestTransactionHeightByID method.
		NewestTransactionHeightByID []struct {
			// In1 is the in1 argument value.
			In1 []byte
		}
		// RetrieveNewestBinaryEntry holds details about calls to the RetrieveNewestBinaryEntry method.
		RetrieveNewestBinaryEntry []struct {
			// Account is the account argument value.
			Account proto.Recipient
			// Key is the key argument value.
			Key string
		}
		// RetrieveNewestBooleanEntry holds details about calls to the RetrieveNewestBooleanEntry method.
		RetrieveNewestBooleanEntry []struct {
			// Account is the account argument value.
			Account proto.Recipient
			// Key is the key argument value.
			Key string
		}
		// RetrieveNewestIntegerEntry holds details about calls to the RetrieveNewestIntegerEntry method.
		RetrieveNewestIntegerEntry []struct {
			// Account is the account argument value.
			Account proto.Recipient
			// Key is the key argument value.
			Key string
		}
		// RetrieveNewestStringEntry holds details about calls to the RetrieveNewestStringEntry method.
		RetrieveNewestStringEntry []struct {
			// Account is the account argument value.
			Account proto.Recipient
			// Key is the key argument value.
			Key string
		}
	}
	lockAddingBlockHeight           sync.RWMutex
	lockBlockVRF                    sync.RWMutex
	lockEstimatorVersion            sync.RWMutex
	lockIsNotFound                  sync.RWMutex
	lockNewestAccountBalance        sync.RWMutex
	lockNewestAddrByAlias           sync.RWMutex
	lockNewestAssetInfo             sync.RWMutex
	lockNewestAssetIsSponsored      sync.RWMutex
	lockNewestFullAssetInfo         sync.RWMutex
	lockNewestFullWavesBalance      sync.RWMutex
	lockNewestHeaderByHeight        sync.RWMutex
	lockNewestTransactionByID       sync.RWMutex
	lockNewestTransactionHeightByID sync.RWMutex
	lockRetrieveNewestBinaryEntry   sync.RWMutex
	lockRetrieveNewestBooleanEntry  sync.RWMutex
	lockRetrieveNewestIntegerEntry  sync.RWMutex
	lockRetrieveNewestStringEntry   sync.RWMutex
}

// AddingBlockHeight calls AddingBlockHeightFunc.
func (mock *MockSmartState) AddingBlockHeight() (uint64, error) {
	if mock.AddingBlockHeightFunc == nil {
		panic("MockSmartState.AddingBlockHeightFunc: method is nil but SmartState.AddingBlockHeight was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAddingBlockHeight.Lock()
	mock.calls.AddingBlockHeight = append(mock.calls.AddingBlockHeight, callInfo)
	mock.lockAddingBlockHeight.Unlock()
	return mock.AddingBlockHeightFunc()
}

// AddingBlockHeightCalls gets all the calls that were made to AddingBlockHeight.
// Check the length with:
//     len(mockedSmartState.AddingBlockHeightCalls())
func (mock *MockSmartState) AddingBlockHeightCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAddingBlockHeight.RLock()
	calls = mock.calls.AddingBlockHeight
	mock.lockAddingBlockHeight.RUnlock()
	return calls
}

// BlockVRF calls BlockVRFFunc.
func (mock *MockSmartState) BlockVRF(blockHeader *proto.BlockHeader, height uint64) ([]byte, error) {
	if mock.BlockVRFFunc == nil {
		panic("MockSmartState.BlockVRFFunc: method is nil but SmartState.BlockVRF was just called")
	}
	callInfo := struct {
		BlockHeader *proto.BlockHeader
		Height      uint64
	}{
		BlockHeader: blockHeader,
		Height:      height,
	}
	mock.lockBlockVRF.Lock()
	mock.calls.BlockVRF = append(mock.calls.BlockVRF, callInfo)
	mock.lockBlockVRF.Unlock()
	return mock.BlockVRFFunc(blockHeader, height)
}

// BlockVRFCalls gets all the calls that were made to BlockVRF.
// Check the length with:
//     len(mockedSmartState.BlockVRFCalls())
func (mock *MockSmartState) BlockVRFCalls() []struct {
	BlockHeader *proto.BlockHeader
	Height      uint64
} {
	var calls []struct {
		BlockHeader *proto.BlockHeader
		Height      uint64
	}
	mock.lockBlockVRF.RLock()
	calls = mock.calls.BlockVRF
	mock.lockBlockVRF.RUnlock()
	return calls
}

// EstimatorVersion calls EstimatorVersionFunc.
func (mock *MockSmartState) EstimatorVersion() (int, error) {
	if mock.EstimatorVersionFunc == nil {
		panic("MockSmartState.EstimatorVersionFunc: method is nil but SmartState.EstimatorVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockEstimatorVersion.Lock()
	mock.calls.EstimatorVersion = append(mock.calls.EstimatorVersion, callInfo)
	mock.lockEstimatorVersion.Unlock()
	return mock.EstimatorVersionFunc()
}

// EstimatorVersionCalls gets all the calls that were made to EstimatorVersion.
// Check the length with:
//     len(mockedSmartState.EstimatorVersionCalls())
func (mock *MockSmartState) EstimatorVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockEstimatorVersion.RLock()
	calls = mock.calls.EstimatorVersion
	mock.lockEstimatorVersion.RUnlock()
	return calls
}

// IsNotFound calls IsNotFoundFunc.
func (mock *MockSmartState) IsNotFound(err error) bool {
	if mock.IsNotFoundFunc == nil {
		panic("MockSmartState.IsNotFoundFunc: method is nil but SmartState.IsNotFound was just called")
	}
	callInfo := struct {
		Err error
	}{
		Err: err,
	}
	mock.lockIsNotFound.Lock()
	mock.calls.IsNotFound = append(mock.calls.IsNotFound, callInfo)
	mock.lockIsNotFound.Unlock()
	return mock.IsNotFoundFunc(err)
}

// IsNotFoundCalls gets all the calls that were made to IsNotFound.
// Check the length with:
//     len(mockedSmartState.IsNotFoundCalls())
func (mock *MockSmartState) IsNotFoundCalls() []struct {
	Err error
} {
	var calls []struct {
		Err error
	}
	mock.lockIsNotFound.RLock()
	calls = mock.calls.IsNotFound
	mock.lockIsNotFound.RUnlock()
	return calls
}

// NewestAccountBalance calls NewestAccountBalanceFunc.
func (mock *MockSmartState) NewestAccountBalance(account proto.Recipient, asset []byte) (uint64, error) {
	if mock.NewestAccountBalanceFunc == nil {
		panic("MockSmartState.NewestAccountBalanceFunc: method is nil but SmartState.NewestAccountBalance was just called")
	}
	callInfo := struct {
		Account proto.Recipient
		Asset   []byte
	}{
		Account: account,
		Asset:   asset,
	}
	mock.lockNewestAccountBalance.Lock()
	mock.calls.NewestAccountBalance = append(mock.calls.NewestAccountBalance, callInfo)
	mock.lockNewestAccountBalance.Unlock()
	return mock.NewestAccountBalanceFunc(account, asset)
}

// NewestAccountBalanceCalls gets all the calls that were made to NewestAccountBalance.
// Check the length with:
//     len(mockedSmartState.NewestAccountBalanceCalls())
func (mock *MockSmartState) NewestAccountBalanceCalls() []struct {
	Account proto.Recipient
	Asset   []byte
} {
	var calls []struct {
		Account proto.Recipient
		Asset   []byte
	}
	mock.lockNewestAccountBalance.RLock()
	calls = mock.calls.NewestAccountBalance
	mock.lockNewestAccountBalance.RUnlock()
	return calls
}

// NewestAddrByAlias calls NewestAddrByAliasFunc.
func (mock *MockSmartState) NewestAddrByAlias(alias proto.Alias) (proto.Address, error) {
	if mock.NewestAddrByAliasFunc == nil {
		panic("MockSmartState.NewestAddrByAliasFunc: method is nil but SmartState.NewestAddrByAlias was just called")
	}
	callInfo := struct {
		Alias proto.Alias
	}{
		Alias: alias,
	}
	mock.lockNewestAddrByAlias.Lock()
	mock.calls.NewestAddrByAlias = append(mock.calls.NewestAddrByAlias, callInfo)
	mock.lockNewestAddrByAlias.Unlock()
	return mock.NewestAddrByAliasFunc(alias)
}

// NewestAddrByAliasCalls gets all the calls that were made to NewestAddrByAlias.
// Check the length with:
//     len(mockedSmartState.NewestAddrByAliasCalls())
func (mock *MockSmartState) NewestAddrByAliasCalls() []struct {
	Alias proto.Alias
} {
	var calls []struct {
		Alias proto.Alias
	}
	mock.lockNewestAddrByAlias.RLock()
	calls = mock.calls.NewestAddrByAlias
	mock.lockNewestAddrByAlias.RUnlock()
	return calls
}

// NewestAssetInfo calls NewestAssetInfoFunc.
func (mock *MockSmartState) NewestAssetInfo(assetID crypto.Digest) (*proto.AssetInfo, error) {
	if mock.NewestAssetInfoFunc == nil {
		panic("MockSmartState.NewestAssetInfoFunc: method is nil but SmartState.NewestAssetInfo was just called")
	}
	callInfo := struct {
		AssetID crypto.Digest
	}{
		AssetID: assetID,
	}
	mock.lockNewestAssetInfo.Lock()
	mock.calls.NewestAssetInfo = append(mock.calls.NewestAssetInfo, callInfo)
	mock.lockNewestAssetInfo.Unlock()
	return mock.NewestAssetInfoFunc(assetID)
}

// NewestAssetInfoCalls gets all the calls that were made to NewestAssetInfo.
// Check the length with:
//     len(mockedSmartState.NewestAssetInfoCalls())
func (mock *MockSmartState) NewestAssetInfoCalls() []struct {
	AssetID crypto.Digest
} {
	var calls []struct {
		AssetID crypto.Digest
	}
	mock.lockNewestAssetInfo.RLock()
	calls = mock.calls.NewestAssetInfo
	mock.lockNewestAssetInfo.RUnlock()
	return calls
}

// NewestAssetIsSponsored calls NewestAssetIsSponsoredFunc.
func (mock *MockSmartState) NewestAssetIsSponsored(assetID crypto.Digest) (bool, error) {
	if mock.NewestAssetIsSponsoredFunc == nil {
		panic("MockSmartState.NewestAssetIsSponsoredFunc: method is nil but SmartState.NewestAssetIsSponsored was just called")
	}
	callInfo := struct {
		AssetID crypto.Digest
	}{
		AssetID: assetID,
	}
	mock.lockNewestAssetIsSponsored.Lock()
	mock.calls.NewestAssetIsSponsored = append(mock.calls.NewestAssetIsSponsored, callInfo)
	mock.lockNewestAssetIsSponsored.Unlock()
	return mock.NewestAssetIsSponsoredFunc(assetID)
}

// NewestAssetIsSponsoredCalls gets all the calls that were made to NewestAssetIsSponsored.
// Check the length with:
//     len(mockedSmartState.NewestAssetIsSponsoredCalls())
func (mock *MockSmartState) NewestAssetIsSponsoredCalls() []struct {
	AssetID crypto.Digest
} {
	var calls []struct {
		AssetID crypto.Digest
	}
	mock.lockNewestAssetIsSponsored.RLock()
	calls = mock.calls.NewestAssetIsSponsored
	mock.lockNewestAssetIsSponsored.RUnlock()
	return calls
}

// NewestFullAssetInfo calls NewestFullAssetInfoFunc.
func (mock *MockSmartState) NewestFullAssetInfo(assetID crypto.Digest) (*proto.FullAssetInfo, error) {
	if mock.NewestFullAssetInfoFunc == nil {
		panic("MockSmartState.NewestFullAssetInfoFunc: method is nil but SmartState.NewestFullAssetInfo was just called")
	}
	callInfo := struct {
		AssetID crypto.Digest
	}{
		AssetID: assetID,
	}
	mock.lockNewestFullAssetInfo.Lock()
	mock.calls.NewestFullAssetInfo = append(mock.calls.NewestFullAssetInfo, callInfo)
	mock.lockNewestFullAssetInfo.Unlock()
	return mock.NewestFullAssetInfoFunc(assetID)
}

// NewestFullAssetInfoCalls gets all the calls that were made to NewestFullAssetInfo.
// Check the length with:
//     len(mockedSmartState.NewestFullAssetInfoCalls())
func (mock *MockSmartState) NewestFullAssetInfoCalls() []struct {
	AssetID crypto.Digest
} {
	var calls []struct {
		AssetID crypto.Digest
	}
	mock.lockNewestFullAssetInfo.RLock()
	calls = mock.calls.NewestFullAssetInfo
	mock.lockNewestFullAssetInfo.RUnlock()
	return calls
}

// NewestFullWavesBalance calls NewestFullWavesBalanceFunc.
func (mock *MockSmartState) NewestFullWavesBalance(account proto.Recipient) (*proto.FullWavesBalance, error) {
	if mock.NewestFullWavesBalanceFunc == nil {
		panic("MockSmartState.NewestFullWavesBalanceFunc: method is nil but SmartState.NewestFullWavesBalance was just called")
	}
	callInfo := struct {
		Account proto.Recipient
	}{
		Account: account,
	}
	mock.lockNewestFullWavesBalance.Lock()
	mock.calls.NewestFullWavesBalance = append(mock.calls.NewestFullWavesBalance, callInfo)
	mock.lockNewestFullWavesBalance.Unlock()
	return mock.NewestFullWavesBalanceFunc(account)
}

// NewestFullWavesBalanceCalls gets all the calls that were made to NewestFullWavesBalance.
// Check the length with:
//     len(mockedSmartState.NewestFullWavesBalanceCalls())
func (mock *MockSmartState) NewestFullWavesBalanceCalls() []struct {
	Account proto.Recipient
} {
	var calls []struct {
		Account proto.Recipient
	}
	mock.lockNewestFullWavesBalance.RLock()
	calls = mock.calls.NewestFullWavesBalance
	mock.lockNewestFullWavesBalance.RUnlock()
	return calls
}

// NewestHeaderByHeight calls NewestHeaderByHeightFunc.
func (mock *MockSmartState) NewestHeaderByHeight(height uint64) (*proto.BlockHeader, error) {
	if mock.NewestHeaderByHeightFunc == nil {
		panic("MockSmartState.NewestHeaderByHeightFunc: method is nil but SmartState.NewestHeaderByHeight was just called")
	}
	callInfo := struct {
		Height uint64
	}{
		Height: height,
	}
	mock.lockNewestHeaderByHeight.Lock()
	mock.calls.NewestHeaderByHeight = append(mock.calls.NewestHeaderByHeight, callInfo)
	mock.lockNewestHeaderByHeight.Unlock()
	return mock.NewestHeaderByHeightFunc(height)
}

// NewestHeaderByHeightCalls gets all the calls that were made to NewestHeaderByHeight.
// Check the length with:
//     len(mockedSmartState.NewestHeaderByHeightCalls())
func (mock *MockSmartState) NewestHeaderByHeightCalls() []struct {
	Height uint64
} {
	var calls []struct {
		Height uint64
	}
	mock.lockNewestHeaderByHeight.RLock()
	calls = mock.calls.NewestHeaderByHeight
	mock.lockNewestHeaderByHeight.RUnlock()
	return calls
}

// NewestTransactionByID calls NewestTransactionByIDFunc.
func (mock *MockSmartState) NewestTransactionByID(in1 []byte) (proto.Transaction, error) {
	if mock.NewestTransactionByIDFunc == nil {
		panic("MockSmartState.NewestTransactionByIDFunc: method is nil but SmartState.NewestTransactionByID was just called")
	}
	callInfo := struct {
		In1 []byte
	}{
		In1: in1,
	}
	mock.lockNewestTransactionByID.Lock()
	mock.calls.NewestTransactionByID = append(mock.calls.NewestTransactionByID, callInfo)
	mock.lockNewestTransactionByID.Unlock()
	return mock.NewestTransactionByIDFunc(in1)
}

// NewestTransactionByIDCalls gets all the calls that were made to NewestTransactionByID.
// Check the length with:
//     len(mockedSmartState.NewestTransactionByIDCalls())
func (mock *MockSmartState) NewestTransactionByIDCalls() []struct {
	In1 []byte
} {
	var calls []struct {
		In1 []byte
	}
	mock.lockNewestTransactionByID.RLock()
	calls = mock.calls.NewestTransactionByID
	mock.lockNewestTransactionByID.RUnlock()
	return calls
}

// NewestTransactionHeightByID calls NewestTransactionHeightByIDFunc.
func (mock *MockSmartState) NewestTransactionHeightByID(in1 []byte) (uint64, error) {
	if mock.NewestTransactionHeightByIDFunc == nil {
		panic("MockSmartState.NewestTransactionHeightByIDFunc: method is nil but SmartState.NewestTransactionHeightByID was just called")
	}
	callInfo := struct {
		In1 []byte
	}{
		In1: in1,
	}
	mock.lockNewestTransactionHeightByID.Lock()
	mock.calls.NewestTransactionHeightByID = append(mock.calls.NewestTransactionHeightByID, callInfo)
	mock.lockNewestTransactionHeightByID.Unlock()
	return mock.NewestTransactionHeightByIDFunc(in1)
}

// NewestTransactionHeightByIDCalls gets all the calls that were made to NewestTransactionHeightByID.
// Check the length with:
//     len(mockedSmartState.NewestTransactionHeightByIDCalls())
func (mock *MockSmartState) NewestTransactionHeightByIDCalls() []struct {
	In1 []byte
} {
	var calls []struct {
		In1 []byte
	}
	mock.lockNewestTransactionHeightByID.RLock()
	calls = mock.calls.NewestTransactionHeightByID
	mock.lockNewestTransactionHeightByID.RUnlock()
	return calls
}

// RetrieveNewestBinaryEntry calls RetrieveNewestBinaryEntryFunc.
func (mock *MockSmartState) RetrieveNewestBinaryEntry(account proto.Recipient, key string) (*proto.BinaryDataEntry, error) {
	if mock.RetrieveNewestBinaryEntryFunc == nil {
		panic("MockSmartState.RetrieveNewestBinaryEntryFunc: method is nil but SmartState.RetrieveNewestBinaryEntry was just called")
	}
	callInfo := struct {
		Account proto.Recipient
		Key     string
	}{
		Account: account,
		Key:     key,
	}
	mock.lockRetrieveNewestBinaryEntry.Lock()
	mock.calls.RetrieveNewestBinaryEntry = append(mock.calls.RetrieveNewestBinaryEntry, callInfo)
	mock.lockRetrieveNewestBinaryEntry.Unlock()
	return mock.RetrieveNewestBinaryEntryFunc(account, key)
}

// RetrieveNewestBinaryEntryCalls gets all the calls that were made to RetrieveNewestBinaryEntry.
// Check the length with:
//     len(mockedSmartState.RetrieveNewestBinaryEntryCalls())
func (mock *MockSmartState) RetrieveNewestBinaryEntryCalls() []struct {
	Account proto.Recipient
	Key     string
} {
	var calls []struct {
		Account proto.Recipient
		Key     string
	}
	mock.lockRetrieveNewestBinaryEntry.RLock()
	calls = mock.calls.RetrieveNewestBinaryEntry
	mock.lockRetrieveNewestBinaryEntry.RUnlock()
	return calls
}

// RetrieveNewestBooleanEntry calls RetrieveNewestBooleanEntryFunc.
func (mock *MockSmartState) RetrieveNewestBooleanEntry(account proto.Recipient, key string) (*proto.BooleanDataEntry, error) {
	if mock.RetrieveNewestBooleanEntryFunc == nil {
		panic("MockSmartState.RetrieveNewestBooleanEntryFunc: method is nil but SmartState.RetrieveNewestBooleanEntry was just called")
	}
	callInfo := struct {
		Account proto.Recipient
		Key     string
	}{
		Account: account,
		Key:     key,
	}
	mock.lockRetrieveNewestBooleanEntry.Lock()
	mock.calls.RetrieveNewestBooleanEntry = append(mock.calls.RetrieveNewestBooleanEntry, callInfo)
	mock.lockRetrieveNewestBooleanEntry.Unlock()
	return mock.RetrieveNewestBooleanEntryFunc(account, key)
}

// RetrieveNewestBooleanEntryCalls gets all the calls that were made to RetrieveNewestBooleanEntry.
// Check the length with:
//     len(mockedSmartState.RetrieveNewestBooleanEntryCalls())
func (mock *MockSmartState) RetrieveNewestBooleanEntryCalls() []struct {
	Account proto.Recipient
	Key     string
} {
	var calls []struct {
		Account proto.Recipient
		Key     string
	}
	mock.lockRetrieveNewestBooleanEntry.RLock()
	calls = mock.calls.RetrieveNewestBooleanEntry
	mock.lockRetrieveNewestBooleanEntry.RUnlock()
	return calls
}

// RetrieveNewestIntegerEntry calls RetrieveNewestIntegerEntryFunc.
func (mock *MockSmartState) RetrieveNewestIntegerEntry(account proto.Recipient, key string) (*proto.IntegerDataEntry, error) {
	if mock.RetrieveNewestIntegerEntryFunc == nil {
		panic("MockSmartState.RetrieveNewestIntegerEntryFunc: method is nil but SmartState.RetrieveNewestIntegerEntry was just called")
	}
	callInfo := struct {
		Account proto.Recipient
		Key     string
	}{
		Account: account,
		Key:     key,
	}
	mock.lockRetrieveNewestIntegerEntry.Lock()
	mock.calls.RetrieveNewestIntegerEntry = append(mock.calls.RetrieveNewestIntegerEntry, callInfo)
	mock.lockRetrieveNewestIntegerEntry.Unlock()
	return mock.RetrieveNewestIntegerEntryFunc(account, key)
}

// RetrieveNewestIntegerEntryCalls gets all the calls that were made to RetrieveNewestIntegerEntry.
// Check the length with:
//     len(mockedSmartState.RetrieveNewestIntegerEntryCalls())
func (mock *MockSmartState) RetrieveNewestIntegerEntryCalls() []struct {
	Account proto.Recipient
	Key     string
} {
	var calls []struct {
		Account proto.Recipient
		Key     string
	}
	mock.lockRetrieveNewestIntegerEntry.RLock()
	calls = mock.calls.RetrieveNewestIntegerEntry
	mock.lockRetrieveNewestIntegerEntry.RUnlock()
	return calls
}

// RetrieveNewestStringEntry calls RetrieveNewestStringEntryFunc.
func (mock *MockSmartState) RetrieveNewestStringEntry(account proto.Recipient, key string) (*proto.StringDataEntry, error) {
	if mock.RetrieveNewestStringEntryFunc == nil {
		panic("MockSmartState.RetrieveNewestStringEntryFunc: method is nil but SmartState.RetrieveNewestStringEntry was just called")
	}
	callInfo := struct {
		Account proto.Recipient
		Key     string
	}{
		Account: account,
		Key:     key,
	}
	mock.lockRetrieveNewestStringEntry.Lock()
	mock.calls.RetrieveNewestStringEntry = append(mock.calls.RetrieveNewestStringEntry, callInfo)
	mock.lockRetrieveNewestStringEntry.Unlock()
	return mock.RetrieveNewestStringEntryFunc(account, key)
}

// RetrieveNewestStringEntryCalls gets all the calls that were made to RetrieveNewestStringEntry.
// Check the length with:
//     len(mockedSmartState.RetrieveNewestStringEntryCalls())
func (mock *MockSmartState) RetrieveNewestStringEntryCalls() []struct {
	Account proto.Recipient
	Key     string
} {
	var calls []struct {
		Account proto.Recipient
		Key     string
	}
	mock.lockRetrieveNewestStringEntry.RLock()
	calls = mock.calls.RetrieveNewestStringEntry
	mock.lockRetrieveNewestStringEntry.RUnlock()
	return calls
}
