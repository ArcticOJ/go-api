package arctic

import (
	"errors"
	"github.com/ArcticOJ/go-api-bindings/v0/types"
	"github.com/ArcticOJ/go-api-bindings/v0/types/common"
	"github.com/ArcticOJ/go-api-bindings/v0/types/common/auth"
	"net/http"
	"slices"
	"strings"
)

func (c *Client) IsAuthenticated() bool {
	return slices.ContainsFunc(c.c.Cookies, func(cookie *http.Cookie) bool {
		return cookie.Name == "session" && strings.TrimSpace(cookie.Value) != ""
	})
}

func (c *Client) Auth(cookie *http.Cookie) {
	c.c.SetCommonCookies(cookie)
}

func (c *Client) Login(handle string, password string, remember bool) (*http.Cookie, error) {
	var err common.Error
	res, e := c.c.R().
		SetErrorResult(&err).
		SetBody(auth.LoginForm{
			Handle:     handle,
			Password:   password,
			RememberMe: remember,
		}).
		Post("/auth/login")
	if e != nil {
		return nil, e
	}
	if err.Code >= http.StatusBadRequest {
		return nil, errors.New(err.Message)
	}
	cookie := slices.IndexFunc(res.Cookies(), func(cookie *http.Cookie) bool {
		return cookie.Name == "session" && strings.TrimSpace(cookie.Value) != ""
	})
	if cookie == -1 {
		return nil, types.ErrInvalidCookie
	}
	return res.Cookies()[cookie], nil
}
