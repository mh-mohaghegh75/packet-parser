package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {

	start := time.Now()
	if len(os.Args) < 2 {
		fmt.Println("Please provide a pcap file to read!")
		os.Exit(1)
	}

	handle, err := pcap.OpenOffline(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet.String())

		ethLayer := packet.Layer(layers.LayerTypeEthernet)
		if ethLayer != nil {
			ethPacket, _ := ethLayer.(*layers.Ethernet)
			fmt.Println("Ethernet source MAC address:", ethPacket.SrcMAC)
			fmt.Println("Ethernet destination MAC address:", ethPacket.DstMAC)
		}

		ipLayer := packet.Layer(layers.LayerTypeIPv4)
		if ipLayer != nil {
			ipPacket, _ := ipLayer.(*layers.IPv4)
			fmt.Println("IP source address:", ipPacket.SrcIP)
			fmt.Println("IP destination address:", ipPacket.DstIP)
		}
		fmt.Println("--------------------------------")

	}

	fmt.Println("Hello again")
	end := time.Since(start) / 1000
	fmt.Printf("The execition time is %d", end)
}
