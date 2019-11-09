package memory

import (
	"errors"
	"net/http"
	"time"

	"github.com/jonfriesen/simple-sso/pkg/util"

	"github.com/jonfriesen/simple-sso/pkg/sso"
)

type MemorySSO struct {
	Cookie *sso.CookieConfig
	User   string
	Pass   string
	Groups *[]string
}

// NewMemorySSO creates a in-memory SSO provider
func NewMemorySSO() (*MemorySSO, error) {

	setupBaseConfig()

	c, err := sso.SetupCookieConfig()
	if err != nil {
		return nil, err
	}

	return &MemorySSO{
		Cookie: c,
		User:   "alice",
		Pass:   "password123",
		Groups: &[]string{"admin", "moderator", "super"},
	}, nil

}

// Auth takes user,password strings as arguments and returns the user, user roles (e.g ldap groups)
// (string slice) if the call succeeds. Auth should return the ErrUnAuthorized or ErrUserNotFound error if
// auth fails or if the user is not found respectively.
func (m *MemorySSO) Auth(username string, password string) (*string, *[]string, error) {
	if m.User == username && m.Pass == password {
		return &m.User, m.Groups, nil
	}

	return nil, nil, errors.New("Bad User or Password")
}

// CTValidHours returns the cookie/jwt token validity in hours.
func (m *MemorySSO) CTValidHours() int64 {
	return m.Cookie.ValidHours
}

func (m *MemorySSO) CookieName() string {
	return m.Cookie.Name
}

func (m *MemorySSO) CookieDomain() string {
	return m.Cookie.Domain
}

// BuildJWTToken takes the user and the user roles info which is then signed by the private
// key of the login server. The expiry of the token is set per the third argument.
func (m *MemorySSO) BuildJWTToken(username string, groups []string, _ time.Time) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(m.CTValidHours())).UTC()
	return util.GenJWT(username, groups, PrivateKey, exp.Unix())
}

// BuildCookie takes the jwt token and returns a cookie and sets the expiration time of the same to that of
// the second arg.
func (m *MemorySSO) BuildCookie(s string, exp time.Time) http.Cookie {
	return http.Cookie{
		Name:     m.Cookie.Name,
		Value:    s,
		Domain:   m.Cookie.Domain,
		Path:     "/",
		Expires:  exp,
		MaxAge:   int(m.Cookie.ValidHours * 3600),
		Secure:   true,
		HttpOnly: true,
	}
}

// Logout sets the expiration time of the cookie in the past rendering it unusable.
func (m *MemorySSO) Logout(expT time.Time) http.Cookie {
	return http.Cookie{
		Name:     m.Cookie.Name,
		Value:    "",
		Domain:   m.Cookie.Domain,
		Path:     "/",
		Expires:  expT,
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
	}
}
