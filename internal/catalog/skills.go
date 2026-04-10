package catalog

import "github.com/gentleman-programming/gentle-ai/internal/model"

type Skill struct {
	ID       model.SkillID
	Name     string
	Category string
	Priority string
}

var mvpSkills = []Skill{
	// SDD skills anchor the engagement workflow across all pentesting domains.
	{ID: model.SkillSDDInit, Name: "sdd-init", Category: "sdd", Priority: "p0"},

	{ID: model.SkillSDDApply, Name: "sdd-apply", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDVerify, Name: "sdd-verify", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDExplore, Name: "sdd-explore", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDPropose, Name: "sdd-propose", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDSpec, Name: "sdd-spec", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDDesign, Name: "sdd-design", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDTasks, Name: "sdd-tasks", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDArchive, Name: "sdd-archive", Category: "sdd", Priority: "p0"},
	{ID: model.SkillSDDOnboard, Name: "sdd-onboard", Category: "sdd", Priority: "p0"},
	// Foundation skills support testing, workflow, registry, and review across capability layers.
	{ID: model.SkillGoTesting, Name: "go-testing", Category: "testing", Priority: "p0"},
	{ID: model.SkillCreator, Name: "skill-creator", Category: "workflow", Priority: "p0"},
	{ID: model.SkillJudgmentDay, Name: "judgment-day", Category: "workflow", Priority: "p0"},
	{ID: model.SkillBranchPR, Name: "branch-pr", Category: "workflow", Priority: "p0"},
	{ID: model.SkillIssueCreation, Name: "issue-creation", Category: "workflow", Priority: "p0"},
	{ID: model.SkillSkillRegistry, Name: "skill-registry", Category: "workflow", Priority: "p0"},
	// Domain packs expose post-install pentesting capability layers.
	{ID: model.SkillDomainWeb, Name: "domain-web", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainAPI, Name: "domain-api", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainMobile, Name: "domain-mobile", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainADInternal, Name: "domain-ad-internal", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainWiFi, Name: "domain-wifi-wireless", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainCloud, Name: "domain-cloud", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainRecon, Name: "domain-recon", Category: "domain", Priority: "p1"},
	{ID: model.SkillDomainReporting, Name: "domain-reporting", Category: "domain", Priority: "p1"},
	// Workflow packs expose post-install engagement-phase guidance.
	{ID: model.SkillWorkflowScoping, Name: "workflow-scoping", Category: "workflow-pack", Priority: "p1"},
	{ID: model.SkillWorkflowRecon, Name: "workflow-recon", Category: "workflow-pack", Priority: "p1"},
	{ID: model.SkillWorkflowEnumeration, Name: "workflow-enumeration", Category: "workflow-pack", Priority: "p1"},
	{ID: model.SkillWorkflowExploitation, Name: "workflow-exploitation", Category: "workflow-pack", Priority: "p1"},
	{ID: model.SkillWorkflowPostExploitation, Name: "workflow-post-exploitation", Category: "workflow-pack", Priority: "p1"},
	{ID: model.SkillWorkflowEvidence, Name: "workflow-evidence", Category: "workflow-pack", Priority: "p1"},
	{ID: model.SkillWorkflowReporting, Name: "workflow-reporting", Category: "workflow-pack", Priority: "p1"},
}

func MVPSkills() []Skill {
	skills := make([]Skill, len(mvpSkills))
	copy(skills, mvpSkills)
	return skills
}
