package __obf_be728ac087d82248

import (
	"flag"
)

var (
	__obf_35c2f92426e7679c = flag.String("dir", "tls", "Directory to store certificates")
	__obf_2017486f76444d49 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_2a90108bf71837ff = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_f272f57faa9e35d7 = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_da026a6582b22c97 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	__obf_534de083cb38fe22 = flag.String("root-org", "root", "root organization name")
	__obf_14b63fdf42420081 = flag.String("server-org", "server", "server organization name")
	__obf_4ac727a060e9a62b = flag.String("client-org", "client", "client organization name")
	__obf_58c8553c2ef7748c = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_c64022213fc6de55 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_82052cf1888686b0 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_fce9244442953147 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_e5be54f6dccae7ba = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_ee52aae87c9ea350 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func __obf_be728ac087d82248() {

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
