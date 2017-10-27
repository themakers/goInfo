# Introduction

Package provides single function to obtain OS-specific information (see OSInfo struct) on major systems and 'os-release' information on
Linux (i.e. distro name/version)

## Futures

* Obtain Windows, Linux, macOS, FreeBSD OS/Kernel information
* Support for os-release on Linux systems

## Install

```sh
go get github.com/themakers/osinfo
```

## Example

```go
package main

import (
  "fmt"
  "github.com/themakers/osinfo"
)

func main() {
  osi := osinfo.GetInfo()
  fmt.Println(osi.AsJSON())
}
```

## Copyright

2012-2014 Matis Hsiao

2017 The Architect
