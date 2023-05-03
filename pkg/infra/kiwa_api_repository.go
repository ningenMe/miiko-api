package infra

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const CookieName = "NINGENME_NET_SESSION"

type KiwaApiRepository struct{}

type UsersMeGetResponse struct {
	UserId    string                      `json:"userId"`
	Authority UsersMeGetResponseAuthority `json:"authority"`
}
type UsersMeGetResponseAuthority struct {
	ComproCategory bool `json:"comproCategory"`
}

func (KiwaApiRepository) IsLoggedIn(cookie *http.Cookie) error {
	req, _ := http.NewRequest("GET", "https://kiwa-api.ningenme.net/users/me", nil)
	req.AddCookie(cookie)

	client := &http.Client{Transport: http.DefaultTransport}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("kiwa-api response status was not 200")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var usersMeGetResponse UsersMeGetResponse
	if err := json.Unmarshal(body, &usersMeGetResponse); err != nil {
		return err
	}

	if !usersMeGetResponse.Authority.ComproCategory {
		return fmt.Errorf("not authorized")
	}

	return nil
}
