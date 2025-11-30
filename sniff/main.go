package __obf_c0932c9a02e09beb

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var __obf_8500954641377fc3 = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_8cbe1dca1f0589a9 = flag.String("f", "", "BPF filter")
	var __obf_eaf3a48f3df4cff7 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_0f09f1bcb752521d = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()
	__obf_d52c8d344def5163, __obf_e1dac77a901a86c3 := pcap.OpenLive(*__obf_8500954641377fc3, int32(*__obf_eaf3a48f3df4cff7), true, pcap.BlockForever)
	if __obf_e1dac77a901a86c3 != nil {
		log.Printf("Error: %s\n", __obf_e1dac77a901a86c3)
		return
	}
	defer __obf_d52c8d344def5163.Close()

	if *__obf_8cbe1dca1f0589a9 != "" {
		if __obf_e1dac77a901a86c3 := __obf_d52c8d344def5163.SetBPFFilter(*__obf_8cbe1dca1f0589a9); __obf_e1dac77a901a86c3 != nil {
			panic(__obf_e1dac77a901a86c3)
		}
	}
	__obf_af5af8a9de29aa63 := //Create a new PacketDataSource
		gopacket.NewPacketSource(__obf_d52c8d344def5163, layers.LayerTypeEthernet)
	__obf_7ce1a93b426ec51b := //Packets returns a channel of packets
		__obf_af5af8a9de29aa63.Packets()

	for {
		__obf_0d236e962a9f1f4b := // var packet gopacket.Packet
			<-__obf_7ce1a93b426ec51b
		if *__obf_0f09f1bcb752521d {
			log.Println(__obf_0d236e962a9f1f4b)
		}
		DstIpProxy(__obf_d52c8d344def5163, __obf_0d236e962a9f1f4b)
	}
}

func DstIpProxy(__obf_d52c8d344def5163 *pcap.Handle, __obf_0d236e962a9f1f4b gopacket.Packet) {
	__obf_feedb44685fd04bc := // dstIp, err := net.LookupHost(domain)
		// if err != nil {
		// 	return nil, err
		// }

		// get all layers
		__obf_0d236e962a9f1f4b.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_e1dac77a901a86c3 error
	for _, __obf_11742251f9751dc3 := range __obf_feedb44685fd04bc {
		switch __obf_11742251f9751dc3.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_4ef7628fd581a1dd := __obf_11742251f9751dc3.(*layers.IPv4)
			__obf_4ef7628fd581a1dd.
				DstIP = net.IPv4(127, 0, 0, 1)
			__obf_4c24e38810af71e1 := gopacket.NewSerializeBuffer()
			__obf_e1dac77a901a86c3 = __obf_4ef7628fd581a1dd.SerializeTo(__obf_4c24e38810af71e1, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_e1dac77a901a86c3 != nil {
				log.Printf("Error: %s\n", __obf_e1dac77a901a86c3)
			}
			__obf_e1dac77a901a86c3 = // send the packet
				__obf_d52c8d344def5163.WritePacketData(__obf_4c24e38810af71e1.Bytes())
			if __obf_e1dac77a901a86c3 != nil {
				log.Printf("Error: %s\n", __obf_e1dac77a901a86c3)
			}

		case layers.LayerTypeIPv6:
			__obf_4ef7628fd581a1dd := __obf_11742251f9751dc3.(*layers.IPv6)
			__obf_4ef7628fd581a1dd.
				DstIP = net.IPv6loopback
			__obf_4c24e38810af71e1 := gopacket.NewSerializeBuffer()
			__obf_e1dac77a901a86c3 = __obf_4ef7628fd581a1dd.SerializeTo(__obf_4c24e38810af71e1, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_e1dac77a901a86c3 != nil {
				log.Printf("Error: %s\n", __obf_e1dac77a901a86c3)
			}
			__obf_e1dac77a901a86c3 = // send the packet
				__obf_d52c8d344def5163.WritePacketData(__obf_4c24e38810af71e1.Bytes())
			if __obf_e1dac77a901a86c3 != nil {
				log.Printf("Error: %s\n", __obf_e1dac77a901a86c3)
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
