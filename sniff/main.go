package __obf_f8061b98fd811396

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func __obf_69429786061db9bc() {
	var __obf_72d0c560fa78901e = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_fa2874f860a4103a = flag.String("f", "", "BPF filter")
	var __obf_b29b2c486f91f6db = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_354a9768ccb383c3 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()

	__obf_b8981c7582ef6b51, __obf_64ed9cb94c6de02d := pcap.OpenLive(*__obf_72d0c560fa78901e, int32(*__obf_b29b2c486f91f6db), true, pcap.BlockForever)
	if __obf_64ed9cb94c6de02d != nil {
		log.Printf("Error: %s\n", __obf_64ed9cb94c6de02d)
		return
	}
	defer __obf_b8981c7582ef6b51.Close()

	if *__obf_fa2874f860a4103a != "" {
		if __obf_64ed9cb94c6de02d := __obf_b8981c7582ef6b51.SetBPFFilter(*__obf_fa2874f860a4103a); __obf_64ed9cb94c6de02d != nil {
			panic(__obf_64ed9cb94c6de02d)
		}
	}

	//Create a new PacketDataSource
	__obf_92025a797148dc14 := gopacket.NewPacketSource(__obf_b8981c7582ef6b51, layers.LayerTypeEthernet)
	//Packets returns a channel of packets
	__obf_f04a2fdbc1485165 := __obf_92025a797148dc14.Packets()

	for {
		// var packet gopacket.Packet
		__obf_a01de5befe4bcbe7 := <-__obf_f04a2fdbc1485165
		if *__obf_354a9768ccb383c3 {
			log.Println(__obf_a01de5befe4bcbe7)
		}
		DstIpProxy(__obf_b8981c7582ef6b51, __obf_a01de5befe4bcbe7)
	}
}

func DstIpProxy(__obf_b8981c7582ef6b51 *pcap.Handle, __obf_a01de5befe4bcbe7 gopacket.Packet) {
	// dstIp, err := net.LookupHost(domain)
	// if err != nil {
	// 	return nil, err
	// }

	// get all layers
	__obf_2c1030d46fef7573 := __obf_a01de5befe4bcbe7.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_64ed9cb94c6de02d error
	for _, __obf_b0df9312dc52ca6d := range __obf_2c1030d46fef7573 {
		switch __obf_b0df9312dc52ca6d.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_e9afeb29d0da450d := __obf_b0df9312dc52ca6d.(*layers.IPv4)
			__obf_e9afeb29d0da450d.DstIP = net.IPv4(127, 0, 0, 1)

			__obf_8f2db06dd8b4d1fc := gopacket.NewSerializeBuffer()
			__obf_64ed9cb94c6de02d = __obf_e9afeb29d0da450d.SerializeTo(__obf_8f2db06dd8b4d1fc, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_64ed9cb94c6de02d != nil {
				log.Printf("Error: %s\n", __obf_64ed9cb94c6de02d)
			}
			// send the packet
			__obf_64ed9cb94c6de02d = __obf_b8981c7582ef6b51.WritePacketData(__obf_8f2db06dd8b4d1fc.Bytes())
			if __obf_64ed9cb94c6de02d != nil {
				log.Printf("Error: %s\n", __obf_64ed9cb94c6de02d)
			}

		case layers.LayerTypeIPv6:
			__obf_e9afeb29d0da450d := __obf_b0df9312dc52ca6d.(*layers.IPv6)
			__obf_e9afeb29d0da450d.DstIP = net.IPv6loopback

			__obf_8f2db06dd8b4d1fc := gopacket.NewSerializeBuffer()
			__obf_64ed9cb94c6de02d = __obf_e9afeb29d0da450d.SerializeTo(__obf_8f2db06dd8b4d1fc, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_64ed9cb94c6de02d != nil {
				log.Printf("Error: %s\n", __obf_64ed9cb94c6de02d)
			}
			// send the packet
			__obf_64ed9cb94c6de02d = __obf_b8981c7582ef6b51.WritePacketData(__obf_8f2db06dd8b4d1fc.Bytes())
			if __obf_64ed9cb94c6de02d != nil {
				log.Printf("Error: %s\n", __obf_64ed9cb94c6de02d)
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
