package __obf_3d5f73927e7527a3

import (
	"flag"
)

var (
	__obf_99d81957a2a2485e = flag.String("dir", "tls", "Directory to store certificates")
	__obf_1a8588b7039b1147 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_d20d08b4ac781958 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_ca39108216f748ac = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_3ac46f69661ae0d8 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	__obf_b5bc963dd5e57e51 = flag.String("root-org", "root", "root organization name")
	__obf_d03c4ee6c7e51cb5 = flag.String("server-org", "server", "server organization name")
	__obf_5aaf18bf1be127ce = flag.String("client-org", "client", "client organization name")
	__obf_5670263c37016623 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_3dfc5003a42402a4 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_a10e2c032b2a681d = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_697861048accd993 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_a7d625f3ccb24dd6 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_7350e0f0927a0cf7 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func __obf_3d5f73927e7527a3() {

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
