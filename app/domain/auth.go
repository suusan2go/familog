package domain

// CloudEndpointから送られてくる認証情報のストラクト
type Auth struct {
	Issuer string `json:"issuer"`
	ID     string `json:"id"`
	Email  string `json:"email"`
	Alg    string `json:"alg"`
}
