package caddyhttp

import (
	// http.prometheus
	_ "github.com/miekg/caddy-prometheus"
	// http.ipfilter
	_ "github.com/pyed/ipfilter"
	_ "github.com/caddyserver/dnsproviders/cloudflare"
	_ "github.com/caddyserver/dnsproviders/digitalocean"
	_ "github.com/caddyserver/dnsproviders/dnsimple"
	_ "github.com/caddyserver/dnsproviders/godaddy"
	_ "github.com/caddyserver/dnsproviders/route53"
)
