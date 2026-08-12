package jwt_test

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestNumericDateUnmarshalLargeSeconds(t *testing.T) {
	oldPrecision := jwt.TimePrecision
	t.Cleanup(func() { jwt.TimePrecision = oldPrecision })
	jwt.TimePrecision = time.Nanosecond

	var date jwt.NumericDate
	value := "9223372036.5"
	if err := json.Unmarshal([]byte(value), &date); err != nil {
		t.Fatal(err)
	}
	if got := date.Unix(); got != int64(math.Floor(9223372036.5)) {
		t.Fatalf("unexpected Unix seconds: got %d want 9223372036", got)
	}
	if got := date.Nanosecond(); got != 500000000 {
		t.Fatalf("unexpected nanoseconds: got %d want 500000000", got)
	}
}

func TestNumericDate(t *testing.T) {
	var s struct {
		Iat jwt.NumericDate `json:"iat"`
		Exp jwt.NumericDate `json:"exp"`
	}

	oldPrecision := jwt.TimePrecision

	jwt.TimePrecision = time.Microsecond

	raw := `{"iat":1516239022,"exp":1516239022.12345}`

	err := json.Unmarshal([]byte(raw), &s)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	b, _ := json.Marshal(s)

	if raw != string(b) {
		t.Errorf("Serialized format of numeric date mismatch. Expecting: %s  Got: %s", string(raw), string(b))
	}

	jwt.TimePrecision = oldPrecision
}

func TestSingleArrayMarshal(t *testing.T) {
	jwt.MarshalSingleStringAsArray = false

	s := jwt.ClaimStrings{"test"}
	expected := `"test"`

	b, err := json.Marshal(s)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if expected != string(b) {
		t.Errorf("Serialized format of string array mismatch. Expecting: %s  Got: %s", string(expected), string(b))
	}

	jwt.MarshalSingleStringAsArray = true

	expected = `["test"]`

	b, err = json.Marshal(s)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if expected != string(b) {
		t.Errorf("Serialized format of string array mismatch. Expecting: %s  Got: %s", string(expected), string(b))
	}
}
