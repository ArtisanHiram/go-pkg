package __obf_a7fac689e11862d9

import (
	"flag"
)

var (
	__obf_d04fc7665a661d58 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_757dc2f64e06a3fa = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_b416844f5e557892 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_f279281a432f941c = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_8f522c929d5b2e06 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	__obf_b0f27e785a5d7613 = flag.String("root-org", "root", "root organization name")
	__obf_8c0bc0ba949334fc = flag.String("server-org", "server", "server organization name")
	__obf_0e47c303a8f77c9f = flag.String("client-org", "client", "client organization name")
	__obf_115c15a762a4cfbe = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_64b88161bfe6ddb6 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_e85a2925f0461119 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_f0a9e6b00a4f1eea = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_7497a241ee847b27 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_6558bea5d760a593 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func __obf_a7fac689e11862d9() {

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
