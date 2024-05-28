package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeNow(t *testing.T) {
	fmt.Println(time.Now().Local().UTC())
}
