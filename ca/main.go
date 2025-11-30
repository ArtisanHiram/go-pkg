package main

import (
	"flag"
)

var (
	__obf_a210ca1feaeefbbf = flag.String("dir", "tls", "Directory to store certificates")
	__obf_c60afb3a5311a7bd = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_676c29a914cffb78 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_65d9b57be59fef7f = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_bda93e55c59138d6 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")
	__obf_c9dfe7ad685e91c6 =

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	flag.String("root-org", "root", "root organization name")
	__obf_d66092b40e18dd97 = flag.String("server-org", "server", "server organization name")
	__obf_bf1a0d93435070a5 = flag.String("client-org", "client", "client organization name")
	__obf_2010a1d7a6b43947 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_77175f9ff579aa0c = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_cfa66aa3eaf60752 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_04ef712d5ae945e6 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_3641c7ed101c51be = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_60288d6a3a5662c8 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func main() {

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
