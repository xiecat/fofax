package retryabledns

import (
	"net"
	"strings"

	"github.com/projectdiscovery/stringsutil"
)

type Protocol int

const (
	UDP Protocol = iota
	TCP
	DOH
)

func (p Protocol) String() string {
	switch p {
	case DOH:
		return "doh"
	case UDP:
		return "udp"
	case TCP:
		return "tcp"
	}

	return ""
}

func (p Protocol) StringWithSemicolon() string {
	return p.String() + ":"
}

type DohProtocol int

const (
	JsonAPI DohProtocol = iota
	GET
	POST
)

func (p DohProtocol) String() string {
	switch p {
	case JsonAPI:
		return "jsonapi"
	case GET:
		return "get"
	case POST:
		return "post"
	}

	return ""
}

func (p DohProtocol) StringWithSemicolon() string {
	return ":" + p.String()
}

type Resolver interface {
	String() string
}

type NetworkResolver struct {
	Protocol Protocol
	Host     string
	Port     string
}

func (r NetworkResolver) String() string {
	return net.JoinHostPort(r.Host, r.Port)
}

type DohResolver struct {
	Protocol DohProtocol
	URL      string
}

func (r DohResolver) Method() string {
	if r.Protocol == POST {
		return POST.String()
	}

	return GET.String()
}

func (r DohResolver) String() string {
	return r.URL
}

func parseResolver(r string) (resolver Resolver) {
	isTcp, isUDP, isDoh := hasProtocol(r, TCP.StringWithSemicolon()), hasProtocol(r, UDP.StringWithSemicolon()), hasProtocol(r, DOH.StringWithSemicolon())
	rNetworkTokens := trimProtocol(r)
	if isTcp || isUDP {
		networkResolver := &NetworkResolver{Protocol: UDP}
		if isTcp {
			networkResolver.Protocol = TCP
		}
		parseHostPort(networkResolver, rNetworkTokens)
		resolver = networkResolver
	} else if isDoh {
		isJsonApi, isGet := hasDohProtocol(r, JsonAPI.StringWithSemicolon()), hasDohProtocol(r, GET.StringWithSemicolon())
		URL := trimDohProtocol(rNetworkTokens)
		dohResolver := &DohResolver{URL: URL, Protocol: POST}
		if isJsonApi {
			dohResolver.Protocol = JsonAPI
		} else if isGet {
			dohResolver.Protocol = GET
		}
		resolver = dohResolver
	} else {
		networkResolver := &NetworkResolver{Protocol: UDP}
		parseHostPort(networkResolver, rNetworkTokens)
		resolver = networkResolver
	}

	return
}

func parseHostPort(networkResolver *NetworkResolver, r string) {
	if host, port, err := net.SplitHostPort(r); err == nil {
		networkResolver.Host = host
		networkResolver.Port = port
	} else {
		networkResolver.Host = r
		networkResolver.Port = "53"
	}
}

func hasProtocol(resolver, protocol string) bool {
	return strings.HasPrefix(resolver, protocol)
}

func hasDohProtocol(resolver, protocol string) bool {
	return strings.HasSuffix(resolver, protocol)
}

func trimProtocol(resolver string) string {
	return stringsutil.TrimPrefixAny(resolver, TCP.StringWithSemicolon(), UDP.StringWithSemicolon(), DOH.StringWithSemicolon())
}

func trimDohProtocol(resolver string) string {
	return stringsutil.TrimSuffixAny(resolver, GET.StringWithSemicolon(), POST.StringWithSemicolon(), JsonAPI.StringWithSemicolon())
}

func parseResolvers(resolvers []string) []Resolver {
	var parsedResolvers []Resolver
	for _, resolver := range resolvers {
		parsedResolvers = append(parsedResolvers, parseResolver(resolver))
	}
	return parsedResolvers
}
