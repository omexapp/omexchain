package client

import (
	"github.com/omexapp/omexchain/x/dex/client/cli"
	"github.com/omexapp/omexchain/x/dex/client/rest"
	govclient "github.com/omexapp/omexchain/x/gov/client"
)

// param change proposal handler
var (
	// DelistProposalHandler alias gov NewProposalHandler
	DelistProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDelistProposal, rest.DelistProposalRESTHandler)
)
