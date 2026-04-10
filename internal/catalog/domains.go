package catalog

type PentestingDomain struct {
	Name        string
	Description string
}

type EngagementPhase struct {
	Name        string
	Description string
}

var pentestingDomains = []PentestingDomain{
	{Name: "AD/Internal", Description: "Internal network, Active Directory, identity, privilege and lateral movement assessments."},
	{Name: "Mobile", Description: "Mobile application and device assessment workflows for iOS and Android targets."},
	{Name: "Web", Description: "Web application assessment workflows for authenticated and unauthenticated surfaces."},
	{Name: "API", Description: "API discovery, authorization, business-logic and abuse-case assessment workflows."},
	{Name: "WiFi/Wireless", Description: "Wireless reconnaissance, authentication, access, and radio-layer attack workflows."},
	{Name: "Cloud", Description: "Cloud control-plane, identity, storage, secrets, and service-configuration assessment workflows."},
	{Name: "Reporting", Description: "Evidence capture, reproducibility, remediation framing, and stakeholder-ready reporting workflows."},
	{Name: "Recon", Description: "Reconnaissance and OSINT workflows that feed later engagement phases."},
}

var engagementPhases = []EngagementPhase{
	{Name: "Scoping", Description: "Define target boundaries, constraints, assumptions, and success criteria."},
	{Name: "Recon", Description: "Gather external and internal context before active validation."},
	{Name: "Enumeration", Description: "Enumerate services, identities, permissions, attack paths, and exposed behaviors."},
	{Name: "Exploitation", Description: "Validate exploitability and demonstrate impact safely within scope."},
	{Name: "Post-exploitation", Description: "Assess pivoting, escalation, persistence implications, and blast radius."},
	{Name: "Evidence", Description: "Capture artifacts, commands, screenshots, indicators, and validation data."},
	{Name: "Reporting", Description: "Translate findings into narrative, severity, remediation, and executive communication."},
}

func PentestingDomains() []PentestingDomain {
	domains := make([]PentestingDomain, len(pentestingDomains))
	copy(domains, pentestingDomains)
	return domains
}

func EngagementPhases() []EngagementPhase {
	phases := make([]EngagementPhase, len(engagementPhases))
	copy(phases, engagementPhases)
	return phases
}
