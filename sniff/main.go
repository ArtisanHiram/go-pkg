package __obf_e49a3ffbc913c7ba

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var __obf_b197ee60d891e76f = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_4dfffe3c7dce124f = flag.String("f", "", "BPF filter")
	var __obf_65e01100635b2be7 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_abae2b059a56bfb9 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()
	__obf_62ccd8959e87f5ff, __obf_7811bf5a95f56bea := pcap.OpenLive(*__obf_b197ee60d891e76f, int32(*__obf_65e01100635b2be7), true, pcap.BlockForever)
	if __obf_7811bf5a95f56bea != nil {
		log.Printf("Error: %s\n", __obf_7811bf5a95f56bea)
		return
	}
	defer __obf_62ccd8959e87f5ff.Close()

	if *__obf_4dfffe3c7dce124f != "" {
		if __obf_7811bf5a95f56bea := __obf_62ccd8959e87f5ff.SetBPFFilter(*__obf_4dfffe3c7dce124f); __obf_7811bf5a95f56bea != nil {
			panic(__obf_7811bf5a95f56bea)
		}
	}
	__obf_d44f86ad42838fb2 := //Create a new PacketDataSource
		gopacket.NewPacketSource(__obf_62ccd8959e87f5ff, layers.LayerTypeEthernet)
	__obf_33c648581f5c3e21 := //Packets returns a channel of packets
		__obf_d44f86ad42838fb2.Packets()

	for {
		__obf_3e00a507b3752304 := // var packet gopacket.Packet
			<-__obf_33c648581f5c3e21
		if *__obf_abae2b059a56bfb9 {
			log.Println(__obf_3e00a507b3752304)
		}
		DstIpProxy(__obf_62ccd8959e87f5ff, __obf_3e00a507b3752304)
	}
}

func DstIpProxy(__obf_62ccd8959e87f5ff *pcap.Handle, __obf_3e00a507b3752304 gopacket.Packet) {
	__obf_7a97d2ec1d27f498 := // dstIp, err := net.LookupHost(domain)
		// if err != nil {
		// 	return nil, err
		// }

		// get all layers
		__obf_3e00a507b3752304.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_7811bf5a95f56bea error
	for _, __obf_1ca743e42002c489 := range __obf_7a97d2ec1d27f498 {
		switch __obf_1ca743e42002c489.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_fef653716f0086b7 := __obf_1ca743e42002c489.(*layers.IPv4)
			__obf_fef653716f0086b7.
				DstIP = net.IPv4(127, 0, 0, 1)
			__obf_5a922ce9f88e4bf7 := gopacket.NewSerializeBuffer()
			__obf_7811bf5a95f56bea = __obf_fef653716f0086b7.SerializeTo(__obf_5a922ce9f88e4bf7, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_7811bf5a95f56bea != nil {
				log.Printf("Error: %s\n", __obf_7811bf5a95f56bea)
			}
			__obf_7811bf5a95f56bea = // send the packet
				__obf_62ccd8959e87f5ff.WritePacketData(__obf_5a922ce9f88e4bf7.Bytes())
			if __obf_7811bf5a95f56bea != nil {
				log.Printf("Error: %s\n", __obf_7811bf5a95f56bea)
			}

		case layers.LayerTypeIPv6:
			__obf_fef653716f0086b7 := __obf_1ca743e42002c489.(*layers.IPv6)
			__obf_fef653716f0086b7.
				DstIP = net.IPv6loopback
			__obf_5a922ce9f88e4bf7 := gopacket.NewSerializeBuffer()
			__obf_7811bf5a95f56bea = __obf_fef653716f0086b7.SerializeTo(__obf_5a922ce9f88e4bf7, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_7811bf5a95f56bea != nil {
				log.Printf("Error: %s\n", __obf_7811bf5a95f56bea)
			}
			__obf_7811bf5a95f56bea = // send the packet
				__obf_62ccd8959e87f5ff.WritePacketData(__obf_5a922ce9f88e4bf7.Bytes())
			if __obf_7811bf5a95f56bea != nil {
				log.Printf("Error: %s\n", __obf_7811bf5a95f56bea)
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
