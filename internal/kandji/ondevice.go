package kandji

import (
	"errors"
	"fmt"
	"os"

	"howett.net/plist"
)

type KandjiProfileVars struct {
	// Needed Kandji variables from variables plist:
	// https://support.kandji.io/support/solutions/articles/72000560519-global-variables
	DeviceID    string `plist:"DEVICE_ID"`
	FullName    string `plist:"FULL_NAME"`
	EmailPrefix string `plist:"EMAIL_PREFIX"`
}

func GetPlistInfo() (KandjiProfileVars, error) {
	// Get the plist file from the Kandji variables
	plistPath := "/Library/Managed Preferences/io.kandji.globalvariables.plist"

	// Check if the file exists
	_, err := os.Stat(plistPath)
	if err != nil {
		return KandjiProfileVars{}, errors.New("kandji global variables plist does not exist")
	}

	// Open the file
	file, err := os.Open(plistPath)
	if err != nil {
		return KandjiProfileVars{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Decode the plist file
	var kVars KandjiProfileVars
	decoder := plist.NewDecoder(file)
	err = decoder.Decode(&kVars)
	if err != nil {
		return KandjiProfileVars{}, fmt.Errorf("error decoding plist: %v", err)
	}

	return kVars, nil
}