package __obf_a050c10b70e794fe

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var __obf_08bdcdfd6eba138f = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_ca610a5d217e6de2 = flag.String("f", "", "BPF filter")
	var __obf_ddfb39de9a43fa4d = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_f01a27256244d55d = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()
	__obf_67ab146b579ef3a1, __obf_38fcaab76f2926d6 := pcap.OpenLive(*__obf_08bdcdfd6eba138f, int32(*__obf_ddfb39de9a43fa4d), true, pcap.BlockForever)
	if __obf_38fcaab76f2926d6 != nil {
		log.Printf("Error: %s\n", __obf_38fcaab76f2926d6)
		return
	}
	defer __obf_67ab146b579ef3a1.Close()

	if *__obf_ca610a5d217e6de2 != "" {
		if __obf_38fcaab76f2926d6 := __obf_67ab146b579ef3a1.SetBPFFilter(*__obf_ca610a5d217e6de2); __obf_38fcaab76f2926d6 != nil {
			panic(__obf_38fcaab76f2926d6)
		}
	}
	__obf_b62d148cf92b27fb := //Create a new PacketDataSource
		gopacket.NewPacketSource(__obf_67ab146b579ef3a1, layers.LayerTypeEthernet)
	__obf_04fe5771d518bfdc := //Packets returns a channel of packets
		__obf_b62d148cf92b27fb.Packets()

	for {
		__obf_145f099276e96e3a := // var packet gopacket.Packet
			<-__obf_04fe5771d518bfdc
		if *__obf_f01a27256244d55d {
			log.Println(__obf_145f099276e96e3a)
		}
		DstIpProxy(__obf_67ab146b579ef3a1, __obf_145f099276e96e3a)
	}
}

func DstIpProxy(__obf_67ab146b579ef3a1 *pcap.Handle, __obf_145f099276e96e3a gopacket.Packet) {
	__obf_7f1075602e2e13f5 := // dstIp, err := net.LookupHost(domain)
		// if err != nil {
		// 	return nil, err
		// }

		// get all layers
		__obf_145f099276e96e3a.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_38fcaab76f2926d6 error
	for _, __obf_961a1e69965946a7 := range __obf_7f1075602e2e13f5 {
		switch __obf_961a1e69965946a7.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_36906111bf41a59e := __obf_961a1e69965946a7.(*layers.IPv4)
			__obf_36906111bf41a59e.
				DstIP = net.IPv4(127, 0, 0, 1)
			__obf_361638f5adeb38ab := gopacket.NewSerializeBuffer()
			__obf_38fcaab76f2926d6 = __obf_36906111bf41a59e.SerializeTo(__obf_361638f5adeb38ab, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_38fcaab76f2926d6 != nil {
				log.Printf("Error: %s\n", __obf_38fcaab76f2926d6)
			}
			__obf_38fcaab76f2926d6 = // send the packet
				__obf_67ab146b579ef3a1.WritePacketData(__obf_361638f5adeb38ab.Bytes())
			if __obf_38fcaab76f2926d6 != nil {
				log.Printf("Error: %s\n", __obf_38fcaab76f2926d6)
			}

		case layers.LayerTypeIPv6:
			__obf_36906111bf41a59e := __obf_961a1e69965946a7.(*layers.IPv6)
			__obf_36906111bf41a59e.
				DstIP = net.IPv6loopback
			__obf_361638f5adeb38ab := gopacket.NewSerializeBuffer()
			__obf_38fcaab76f2926d6 = __obf_36906111bf41a59e.SerializeTo(__obf_361638f5adeb38ab, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_38fcaab76f2926d6 != nil {
				log.Printf("Error: %s\n", __obf_38fcaab76f2926d6)
			}
			__obf_38fcaab76f2926d6 = // send the packet
				__obf_67ab146b579ef3a1.WritePacketData(__obf_361638f5adeb38ab.Bytes())
			if __obf_38fcaab76f2926d6 != nil {
				log.Printf("Error: %s\n", __obf_38fcaab76f2926d6)
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
