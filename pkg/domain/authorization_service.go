package domain

import (
	"fmt"
	"github.com/ningenMe/miiko-api/pkg/infra"
	"net/http"
)

type AuthorizationService struct{}

var kiwaApiRepository = infra.KiwaApiRepository{}

func (AuthorizationService) CheckLoggedIn(header http.Header) error {
	req := http.Request{Header: header}
	cookie, err := req.Cookie(infra.CookieName)
	if err != nil {
		return err
	}

	dto, err := kiwaApiRepository.GetUsersMe(cookie)
	if err != nil {
		return err
	}

	if !dto.Authority.ComproCategory {
		return fmt.Errorf("not authorized")
	}

	//認可があるときはnilを返す
	return nil
}
