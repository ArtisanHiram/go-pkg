package __obf_b6eef43f2d5b2785

import (
	"flag"
)

var (
	__obf_7662752a24f6c29a = flag.String("dir", "tls", "Directory to store certificates")
	__obf_18118f4267561011 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_6cac191609a547f0 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_bba0c595ec2b1e3d = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_8188aca91d5a646c = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	__obf_370baf1383967997 = flag.String("root-org", "root", "root organization name")
	__obf_5465625a1ceeaf75 = flag.String("server-org", "server", "server organization name")
	__obf_4d65ba9ada5d6594 = flag.String("client-org", "client", "client organization name")
	__obf_f3d20ae37848f3fc = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_ece9c7189be6d320 = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_d8c4f6c2d7602f25 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_b42de427a3167a74 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_f625282011c32525 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_81b901915e137680 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
)

func __obf_b6eef43f2d5b2785() {

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
