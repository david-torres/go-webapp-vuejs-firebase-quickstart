package models

// FirebaseUser models an authenticated user response from firebase
type FirebaseUser struct {
	UID           string `json:"uid"`
	DisplayName   string `json:"displayName"`
	PhotoURL      string `json:"photoURL"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"emailVerified"`
	PhoneNumber   string `json:"phoneNumber"`
	IsAnonymous   bool   `json:"isAnonymous"`
	ProviderData  []struct {
		UID         string `json:"uid"`
		DisplayName string `json:"displayName"`
		PhotoURL    string `json:"photoURL"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
		ProviderID  string `json:"providerId"`
	} `json:"providerData"`
	APIKey          string `json:"apiKey"`
	AppName         string `json:"appName"`
	AuthDomain      string `json:"authDomain"`
	StsTokenManager struct {
		APIKey         string `json:"apiKey"`
		RefreshToken   string `json:"refreshToken"`
		AccessToken    string `json:"accessToken"`
		ExpirationTime int64  `json:"expirationTime"`
	} `json:"stsTokenManager"`
	RedirectEventID string `json:"redirectEventId"`
	LastLoginAt     string `json:"lastLoginAt"`
	CreatedAt       string `json:"createdAt"`
}
