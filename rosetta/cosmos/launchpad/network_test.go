package launchpad

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta"
	client "github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/generated"
	"github.com/tendermint/cosmos-rosetta-gateway/rosetta/cosmos/launchpad/client/cosmos/mocks"
)

func TestLaunchpad_NetworkList(t *testing.T) {
	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
	}

	adapter := NewLaunchpad(nil, API{}, "http://the-url", properties)

	list, err := adapter.NetworkList(context.Background(), nil)
	require.Nil(t, err)

	require.Len(t, list.NetworkIdentifiers, 1)
	require.Equal(t, list.NetworkIdentifiers[0].Network, "TheNetwork")
	require.Equal(t, list.NetworkIdentifiers[0].Blockchain, "TheBlockchain")
}

func TestLaunchpad_NetworkOptions(t *testing.T) {
	m := &mocks.TendermintAPI{}
	defer m.AssertExpectations(t)

	m.
		On("NodeInfoGet", mock.Anything).
		Return(client.InlineResponse200{
			NodeInfo: client.InlineResponse200NodeInfo{
				Version: "5",
			},
		}, nil, nil).
		Once()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(http.DefaultClient, API{Tendermint: m}, "", properties)

	options, err := adapter.NetworkOptions(context.Background(), nil)
	require.Nil(t, err)
	require.NotNil(t, options)

	require.Equal(t, &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: "1.2.5",
			NodeVersion:    "5",
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				{
					Status:     "SUCCESS",
					Successful: true,
				},
			},
			OperationTypes: properties.SupportedOperations,
		},
	}, options)
}

func TestLaunchpad_NetworkStatus(t *testing.T) {
	tsTendermint := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/net_info":
			peersContent := getContentsFromFile(t, "testdata/peers.json")
			_, err := w.Write(peersContent)
			require.NoError(t, err)
		case "/block":
			callingGenesis := r.URL.Query().Get("height") == "1"
			if callingGenesis {
				genesisContent := getContentsFromFile(t, "testdata/genesis-block.json")
				_, err := w.Write(genesisContent)
				require.NoError(t, err)
			} else {
				latestContent := getContentsFromFile(t, "testdata/latest-block.json")
				_, err := w.Write(latestContent)
				require.NoError(t, err)
			}
		}
	}))
	defer tsTendermint.Close()

	properties := rosetta.NetworkProperties{
		Blockchain: "TheBlockchain",
		Network:    "TheNetwork",
		SupportedOperations: []string{
			"Transfer",
			"Reward",
		},
	}

	adapter := NewLaunchpad(http.DefaultClient, API{}, tsTendermint.URL, properties)

	status, adapterErr := adapter.NetworkStatus(context.Background(), nil)
	require.Nil(t, adapterErr)
	require.NotNil(t, status)

	require.Equal(t, &types.NetworkStatusResponse{
		CurrentBlockIdentifier: &types.BlockIdentifier{
			Index: 1230,
			Hash:  "8FEB56E18A7B5FE53C42EEB43CD0113D24BB1B2DCEA4747004887A1464E5826C",
		},
		CurrentBlockTimestamp: 1597325577228,
		GenesisBlockIdentifier: &types.BlockIdentifier{
			Hash:  "360A1DED0DEE79A8A28FBD88517EA3B6A9719460A9BE30D8E8D786D5AD79127B",
			Index: 1,
		},
		Peers: []*types.Peer{
			{
				PeerID: "5576458aef205977e18fd50b274e9b5d9014525a",
			},
		},
	}, status)
}

func getContentsFromFile(t *testing.T, filename string) []byte {
	file, err := os.Open(filename)
	require.NoError(t, err)

	genesisContent, err := ioutil.ReadAll(file)
	return genesisContent
}
