package main

import (
	"bytes"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	snapshot_len     int32         = 64 // limit of length of packets captured
	promiscuous      bool          = true
	timeout          time.Duration = 30 * time.Second
	targetMACAddress []byte        = MACParser(mac)
)

func main() {
	// Open device
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet
		arpLayer := packet.Layer(layers.LayerTypeARP)
		if arpLayer != nil {
			arpPacket, _ := arpLayer.(*layers.ARP)
			if bytes.Equal(arpPacket.SourceHwAddress, targetMACAddress) {
				routine()
			}
		}
	}
}

func MACParser(m string) []byte {
	bs := make([]byte, 6)
	for i, v := range strings.Split(m, ":") {
		sv, _ := strconv.ParseInt(v, 16, 64)
		bs[i] = byte(sv)
	}
	return bs
}
