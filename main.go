package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"net"
	"os"
	"time"
)

var (
	VerboseMode bool

	snapshot_len int32         = 64 // limit of length of packets captured
	promiscuous  bool          = true
	timeout      time.Duration = 30 * time.Second

	targetMACAddress net.HardwareAddr
)

func main() {
	flag.BoolVar(&VerboseMode, "v", false, "verbose output")
	flag.Parse()

	packetCapture()
}

func packetCapture() {
	// Open device
	handle, err := pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet
		ethernetPacket, _ := packet.Layer(layers.LayerTypeEthernet).(*layers.Ethernet)
		if bytes.Equal(ethernetPacket.SrcMAC, targetMACAddress) {
			LLCLayer := packet.Layer(layers.LayerTypeLLC)
			if LLCLayer != nil {
				if VerboseMode {
					fmt.Println(packet)
				}
				routine()
			}
		}
	}
}

func init() {
	var err error
	targetMACAddress, err = net.ParseMAC(mac)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
