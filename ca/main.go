package __obf_cad85c1a84aaff79

import (
	"flag"
)

var (
	__obf_d7fd1eaf7db608e6 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_967bcefeb68b3bc2 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_3a3f21e4c08578dc = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_31a37a320eb83526 = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_d8991a976dff20bf = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	__obf_dbbc139bf2b03050 = flag.String("root-org", "root", "root organization name")
	__obf_f782959824cfef0c = flag.String("server-org", "server", "server organization name")
	__obf_5916aadfe17ba1cd = flag.String("client-org", "client", "client organization name")
	__obf_f62a2d95847ff0f5 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_eb61da3b8dc27fa4 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_c1cc1283cf2bcf00 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_6dc1feab678126da = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_29da549346963415 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_3800c25206840554 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func __obf_cad85c1a84aaff79() {

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
