package main

import (
	"flag"
)

var (
	__obf_527251753bbe5828 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_e31bc42d46a4cadc = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_ed53a6cf1d81a796 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_9c8c2e6a9197c48c = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_88fc8a0958f00870 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")
	__obf_c29edf7ae9e78e22 =

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	flag.String("root-org", "root", "root organization name")
	__obf_18060d85d772c392 = flag.String("server-org", "server", "server organization name")
	__obf_a0186db0e5770e5d = flag.String("client-org", "client", "client organization name")
	__obf_07179264e258cdb6 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_879e45fa0c7555fa = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_7b61af43155d6a85 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_f035b210ead5b9d7 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_f442c10694402f79 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_09177447b8c2c966 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
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
