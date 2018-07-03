package utils

import (
	"testing"
)

var openldap = &LDAP_CONFIG{
	Addr:       "ldap.example.org:389",
	BaseDn:     "dc=example,dc=org",
	BindDn:     "cn=manager,dc=example,dc=org",
	BindPass:   "password",
	AuthFilter: "(&(uid=%s))",
	Attributes: []string{},
	TLS:        false,
	StartTLS:   false,
}

var ad = &LDAP_CONFIG{
	Addr:       "ldap.example.org:636",
	BaseDn:     "dc=example.dc=org",
	BindDn:     "manager@example.org",
	BindPass:   "password",
	AuthFilter: "(&(sAMAccountName=%s))",
	Attributes: []string{"sAMAccountName", "displayName", "mail"},
	TLS:        true,
	StartTLS:   false,
}

func Test_ldap_auth_ad(t *testing.T) {
	err := ad.Connect()
	defer ad.Close()
	if err != nil {
		t.Error(err)
		return
	}

	err = ad.Auth("user", "pass")
	t.Log(err)
	err = ad.Auth("user2", "pass")
	t.Log(err)

}

func Test_ldap_auth_openldap(t *testing.T) {
	err := openldap.Connect()
	defer openldap.Close()
	if err != nil {
		t.Error(err)
		return
	}
	err = openldap.Auth("user", "pass")
	t.Log(err)
	err = openldap.Auth("user2", "pass")
	t.Log(err)
	openldap.Close()
}
