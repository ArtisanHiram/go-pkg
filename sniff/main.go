package __obf_612b7307028a9582

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var __obf_626395acfe95b68d = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_10393cfff8a28528 = flag.String("f", "", "BPF filter")
	var __obf_a3060fd4f94cd7c3 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_162dea07e491df8d = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()
	__obf_23e5ada77eb9ab98, __obf_8cc2bee0071b574d := pcap.OpenLive(*__obf_626395acfe95b68d, int32(*__obf_a3060fd4f94cd7c3), true, pcap.BlockForever)
	if __obf_8cc2bee0071b574d != nil {
		log.Printf("Error: %s\n", __obf_8cc2bee0071b574d)
		return
	}
	defer __obf_23e5ada77eb9ab98.Close()

	if *__obf_10393cfff8a28528 != "" {
		if __obf_8cc2bee0071b574d := __obf_23e5ada77eb9ab98.SetBPFFilter(*__obf_10393cfff8a28528); __obf_8cc2bee0071b574d != nil {
			panic(__obf_8cc2bee0071b574d)
		}
	}
	__obf_a54ccb02379def65 := //Create a new PacketDataSource
		gopacket.NewPacketSource(__obf_23e5ada77eb9ab98, layers.LayerTypeEthernet)
	__obf_85ed0eb8c2bc19fc := //Packets returns a channel of packets
		__obf_a54ccb02379def65.Packets()

	for {
		__obf_f8dd1178274a705e := // var packet gopacket.Packet
			<-__obf_85ed0eb8c2bc19fc
		if *__obf_162dea07e491df8d {
			log.Println(__obf_f8dd1178274a705e)
		}
		DstIpProxy(__obf_23e5ada77eb9ab98, __obf_f8dd1178274a705e)
	}
}

func DstIpProxy(__obf_23e5ada77eb9ab98 *pcap.Handle, __obf_f8dd1178274a705e gopacket.Packet) {
	__obf_c9bfcac6e09e58d7 := // dstIp, err := net.LookupHost(domain)
		// if err != nil {
		// 	return nil, err
		// }

		// get all layers
		__obf_f8dd1178274a705e.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_8cc2bee0071b574d error
	for _, __obf_95b2b4a1f00593cd := range __obf_c9bfcac6e09e58d7 {
		switch __obf_95b2b4a1f00593cd.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_74df6b618c813eb3 := __obf_95b2b4a1f00593cd.(*layers.IPv4)
			__obf_74df6b618c813eb3.
				DstIP = net.IPv4(127, 0, 0, 1)
			__obf_0354db90844a6fd8 := gopacket.NewSerializeBuffer()
			__obf_8cc2bee0071b574d = __obf_74df6b618c813eb3.SerializeTo(__obf_0354db90844a6fd8, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_8cc2bee0071b574d != nil {
				log.Printf("Error: %s\n", __obf_8cc2bee0071b574d)
			}
			__obf_8cc2bee0071b574d = // send the packet
				__obf_23e5ada77eb9ab98.WritePacketData(__obf_0354db90844a6fd8.Bytes())
			if __obf_8cc2bee0071b574d != nil {
				log.Printf("Error: %s\n", __obf_8cc2bee0071b574d)
			}

		case layers.LayerTypeIPv6:
			__obf_74df6b618c813eb3 := __obf_95b2b4a1f00593cd.(*layers.IPv6)
			__obf_74df6b618c813eb3.
				DstIP = net.IPv6loopback
			__obf_0354db90844a6fd8 := gopacket.NewSerializeBuffer()
			__obf_8cc2bee0071b574d = __obf_74df6b618c813eb3.SerializeTo(__obf_0354db90844a6fd8, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_8cc2bee0071b574d != nil {
				log.Printf("Error: %s\n", __obf_8cc2bee0071b574d)
			}
			__obf_8cc2bee0071b574d = // send the packet
				__obf_23e5ada77eb9ab98.WritePacketData(__obf_0354db90844a6fd8.Bytes())
			if __obf_8cc2bee0071b574d != nil {
				log.Printf("Error: %s\n", __obf_8cc2bee0071b574d)
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
