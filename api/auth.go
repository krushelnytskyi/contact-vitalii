package api

type Auth struct {
	IP   string `json:"ip"`
	Word string `json:"word"`
	Used bool   `json:"used"`
}

func Authorize (IP string) *Auth {
	return &Auth{IP: IP, Word: "Guest", Used: false}
}

func (auth *Auth) use () {
	auth.Used = true;
}