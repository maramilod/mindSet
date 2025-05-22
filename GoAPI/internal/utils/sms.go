package utils

import (
	"fmt"
)

func SendSMSCode(phone string, code string) error {

	fmt.Printf("[MOCK SMS] Sending code %s to phone %s\n", code, phone)


	return nil
}
