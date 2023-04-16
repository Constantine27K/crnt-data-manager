///
//  Generated code. Do not modify.
//  source: api/sprint/sprint.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
import '../../google/protobuf/timestamp.pbjson.dart' as $0;

@$core.Deprecated('Use sprintDescriptor instead')
const Sprint$json = const {
  '1': 'Sprint',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'project', '3': 3, '4': 1, '5': 3, '10': 'project'},
    const {'1': 'started_at', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'startedAt'},
    const {'1': 'finished_at', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'finishedAt'},
    const {'1': 'status', '3': 6, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.status.Sprint', '10': 'status'},
    const {'1': 'stories_backlog', '3': 7, '4': 1, '5': 3, '10': 'storiesBacklog'},
    const {'1': 'stories_in_progress', '3': 8, '4': 1, '5': 3, '10': 'storiesInProgress'},
    const {'1': 'stories_done', '3': 9, '4': 1, '5': 3, '10': 'storiesDone'},
    const {'1': 'issues', '3': 10, '4': 3, '5': 3, '10': 'issues'},
  ],
};

/// Descriptor for `Sprint`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintDescriptor = $convert.base64Decode('CgZTcHJpbnQSDgoCaWQYASABKANSAmlkEhIKBG5hbWUYAiABKAlSBG5hbWUSGAoHcHJvamVjdBgDIAEoA1IHcHJvamVjdBI5CgpzdGFydGVkX2F0GAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIJc3RhcnRlZEF0EjsKC2ZpbmlzaGVkX2F0GAUgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcFIKZmluaXNoZWRBdBJYCgZzdGF0dXMYBiABKA4yQC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnN0YXR1cy5TcHJpbnRSBnN0YXR1cxInCg9zdG9yaWVzX2JhY2tsb2cYByABKANSDnN0b3JpZXNCYWNrbG9nEi4KE3N0b3JpZXNfaW5fcHJvZ3Jlc3MYCCABKANSEXN0b3JpZXNJblByb2dyZXNzEiEKDHN0b3JpZXNfZG9uZRgJIAEoA1ILc3Rvcmllc0RvbmUSFgoGaXNzdWVzGAogAygDUgZpc3N1ZXM=');
@$core.Deprecated('Use sprintCreateRequestDescriptor instead')
const SprintCreateRequest$json = const {
  '1': 'SprintCreateRequest',
  '2': const [
    const {'1': 'sprint', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.sprint.Sprint', '10': 'sprint'},
  ],
};

/// Descriptor for `SprintCreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintCreateRequestDescriptor = $convert.base64Decode('ChNTcHJpbnRDcmVhdGVSZXF1ZXN0ElIKBnNwcmludBgBIAEoCzI6LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3ByaW50LlNwcmludFIGc3ByaW50');
@$core.Deprecated('Use sprintCreateResponseDescriptor instead')
const SprintCreateResponse$json = const {
  '1': 'SprintCreateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `SprintCreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintCreateResponseDescriptor = $convert.base64Decode('ChRTcHJpbnRDcmVhdGVSZXNwb25zZRIOCgJpZBgBIAEoA1ICaWQ=');
@$core.Deprecated('Use addIssueRequestDescriptor instead')
const AddIssueRequest$json = const {
  '1': 'AddIssueRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'issue_id', '3': 2, '4': 1, '5': 3, '10': 'issueId'},
  ],
};

/// Descriptor for `AddIssueRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addIssueRequestDescriptor = $convert.base64Decode('Cg9BZGRJc3N1ZVJlcXVlc3QSDgoCaWQYASABKANSAmlkEhkKCGlzc3VlX2lkGAIgASgDUgdpc3N1ZUlk');
@$core.Deprecated('Use addIssueResponseDescriptor instead')
const AddIssueResponse$json = const {
  '1': 'AddIssueResponse',
  '2': const [
    const {'1': 'sprint_id', '3': 1, '4': 1, '5': 3, '10': 'sprintId'},
  ],
};

/// Descriptor for `AddIssueResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List addIssueResponseDescriptor = $convert.base64Decode('ChBBZGRJc3N1ZVJlc3BvbnNlEhsKCXNwcmludF9pZBgBIAEoA1IIc3ByaW50SWQ=');
@$core.Deprecated('Use removeIssueRequestDescriptor instead')
const RemoveIssueRequest$json = const {
  '1': 'RemoveIssueRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'issue_id', '3': 2, '4': 1, '5': 3, '10': 'issueId'},
  ],
};

/// Descriptor for `RemoveIssueRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List removeIssueRequestDescriptor = $convert.base64Decode('ChJSZW1vdmVJc3N1ZVJlcXVlc3QSDgoCaWQYASABKANSAmlkEhkKCGlzc3VlX2lkGAIgASgDUgdpc3N1ZUlk');
@$core.Deprecated('Use removeIssueResponseDescriptor instead')
const RemoveIssueResponse$json = const {
  '1': 'RemoveIssueResponse',
  '2': const [
    const {'1': 'sprint_id', '3': 1, '4': 1, '5': 3, '10': 'sprintId'},
  ],
};

/// Descriptor for `RemoveIssueResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List removeIssueResponseDescriptor = $convert.base64Decode('ChNSZW1vdmVJc3N1ZVJlc3BvbnNlEhsKCXNwcmludF9pZBgBIAEoA1IIc3ByaW50SWQ=');
@$core.Deprecated('Use sprintUpdateRequestDescriptor instead')
const SprintUpdateRequest$json = const {
  '1': 'SprintUpdateRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'sprint', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.sprint.Sprint', '10': 'sprint'},
  ],
};

/// Descriptor for `SprintUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintUpdateRequestDescriptor = $convert.base64Decode('ChNTcHJpbnRVcGRhdGVSZXF1ZXN0Eg4KAmlkGAEgASgDUgJpZBJSCgZzcHJpbnQYAiABKAsyOi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnNwcmludC5TcHJpbnRSBnNwcmludA==');
@$core.Deprecated('Use sprintUpdateResponseDescriptor instead')
const SprintUpdateResponse$json = const {
  '1': 'SprintUpdateResponse',
  '2': const [
    const {'1': 'id', '3': 2, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `SprintUpdateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintUpdateResponseDescriptor = $convert.base64Decode('ChRTcHJpbnRVcGRhdGVSZXNwb25zZRIOCgJpZBgCIAEoA1ICaWQ=');
@$core.Deprecated('Use sprintChangeStatusRequestDescriptor instead')
const SprintChangeStatusRequest$json = const {
  '1': 'SprintChangeStatusRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.status.Sprint', '10': 'status'},
  ],
};

/// Descriptor for `SprintChangeStatusRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintChangeStatusRequestDescriptor = $convert.base64Decode('ChlTcHJpbnRDaGFuZ2VTdGF0dXNSZXF1ZXN0Eg4KAmlkGAEgASgDUgJpZBJYCgZzdGF0dXMYAiABKA4yQC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnN0YXR1cy5TcHJpbnRSBnN0YXR1cw==');
@$core.Deprecated('Use sprintChangeStatusResponseDescriptor instead')
const SprintChangeStatusResponse$json = const {
  '1': 'SprintChangeStatusResponse',
  '2': const [
    const {'1': 'id', '3': 2, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `SprintChangeStatusResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintChangeStatusResponseDescriptor = $convert.base64Decode('ChpTcHJpbnRDaGFuZ2VTdGF0dXNSZXNwb25zZRIOCgJpZBgCIAEoA1ICaWQ=');
@$core.Deprecated('Use sprintGetRequestDescriptor instead')
const SprintGetRequest$json = const {
  '1': 'SprintGetRequest',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 3, '10': 'ids'},
    const {'1': 'names', '3': 2, '4': 3, '5': 9, '10': 'names'},
    const {'1': 'projects', '3': 3, '4': 3, '5': 3, '10': 'projects'},
    const {'1': 'started_at_after', '3': 4, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'startedAtAfter'},
    const {'1': 'finished_at_before', '3': 5, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'finishedAtBefore'},
    const {'1': 'status', '3': 6, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.status.Sprint', '10': 'status'},
  ],
};

/// Descriptor for `SprintGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintGetRequestDescriptor = $convert.base64Decode('ChBTcHJpbnRHZXRSZXF1ZXN0EhAKA2lkcxgBIAMoA1IDaWRzEhQKBW5hbWVzGAIgAygJUgVuYW1lcxIaCghwcm9qZWN0cxgDIAMoA1IIcHJvamVjdHMSRAoQc3RhcnRlZF9hdF9hZnRlchgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSDnN0YXJ0ZWRBdEFmdGVyEkgKEmZpbmlzaGVkX2F0X2JlZm9yZRgFIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSEGZpbmlzaGVkQXRCZWZvcmUSWAoGc3RhdHVzGAYgASgOMkAuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zdGF0ZS5zdGF0dXMuU3ByaW50UgZzdGF0dXM=');
@$core.Deprecated('Use sprintGetResponseDescriptor instead')
const SprintGetResponse$json = const {
  '1': 'SprintGetResponse',
  '2': const [
    const {'1': 'sprints', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.sprint.Sprint', '10': 'sprints'},
  ],
};

/// Descriptor for `SprintGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintGetResponseDescriptor = $convert.base64Decode('ChFTcHJpbnRHZXRSZXNwb25zZRJUCgdzcHJpbnRzGAEgAygLMjouZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50UgdzcHJpbnRz');
@$core.Deprecated('Use sprintGetByIDRequestDescriptor instead')
const SprintGetByIDRequest$json = const {
  '1': 'SprintGetByIDRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `SprintGetByIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintGetByIDRequestDescriptor = $convert.base64Decode('ChRTcHJpbnRHZXRCeUlEUmVxdWVzdBIOCgJpZBgBIAEoA1ICaWQ=');
@$core.Deprecated('Use sprintGetByIDResponseDescriptor instead')
const SprintGetByIDResponse$json = const {
  '1': 'SprintGetByIDResponse',
  '2': const [
    const {'1': 'sprint', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.sprint.Sprint', '10': 'sprint'},
  ],
};

/// Descriptor for `SprintGetByIDResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List sprintGetByIDResponseDescriptor = $convert.base64Decode('ChVTcHJpbnRHZXRCeUlEUmVzcG9uc2USUgoGc3ByaW50GAEgASgLMjouZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50UgZzcHJpbnQ=');
const $core.Map<$core.String, $core.dynamic> SprintRegistryServiceBase$json = const {
  '1': 'SprintRegistry',
  '2': const [
    const {'1': 'CreateSprint', '2': '.github.constantine27k.crnt_data_manager.api.sprint.SprintCreateRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.SprintCreateResponse', '4': const {}},
    const {'1': 'AddIssue', '2': '.github.constantine27k.crnt_data_manager.api.sprint.AddIssueRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.AddIssueResponse', '4': const {}},
    const {'1': 'RemoveIssue', '2': '.github.constantine27k.crnt_data_manager.api.sprint.RemoveIssueRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.RemoveIssueResponse', '4': const {}},
    const {'1': 'UpdateSprint', '2': '.github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateResponse', '4': const {}},
    const {'1': 'SprintChangeStatus', '2': '.github.constantine27k.crnt_data_manager.api.sprint.SprintChangeStatusRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.SprintChangeStatusResponse', '4': const {}},
    const {'1': 'GetSprint', '2': '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetResponse', '4': const {}},
    const {'1': 'GetSprintByID', '2': '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDRequest', '3': '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use sprintRegistryServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> SprintRegistryServiceBase$messageJson = const {
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintCreateRequest': SprintCreateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.Sprint': Sprint$json,
  '.google.protobuf.Timestamp': $0.Timestamp$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintCreateResponse': SprintCreateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.AddIssueRequest': AddIssueRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.AddIssueResponse': AddIssueResponse$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.RemoveIssueRequest': RemoveIssueRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.RemoveIssueResponse': RemoveIssueResponse$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateRequest': SprintUpdateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintUpdateResponse': SprintUpdateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintChangeStatusRequest': SprintChangeStatusRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintChangeStatusResponse': SprintChangeStatusResponse$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetRequest': SprintGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetResponse': SprintGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDRequest': SprintGetByIDRequest$json,
  '.github.constantine27k.crnt_data_manager.api.sprint.SprintGetByIDResponse': SprintGetByIDResponse$json,
};

/// Descriptor for `SprintRegistry`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List sprintRegistryServiceDescriptor = $convert.base64Decode('Cg5TcHJpbnRSZWdpc3RyeRK4AQoMQ3JlYXRlU3ByaW50EkcuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50Q3JlYXRlUmVxdWVzdBpILmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3ByaW50LlNwcmludENyZWF0ZVJlc3BvbnNlIhWC0+STAg86ASoiCi92MS9zcHJpbnQSvwEKCEFkZElzc3VlEkMuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuQWRkSXNzdWVSZXF1ZXN0GkQuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuQWRkSXNzdWVSZXNwb25zZSIogtPkkwIiIiAvdjEvc3ByaW50L3tpZH0vaXNzdWUve2lzc3VlX2lkfRLIAQoLUmVtb3ZlSXNzdWUSRi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnNwcmludC5SZW1vdmVJc3N1ZVJlcXVlc3QaRy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnNwcmludC5SZW1vdmVJc3N1ZVJlc3BvbnNlIiiC0+STAiIqIC92MS9zcHJpbnQve2lkfS9pc3N1ZS97aXNzdWVfaWR9Er0BCgxVcGRhdGVTcHJpbnQSRy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnNwcmludC5TcHJpbnRVcGRhdGVSZXF1ZXN0GkguZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50VXBkYXRlUmVzcG9uc2UiGoLT5JMCFDoBKhoPL3YxL3NwcmludC97aWR9EtYBChJTcHJpbnRDaGFuZ2VTdGF0dXMSTS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnNwcmludC5TcHJpbnRDaGFuZ2VTdGF0dXNSZXF1ZXN0Gk4uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50Q2hhbmdlU3RhdHVzUmVzcG9uc2UiIYLT5JMCGzoBKiIWL3YxL3NwcmludC97aWR9L3N0YXR1cxKtAQoJR2V0U3ByaW50EkQuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50R2V0UmVxdWVzdBpFLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3ByaW50LlNwcmludEdldFJlc3BvbnNlIhOC0+STAg0SCy92MS9zcHJpbnRzEr0BCg1HZXRTcHJpbnRCeUlEEkguZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zcHJpbnQuU3ByaW50R2V0QnlJRFJlcXVlc3QaSS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnNwcmludC5TcHJpbnRHZXRCeUlEUmVzcG9uc2UiF4LT5JMCERIPL3YxL3NwcmludC97aWR9');
