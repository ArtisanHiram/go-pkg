package main

import (
	"flag"
)

var (
	__obf_42e3c6717d306275 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_1d133200ede8a085 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_0cb5ccb4d6fefda5 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_3d01f66db0c6933a = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_762a091af78eacdc = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")
	__obf_76e50b8d8497183b =

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	flag.String("root-org", "root", "root organization name")
	__obf_ca34a921c84f35c2 = flag.String("server-org", "server", "server organization name")
	__obf_e89e84dedb3f1ddb = flag.String("client-org", "client", "client organization name")
	__obf_77b98a74ef4e3ee6 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_7342674c21b8d9fe = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_0ce26a566e9dee32 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_f2a23f9f3af65936 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_7ce6c606cf2b1391 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_626afac41336a16f = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
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
