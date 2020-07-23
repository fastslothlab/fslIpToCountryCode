package fslIpToCountryCode

import (
	"net"
	"sync"
	"sync/atomic"
	"unsafe"
)

func EnsureInit(){
	gEnsureInitOnce.Do(func(){
		thisReader:=getGeoip2Reader()
		SetReader(thisReader)
	})
}

// iso code look like CN, US
func MustGetCountryIsoCode(ip net.IP) (code string) {
	EnsureInit()
	return GetReader().MustGetCountryIsoCode(ip)
}

func MustGetCountryIsoCodeByString(ip string) (code string){
	EnsureInit()
	return GetReader().MustGetCountryIsoCodeByString(ip)
}

func GetCountryCodeList() []string{
	EnsureInit()
	return GetReader().GetAllCountryCode()
}

var gReaderL2 unsafe.Pointer

func SetReader(r *Reader){
	if r==nil{
		panic("[fslIpToCountryCode.SetReader] r==nil")
	}
	r.CompatibleToOldVersion()
	atomic.StorePointer(&gReaderL2,unsafe.Pointer(r))
}

func GetReader() *Reader{
	thisReaderL1:=atomic.LoadPointer(&gReaderL2)
	return (*Reader)(thisReaderL1)
}

var gEnsureInitOnce sync.Once
