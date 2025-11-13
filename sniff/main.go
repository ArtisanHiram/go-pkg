package __obf_cbb2f31b14885fa4

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func __obf_8edc9cd300112d67() {
	var __obf_0569ae7da31d7005 = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_69bc4d8570f784dd = flag.String("f", "", "BPF filter")
	var __obf_774a500a1e06eaf7 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_6f75468396ce7153 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()

	__obf_2830318a0a2bcd63, __obf_d6c922bcc0347b3e := pcap.OpenLive(*__obf_0569ae7da31d7005, int32(*__obf_774a500a1e06eaf7), true, pcap.BlockForever)
	if __obf_d6c922bcc0347b3e != nil {
		log.Printf("Error: %s\n", __obf_d6c922bcc0347b3e)
		return
	}
	defer __obf_2830318a0a2bcd63.Close()

	if *__obf_69bc4d8570f784dd != "" {
		if __obf_d6c922bcc0347b3e := __obf_2830318a0a2bcd63.SetBPFFilter(*__obf_69bc4d8570f784dd); __obf_d6c922bcc0347b3e != nil {
			panic(__obf_d6c922bcc0347b3e)
		}
	}

	//Create a new PacketDataSource
	__obf_0e12eb29d7baf7b3 := gopacket.NewPacketSource(__obf_2830318a0a2bcd63, layers.LayerTypeEthernet)
	//Packets returns a channel of packets
	__obf_fd5757d2e752ef4f := __obf_0e12eb29d7baf7b3.Packets()

	for {
		// var packet gopacket.Packet
		__obf_a654ffa9029aed96 := <-__obf_fd5757d2e752ef4f
		if *__obf_6f75468396ce7153 {
			log.Println(__obf_a654ffa9029aed96)
		}
		DstIpProxy(__obf_2830318a0a2bcd63, __obf_a654ffa9029aed96)
	}
}

func DstIpProxy(__obf_2830318a0a2bcd63 *pcap.Handle, __obf_a654ffa9029aed96 gopacket.Packet) {
	// dstIp, err := net.LookupHost(domain)
	// if err != nil {
	// 	return nil, err
	// }

	// get all layers
	__obf_540b093d8864b201 := __obf_a654ffa9029aed96.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_d6c922bcc0347b3e error
	for _, __obf_238412ca6ffc41da := range __obf_540b093d8864b201 {
		switch __obf_238412ca6ffc41da.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_b31120bbebbb2dc1 := __obf_238412ca6ffc41da.(*layers.IPv4)
			__obf_b31120bbebbb2dc1.DstIP = net.IPv4(127, 0, 0, 1)

			__obf_629e2b0bbaae4993 := gopacket.NewSerializeBuffer()
			__obf_d6c922bcc0347b3e = __obf_b31120bbebbb2dc1.SerializeTo(__obf_629e2b0bbaae4993, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_d6c922bcc0347b3e != nil {
				log.Printf("Error: %s\n", __obf_d6c922bcc0347b3e)
			}
			// send the packet
			__obf_d6c922bcc0347b3e = __obf_2830318a0a2bcd63.WritePacketData(__obf_629e2b0bbaae4993.Bytes())
			if __obf_d6c922bcc0347b3e != nil {
				log.Printf("Error: %s\n", __obf_d6c922bcc0347b3e)
			}

		case layers.LayerTypeIPv6:
			__obf_b31120bbebbb2dc1 := __obf_238412ca6ffc41da.(*layers.IPv6)
			__obf_b31120bbebbb2dc1.DstIP = net.IPv6loopback

			__obf_629e2b0bbaae4993 := gopacket.NewSerializeBuffer()
			__obf_d6c922bcc0347b3e = __obf_b31120bbebbb2dc1.SerializeTo(__obf_629e2b0bbaae4993, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_d6c922bcc0347b3e != nil {
				log.Printf("Error: %s\n", __obf_d6c922bcc0347b3e)
			}
			// send the packet
			__obf_d6c922bcc0347b3e = __obf_2830318a0a2bcd63.WritePacketData(__obf_629e2b0bbaae4993.Bytes())
			if __obf_d6c922bcc0347b3e != nil {
				log.Printf("Error: %s\n", __obf_d6c922bcc0347b3e)
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
