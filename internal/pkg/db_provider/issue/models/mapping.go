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

func MapFromProtoStatus(stat *status.TaskStatus) int32 {
	if stat.GetCommon().GetStatus() != status.TaskCommon_STATUS_COMMON_UNKNOWN {
		statNum := int32(stat.GetCommon().GetStatus().Number())
		return StatusCommon + statNum
	}
	if stat.GetDevelopment().GetStatus() != status.TaskDevelopment_STATUS_DEVELOPMENT_UNKNOWN {
		statNum := int32(stat.GetDevelopment().GetStatus().Number())
		return StatusDevelopment + statNum
	}
	return StatusUnknown
}

func MapToProtoStatus(stat int32) *status.TaskStatus {
	switch stat / 100 {
	case 1:
		statNum := stat % 100
		return &status.TaskStatus{Status: &status.TaskStatus_Common{Common: &status.TaskStatusCommon{Status: status.TaskCommon(statNum)}}}
	case 2:
		statNum := stat % 200
		return &status.TaskStatus{Status: &status.TaskStatus_Development{Development: &status.TaskStatusDevelopment{Status: status.TaskDevelopment(statNum)}}}
	default:
		return &status.TaskStatus{}
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
