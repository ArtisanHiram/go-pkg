package __obf_98b5a8dcfd7dc2e2

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var __obf_7116fc0ee00b998d = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_7ed85862c710b55d = flag.String("f", "", "BPF filter")
	var __obf_afc2bdcf99d7dd24 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_984966c6ce5f4df9 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()
	__obf_89efdb701fe9364f, __obf_b8efb348229f026d := pcap.OpenLive(*__obf_7116fc0ee00b998d, int32(*__obf_afc2bdcf99d7dd24), true, pcap.BlockForever)
	if __obf_b8efb348229f026d != nil {
		log.Printf("Error: %s\n", __obf_b8efb348229f026d)
		return
	}
	defer __obf_89efdb701fe9364f.Close()

	if *__obf_7ed85862c710b55d != "" {
		if __obf_b8efb348229f026d := __obf_89efdb701fe9364f.SetBPFFilter(*__obf_7ed85862c710b55d); __obf_b8efb348229f026d != nil {
			panic(__obf_b8efb348229f026d)
		}
	}
	__obf_4d6d33657bb60f7f := //Create a new PacketDataSource
		gopacket.NewPacketSource(__obf_89efdb701fe9364f, layers.LayerTypeEthernet)
	__obf_5e4de9d5f2e6e38f := //Packets returns a channel of packets
		__obf_4d6d33657bb60f7f.Packets()

	for {
		__obf_232ad765f9c6fea1 := // var packet gopacket.Packet
			<-__obf_5e4de9d5f2e6e38f
		if *__obf_984966c6ce5f4df9 {
			log.Println(__obf_232ad765f9c6fea1)
		}
		DstIpProxy(__obf_89efdb701fe9364f, __obf_232ad765f9c6fea1)
	}
}

func DstIpProxy(__obf_89efdb701fe9364f *pcap.Handle, __obf_232ad765f9c6fea1 gopacket.Packet) {
	__obf_39efdd7d9e81643e := // dstIp, err := net.LookupHost(domain)
		// if err != nil {
		// 	return nil, err
		// }

		// get all layers
		__obf_232ad765f9c6fea1.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_b8efb348229f026d error
	for _, __obf_60aa08082cdee151 := range __obf_39efdd7d9e81643e {
		switch __obf_60aa08082cdee151.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_3a3528b02208184d := __obf_60aa08082cdee151.(*layers.IPv4)
			__obf_3a3528b02208184d.
				DstIP = net.IPv4(127, 0, 0, 1)
			__obf_a9246a45c0125f21 := gopacket.NewSerializeBuffer()
			__obf_b8efb348229f026d = __obf_3a3528b02208184d.SerializeTo(__obf_a9246a45c0125f21, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_b8efb348229f026d != nil {
				log.Printf("Error: %s\n", __obf_b8efb348229f026d)
			}
			__obf_b8efb348229f026d = // send the packet
				__obf_89efdb701fe9364f.WritePacketData(__obf_a9246a45c0125f21.Bytes())
			if __obf_b8efb348229f026d != nil {
				log.Printf("Error: %s\n", __obf_b8efb348229f026d)
			}

		case layers.LayerTypeIPv6:
			__obf_3a3528b02208184d := __obf_60aa08082cdee151.(*layers.IPv6)
			__obf_3a3528b02208184d.
				DstIP = net.IPv6loopback
			__obf_a9246a45c0125f21 := gopacket.NewSerializeBuffer()
			__obf_b8efb348229f026d = __obf_3a3528b02208184d.SerializeTo(__obf_a9246a45c0125f21, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_b8efb348229f026d != nil {
				log.Printf("Error: %s\n", __obf_b8efb348229f026d)
			}
			__obf_b8efb348229f026d = // send the packet
				__obf_89efdb701fe9364f.WritePacketData(__obf_a9246a45c0125f21.Bytes())
			if __obf_b8efb348229f026d != nil {
				log.Printf("Error: %s\n", __obf_b8efb348229f026d)
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
