package __obf_8f6d6401b1ecc0c1

// Config .config for 小程序
type Config struct {
	AppID          string `json:"app_id" yaml:"app_id"`                     // appid
	AppSecret      string `json:"app_secret" yaml:"app_secret"`             // appSecret
	AppKey         string `json:"app_key" yaml:"app_key"`                   // appKey
	OfferID        string `json:"offer_id" yaml:"offer_id"`                 // offerId
	Token          string `json:"token" yaml:"token"`                       // token
	EncodingAESKey string `json:"encoding_aes_key" yaml:"encoding_aes_key"` // EncodingAESKey
	UseStableAK    bool   // use the stable access_token
}
