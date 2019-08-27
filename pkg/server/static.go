package server

import (
	"context"

	"github.com/operator-framework/operator-registry/pkg/api"
	"github.com/operator-framework/operator-registry/pkg/registry"
)

type LogoServer struct {
	store registry.Query
}

func NewLogoServer() *LogoServer {
	return &LogoServer{}
}

// TODO(alecmerdler): Serve logos from static HTTP endpoint
func (s *LogoServer) logoFor(pkgName, channelName, csvName string) {
	bundleString, err := s.store.GetBundle(context.Background(), pkgName, channelName, csvName)
	if err != nil {
		return nil, err
	}
	entry := &registry.ChannelEntry{
		PackageName: pkgName,
		ChannelName: channelName,
	}
	bundle, err := api.BundleStringToAPIBundle(bundleString, entry)

}
