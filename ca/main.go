package main

import (
	"flag"
)

var (
	__obf_4b00946562018abc = flag.String("dir", "tls", "Directory to store certificates")
	__obf_7927c321cb778076 = flag.String("host", "localhost,127.0.0.1,122.51.240.88", "Comma-separated hostnames and IPs to generate a certificate for")
	__obf_9888c4a4b3aa3e56 = flag.Int("rsa-bits", 2048, "Size of RSA key to generate. Ignored if `ecdsa-curve` is set")
	__obf_df6ed5643a391a04 = flag.String("ecdsa-curve", "P256", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521")
	__obf_7202109d1688abad = flag.Bool("use-ed25519", false, "Generate an Ed25519 key or not")
	__obf_b1f0196313c2663e =

	// mode     = flag.String("mode", "flag", "cmd flag or yaml")
	// confFile = flag.String("conf-file", "ca.yml", "path of ca information file")

	flag.String("root-org", "root", "root organization name")
	__obf_46d118ebb28c1025 = flag.String("server-org", "server", "server organization name")
	__obf_64c69c9a9ddb71c1 = flag.String("client-org", "client", "client organization name")
	__obf_6abe834fb3e25469 = flag.String("root-start-date", "2022-01-01", "Creation date of root certificate formatted as 2020-07-22")
	__obf_8e04f1efba12895c = flag.Int("root-expire", 3650, "Days of duration that root certificate is valid")
	__obf_eafe7ab8a89428b8 = flag.String("server-start-date", "2022-01-01", "Creation date server certificate formatted as 2020-07-22")
	__obf_e29bec54c690b2f9 = flag.Int("server-expire", 3650, "Day of duration that server certificate is valid")
	__obf_6315e9ca5d8be1e7 = flag.String("client-start-date", "2022-01-01", "Creation date client certificate formatted as 2020-07-22")
	__obf_892547ae07fe593a = flag.Int("client-expire", 3650, "Day of duration that client certificate is valid")
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
