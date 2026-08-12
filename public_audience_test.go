package jwt

import (
	"errors"
	"testing"
)

func TestValidatorAcceptsAnyExpectedAudience(t *testing.T) {
	claims := RegisteredClaims{Audience: ClaimStrings{"service-b"}}
	validator := NewValidator(WithAudience("service-a", "service-b"))
	for i := 0; i < 100; i++ {
		if err := validator.Validate(claims); err != nil {
			t.Fatalf("Validate returned an error for a matching audience on iteration %d: %v", i, err)
		}
	}
}

func TestValidatorRequiresEveryExpectedAudience(t *testing.T) {
	claims := RegisteredClaims{Audience: ClaimStrings{"service-a"}}
	err := NewValidator(WithAllAudiences("service-a", "service-b")).Validate(claims)
	if !errors.Is(err, ErrTokenInvalidAudience) {
		t.Fatalf("Validate error = %v, want ErrTokenInvalidAudience", err)
	}
}
