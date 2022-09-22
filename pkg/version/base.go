package version

var AgentVersion string

var Commit string

var agentVersionDefault = "0.0.1"

func init() {
	if AgentVersion == "" {
		AgentVersion = agentVersionDefault
	}
}
