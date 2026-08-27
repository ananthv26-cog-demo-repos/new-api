package constant

// FireConnect CLI stamps attribution headers on every harness request so a
// self-hosted gateway can tell which harness and CLI version produced it.
const (
	FireconnectHeaderPrefix  = "x-fireconnect-"
	FireconnectHarnessHeader = "X-FireConnect-Harness"
	FireconnectVersionHeader = "X-FireConnect-Version"
)
