package gapi

import (
	"context"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcGatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP  string
}

// extractMetadata takes the Context objects and returns  a Metadata object
func (server *Server) extractMetadata(ctx context.Context) *Metadata {

	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {

		// getting the user agent header [HTTP]
		if userAgents := md.Get(grpcGatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// getting the user agent header [gRPC]
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}

		// getting the ip address [HTTP]
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.ClientIP = clientIPs[0]
		}
	}

	// getting the ip address [gRPC]
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = p.Addr.String()
	}

	return mtdt
}
