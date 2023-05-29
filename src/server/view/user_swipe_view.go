package view

type UserSwipeResponse struct {
	UserTarget string `json:"user_target"`
	Match      bool   `json:"match"`
}
