package contracts

import (
	"fmt"
	"time"
)

type OnlyDate struct {
	time.Time
}

func (m *OnlyDate) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.DateOnly+`"`, string(data))
	fmt.Println("OnlyDate", tt)
	*m = OnlyDate{tt}
	return err
}
