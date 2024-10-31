package myzap

const (
	startInstanceEndpoint  = "start"         // POST
	deleteInstanceEndpoint = "deleteSession" // POST

	sendTextEndpoint  = "sendText"  // POST
	sendImageEndpoint = "sendImage" // POST
	sendFileEndpoint  = "sendFile"  // POST

	tagsEndpoint           = "getAllLabels"      // POST
	createTagEndpoint      = "addNewLabel"       // POST
	addOrRemoveTagEndpoint = "addOrRemoveLabels" // POST

	getConnectionStatusEndpoint = "getConnectionStatus"
	getQrCodeEndpoint           = "getQrCodeString"
)
