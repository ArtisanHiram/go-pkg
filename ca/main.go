package main

import (
	"flag"
)

var (
	__obf_b832ed210a876042 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_fb9853e4bd03851d = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_63740d0c822b9e42 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_340232d3b11c627a = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_47b5e014df570fef = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")
	__obf_9f01f358f5287bb0 =

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	flag.String("root-org", "root", "root organization name")
	__obf_83f5e2108cdf9dd8 = flag.String("server-org", "server", "server organization name")
	__obf_302bcd74543539f2 = flag.String("client-org", "client", "client organization name")
	__obf_f179418df50a4ef1 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_013e90306300e806 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_5dae3cf5185920ed = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_14693b8ee36fd73b = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_f8a1ce35efbbcd8b = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_6a6634cefa52690f = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
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
