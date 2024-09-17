package models

type LoginData struct {
	ID           uint64 `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Nick         string `json:"nick,omitempty"`
	Email        string `json:"email,omitempty"`
	UserImage    string `json:"UserImage,omitempty"`
	RefreshToken string `json:"token,omitempty"`
}
