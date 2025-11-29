package main

import (
	"flag"
)

var (
	__obf_e7dc96daf488b4d7 = flag.String("dir", "tls", "Directory to store certificates")
	__obf_7c9a6812bf2314f9 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_244f5fea7ddf0c82 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_f5980e672145c29f = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_687f16797dbe5480 = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")
	__obf_5723a8a1754a6ff7 =

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	flag.String("root-org", "root", "root organization name")
	__obf_fe2344381e73ffee = flag.String("server-org", "server", "server organization name")
	__obf_dc0cb0d6e620a6ab = flag.String("client-org", "client", "client organization name")
	__obf_bf317b5f1f2b6382 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_c82831f575e7670b = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_28837c6c18a9b3d9 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_e1e272065c9d4c14 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_af7fc1426a450125 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_b31dfc092fcc3a05 = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
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
