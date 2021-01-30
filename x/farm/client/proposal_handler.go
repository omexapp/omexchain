package client

import (
	"github.com/omexapp/omexchain/x/farm/client/cli"
	"github.com/omexapp/omexchain/x/farm/client/rest"
	govcli "github.com/omexapp/omexchain/x/gov/client"
)

var (
	// ManageWhiteListProposalHandler alias gov NewProposalHandler
	ManageWhiteListProposalHandler = govcli.NewProposalHandler(cli.GetCmdManageWhiteListProposal, rest.ManageWhiteListProposalRESTHandler)
)
