package __obf_514f52764222e263

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func __obf_9bff493e296f3514() {
	var __obf_3ef3afd9b0612832 = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_aeb81d7f656295d2 = flag.String("f", "", "BPF filter")
	var __obf_ca1c9670be624c55 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_204931d1c28353cc = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()

	__obf_f677e1d7ceb95fae, __obf_ca77dfc0c417cd09 := pcap.OpenLive(*__obf_3ef3afd9b0612832, int32(*__obf_ca1c9670be624c55), true, pcap.BlockForever)
	if __obf_ca77dfc0c417cd09 != nil {
		log.Printf("Error: %s\n", __obf_ca77dfc0c417cd09)
		return
	}
	defer __obf_f677e1d7ceb95fae.Close()

	if *__obf_aeb81d7f656295d2 != "" {
		if __obf_ca77dfc0c417cd09 := __obf_f677e1d7ceb95fae.SetBPFFilter(*__obf_aeb81d7f656295d2); __obf_ca77dfc0c417cd09 != nil {
			panic(__obf_ca77dfc0c417cd09)
		}
	}

	//Create a new PacketDataSource
	__obf_bdf4f00587266a50 := gopacket.NewPacketSource(__obf_f677e1d7ceb95fae, layers.LayerTypeEthernet)
	//Packets returns a channel of packets
	__obf_005c5905bf14a0d4 := __obf_bdf4f00587266a50.Packets()

	for {
		// var packet gopacket.Packet
		__obf_6963a2501bd0b5d1 := <-__obf_005c5905bf14a0d4
		if *__obf_204931d1c28353cc {
			log.Println(__obf_6963a2501bd0b5d1)
		}
		DstIpProxy(__obf_f677e1d7ceb95fae, __obf_6963a2501bd0b5d1)
	}
}

func DstIpProxy(__obf_f677e1d7ceb95fae *pcap.Handle, __obf_6963a2501bd0b5d1 gopacket.Packet) {
	// dstIp, err := net.LookupHost(domain)
	// if err != nil {
	// 	return nil, err
	// }

	// get all layers
	__obf_ac3801b02fad489d := __obf_6963a2501bd0b5d1.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_ca77dfc0c417cd09 error
	for _, __obf_06819c6414cdd56a := range __obf_ac3801b02fad489d {
		switch __obf_06819c6414cdd56a.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_68abc112f2de92ba := __obf_06819c6414cdd56a.(*layers.IPv4)
			__obf_68abc112f2de92ba.DstIP = net.IPv4(127, 0, 0, 1)

			__obf_e0fb402aaa51eff6 := gopacket.NewSerializeBuffer()
			__obf_ca77dfc0c417cd09 = __obf_68abc112f2de92ba.SerializeTo(__obf_e0fb402aaa51eff6, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_ca77dfc0c417cd09 != nil {
				log.Printf("Error: %s\n", __obf_ca77dfc0c417cd09)
			}
			// send the packet
			__obf_ca77dfc0c417cd09 = __obf_f677e1d7ceb95fae.WritePacketData(__obf_e0fb402aaa51eff6.Bytes())
			if __obf_ca77dfc0c417cd09 != nil {
				log.Printf("Error: %s\n", __obf_ca77dfc0c417cd09)
			}

		case layers.LayerTypeIPv6:
			__obf_68abc112f2de92ba := __obf_06819c6414cdd56a.(*layers.IPv6)
			__obf_68abc112f2de92ba.DstIP = net.IPv6loopback

			__obf_e0fb402aaa51eff6 := gopacket.NewSerializeBuffer()
			__obf_ca77dfc0c417cd09 = __obf_68abc112f2de92ba.SerializeTo(__obf_e0fb402aaa51eff6, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_ca77dfc0c417cd09 != nil {
				log.Printf("Error: %s\n", __obf_ca77dfc0c417cd09)
			}
			// send the packet
			__obf_ca77dfc0c417cd09 = __obf_f677e1d7ceb95fae.WritePacketData(__obf_e0fb402aaa51eff6.Bytes())
			if __obf_ca77dfc0c417cd09 != nil {
				log.Printf("Error: %s\n", __obf_ca77dfc0c417cd09)
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
