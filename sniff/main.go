package __obf_9ea02814b17c6f40

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func __obf_56adab93ce7b3a78() {
	var __obf_c0409b6da6f1dd59 = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_815955505907a823 = flag.String("f", "", "BPF filter")
	var __obf_7729ac968f5c6c84 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_c1f1aff67ce62547 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()

	__obf_cc11b74613ddb15b, __obf_4cdfc3b251a33a87 := pcap.OpenLive(*__obf_c0409b6da6f1dd59, int32(*__obf_7729ac968f5c6c84), true, pcap.BlockForever)
	if __obf_4cdfc3b251a33a87 != nil {
		log.Printf("Error: %s\n", __obf_4cdfc3b251a33a87)
		return
	}
	defer __obf_cc11b74613ddb15b.Close()

	if *__obf_815955505907a823 != "" {
		if __obf_4cdfc3b251a33a87 := __obf_cc11b74613ddb15b.SetBPFFilter(*__obf_815955505907a823); __obf_4cdfc3b251a33a87 != nil {
			panic(__obf_4cdfc3b251a33a87)
		}
	}

	//Create a new PacketDataSource
	__obf_264d408323a8f99b := gopacket.NewPacketSource(__obf_cc11b74613ddb15b, layers.LayerTypeEthernet)
	//Packets returns a channel of packets
	__obf_a48c7ca1d0c17d6a := __obf_264d408323a8f99b.Packets()

	for {
		// var packet gopacket.Packet
		__obf_5904cc4704e24a64 := <-__obf_a48c7ca1d0c17d6a
		if *__obf_c1f1aff67ce62547 {
			log.Println(__obf_5904cc4704e24a64)
		}
		DstIpProxy(__obf_cc11b74613ddb15b, __obf_5904cc4704e24a64)
	}
}

func DstIpProxy(__obf_cc11b74613ddb15b *pcap.Handle, __obf_5904cc4704e24a64 gopacket.Packet) {
	// dstIp, err := net.LookupHost(domain)
	// if err != nil {
	// 	return nil, err
	// }

	// get all layers
	__obf_9fc6eeb35c1e5bae := __obf_5904cc4704e24a64.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_4cdfc3b251a33a87 error
	for _, __obf_b1bda9c632cc7bdb := range __obf_9fc6eeb35c1e5bae {
		switch __obf_b1bda9c632cc7bdb.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_fa7baebc62e708a3 := __obf_b1bda9c632cc7bdb.(*layers.IPv4)
			__obf_fa7baebc62e708a3.DstIP = net.IPv4(127, 0, 0, 1)

			__obf_4a247b137f5f72ab := gopacket.NewSerializeBuffer()
			__obf_4cdfc3b251a33a87 = __obf_fa7baebc62e708a3.SerializeTo(__obf_4a247b137f5f72ab, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_4cdfc3b251a33a87 != nil {
				log.Printf("Error: %s\n", __obf_4cdfc3b251a33a87)
			}
			// send the packet
			__obf_4cdfc3b251a33a87 = __obf_cc11b74613ddb15b.WritePacketData(__obf_4a247b137f5f72ab.Bytes())
			if __obf_4cdfc3b251a33a87 != nil {
				log.Printf("Error: %s\n", __obf_4cdfc3b251a33a87)
			}

		case layers.LayerTypeIPv6:
			__obf_fa7baebc62e708a3 := __obf_b1bda9c632cc7bdb.(*layers.IPv6)
			__obf_fa7baebc62e708a3.DstIP = net.IPv6loopback

			__obf_4a247b137f5f72ab := gopacket.NewSerializeBuffer()
			__obf_4cdfc3b251a33a87 = __obf_fa7baebc62e708a3.SerializeTo(__obf_4a247b137f5f72ab, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_4cdfc3b251a33a87 != nil {
				log.Printf("Error: %s\n", __obf_4cdfc3b251a33a87)
			}
			// send the packet
			__obf_4cdfc3b251a33a87 = __obf_cc11b74613ddb15b.WritePacketData(__obf_4a247b137f5f72ab.Bytes())
			if __obf_4cdfc3b251a33a87 != nil {
				log.Printf("Error: %s\n", __obf_4cdfc3b251a33a87)
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
