package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
)

func TestMain_HandleRequest(t *testing.T) {
	ret, err := handleRequest(nil, events.APIGatewayProxyRequest{Path: "/cothority/"})
	require.NoError(t, err)
	require.Contains(t, ret.Body, `<meta http-equiv="refresh" content="0; url=https://godoc.org/go2.dedis.ch/cothority/">`)

	ret, err = handleRequest(nil, events.APIGatewayProxyRequest{Path: "/cothority/v3/blscosi"})
	require.NoError(t, err)
	require.Contains(t, ret.Body, `<meta http-equiv="refresh" content="0; url=https://godoc.org/go2.dedis.ch/cothority/blscosi">`)

	ret, err = handleRequest(nil, events.APIGatewayProxyRequest{Path: "/abc"})
	require.NoError(t, err)
	require.Contains(t, ret.Body, "<h1>404 Page Not Found</h1>")

	ret, err = handleRequest(nil, events.APIGatewayProxyRequest{Path: "///"})
	require.NoError(t, err)
	require.Contains(t, ret.Body, `<meta http-equiv="refresh" content="0; url=https://dedis.epfl.ch">`)

	ret, err = handleRequest(nil, events.APIGatewayProxyRequest{Path: ""})
	require.NoError(t, err)
	require.Contains(t, ret.Body, `<meta http-equiv="refresh" content="0; url=https://dedis.epfl.ch">`)
}
