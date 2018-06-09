ovpnstats [![GoDoc](https://godoc.org/github.com/Thor77/ovpnstats?status.svg)](https://godoc.org/github.com/Thor77/ovpnstats)
=========
Parse OpenVPNs `openvpn-status.log`.

# Installation
`$ go get github.com/thor77/ovpnstats`

# Example
```go
package main

import (
        "fmt"
        "github.com/thor77/ovpnstats"
)

func main() {
        clients, routes, _g := ovpnstats.ParseStatusFile("openvpn-status.log")
        // TODO: check error
        fmt.Printf("%#v\n", clients)
        /*
        []ovpnstats.ClientInfo{
                ovpnstats.ClientInfo{
                        Name:"client1",
                        RealAddress:"2001::1",
                        VirtualAddress:"10.0.0.2",
                        VirtualV6Address:"2002::2",
                        BytesReceived:0,
                        BytesSent:0,
                        ConnectedSince:time.Time{wall:0x0, ext:0, loc:(*time.Location)(0x5410a0)},
                        Username:"UNDEF",
                        ClientID:1,
                        PeerID:0
                }
        }
        */
        fmt.Printf("%#v\n", routes)?
        /*
        []ovpnstats.RoutingInfo{
                ovpnstats.RoutingInfo{
                        VirtualAddress:"10.0.0.2",
                        CommonName:"client1",
                        RealAddress:"2001::1",
                        LastRef:time.Time{wall:0x0, ext:0, loc:(*time.Location)(0x5410a0)
                }
        },
        */
}

```