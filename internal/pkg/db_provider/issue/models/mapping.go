package models

import (
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/priority"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/status"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/state/template"
	"github.com/Constantine27K/crnt-data-manager/pkg/api/tasks/issue"
)

func MapFromProtoIssueType(issueType issue.IssueType) IssueType {
	switch issueType {
	case issue.IssueType_ISSUE_TYPE_STORY:
		return IssueTypeStory
	case issue.IssueType_ISSUE_TYPE_TASK:
		return IssueTypeTask
	case issue.IssueType_ISSUE_TYPE_SUBTASK:
		return IssueTypeSubtask
	default:
		return IssueTypeUnknown
	}
}

func MapToProtoIssueType(issueType IssueType) issue.IssueType {
	switch issueType {
	case IssueTypeStory:
		return issue.IssueType_ISSUE_TYPE_STORY
	case IssueTypeTask:
		return issue.IssueType_ISSUE_TYPE_TASK
	case IssueTypeSubtask:
		return issue.IssueType_ISSUE_TYPE_SUBTASK
	default:
		return issue.IssueType_ISSUE_TYPE_UNKNOWN
	}
}

func MapFromProtoTemplate(temp template.Template) Template {
	switch temp {
	case template.Template_TEMPLATE_COMMON:
		return TemplateCommon
	case template.Template_TEMPLATE_DEVELOPMENT:
		return TemplateDevelopment
	default:
		return TemplateUnknown
	}
}

func MapToProtoTemplate(temp Template) template.Template {
	switch temp {
	case TemplateCommon:
		return template.Template_TEMPLATE_COMMON
	case TemplateDevelopment:
		return template.Template_TEMPLATE_DEVELOPMENT
	default:
		return template.Template_TEMPLATE_UNKNOWN
	}
}

func MapFromProtoStatus(stat *status.IssueStatus) int32 {
	if stat.GetCommon().GetStatus() != status.Common_STATUS_COMMON_UNKNOWN {
		statNum := int32(stat.GetCommon().GetStatus().Number())
		return StatusCommon + statNum
	}
	if stat.GetDevelopment().GetStatus() != status.Development_STATUS_DEVELOPMENT_UNKNOWN {
		statNum := int32(stat.GetDevelopment().GetStatus().Number())
		return StatusDevelopment + statNum
	}
	if stat.GetEpic().GetStatus() != status.Epic_STATUS_EPIC_UNKNOWN {
		statNum := int32(stat.GetEpic().GetStatus().Number())
		return StatusEpic + statNum
	}
	return StatusUnknown
}

func MapToProtoStatus(stat int32) *status.IssueStatus {
	switch stat / 100 {
	case 1:
		statNum := stat % 100
		return &status.IssueStatus{Status: &status.IssueStatus_Common{Common: &status.IssueStatusCommon{Status: status.Common(statNum)}}}
	case 2:
		statNum := stat % 200
		return &status.IssueStatus{Status: &status.IssueStatus_Development{Development: &status.IssueStatusDevelopment{Status: status.Development(statNum)}}}
	case 3:
		statNum := stat % 200
		return &status.IssueStatus{Status: &status.IssueStatus_Epic{Epic: &status.IssueStatusEpic{Status: status.Epic(statNum)}}}
	default:
		return &status.IssueStatus{}
	}
}

func MapFromProtoPriority(pripor priority.Priority) Priority {
	switch pripor {
	case priority.Priority_PRIORITY_CRITICAL:
		return PriorityCritical
	case priority.Priority_PRIORITY_HIGH:
		return PriorityHigh
	case priority.Priority_PRIORITY_MEDIUM:
		return PriorityMedium
	case priority.Priority_PRIORITY_LOW:
		return PriorityLow
	default:
		return PriorityUnknown
	}
}

func MapToProtoPriority(prior Priority) priority.Priority {
	switch prior {
	case PriorityCritical:
		return priority.Priority_PRIORITY_CRITICAL
	case PriorityHigh:
		return priority.Priority_PRIORITY_HIGH
	case PriorityMedium:
		return priority.Priority_PRIORITY_MEDIUM
	case PriorityLow:
		return priority.Priority_PRIORITY_LOW
	default:
		return priority.Priority_PRIORITY_UNKNOWN
	}
}
