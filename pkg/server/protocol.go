package server

// Naming convention in the package.
// When a action is required we prefix it with CMD (Command)
// When a data or information is passed we prefix variable as MSG(Message)
// Telemetry command to send some telemetry data
const CMD_Telemetry = "Telemetry"
const CMD_Auth = "Auth"
const CMD_Ping = "Ping"

//User commands
const (

	)

// Admin commands
const(
     CMD_MonitorConversation = "MonitorSession"
     CMD_JoinConversation = "JoinSession" // Join to an existing user session
)

// Server commands
const (
	CMD_ReceiveSessionId = "ReceiveSessionId"
	CMD_ReceiveBinaryData = "ReceiveData"
)


type ClientMsg struct {
	Command string
	SessionId string
	AuthToken string
	Data string
}

type ServerMsg struct {
	Command string
	Data string
	SessionId string
}

type TelemetryData struct {
	Type string
	Data string

}
