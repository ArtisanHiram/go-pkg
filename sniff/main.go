package __obf_47ac6ae7cd6db676

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func __obf_d1d4b31b449cf9f0() {
	var __obf_5567acb01b0d41d7 = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_2cc0d88e990bd5c4 = flag.String("f", "", "BPF filter")
	var __obf_160c9e097ba6150a = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_d8ea0f4e49f6ef91 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()

	__obf_21033b5ec05bcc05, __obf_75ff17729937c31c := pcap.OpenLive(*__obf_5567acb01b0d41d7, int32(*__obf_160c9e097ba6150a), true, pcap.BlockForever)
	if __obf_75ff17729937c31c != nil {
		log.Printf("Error: %s\n", __obf_75ff17729937c31c)
		return
	}
	defer __obf_21033b5ec05bcc05.Close()

	if *__obf_2cc0d88e990bd5c4 != "" {
		if __obf_75ff17729937c31c := __obf_21033b5ec05bcc05.SetBPFFilter(*__obf_2cc0d88e990bd5c4); __obf_75ff17729937c31c != nil {
			panic(__obf_75ff17729937c31c)
		}
	}

	//Create a new PacketDataSource
	__obf_f3074cfce13e8eae := gopacket.NewPacketSource(__obf_21033b5ec05bcc05, layers.LayerTypeEthernet)
	//Packets returns a channel of packets
	__obf_a6206a0c25d0cb36 := __obf_f3074cfce13e8eae.Packets()

	for {
		// var packet gopacket.Packet
		__obf_2e504760e0e7217a := <-__obf_a6206a0c25d0cb36
		if *__obf_d8ea0f4e49f6ef91 {
			log.Println(__obf_2e504760e0e7217a)
		}
		DstIpProxy(__obf_21033b5ec05bcc05, __obf_2e504760e0e7217a)
	}
}

func DstIpProxy(__obf_21033b5ec05bcc05 *pcap.Handle, __obf_2e504760e0e7217a gopacket.Packet) {
	// dstIp, err := net.LookupHost(domain)
	// if err != nil {
	// 	return nil, err
	// }

	// get all layers
	__obf_342346aab973413f := __obf_2e504760e0e7217a.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_75ff17729937c31c error
	for _, __obf_26095da8239d84a3 := range __obf_342346aab973413f {
		switch __obf_26095da8239d84a3.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_f1bc79b7b13e8ed7 := __obf_26095da8239d84a3.(*layers.IPv4)
			__obf_f1bc79b7b13e8ed7.DstIP = net.IPv4(127, 0, 0, 1)

			__obf_0d718278db197660 := gopacket.NewSerializeBuffer()
			__obf_75ff17729937c31c = __obf_f1bc79b7b13e8ed7.SerializeTo(__obf_0d718278db197660, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_75ff17729937c31c != nil {
				log.Printf("Error: %s\n", __obf_75ff17729937c31c)
			}
			// send the packet
			__obf_75ff17729937c31c = __obf_21033b5ec05bcc05.WritePacketData(__obf_0d718278db197660.Bytes())
			if __obf_75ff17729937c31c != nil {
				log.Printf("Error: %s\n", __obf_75ff17729937c31c)
			}

		case layers.LayerTypeIPv6:
			__obf_f1bc79b7b13e8ed7 := __obf_26095da8239d84a3.(*layers.IPv6)
			__obf_f1bc79b7b13e8ed7.DstIP = net.IPv6loopback

			__obf_0d718278db197660 := gopacket.NewSerializeBuffer()
			__obf_75ff17729937c31c = __obf_f1bc79b7b13e8ed7.SerializeTo(__obf_0d718278db197660, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_75ff17729937c31c != nil {
				log.Printf("Error: %s\n", __obf_75ff17729937c31c)
			}
			// send the packet
			__obf_75ff17729937c31c = __obf_21033b5ec05bcc05.WritePacketData(__obf_0d718278db197660.Bytes())
			if __obf_75ff17729937c31c != nil {
				log.Printf("Error: %s\n", __obf_75ff17729937c31c)
			}
		}

		// if l, ok := v.(gopacket.SerializableLayer); ok {
		// 	replaceLayer = append(replaceLayer, l)
		// }
	}

	// buffer := gopacket.NewSerializeBuffer()
	// err := gopacket.SerializeLayers(buffer, gopacket.SerializeOptions{}, replaceLayer...)
	// if err != nil {
	// 	return nil, err
	// }
	// p := gopacket.NewPacket(buffer.Bytes(), layers.LayerTypeIPv4, gopacket.Default)
	// if p.ErrorLayer() != nil {
	// 	return nil, err
	// }
}
