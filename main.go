// Package ovpnstats provides an interface to parse the openvpn-status.log file
package ovpnstats

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// ClientInfo represents a CLIENT_LIST entry
// HEADER,CLIENT_LIST,Common Name,Real Address,Virtual Address,Virtual IPv6 Address,Bytes Received,Bytes Sent,Connected Since,Connected Since (time_t),Username,Client ID,Peer ID
type ClientInfo struct {
	Name             string
	RealAddress      string
	VirtualAddress   string
	VirtualV6Address string
	BytesReceived    int64
	BytesSent        int64
	ConnectedSince   time.Time
	Username         string
	ClientID         int32
	PeerID           int32
}

// RoutingInfo represents a ROUTING_TABLE entry
// HEADER,ROUTING_TABLE,Virtual Address,Common Name,Real Address,Last Ref,Last Ref (time_t)
type RoutingInfo struct {
	VirtualAddress string
	CommonName     string
	RealAddress    string
	LastRef        time.Time
}

func parseClientListEntry(line string) (ClientInfo, error) {
	parts := strings.Split(line, ",")
	bytesReceived, err := strconv.Atoi(parts[5])
	if err != nil {
		return ClientInfo{}, err
	}
	bytesSent, err := strconv.Atoi(parts[6])
	if err != nil {
		return ClientInfo{}, err
	}
	connectedSinceUnix, err := strconv.Atoi(parts[8])
	if err != nil {
		return ClientInfo{}, err
	}
	clientID, err := strconv.Atoi(parts[10])
	if err != nil {
		return ClientInfo{}, err
	}
	peerID, err := strconv.Atoi(parts[11])
	if err != nil {
		return ClientInfo{}, err
	}
	info := ClientInfo{
		Name:             parts[1],
		RealAddress:      parts[2],
		VirtualAddress:   parts[3],
		VirtualV6Address: parts[4],
		BytesReceived:    int64(bytesReceived),
		BytesSent:        int64(bytesSent),
		ConnectedSince:   time.Unix(int64(connectedSinceUnix), 0),
		Username:         parts[9],
		ClientID:         int32(clientID),
		PeerID:           int32(peerID),
	}
	return info, nil
}

func parseRoutingTableEntry(line string) (RoutingInfo, error) {
	parts := strings.Split(line, ",")
	lastRefUnix, err := strconv.Atoi(parts[5])
	if err != nil {
		return RoutingInfo{}, err
	}
	info := RoutingInfo{
		VirtualAddress: parts[1],
		CommonName:     parts[2],
		RealAddress:    parts[3],
		LastRef:        time.Unix(int64(lastRefUnix), 0),
	}
	return info, nil
}

// ParseStatusFile parses the openvpn-status.log file at `filename` and returns a corresponding slice of ClientInfo and RoutingInfo objects
func ParseStatusFile(filename string) ([]ClientInfo, []RoutingInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var clients []ClientInfo
	var routes []RoutingInfo

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch parts := strings.Split(line, ","); parts[0] {
		case "HEADER":
		case "END":
			break
		default:
			switch statusType := parts[0]; statusType {
			case "CLIENT_LIST":
				info, err := parseClientListEntry(line)
				if err != nil {
					return nil, nil, err
				}
				clients = append(clients, info)
			case "ROUTING_TABLE":
				info, err := parseRoutingTableEntry(line)
				if err != nil {
					return nil, nil, err
				}
				routes = append(routes, info)
			}
		}
	}
	return clients, routes, nil
}
