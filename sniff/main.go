package __obf_47f69d38d23d3d12

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func __obf_5634b2a25f68a229() {
	var __obf_5bc3299c3bfb8d9f = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_51c5459d20026c44 = flag.String("f", "", "BPF filter")
	var __obf_65ece0a72ad16b74 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_4f388560bdc9191f = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()

	__obf_e1ac535169807a33, __obf_cd0729cf7c05d2d9 := pcap.OpenLive(*__obf_5bc3299c3bfb8d9f, int32(*__obf_65ece0a72ad16b74), true, pcap.BlockForever)
	if __obf_cd0729cf7c05d2d9 != nil {
		log.Printf("Error: %s\n", __obf_cd0729cf7c05d2d9)
		return
	}
	defer __obf_e1ac535169807a33.Close()

	if *__obf_51c5459d20026c44 != "" {
		if __obf_cd0729cf7c05d2d9 := __obf_e1ac535169807a33.SetBPFFilter(*__obf_51c5459d20026c44); __obf_cd0729cf7c05d2d9 != nil {
			panic(__obf_cd0729cf7c05d2d9)
		}
	}

	//Create a new PacketDataSource
	__obf_0fc1fbac48fb6f6f := gopacket.NewPacketSource(__obf_e1ac535169807a33, layers.LayerTypeEthernet)
	//Packets returns a channel of packets
	__obf_d12289ef8f470382 := __obf_0fc1fbac48fb6f6f.Packets()

	for {
		// var packet gopacket.Packet
		__obf_557dcd4297eff908 := <-__obf_d12289ef8f470382
		if *__obf_4f388560bdc9191f {
			log.Println(__obf_557dcd4297eff908)
		}
		DstIpProxy(__obf_e1ac535169807a33, __obf_557dcd4297eff908)
	}
}

func DstIpProxy(__obf_e1ac535169807a33 *pcap.Handle, __obf_557dcd4297eff908 gopacket.Packet) {
	// dstIp, err := net.LookupHost(domain)
	// if err != nil {
	// 	return nil, err
	// }

	// get all layers
	__obf_46cc181d88008a5a := __obf_557dcd4297eff908.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_cd0729cf7c05d2d9 error
	for _, __obf_d555be62766d3836 := range __obf_46cc181d88008a5a {
		switch __obf_d555be62766d3836.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_40966f31ce98054e := __obf_d555be62766d3836.(*layers.IPv4)
			__obf_40966f31ce98054e.DstIP = net.IPv4(127, 0, 0, 1)

			__obf_650682206bd3604a := gopacket.NewSerializeBuffer()
			__obf_cd0729cf7c05d2d9 = __obf_40966f31ce98054e.SerializeTo(__obf_650682206bd3604a, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_cd0729cf7c05d2d9 != nil {
				log.Printf("Error: %s\n", __obf_cd0729cf7c05d2d9)
			}
			// send the packet
			__obf_cd0729cf7c05d2d9 = __obf_e1ac535169807a33.WritePacketData(__obf_650682206bd3604a.Bytes())
			if __obf_cd0729cf7c05d2d9 != nil {
				log.Printf("Error: %s\n", __obf_cd0729cf7c05d2d9)
			}

		case layers.LayerTypeIPv6:
			__obf_40966f31ce98054e := __obf_d555be62766d3836.(*layers.IPv6)
			__obf_40966f31ce98054e.DstIP = net.IPv6loopback

			__obf_650682206bd3604a := gopacket.NewSerializeBuffer()
			__obf_cd0729cf7c05d2d9 = __obf_40966f31ce98054e.SerializeTo(__obf_650682206bd3604a, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_cd0729cf7c05d2d9 != nil {
				log.Printf("Error: %s\n", __obf_cd0729cf7c05d2d9)
			}
			// send the packet
			__obf_cd0729cf7c05d2d9 = __obf_e1ac535169807a33.WritePacketData(__obf_650682206bd3604a.Bytes())
			if __obf_cd0729cf7c05d2d9 != nil {
				log.Printf("Error: %s\n", __obf_cd0729cf7c05d2d9)
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
