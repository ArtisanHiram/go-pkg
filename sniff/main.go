package __obf_17ead779a85eda96

import (
	"flag"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	var __obf_2872469b8477f1e5 = flag.String("i", "eth0", "Interface to get packets from")
	var __obf_d092f39376feafe8 = flag.String("f", "", "BPF filter")
	var __obf_4f3cfeb93bb47547 = flag.Int("s", 16<<10, "SnapLen for pcap packet capture")
	var __obf_461c961b00616a98 = flag.Bool("v", false, "Logs every packet in great detail")

	flag.Parse()
	__obf_1da25b2adb13968d, __obf_ace39ef064ca7954 := pcap.OpenLive(*__obf_2872469b8477f1e5, int32(*__obf_4f3cfeb93bb47547), true, pcap.BlockForever)
	if __obf_ace39ef064ca7954 != nil {
		log.Printf("Error: %s\n", __obf_ace39ef064ca7954)
		return
	}
	defer __obf_1da25b2adb13968d.Close()

	if *__obf_d092f39376feafe8 != "" {
		if __obf_ace39ef064ca7954 := __obf_1da25b2adb13968d.SetBPFFilter(*__obf_d092f39376feafe8); __obf_ace39ef064ca7954 != nil {
			panic(__obf_ace39ef064ca7954)
		}
	}
	__obf_01933f6b60862535 := //Create a new PacketDataSource
		gopacket.NewPacketSource(__obf_1da25b2adb13968d, layers.LayerTypeEthernet)
	__obf_f9b6fe346ec835b8 := //Packets returns a channel of packets
		__obf_01933f6b60862535.Packets()

	for {
		__obf_2ced1c32bee17307 := // var packet gopacket.Packet
			<-__obf_f9b6fe346ec835b8
		if *__obf_461c961b00616a98 {
			log.Println(__obf_2ced1c32bee17307)
		}
		DstIpProxy(__obf_1da25b2adb13968d, __obf_2ced1c32bee17307)
	}
}

func DstIpProxy(__obf_1da25b2adb13968d *pcap.Handle, __obf_2ced1c32bee17307 gopacket.Packet) {
	__obf_8270abe0d1182f4d := // dstIp, err := net.LookupHost(domain)
		// if err != nil {
		// 	return nil, err
		// }

		// get all layers
		__obf_2ced1c32bee17307.Layers()
	// replaceLayer := make([]gopacket.SerializableLayer, 0, len(allLayers))
	var __obf_ace39ef064ca7954 error
	for _, __obf_9d45b164a732a74a := range __obf_8270abe0d1182f4d {
		switch __obf_9d45b164a732a74a.LayerType() {
		case layers.LayerTypeIPv4:
			__obf_690fdeb6318f72db := __obf_9d45b164a732a74a.(*layers.IPv4)
			__obf_690fdeb6318f72db.
				DstIP = net.IPv4(127, 0, 0, 1)
			__obf_8a0a3260d1bbcfe4 := gopacket.NewSerializeBuffer()
			__obf_ace39ef064ca7954 = __obf_690fdeb6318f72db.SerializeTo(__obf_8a0a3260d1bbcfe4, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_ace39ef064ca7954 != nil {
				log.Printf("Error: %s\n", __obf_ace39ef064ca7954)
			}
			__obf_ace39ef064ca7954 = // send the packet
				__obf_1da25b2adb13968d.WritePacketData(__obf_8a0a3260d1bbcfe4.Bytes())
			if __obf_ace39ef064ca7954 != nil {
				log.Printf("Error: %s\n", __obf_ace39ef064ca7954)
			}

		case layers.LayerTypeIPv6:
			__obf_690fdeb6318f72db := __obf_9d45b164a732a74a.(*layers.IPv6)
			__obf_690fdeb6318f72db.
				DstIP = net.IPv6loopback
			__obf_8a0a3260d1bbcfe4 := gopacket.NewSerializeBuffer()
			__obf_ace39ef064ca7954 = __obf_690fdeb6318f72db.SerializeTo(__obf_8a0a3260d1bbcfe4, gopacket.SerializeOptions{ComputeChecksums: true, FixLengths: true})
			if __obf_ace39ef064ca7954 != nil {
				log.Printf("Error: %s\n", __obf_ace39ef064ca7954)
			}
			__obf_ace39ef064ca7954 = // send the packet
				__obf_1da25b2adb13968d.WritePacketData(__obf_8a0a3260d1bbcfe4.Bytes())
			if __obf_ace39ef064ca7954 != nil {
				log.Printf("Error: %s\n", __obf_ace39ef064ca7954)
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
