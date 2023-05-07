package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const CookieName = "NINGENME_NET_SESSION"

type KiwaApiRepository struct{}

type UsersMeGetResponseDto struct {
	UserId    string                         `json:"userId"`
	Authority UsersMeGetResponseAuthorityDto `json:"authority"`
}
type UsersMeGetResponseAuthorityDto struct {
	ComproCategory bool `json:"comproCategory"`
}

func (KiwaApiRepository) GetUsersMe(cookie *http.Cookie) (*UsersMeGetResponseDto, error) {
	req, _ := http.NewRequest("GET", "https://kiwa-api.ningenme.net/users/me", nil)
	req.AddCookie(cookie)

	client := &http.Client{Transport: http.DefaultTransport}
	res, err := client.Do(req)
	if err != nil {
		return &UsersMeGetResponseDto{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return &UsersMeGetResponseDto{}, fmt.Errorf("kiwa-api response status was not 200")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return &UsersMeGetResponseDto{}, err
	}
	var usersMeGetResponse *UsersMeGetResponseDto
	if err := json.Unmarshal(body, &usersMeGetResponse); err != nil {
		return &UsersMeGetResponseDto{}, err
	}
	
	return usersMeGetResponse, nil
}
