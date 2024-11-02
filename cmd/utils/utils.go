package utils

import (
	"github.com/XinFinOrg/XDC-Subnet/DCx"
	"github.com/XinFinOrg/XDC-Subnet/DCxlending"
	"github.com/XinFinOrg/XDC-Subnet/eth"
	"github.com/XinFinOrg/XDC-Subnet/eth/downloader"
	"github.com/XinFinOrg/XDC-Subnet/ethstats"
	"github.com/XinFinOrg/XDC-Subnet/les"
	"github.com/XinFinOrg/XDC-Subnet/node"
	whisper "github.com/XinFinOrg/XDC-Subnet/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var XDCXServ *DCx.XDCX
			ctx.Service(&XDCXServ)
			var lendingServ *DCxlending.Lending
			ctx.Service(&lendingServ)
			fullNode, err := eth.New(ctx, cfg, XDCXServ, lendingServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterXDCXService(stack *node.Node, cfg *DCx.Config) {
	XDCX := DCx.New(cfg)
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return XDCX, nil
	}); err != nil {
		Fatalf("Failed to register the XDCX service: %v", err)
	}

	// register DCxlending service
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return DCxlending.New(XDCX), nil
	}); err != nil {
		Fatalf("Failed to register the XDCXLending service: %v", err)
	}
}
