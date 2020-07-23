fslIpToCountryCode
===============
* A golang library that embed maxmind GeoLite2 Country database. You do not need to download some files to your server, just import this library.

## install
```
go get github.com/fastslothlab/fslIpToCountryCode
```

## usage example
```
package main

import (
	"fmt"
	"github.com/fastslothlab/fslIpToCountryCode"
)

func main(){
	fmt.Println(fslIpToCountryCode.MustGetCountryIsoCodeByString("74.125.204.103"))
}
```

## reference
This product includes GeoLite2 data created by MaxMind, available from
<a href="https://www.maxmind.com">https://www.maxmind.com</a>.