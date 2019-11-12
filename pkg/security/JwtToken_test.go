package security


import (
	"testing"
	"time"
)

func TestTokenFactory_MakeTokens(t *testing.T) {
	tf := &TokenFactory{
		secret:                []byte("thisissomesecret"),
		AccessExpireDuration:  3 * time.Minute,
		RefreshExpireDuration: 10 * time.Minute,
		Issuer:                "Issuer",
		Audience:              []string{"Audience"},
	}

	source := make(map[string]string)
	source["One"] = "ONE"
	source["Two"] = "TWO"

	access, refresh, err := tf.MakeTokens(source)
	if err != nil {
		t.Errorf("Got %v", err)
		t.Fail()
	}

	time.Sleep(500 * time.Millisecond)

	if valid, claims := tf.ValidateToken(access); !valid {
		t.Errorf("Not valid")
		t.Fail()
	} else {
		if claims.Get("TYPE") != "ACCESS" {
			t.Errorf("Not Access")
			t.Fail()
		}
		if claims.Get("One") != "ONE" {
			t.Errorf("Not ONE")
			t.Fail()
		}
	}

	if valid, claims := tf.ValidateToken(refresh); !valid {
		t.Errorf("Not valid")
		t.Fail()
	} else {
		if claims.Get("TYPE") != "REFRESH" {
			t.Errorf("Not refresh")
			t.Fail()
		}
		if claims.Get("One") != "ONE" {
			t.Errorf("Not ONE")
			t.Fail()
		}
	}

}

