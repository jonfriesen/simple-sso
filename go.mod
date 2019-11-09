module github.com/jonfriesen/simple-sso

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/samitpal/goProbe v0.5.1
	gopkg.in/asn1-ber.v1 v1.0.0-00010101000000-000000000000 // indirect
	gopkg.in/ldap.v2 v2.5.1
)

// Hacks
replace gopkg.in/asn1-ber.v1 => github.com/go-asn1-ber/asn1-ber v1.3.1
