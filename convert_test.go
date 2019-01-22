package struct2map

import (
	"testing"
)

type User struct {
	ID       int64  `field:"id"`
	Username string `field:"username"`
	Gender   bool   `field:"gender"`
}

func TestConvert(t *testing.T) {
	user := User{
		ID:       111,
		Username: "dxvgef",
		Gender:   true,
	}
	mapData, err := Convert(&user, []string{"id", "username"}, "field")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(mapData)
}

func TestConvertSlice(t *testing.T) {
	users := make([]User, 2)
	users[0] = User{
		ID:       111,
		Username: "dxvgef",
		Gender:   true,
	}
	users[1] = User{
		ID:       222,
		Username: "test",
		Gender:   true,
	}

	mapData, err := ConvertSlice(users, []string{"id", "username"}, "field")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(mapData)
}
