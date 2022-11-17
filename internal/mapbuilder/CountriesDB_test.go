package mapbuilder

import (
	"fmt"
	"testing"
	"time"
)

func TestLoadDB(t *testing.T) {
	fmt.Println(time.Now())
	err := LoadCountriesDB()
	fmt.Println(time.Now())
	if err != nil {
		t.Errorf("Could not read countries folder. Error: '%s'", err.Error())
	}
}
