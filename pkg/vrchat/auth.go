package vrchat

import (
	"errors"
)

func (v *VRChatClient) Logout() error {
	_, err := v.client.R().Put("logout")
	return err
}

func (v *VRChatClient) Verify2FA(code string) error {
	var info struct {
		Verified bool `json:"verified"`
	}
	_, err := v.client.R().
		SetBody(map[string]string{"code": code}).
		SetResult(&info).
		Post("auth/twofactorauth/totp/verify")
	if err != nil {
		return err
	} else if !info.Verified {
		return errors.New("2FA verification failed")
	}
	return nil
}

func (v *VRChatClient) AuthUser() (err error) {
	_, err = v.client.R().
		SetResult(&v.Auth).
		Get("auth/user")
	if err != nil {
		return
	}
	return
}