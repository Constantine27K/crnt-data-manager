///
//  Generated code. Do not modify.
//  source: api/tasks/issue/issue.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../comments/comments.pbjson.dart' as $3;
import '../../../google/protobuf/timestamp.pbjson.dart' as $0;
import '../../state/status/status.pbjson.dart' as $4;

@$core.Deprecated('Use issueTypeDescriptor instead')
const IssueType$json = const {
  '1': 'IssueType',
  '2': const [
    const {'1': 'ISSUE_TYPE_UNKNOWN', '2': 0},
    const {'1': 'ISSUE_TYPE_EPIC', '2': 1},
    const {'1': 'ISSUE_TYPE_STORY', '2': 2},
    const {'1': 'ISSUE_TYPE_TASK', '2': 3},
    const {'1': 'ISSUE_TYPE_SUBTASK', '2': 4},
  ],
};

/// Descriptor for `IssueType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List issueTypeDescriptor = $convert.base64Decode('CglJc3N1ZVR5cGUSFgoSSVNTVUVfVFlQRV9VTktOT1dOEAASEwoPSVNTVUVfVFlQRV9FUElDEAESFAoQSVNTVUVfVFlQRV9TVE9SWRACEhMKD0lTU1VFX1RZUEVfVEFTSxADEhYKEklTU1VFX1RZUEVfU1VCVEFTSxAE');
@$core.Deprecated('Use issueDescriptor instead')
const Issue$json = const {
  '1': 'Issue',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'composite_name', '3': 2, '4': 1, '5': 9, '10': 'compositeName'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'type', '3': 4, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueType', '10': 'type'},
    const {'1': 'parent_id', '3': 5, '4': 1, '5': 3, '10': 'parentId'},
    const {'1': 'description', '3': 6, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'comments', '3': 7, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.comments.Comments', '10': 'comments'},
    const {'1': 'author', '3': 8, '4': 1, '5': 9, '10': 'author'},
    const {'1': 'assigned', '3': 9, '4': 1, '5': 9, '10': 'assigned'},
    const {'1': 'qa', '3': 10, '4': 1, '5': 9, '10': 'qa'},
    const {'1': 'reviewer', '3': 11, '4': 1, '5': 9, '10': 'reviewer'},
    const {'1': 'template', '3': 12, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.template.Template', '10': 'template'},
    const {'1': 'created_at', '3': 13, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'createdAt'},
    const {'1': 'deadline', '3': 14, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'deadline'},
    const {'1': 'updated_at', '3': 15, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    const {'1': 'status', '3': 16, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatus', '10': 'status'},
    const {'1': 'priority', '3': 17, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.priority.Priority', '10': 'priority'},
    const {'1': 'sprint_id', '3': 18, '4': 1, '5': 3, '10': 'sprintId'},
    const {'1': 'project_id', '3': 19, '4': 1, '5': 3, '10': 'projectId'},
    const {'1': 'component_ids', '3': 20, '4': 3, '5': 3, '10': 'componentIds'},
    const {'1': 'story_points', '3': 21, '4': 1, '5': 3, '10': 'storyPoints'},
    const {'1': 'payment', '3': 22, '4': 1, '5': 1, '10': 'payment'},
    const {'1': 'time_spent', '3': 23, '4': 1, '5': 3, '10': 'timeSpent'},
    const {'1': 'children', '3': 24, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue', '10': 'children'},
  ],
};

/// Descriptor for `Issue`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueDescriptor = $convert.base64Decode('CgVJc3N1ZRIOCgJpZBgBIAEoA1ICaWQSJQoOY29tcG9zaXRlX25hbWUYAiABKAlSDWNvbXBvc2l0ZU5hbWUSEgoEbmFtZRgDIAEoCVIEbmFtZRJWCgR0eXBlGAQgASgOMkIuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZVR5cGVSBHR5cGUSGwoJcGFyZW50X2lkGAUgASgDUghwYXJlbnRJZBIgCgtkZXNjcmlwdGlvbhgGIAEoCVILZGVzY3JpcHRpb24SWgoIY29tbWVudHMYByABKAsyPi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmNvbW1lbnRzLkNvbW1lbnRzUghjb21tZW50cxIWCgZhdXRob3IYCCABKAlSBmF1dGhvchIaCghhc3NpZ25lZBgJIAEoCVIIYXNzaWduZWQSDgoCcWEYCiABKAlSAnFhEhoKCHJldmlld2VyGAsgASgJUghyZXZpZXdlchJgCgh0ZW1wbGF0ZRgMIAEoDjJELmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3RhdGUudGVtcGxhdGUuVGVtcGxhdGVSCHRlbXBsYXRlEjkKCmNyZWF0ZWRfYXQYDSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgljcmVhdGVkQXQSNgoIZGVhZGxpbmUYDiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUghkZWFkbGluZRI5Cgp1cGRhdGVkX2F0GA8gASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJdXBkYXRlZEF0El0KBnN0YXR1cxgQIAEoCzJFLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3RhdGUuc3RhdHVzLklzc3VlU3RhdHVzUgZzdGF0dXMSYAoIcHJpb3JpdHkYESABKA4yRC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnByaW9yaXR5LlByaW9yaXR5Ughwcmlvcml0eRIbCglzcHJpbnRfaWQYEiABKANSCHNwcmludElkEh0KCnByb2plY3RfaWQYEyABKANSCXByb2plY3RJZBIjCg1jb21wb25lbnRfaWRzGBQgAygDUgxjb21wb25lbnRJZHMSIQoMc3RvcnlfcG9pbnRzGBUgASgDUgtzdG9yeVBvaW50cxIYCgdwYXltZW50GBYgASgBUgdwYXltZW50Eh0KCnRpbWVfc3BlbnQYFyABKANSCXRpbWVTcGVudBJaCghjaGlsZHJlbhgYIAMoCzI+LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVSCGNoaWxkcmVu');
@$core.Deprecated('Use issueInfoDescriptor instead')
const IssueInfo$json = const {
  '1': 'IssueInfo',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'composite_name', '3': 2, '4': 1, '5': 9, '10': 'compositeName'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'type', '3': 4, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueType', '10': 'type'},
    const {'1': 'assigned', '3': 5, '4': 1, '5': 9, '10': 'assigned'},
    const {'1': 'priority', '3': 6, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.priority.Priority', '10': 'priority'},
    const {'1': 'status', '3': 7, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatus', '10': 'status'},
    const {'1': 'story_points', '3': 8, '4': 1, '5': 3, '10': 'storyPoints'},
  ],
};

/// Descriptor for `IssueInfo`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueInfoDescriptor = $convert.base64Decode('CglJc3N1ZUluZm8SDgoCaWQYASABKANSAmlkEiUKDmNvbXBvc2l0ZV9uYW1lGAIgASgJUg1jb21wb3NpdGVOYW1lEhIKBG5hbWUYAyABKAlSBG5hbWUSVgoEdHlwZRgEIAEoDjJCLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVUeXBlUgR0eXBlEhoKCGFzc2lnbmVkGAUgASgJUghhc3NpZ25lZBJgCghwcmlvcml0eRgGIAEoDjJELmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3RhdGUucHJpb3JpdHkuUHJpb3JpdHlSCHByaW9yaXR5El0KBnN0YXR1cxgHIAEoCzJFLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3RhdGUuc3RhdHVzLklzc3VlU3RhdHVzUgZzdGF0dXMSIQoMc3RvcnlfcG9pbnRzGAggASgDUgtzdG9yeVBvaW50cw==');
@$core.Deprecated('Use issueCreateRequestDescriptor instead')
const IssueCreateRequest$json = const {
  '1': 'IssueCreateRequest',
  '2': const [
    const {'1': 'issue', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue', '10': 'issue'},
  ],
};

/// Descriptor for `IssueCreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueCreateRequestDescriptor = $convert.base64Decode('ChJJc3N1ZUNyZWF0ZVJlcXVlc3QSVAoFaXNzdWUYASABKAsyPi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRhc2tzLmlzc3VlLklzc3VlUgVpc3N1ZQ==');
@$core.Deprecated('Use issueCreateResponseDescriptor instead')
const IssueCreateResponse$json = const {
  '1': 'IssueCreateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `IssueCreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueCreateResponseDescriptor = $convert.base64Decode('ChNJc3N1ZUNyZWF0ZVJlc3BvbnNlEg4KAmlkGAEgASgDUgJpZA==');
@$core.Deprecated('Use issueCreateSubtaskRequestDescriptor instead')
const IssueCreateSubtaskRequest$json = const {
  '1': 'IssueCreateSubtaskRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'child', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue', '10': 'child'},
  ],
};

/// Descriptor for `IssueCreateSubtaskRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueCreateSubtaskRequestDescriptor = $convert.base64Decode('ChlJc3N1ZUNyZWF0ZVN1YnRhc2tSZXF1ZXN0Eg4KAmlkGAEgASgDUgJpZBJUCgVjaGlsZBgCIAEoCzI+LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVSBWNoaWxk');
@$core.Deprecated('Use issueCreateSubtaskResponseDescriptor instead')
const IssueCreateSubtaskResponse$json = const {
  '1': 'IssueCreateSubtaskResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `IssueCreateSubtaskResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueCreateSubtaskResponseDescriptor = $convert.base64Decode('ChpJc3N1ZUNyZWF0ZVN1YnRhc2tSZXNwb25zZRIOCgJpZBgBIAEoA1ICaWQ=');
@$core.Deprecated('Use issueAddCommentRequestDescriptor instead')
const IssueAddCommentRequest$json = const {
  '1': 'IssueAddCommentRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'comment', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.comments.Comment', '10': 'comment'},
  ],
};

/// Descriptor for `IssueAddCommentRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueAddCommentRequestDescriptor = $convert.base64Decode('ChZJc3N1ZUFkZENvbW1lbnRSZXF1ZXN0Eg4KAmlkGAEgASgDUgJpZBJXCgdjb21tZW50GAIgASgLMj0uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5jb21tZW50cy5Db21tZW50Ugdjb21tZW50');
@$core.Deprecated('Use issueAddCommentResponseDescriptor instead')
const IssueAddCommentResponse$json = const {
  '1': 'IssueAddCommentResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `IssueAddCommentResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueAddCommentResponseDescriptor = $convert.base64Decode('ChdJc3N1ZUFkZENvbW1lbnRSZXNwb25zZRIOCgJpZBgBIAEoA1ICaWQ=');
@$core.Deprecated('Use issueUpdateRequestDescriptor instead')
const IssueUpdateRequest$json = const {
  '1': 'IssueUpdateRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'issue', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue', '10': 'issue'},
  ],
};

/// Descriptor for `IssueUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueUpdateRequestDescriptor = $convert.base64Decode('ChJJc3N1ZVVwZGF0ZVJlcXVlc3QSDgoCaWQYASABKANSAmlkElQKBWlzc3VlGAIgASgLMj4uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZVIFaXNzdWU=');
@$core.Deprecated('Use issueUpdateResponseDescriptor instead')
const IssueUpdateResponse$json = const {
  '1': 'IssueUpdateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `IssueUpdateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueUpdateResponseDescriptor = $convert.base64Decode('ChNJc3N1ZVVwZGF0ZVJlc3BvbnNlEg4KAmlkGAEgASgDUgJpZA==');
@$core.Deprecated('Use issueGetRequestDescriptor instead')
const IssueGetRequest$json = const {
  '1': 'IssueGetRequest',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 3, '10': 'ids'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'parent_id', '3': 3, '4': 1, '5': 3, '10': 'parentId'},
    const {'1': 'type', '3': 4, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueType', '10': 'type'},
    const {'1': 'author', '3': 5, '4': 1, '5': 9, '10': 'author'},
    const {'1': 'assigned', '3': 6, '4': 1, '5': 9, '10': 'assigned'},
    const {'1': 'sprint_id', '3': 7, '4': 1, '5': 3, '10': 'sprintId'},
    const {'1': 'project_id', '3': 8, '4': 1, '5': 3, '10': 'projectId'},
    const {'1': 'status', '3': 9, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatus', '10': 'status'},
  ],
};

/// Descriptor for `IssueGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueGetRequestDescriptor = $convert.base64Decode('Cg9Jc3N1ZUdldFJlcXVlc3QSEAoDaWRzGAEgAygDUgNpZHMSEgoEbmFtZRgCIAEoCVIEbmFtZRIbCglwYXJlbnRfaWQYAyABKANSCHBhcmVudElkElYKBHR5cGUYBCABKA4yQi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRhc2tzLmlzc3VlLklzc3VlVHlwZVIEdHlwZRIWCgZhdXRob3IYBSABKAlSBmF1dGhvchIaCghhc3NpZ25lZBgGIAEoCVIIYXNzaWduZWQSGwoJc3ByaW50X2lkGAcgASgDUghzcHJpbnRJZBIdCgpwcm9qZWN0X2lkGAggASgDUglwcm9qZWN0SWQSXQoGc3RhdHVzGAkgASgLMkUuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zdGF0ZS5zdGF0dXMuSXNzdWVTdGF0dXNSBnN0YXR1cw==');
@$core.Deprecated('Use issueGetResponseDescriptor instead')
const IssueGetResponse$json = const {
  '1': 'IssueGetResponse',
  '2': const [
    const {'1': 'issues', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue', '10': 'issues'},
  ],
};

/// Descriptor for `IssueGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueGetResponseDescriptor = $convert.base64Decode('ChBJc3N1ZUdldFJlc3BvbnNlElYKBmlzc3VlcxgBIAMoCzI+LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVSBmlzc3Vlcw==');
@$core.Deprecated('Use issueInfoGetRequestDescriptor instead')
const IssueInfoGetRequest$json = const {
  '1': 'IssueInfoGetRequest',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 3, '10': 'ids'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'parent_id', '3': 3, '4': 1, '5': 3, '10': 'parentId'},
    const {'1': 'type', '3': 4, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueType', '10': 'type'},
    const {'1': 'author', '3': 5, '4': 1, '5': 9, '10': 'author'},
    const {'1': 'assigned', '3': 6, '4': 1, '5': 9, '10': 'assigned'},
    const {'1': 'sprint_id', '3': 7, '4': 1, '5': 3, '10': 'sprintId'},
    const {'1': 'project_id', '3': 8, '4': 1, '5': 3, '10': 'projectId'},
    const {'1': 'status', '3': 9, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatus', '10': 'status'},
  ],
};

/// Descriptor for `IssueInfoGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueInfoGetRequestDescriptor = $convert.base64Decode('ChNJc3N1ZUluZm9HZXRSZXF1ZXN0EhAKA2lkcxgBIAMoA1IDaWRzEhIKBG5hbWUYAiABKAlSBG5hbWUSGwoJcGFyZW50X2lkGAMgASgDUghwYXJlbnRJZBJWCgR0eXBlGAQgASgOMkIuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZVR5cGVSBHR5cGUSFgoGYXV0aG9yGAUgASgJUgZhdXRob3ISGgoIYXNzaWduZWQYBiABKAlSCGFzc2lnbmVkEhsKCXNwcmludF9pZBgHIAEoA1IIc3ByaW50SWQSHQoKcHJvamVjdF9pZBgIIAEoA1IJcHJvamVjdElkEl0KBnN0YXR1cxgJIAEoCzJFLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3RhdGUuc3RhdHVzLklzc3VlU3RhdHVzUgZzdGF0dXM=');
@$core.Deprecated('Use issueInfoGetResponseDescriptor instead')
const IssueInfoGetResponse$json = const {
  '1': 'IssueInfoGetResponse',
  '2': const [
    const {'1': 'issue_info', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfo', '10': 'issueInfo'},
  ],
};

/// Descriptor for `IssueInfoGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueInfoGetResponseDescriptor = $convert.base64Decode('ChRJc3N1ZUluZm9HZXRSZXNwb25zZRJhCgppc3N1ZV9pbmZvGAEgAygLMkIuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUluZm9SCWlzc3VlSW5mbw==');
@$core.Deprecated('Use issueGetByIDRequestDescriptor instead')
const IssueGetByIDRequest$json = const {
  '1': 'IssueGetByIDRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `IssueGetByIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueGetByIDRequestDescriptor = $convert.base64Decode('ChNJc3N1ZUdldEJ5SURSZXF1ZXN0Eg4KAmlkGAEgASgDUgJpZA==');
@$core.Deprecated('Use issueGetByIDResponseDescriptor instead')
const IssueGetByIDResponse$json = const {
  '1': 'IssueGetByIDResponse',
  '2': const [
    const {'1': 'issue', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue', '10': 'issue'},
  ],
};

/// Descriptor for `IssueGetByIDResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueGetByIDResponseDescriptor = $convert.base64Decode('ChRJc3N1ZUdldEJ5SURSZXNwb25zZRJUCgVpc3N1ZRgBIAEoCzI+LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVSBWlzc3Vl');
@$core.Deprecated('Use issueInfoGetByIDRequestDescriptor instead')
const IssueInfoGetByIDRequest$json = const {
  '1': 'IssueInfoGetByIDRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `IssueInfoGetByIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueInfoGetByIDRequestDescriptor = $convert.base64Decode('ChdJc3N1ZUluZm9HZXRCeUlEUmVxdWVzdBIOCgJpZBgBIAEoA1ICaWQ=');
@$core.Deprecated('Use issueInfoGetByIDResponseDescriptor instead')
const IssueInfoGetByIDResponse$json = const {
  '1': 'IssueInfoGetByIDResponse',
  '2': const [
    const {'1': 'issue_info', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfo', '10': 'issueInfo'},
  ],
};

/// Descriptor for `IssueInfoGetByIDResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueInfoGetByIDResponseDescriptor = $convert.base64Decode('ChhJc3N1ZUluZm9HZXRCeUlEUmVzcG9uc2USYQoKaXNzdWVfaW5mbxgBIAEoCzJCLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVJbmZvUglpc3N1ZUluZm8=');
@$core.Deprecated('Use issuePaymentGetRequestDescriptor instead')
const IssuePaymentGetRequest$json = const {
  '1': 'IssuePaymentGetRequest',
};

/// Descriptor for `IssuePaymentGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issuePaymentGetRequestDescriptor = $convert.base64Decode('ChZJc3N1ZVBheW1lbnRHZXRSZXF1ZXN0');
@$core.Deprecated('Use issuePaymentGetResponseDescriptor instead')
const IssuePaymentGetResponse$json = const {
  '1': 'IssuePaymentGetResponse',
  '2': const [
    const {'1': 'user_to_payment', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssuePaymentGetResponse.UserToPaymentEntry', '10': 'userToPayment'},
  ],
  '3': const [IssuePaymentGetResponse_UserToPaymentEntry$json],
};

@$core.Deprecated('Use issuePaymentGetResponseDescriptor instead')
const IssuePaymentGetResponse_UserToPaymentEntry$json = const {
  '1': 'UserToPaymentEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 1, '10': 'value'},
  ],
  '7': const {'7': true},
};

/// Descriptor for `IssuePaymentGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issuePaymentGetResponseDescriptor = $convert.base64Decode('ChdJc3N1ZVBheW1lbnRHZXRSZXNwb25zZRKLAQoPdXNlcl90b19wYXltZW50GAEgAygLMmMuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZVBheW1lbnRHZXRSZXNwb25zZS5Vc2VyVG9QYXltZW50RW50cnlSDXVzZXJUb1BheW1lbnQaQAoSVXNlclRvUGF5bWVudEVudHJ5EhAKA2tleRgBIAEoCVIDa2V5EhQKBXZhbHVlGAIgASgBUgV2YWx1ZToCOAE=');
const $core.Map<$core.String, $core.dynamic> IssueRegistryServiceBase$json = const {
  '1': 'IssueRegistry',
  '2': const [
    const {'1': 'CreateIssue', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateResponse', '4': const {}},
    const {'1': 'CreateSubtask', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateSubtaskRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateSubtaskResponse', '4': const {}},
    const {'1': 'AddComment', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueAddCommentRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueAddCommentResponse', '4': const {}},
    const {'1': 'UpdateIssue', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueUpdateRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueUpdateResponse', '4': const {}},
    const {'1': 'GetIssues', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetResponse', '4': const {}},
    const {'1': 'GetIssueByID', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetByIDRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetByIDResponse', '4': const {}},
    const {'1': 'GetIssueInfo', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetResponse', '4': const {}},
    const {'1': 'GetIssueInfoByID', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetByIDRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetByIDResponse', '4': const {}},
    const {'1': 'GetUserPayment', '2': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssuePaymentGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssuePaymentGetResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use issueRegistryServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> IssueRegistryServiceBase$messageJson = const {
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateRequest': IssueCreateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.Issue': Issue$json,
  '.github.constantine27k.crnt_data_manager.api.comments.Comments': $3.Comments$json,
  '.github.constantine27k.crnt_data_manager.api.comments.Comment': $3.Comment$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
  '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatus': $4.IssueStatus$json,
  '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatusCommon': $4.IssueStatusCommon$json,
  '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatusDevelopment': $4.IssueStatusDevelopment$json,
  '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatusEpic': $4.IssueStatusEpic$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateResponse': IssueCreateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateSubtaskRequest': IssueCreateSubtaskRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueCreateSubtaskResponse': IssueCreateSubtaskResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueAddCommentRequest': IssueAddCommentRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueAddCommentResponse': IssueAddCommentResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueUpdateRequest': IssueUpdateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueUpdateResponse': IssueUpdateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetRequest': IssueGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetResponse': IssueGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetByIDRequest': IssueGetByIDRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueGetByIDResponse': IssueGetByIDResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetRequest': IssueInfoGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetResponse': IssueInfoGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfo': IssueInfo$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetByIDRequest': IssueInfoGetByIDRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssueInfoGetByIDResponse': IssueInfoGetByIDResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssuePaymentGetRequest': IssuePaymentGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssuePaymentGetResponse': IssuePaymentGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.tasks.issue.IssuePaymentGetResponse.UserToPaymentEntry': IssuePaymentGetResponse_UserToPaymentEntry$json,
};

/// Descriptor for `IssueRegistry`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List issueRegistryServiceDescriptor = $convert.base64Decode('Cg1Jc3N1ZVJlZ2lzdHJ5Er4BCgtDcmVhdGVJc3N1ZRJLLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVDcmVhdGVSZXF1ZXN0GkwuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUNyZWF0ZVJlc3BvbnNlIhSC0+STAg46ASoiCS92MS9pc3N1ZRLbAQoNQ3JlYXRlU3VidGFzaxJSLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVDcmVhdGVTdWJ0YXNrUmVxdWVzdBpTLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVDcmVhdGVTdWJ0YXNrUmVzcG9uc2UiIYLT5JMCGzoBKiIWL3YxL2lzc3VlL3tpZH0vc3VidGFzaxLSAQoKQWRkQ29tbWVudBJPLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVBZGRDb21tZW50UmVxdWVzdBpQLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVBZGRDb21tZW50UmVzcG9uc2UiIYLT5JMCGzoBKiIWL3YxL2lzc3VlL3tpZH0vY29tbWVudBLDAQoLVXBkYXRlSXNzdWUSSy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRhc2tzLmlzc3VlLklzc3VlVXBkYXRlUmVxdWVzdBpMLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVVcGRhdGVSZXNwb25zZSIZgtPkkwITOgEqGg4vdjEvaXNzdWUve2lkfRK0AQoJR2V0SXNzdWVzEkguZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUdldFJlcXVlc3QaSS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRhc2tzLmlzc3VlLklzc3VlR2V0UmVzcG9uc2UiEoLT5JMCDBIKL3YxL2lzc3VlcxLDAQoMR2V0SXNzdWVCeUlEEkwuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUdldEJ5SURSZXF1ZXN0Gk0uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUdldEJ5SURSZXNwb25zZSIWgtPkkwIQEg4vdjEvaXNzdWUve2lkfRLDAQoMR2V0SXNzdWVJbmZvEkwuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUluZm9HZXRSZXF1ZXN0Gk0uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50YXNrcy5pc3N1ZS5Jc3N1ZUluZm9HZXRSZXNwb25zZSIWgtPkkwIQEg4vdjEvaXNzdWVfaW5mbxLUAQoQR2V0SXNzdWVJbmZvQnlJRBJQLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVJbmZvR2V0QnlJRFJlcXVlc3QaUS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRhc2tzLmlzc3VlLklzc3VlSW5mb0dldEJ5SURSZXNwb25zZSIbgtPkkwIVEhMvdjEvaXNzdWVfaW5mby97aWR9Es8BCg5HZXRVc2VyUGF5bWVudBJPLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVQYXltZW50R2V0UmVxdWVzdBpQLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGFza3MuaXNzdWUuSXNzdWVQYXltZW50R2V0UmVzcG9uc2UiGoLT5JMCFBISL3YxL2lzc3Vlcy9wYXltZW50');
