package tools

import (
	"os"
	"testing"
)

func TestGetDcapConfigJsonString(t *testing.T) {

	identityJsonStr, tcbJsonStr, trustedJsonStr, err := GetDcapConfigJsonString()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if identityJsonStr == "" {
		t.Errorf("Error: identityJsonStr is empty")
	}
	if tcbJsonStr == "" {
		t.Errorf("Error: tcbJsonStr is empty")
	}
	if trustedJsonStr == "" {
		t.Errorf("Error: trustedJsonStr is empty")
	}

	// get config from configs directions
	identityJsonStrFromDir, err := os.ReadFile("../../configs/identity.json")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if identityJsonStr != string(identityJsonStrFromDir) {
		t.Errorf("Error: identityJsonStr is not equal to identity.json")
	}

	tcbJsonStrFromDir, err := os.ReadFile("../../configs/tcb.json")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if tcbJsonStr != string(tcbJsonStrFromDir) {
		t.Errorf("Error: tcbJsonStr is not equal to tcb.json")
	}

	trustedJsonStrFromDir, err := os.ReadFile("../../configs/trusted.json")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if trustedJsonStr != string(trustedJsonStrFromDir) {
		t.Errorf("Error: trustedJsonStr is not equal to trusted.json")
	}
}
