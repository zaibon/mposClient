package mposClient

import (
	"testing"
)

func TestPublic(t *testing.T) {
	c := NewClient("http://ebt.laminerie.eu/", "0213e2df691e170d9941d42f5a8124c96d7c79406ca1549429affdf0aa6d2e96")
	info, err := c.PublicInfo()

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if info.PoolName != "La Minerie" {
		t.Fail()
	}
	t.Log(info)
}
