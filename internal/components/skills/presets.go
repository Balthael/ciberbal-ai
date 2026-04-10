package skills

import "github.com/gentleman-programming/gentle-ai/internal/model"

// sddSkills are the SDD orchestrator skills — always included.
var sddSkills = []model.SkillID{
	model.SkillSDDInit,
	model.SkillSDDExplore,
	model.SkillSDDPropose,
	model.SkillSDDSpec,
	model.SkillSDDDesign,
	model.SkillSDDTasks,
	model.SkillSDDApply,
	model.SkillSDDVerify,
	model.SkillSDDArchive,
	model.SkillSDDOnboard,
	model.SkillJudgmentDay,
}

// foundationSkills are baseline learning skills for the "recommended" tier.
var foundationSkills = []model.SkillID{
	model.SkillGoTesting,
	model.SkillCreator,
	model.SkillBranchPR,
	model.SkillIssueCreation,
	model.SkillSkillRegistry,
}

var domainSkills = []model.SkillID{
	model.SkillDomainWeb,
	model.SkillDomainAPI,
	model.SkillDomainMobile,
	model.SkillDomainADInternal,
	model.SkillDomainWiFi,
	model.SkillDomainCloud,
	model.SkillDomainRecon,
	model.SkillDomainReporting,
}

var workflowSkills = []model.SkillID{
	model.SkillWorkflowScoping,
	model.SkillWorkflowRecon,
	model.SkillWorkflowEnumeration,
	model.SkillWorkflowExploitation,
	model.SkillWorkflowPostExploitation,
	model.SkillWorkflowEvidence,
	model.SkillWorkflowReporting,
}

// SkillsForPreset returns which skills should be installed for a given preset.
//
//   - "minimal" / PresetMinimal:            SDD skills only
//   - "ecosystem-core" / PresetEcosystemCore: SDD + common workflow skills
//   - "full-pentest" / PresetFullPentest: all available skills
//   - "custom" / PresetCustom:            empty (caller should provide explicit list)
func SkillsForPreset(preset model.PresetID) []model.SkillID {
	switch preset {
	case model.PresetMinimal:
		return copySkills(sddSkills)
	case model.PresetEcosystemCore:
		return copySkills(append(sddSkills, foundationSkills...))
	case model.PresetFullPentest:
		all := make([]model.SkillID, 0, len(sddSkills)+len(foundationSkills)+len(domainSkills)+len(workflowSkills))
		all = append(all, sddSkills...)
		all = append(all, foundationSkills...)
		all = append(all, domainSkills...)
		all = append(all, workflowSkills...)
		return all
	case model.PresetCustom:
		return nil
	default:
		// Unknown preset — default to full.
		all := make([]model.SkillID, 0, len(sddSkills)+len(foundationSkills)+len(domainSkills)+len(workflowSkills))
		all = append(all, sddSkills...)
		all = append(all, foundationSkills...)
		all = append(all, domainSkills...)
		all = append(all, workflowSkills...)
		return all
	}
}

// AllSkillIDs returns every known skill ID.
func AllSkillIDs() []model.SkillID {
	all := make([]model.SkillID, 0, len(sddSkills)+len(foundationSkills)+len(domainSkills)+len(workflowSkills))
	all = append(all, sddSkills...)
	all = append(all, foundationSkills...)
	all = append(all, domainSkills...)
	all = append(all, workflowSkills...)
	return all
}

func copySkills(src []model.SkillID) []model.SkillID {
	dst := make([]model.SkillID, len(src))
	copy(dst, src)
	return dst
}
