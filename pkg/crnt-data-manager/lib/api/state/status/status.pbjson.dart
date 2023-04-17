///
//  Generated code. Do not modify.
//  source: api/state/status/status.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use commonDescriptor instead')
const Common$json = const {
  '1': 'Common',
  '2': const [
    const {'1': 'STATUS_COMMON_UNKNOWN', '2': 0},
    const {'1': 'STATUS_COMMON_BACKLOG', '2': 1},
    const {'1': 'STATUS_COMMON_IN_PROGRESS', '2': 2},
    const {'1': 'STATUS_COMMON_DONE', '2': 3},
    const {'1': 'STATUS_COMMON_READY_FOR_REVIEW', '2': 4},
    const {'1': 'STATUS_COMMON_IN_REVIEW', '2': 5},
    const {'1': 'STATUS_COMMON_REVIEW_FAILED', '2': 6},
    const {'1': 'STATUS_COMMON_REVIEW_PASSED', '2': 7},
    const {'1': 'STATUS_COMMON_GIVEN_TO_CUSTOMER', '2': 8},
    const {'1': 'STATUS_COMMON_CLOSED', '2': 9},
    const {'1': 'STATUS_COMMON_ON_HOLD', '2': 10},
  ],
};

/// Descriptor for `Common`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List commonDescriptor = $convert.base64Decode('CgZDb21tb24SGQoVU1RBVFVTX0NPTU1PTl9VTktOT1dOEAASGQoVU1RBVFVTX0NPTU1PTl9CQUNLTE9HEAESHQoZU1RBVFVTX0NPTU1PTl9JTl9QUk9HUkVTUxACEhYKElNUQVRVU19DT01NT05fRE9ORRADEiIKHlNUQVRVU19DT01NT05fUkVBRFlfRk9SX1JFVklFVxAEEhsKF1NUQVRVU19DT01NT05fSU5fUkVWSUVXEAUSHwobU1RBVFVTX0NPTU1PTl9SRVZJRVdfRkFJTEVEEAYSHwobU1RBVFVTX0NPTU1PTl9SRVZJRVdfUEFTU0VEEAcSIwofU1RBVFVTX0NPTU1PTl9HSVZFTl9UT19DVVNUT01FUhAIEhgKFFNUQVRVU19DT01NT05fQ0xPU0VEEAkSGQoVU1RBVFVTX0NPTU1PTl9PTl9IT0xEEAo=');
@$core.Deprecated('Use developmentDescriptor instead')
const Development$json = const {
  '1': 'Development',
  '2': const [
    const {'1': 'STATUS_DEVELOPMENT_UNKNOWN', '2': 0},
    const {'1': 'STATUS_DEVELOPMENT_BACKLOG', '2': 1},
    const {'1': 'STATUS_DEVELOPMENT_IN_PROGRESS', '2': 2},
    const {'1': 'STATUS_DEVELOPMENT_IN_REVIEW', '2': 3},
    const {'1': 'STATUS_DEVELOPMENT_READY_FOR_TESTING', '2': 4},
    const {'1': 'STATUS_DEVELOPMENT_TESTING', '2': 5},
    const {'1': 'STATUS_DEVELOPMENT_TESTING_PASSED', '2': 6},
    const {'1': 'STATUS_DEVELOPMENT_DONE', '2': 7},
    const {'1': 'STATUS_DEVELOPMENT_READY_TO_DEPLOY', '2': 8},
    const {'1': 'STATUS_DEVELOPMENT_CLOSED', '2': 9},
    const {'1': 'STATUS_DEVELOPMENT_ON_HOLD', '2': 10},
  ],
};

/// Descriptor for `Development`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List developmentDescriptor = $convert.base64Decode('CgtEZXZlbG9wbWVudBIeChpTVEFUVVNfREVWRUxPUE1FTlRfVU5LTk9XThAAEh4KGlNUQVRVU19ERVZFTE9QTUVOVF9CQUNLTE9HEAESIgoeU1RBVFVTX0RFVkVMT1BNRU5UX0lOX1BST0dSRVNTEAISIAocU1RBVFVTX0RFVkVMT1BNRU5UX0lOX1JFVklFVxADEigKJFNUQVRVU19ERVZFTE9QTUVOVF9SRUFEWV9GT1JfVEVTVElORxAEEh4KGlNUQVRVU19ERVZFTE9QTUVOVF9URVNUSU5HEAUSJQohU1RBVFVTX0RFVkVMT1BNRU5UX1RFU1RJTkdfUEFTU0VEEAYSGwoXU1RBVFVTX0RFVkVMT1BNRU5UX0RPTkUQBxImCiJTVEFUVVNfREVWRUxPUE1FTlRfUkVBRFlfVE9fREVQTE9ZEAgSHQoZU1RBVFVTX0RFVkVMT1BNRU5UX0NMT1NFRBAJEh4KGlNUQVRVU19ERVZFTE9QTUVOVF9PTl9IT0xEEAo=');
@$core.Deprecated('Use epicDescriptor instead')
const Epic$json = const {
  '1': 'Epic',
  '2': const [
    const {'1': 'STATUS_EPIC_UNKNOWN', '2': 0},
    const {'1': 'STATUS_EPIC_BACKLOG', '2': 1},
    const {'1': 'STATUS_EPIC_PLANNING', '2': 2},
    const {'1': 'STATUS_EPIC_DESIGNING', '2': 3},
    const {'1': 'STATUS_EPIC_IT', '2': 4},
    const {'1': 'STATUS_EPIC_DONE', '2': 5},
    const {'1': 'STATUS_EPIC_CLOSED', '2': 6},
    const {'1': 'STATUS_EPIC_ON_HOLD', '2': 7},
  ],
};

/// Descriptor for `Epic`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List epicDescriptor = $convert.base64Decode('CgRFcGljEhcKE1NUQVRVU19FUElDX1VOS05PV04QABIXChNTVEFUVVNfRVBJQ19CQUNLTE9HEAESGAoUU1RBVFVTX0VQSUNfUExBTk5JTkcQAhIZChVTVEFUVVNfRVBJQ19ERVNJR05JTkcQAxISCg5TVEFUVVNfRVBJQ19JVBAEEhQKEFNUQVRVU19FUElDX0RPTkUQBRIWChJTVEFUVVNfRVBJQ19DTE9TRUQQBhIXChNTVEFUVVNfRVBJQ19PTl9IT0xEEAc=');
@$core.Deprecated('Use sprintDescriptor instead')
const Sprint$json = const {
  '1': 'Sprint',
  '2': const [
    const {'1': 'STATUS_SPRINT_UNKNOWN', '2': 0},
    const {'1': 'STATUS_SPRINT_BACKLOG', '2': 1},
    const {'1': 'STATUS_SPRINT_ACTIVE', '2': 2},
    const {'1': 'STATUS_SPRINT_FINISHED', '2': 3},
  ],
};

/// Descriptor for `Sprint`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List sprintDescriptor = $convert.base64Decode('CgZTcHJpbnQSGQoVU1RBVFVTX1NQUklOVF9VTktOT1dOEAASGQoVU1RBVFVTX1NQUklOVF9CQUNLTE9HEAESGAoUU1RBVFVTX1NQUklOVF9BQ1RJVkUQAhIaChZTVEFUVVNfU1BSSU5UX0ZJTklTSEVEEAM=');
@$core.Deprecated('Use issueStatusDescriptor instead')
const IssueStatus$json = const {
  '1': 'IssueStatus',
  '2': const [
    const {'1': 'common', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatusCommon', '9': 0, '10': 'common'},
    const {'1': 'development', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatusDevelopment', '9': 0, '10': 'development'},
    const {'1': 'epic', '3': 3, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.state.status.IssueStatusEpic', '9': 0, '10': 'epic'},
  ],
  '8': const [
    const {'1': 'status'},
  ],
};

/// Descriptor for `IssueStatus`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueStatusDescriptor = $convert.base64Decode('CgtJc3N1ZVN0YXR1cxJlCgZjb21tb24YASABKAsySy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnN0YXR1cy5Jc3N1ZVN0YXR1c0NvbW1vbkgAUgZjb21tb24SdAoLZGV2ZWxvcG1lbnQYAiABKAsyUC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnN0YXR1cy5Jc3N1ZVN0YXR1c0RldmVsb3BtZW50SABSC2RldmVsb3BtZW50El8KBGVwaWMYAyABKAsySS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnN0YXR1cy5Jc3N1ZVN0YXR1c0VwaWNIAFIEZXBpY0IICgZzdGF0dXM=');
@$core.Deprecated('Use issueStatusCommonDescriptor instead')
const IssueStatusCommon$json = const {
  '1': 'IssueStatusCommon',
  '2': const [
    const {'1': 'status', '3': 1, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.status.Common', '10': 'status'},
  ],
};

/// Descriptor for `IssueStatusCommon`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueStatusCommonDescriptor = $convert.base64Decode('ChFJc3N1ZVN0YXR1c0NvbW1vbhJYCgZzdGF0dXMYASABKA4yQC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnN0YXRlLnN0YXR1cy5Db21tb25SBnN0YXR1cw==');
@$core.Deprecated('Use issueStatusDevelopmentDescriptor instead')
const IssueStatusDevelopment$json = const {
  '1': 'IssueStatusDevelopment',
  '2': const [
    const {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.status.Development', '10': 'status'},
  ],
};

/// Descriptor for `IssueStatusDevelopment`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueStatusDevelopmentDescriptor = $convert.base64Decode('ChZJc3N1ZVN0YXR1c0RldmVsb3BtZW50El0KBnN0YXR1cxgCIAEoDjJFLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuc3RhdGUuc3RhdHVzLkRldmVsb3BtZW50UgZzdGF0dXM=');
@$core.Deprecated('Use issueStatusEpicDescriptor instead')
const IssueStatusEpic$json = const {
  '1': 'IssueStatusEpic',
  '2': const [
    const {'1': 'status', '3': 2, '4': 1, '5': 14, '6': '.github.constantine27k.crnt_data_manager.api.state.status.Epic', '10': 'status'},
  ],
};

/// Descriptor for `IssueStatusEpic`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List issueStatusEpicDescriptor = $convert.base64Decode('Cg9Jc3N1ZVN0YXR1c0VwaWMSVgoGc3RhdHVzGAIgASgOMj4uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5zdGF0ZS5zdGF0dXMuRXBpY1IGc3RhdHVz');
