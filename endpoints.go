package myzap

const (
	startInstanceEndpoint  = "start"         // POST
	deleteInstanceEndpoint = "deleteSession" // POST

	sendTextEndpoint = "sendText" // POST

	tagsEndpoint           = "getAllLabels"      // POST
	createTagEndpoint      = "addNewLabel"       // POST
	addOrRemoveTagEndpoint = "addOrRemoveLabels" // POST

	getConnectionStatusEndpoint = "getConnectionStatus"
	getQrCodeEndpoint           = "getQrCode"
)
