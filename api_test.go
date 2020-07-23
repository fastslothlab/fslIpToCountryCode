package fslIpToCountryCode

import (
	"net"
	"testing"
	"fmt"
)

func TestGetCountryIsoCode(ot *testing.T) {
	for _,cas:=range []struct{
		ip string
		code string
	}{
		{"74.125.204.103","US"},
		{"127.0.0.1",""},
		{"10.1.1.1",""},
		{"255.255.255.255",""},
	}{
		ok(MustGetCountryIsoCode(net.ParseIP(cas.ip)) == cas.code,cas.ip)
		ok(MustGetCountryIsoCodeByString(cas.ip)== cas.code,cas.ip)
	}
	ok(MustGetCountryIsoCodeByString("")=="")
	ok(MustGetCountryIsoCodeByString(" ")=="")
	ok(MustGetCountryIsoCode(nil)=="")
}

func ok(b bool,objList ...interface{}){
	if b==false{
		panic("hc2tehkbz5 "+fmt.Sprintln(objList...))
	}
}