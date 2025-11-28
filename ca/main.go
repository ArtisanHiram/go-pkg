package __obf_b8dd0558a7deb9fa

import (
	"flag"
)

var (
	__obf_17f9ffe223fe4d58 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_5a5eb3590804c3ce = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_911ff0576318f3c7 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_694c6c8758e0d040 = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_eb00d4b9a27af664 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	__obf_e0ad65cc4a7b9d78 = flag.String("root-org", "root", "root organization name")
	__obf_a0b8eda05f033c7e = flag.String("server-org", "server", "server organization name")
	__obf_d3a01009d8f993f2 = flag.String("client-org", "client", "client organization name")
	__obf_4beb52f8769f794e = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_8793139ab3849466 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_3498d038b99e93bc = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_9a81352c2d8f646d = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_c2b303d5acd15327 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_e2a041976aba5300 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func __obf_b8dd0558a7deb9fa() {

	// c, err := ioutil.ReadFile(*confFile)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// var ci CaInfo
	// err = yaml.UnmarshalStrict(c, &ci)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	flag.Parse()

	GenRootCA()

	GenServerCA()

	GenClientCA()

}
