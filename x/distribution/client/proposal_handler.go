package client

import (
	"github.com/omexapp/omexchain/x/distribution/client/cli"
	"github.com/omexapp/omexchain/x/distribution/client/rest"
	govclient "github.com/omexapp/omexchain/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
