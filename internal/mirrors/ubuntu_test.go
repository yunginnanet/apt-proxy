package mirrors

import "testing"

func TestGetUbuntuMirrorUrlsByGeo(t *testing.T) {
	mirrors, err := GetUbuntuMirrorUrlsByGeo()
	if err != nil {
		t.Fatal(err)
	}
	if len(mirrors) == 0 {
		t.Fatal("get ubuntu get mirrors failed")
	}
}
