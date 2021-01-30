package client

import (
	govclient "github.com/omexapp/omexchain/x/gov/client"
	"github.com/omexapp/omexchain/x/params/client/cli"
	"github.com/omexapp/omexchain/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler in cmsdk
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
