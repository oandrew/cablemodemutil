package cablemodemutil

// CableModemRawStatus contains the raw status retrieved from the cable modem without any parsing.
type CableModemRawStatus map[string]interface{}

// CableModemDeviceInfo contains Cable Modem Device information.
type CableModemDeviceInfo struct {
	// Cable Modem model.
	Model string
	// Cable Modem serial number.
	SerialNumber string
	// Cable Modem MAC address.
	MACAddress string
}

// CableModemDeviceSettings contains Cable Modem Device settings.
type CableModemDeviceSettings struct {
	// True if front panel LED lights are configured to be on, false otherwise.
	FrontPanelLightsOn bool
	// True if energy efficient ethernet setting is turned on, false otherwise.
	EnergyEfficientEthernetOn bool
	// True if ask me later setting has been opted into, false otherwise.
	AskMeLater bool
	// True if never ask setting has been opted into, false otherwise.
	NeverAsk bool
}

// CableModemStatus contains Cable Modem Status.
type CableModemStatus struct {
	// Device related information.
	DeviceInfo CableModemDeviceInfo
	// General settings.
	DeviceSettings CableModemDeviceSettings
}
