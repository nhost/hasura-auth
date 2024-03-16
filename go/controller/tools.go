package controller

import (
	"crypto/md5" //nolint:gosec
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GravatarURLFunc(enabled bool, def string, rating string) func(string) string {
	return func(email string) string {
		if !enabled {
			return ""
		}
		emailHash := md5.Sum([]byte(strings.ToLower(email))) //nolint:gosec
		return fmt.Sprintf("https://www.gravatar.com/avatar/%x?d=%s&r=%s", emailHash, def, rating)
	}
}

type LinkType string

const (
	LinkTypeEmailVerify        LinkType = "emailVerify"
	LinkTypeEmailConfirmChange LinkType = "emailConfirmChange"
	LinkTypePasswordlessEmail  LinkType = "passwordlessEmail"
	LinkTypePasswordReset      LinkType = "passwordReset"
)

func GenLink(serverURL url.URL, typ LinkType, ticket, redirectTo string) (string, error) {
	path, err := url.JoinPath(serverURL.Path, "verify")
	if err != nil {
		return "", fmt.Errorf("problem appending /verify to server url: %w", err)
	}
	serverURL.Path = path

	query := serverURL.Query()
	query.Add("type", string(typ))
	query.Add("ticket", ticket)
	query.Add("redirectTo", redirectTo)
	serverURL.RawQuery = query.Encode()

	return serverURL.String(), nil
}

func verifyHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	return string(hash), nil
}

func hashRefreshToken(token []byte) string {
	hash := sha256.Sum256(token)
	return "\\x" + hex.EncodeToString(hash[:])
}
