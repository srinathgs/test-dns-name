package main

import (
	"testing"

	"golang.org/x/net/idna"
)

func TestIDNA(t *testing.T) {
	dnsName := "ßan-jöśè.com"
	t.Log(idna.Punycode.ToASCII(dnsName))
	t.Log(idna.Registration.ToASCII(dnsName))
	t.Log(idna.Lookup.ToASCII(dnsName))
	lTidn := idna.New(idna.MapForLookup(), idna.Transitional(true))
	t.Log(lTidn.ToASCII(dnsName))
	t.Log(lTidn.ToASCII("faß.de"))
	lNTidn := idna.New(idna.MapForLookup(), idna.Transitional(false))
	t.Log(lNTidn.ToASCII(dnsName))
	t.Log(lNTidn.ToASCII("faß.de"))
}
