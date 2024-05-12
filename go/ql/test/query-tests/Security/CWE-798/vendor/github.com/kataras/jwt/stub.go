// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for github.com/kataras/jwt, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/kataras/jwt (exports: Keys,Alg; functions: Sign,SignEncrypted,SignEncryptedWithHeader,SignWithHeader)

// Package jwt is a stub of github.com/kataras/jwt, generated by depstubber.
package jwt

import (
	time "time"
)

type Alg interface {
	Name() string
	Sign(_ interface{}, _ []byte) ([]byte, error)
	Verify(_ interface{}, _ []byte, _ []byte) error
}

type Audience []string

func (_ *Audience) UnmarshalJSON(_ []byte) error {
	return nil
}

type Claims struct {
	NotBefore int64
	IssuedAt  int64
	Expiry    int64
	ID        string
	OriginID  string
	Issuer    string
	Subject   string
	Audience  Audience
}

func (_ Claims) Age() time.Duration {
	return 0
}

func (_ Claims) ApplyClaims(_ *Claims) {}

func (_ Claims) ExpiresAt() time.Time {
	return time.Time{}
}

func (_ Claims) Timeleft() time.Duration {
	return 0
}

type InjectFunc func([]byte) ([]byte, error)

type Key struct {
	ID      string
	Alg     Alg
	Public  interface{}
	Private interface{}
	MaxAge  time.Duration
	Encrypt InjectFunc
	Decrypt InjectFunc
}

type Keys map[string]*Key

func (_ Keys) Get(_ string) (*Key, bool) {
	return nil, false
}

func (_ Keys) Register(_ Alg, _ string, _ interface{}, _ interface{}) {}

func (_ Keys) SignToken(_ string, _ interface{}, _ ...SignOption) ([]byte, error) {
	return nil, nil
}

func (_ Keys) ValidateHeader(_ string, _ []byte) (Alg, interface{}, InjectFunc, error) {
	return nil, nil, nil, nil
}

func (_ Keys) VerifyToken(_ []byte, _ interface{}, _ ...TokenValidator) error {
	return nil
}

func Sign(_ Alg, _ interface{}, _ interface{}, _ ...SignOption) ([]byte, error) {
	return nil, nil
}

func SignEncrypted(_ Alg, _ interface{}, _ InjectFunc, _ interface{}, _ ...SignOption) ([]byte, error) {
	return nil, nil
}

func SignEncryptedWithHeader(_ Alg, _ interface{}, _ InjectFunc, _ interface{}, _ interface{}, _ ...SignOption) ([]byte, error) {
	return nil, nil
}

type SignOption interface {
	ApplyClaims(_ *Claims)
}

func SignWithHeader(_ Alg, _ interface{}, _ interface{}, _ interface{}, _ ...SignOption) ([]byte, error) {
	return nil, nil
}

type TokenValidator interface {
	ValidateToken(_ []byte, _ Claims, _ error) error
}
