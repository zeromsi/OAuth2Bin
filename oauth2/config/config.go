package config

// Enum for the OAuth 2.0 flows
const (
	AuthCode    = 1
	Implicit    = 2
	ROPC        = 3
	ClientCreds = 4
)

// AuthCodeConfig defines the variables required in the OAuth 2.0 Authorization Code flow
type AuthCodeConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// ImplicitConfig defines the variables required in the OAuth 2.0 Implicit flow
type ImplicitConfig struct {
	ClientID string `json:"clientID"`
}

// ROPCConfig defines the variables required in the OAuth 2.0 Resource Owner Password Credentials flow
type ROPCConfig struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// ClientCredsConfig defines the variables required in the OAuth 2.0 Client Credentials flow
type ClientCredsConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// OA2Config defines the configurations for all the flows in OAuth 2.0
type OA2Config struct {
	BaseURL         string            `json:"baseURL"`
	AuthCodeCnfg    AuthCodeConfig    `json:"authCode"`
	ImplicitCnfg    ImplicitConfig    `json:"implicit"`
	ROPCCnfg        ROPCConfig        `json:"ropc"`
	ClientCredsCnfg ClientCredsConfig `json:"clientCreds"`
}
