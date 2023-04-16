///
//  Generated code. Do not modify.
//  source: api/department/department.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use departmentDescriptor instead')
const Department$json = const {
  '1': 'Department',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'projects', '3': 3, '4': 3, '5': 3, '10': 'projects'},
    const {'1': 'members', '3': 4, '4': 3, '5': 9, '10': 'members'},
  ],
};

/// Descriptor for `Department`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentDescriptor = $convert.base64Decode('CgpEZXBhcnRtZW50Eg4KAmlkGAEgASgDUgJpZBISCgRuYW1lGAIgASgJUgRuYW1lEhoKCHByb2plY3RzGAMgAygDUghwcm9qZWN0cxIYCgdtZW1iZXJzGAQgAygJUgdtZW1iZXJz');
@$core.Deprecated('Use departmentCreateRequestDescriptor instead')
const DepartmentCreateRequest$json = const {
  '1': 'DepartmentCreateRequest',
  '2': const [
    const {'1': 'department', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.department.Department', '10': 'department'},
  ],
};

/// Descriptor for `DepartmentCreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentCreateRequestDescriptor = $convert.base64Decode('ChdEZXBhcnRtZW50Q3JlYXRlUmVxdWVzdBJiCgpkZXBhcnRtZW50GAEgASgLMkIuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5kZXBhcnRtZW50LkRlcGFydG1lbnRSCmRlcGFydG1lbnQ=');
@$core.Deprecated('Use departmentCreateResponseDescriptor instead')
const DepartmentCreateResponse$json = const {
  '1': 'DepartmentCreateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `DepartmentCreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentCreateResponseDescriptor = $convert.base64Decode('ChhEZXBhcnRtZW50Q3JlYXRlUmVzcG9uc2USDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use departmentAddProjectRequestDescriptor instead')
const DepartmentAddProjectRequest$json = const {
  '1': 'DepartmentAddProjectRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'project_id', '3': 2, '4': 1, '5': 3, '10': 'projectId'},
  ],
};

/// Descriptor for `DepartmentAddProjectRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentAddProjectRequestDescriptor = $convert.base64Decode('ChtEZXBhcnRtZW50QWRkUHJvamVjdFJlcXVlc3QSDgoCaWQYASABKANSAmlkEh0KCnByb2plY3RfaWQYAiABKANSCXByb2plY3RJZA==');
@$core.Deprecated('Use departmentAddProjectResponseDescriptor instead')
const DepartmentAddProjectResponse$json = const {
  '1': 'DepartmentAddProjectResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `DepartmentAddProjectResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentAddProjectResponseDescriptor = $convert.base64Decode('ChxEZXBhcnRtZW50QWRkUHJvamVjdFJlc3BvbnNlEg4KAmlkGAEgASgDUgJpZA==');
@$core.Deprecated('Use departmentRemoveProjectRequestDescriptor instead')
const DepartmentRemoveProjectRequest$json = const {
  '1': 'DepartmentRemoveProjectRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'project_id', '3': 2, '4': 1, '5': 3, '10': 'projectId'},
  ],
};

/// Descriptor for `DepartmentRemoveProjectRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentRemoveProjectRequestDescriptor = $convert.base64Decode('Ch5EZXBhcnRtZW50UmVtb3ZlUHJvamVjdFJlcXVlc3QSDgoCaWQYASABKANSAmlkEh0KCnByb2plY3RfaWQYAiABKANSCXByb2plY3RJZA==');
@$core.Deprecated('Use departmentRemoveProjectResponseDescriptor instead')
const DepartmentRemoveProjectResponse$json = const {
  '1': 'DepartmentRemoveProjectResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `DepartmentRemoveProjectResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentRemoveProjectResponseDescriptor = $convert.base64Decode('Ch9EZXBhcnRtZW50UmVtb3ZlUHJvamVjdFJlc3BvbnNlEg4KAmlkGAEgASgDUgJpZA==');
@$core.Deprecated('Use departmentGetRequestDescriptor instead')
const DepartmentGetRequest$json = const {
  '1': 'DepartmentGetRequest',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 3, '10': 'ids'},
    const {'1': 'names', '3': 2, '4': 3, '5': 9, '10': 'names'},
  ],
};

/// Descriptor for `DepartmentGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentGetRequestDescriptor = $convert.base64Decode('ChREZXBhcnRtZW50R2V0UmVxdWVzdBIQCgNpZHMYASADKANSA2lkcxIUCgVuYW1lcxgCIAMoCVIFbmFtZXM=');
@$core.Deprecated('Use departmentGetResponseDescriptor instead')
const DepartmentGetResponse$json = const {
  '1': 'DepartmentGetResponse',
  '2': const [
    const {'1': 'departments', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.department.Department', '10': 'departments'},
  ],
};

/// Descriptor for `DepartmentGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentGetResponseDescriptor = $convert.base64Decode('ChVEZXBhcnRtZW50R2V0UmVzcG9uc2USZAoLZGVwYXJ0bWVudHMYASADKAsyQi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudFILZGVwYXJ0bWVudHM=');
@$core.Deprecated('Use departmentGetByIDRequestDescriptor instead')
const DepartmentGetByIDRequest$json = const {
  '1': 'DepartmentGetByIDRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `DepartmentGetByIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentGetByIDRequestDescriptor = $convert.base64Decode('ChhEZXBhcnRtZW50R2V0QnlJRFJlcXVlc3QSDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use departmentGetByIDResponseDescriptor instead')
const DepartmentGetByIDResponse$json = const {
  '1': 'DepartmentGetByIDResponse',
  '2': const [
    const {'1': 'department', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.department.Department', '10': 'department'},
  ],
};

/// Descriptor for `DepartmentGetByIDResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentGetByIDResponseDescriptor = $convert.base64Decode('ChlEZXBhcnRtZW50R2V0QnlJRFJlc3BvbnNlEmIKCmRlcGFydG1lbnQYASABKAsyQi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudFIKZGVwYXJ0bWVudA==');
@$core.Deprecated('Use departmentUpdateRequestDescriptor instead')
const DepartmentUpdateRequest$json = const {
  '1': 'DepartmentUpdateRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'department', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.department.Department', '10': 'department'},
  ],
};

/// Descriptor for `DepartmentUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentUpdateRequestDescriptor = $convert.base64Decode('ChdEZXBhcnRtZW50VXBkYXRlUmVxdWVzdBIOCgJpZBgBIAEoA1ICaWQSYgoKZGVwYXJ0bWVudBgCIAEoCzJCLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuZGVwYXJ0bWVudC5EZXBhcnRtZW50UgpkZXBhcnRtZW50');
@$core.Deprecated('Use departmentUpdateResponseDescriptor instead')
const DepartmentUpdateResponse$json = const {
  '1': 'DepartmentUpdateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `DepartmentUpdateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List departmentUpdateResponseDescriptor = $convert.base64Decode('ChhEZXBhcnRtZW50VXBkYXRlUmVzcG9uc2USDgoCaWQYASABKANSAmlk');
const $core.Map<$core.String, $core.dynamic> DepartmentRegistryServiceBase$json = const {
  '1': 'DepartmentRegistry',
  '2': const [
    const {'1': 'CreateDepartment', '2': '.github.constantine27k.crnt_data_manager.api.department.DepartmentCreateRequest', '3': '.github.constantine27k.crnt_data_manager.api.department.DepartmentCreateResponse', '4': const {}},
    const {'1': 'DepartmentAddProject', '2': '.github.constantine27k.crnt_data_manager.api.department.DepartmentAddProjectRequest', '3': '.github.constantine27k.crnt_data_manager.api.department.DepartmentAddProjectResponse', '4': const {}},
    const {'1': 'DepartmentRemoveProject', '2': '.github.constantine27k.crnt_data_manager.api.department.DepartmentRemoveProjectRequest', '3': '.github.constantine27k.crnt_data_manager.api.department.DepartmentRemoveProjectResponse', '4': const {}},
    const {'1': 'GetDepartments', '2': '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetResponse', '4': const {}},
    const {'1': 'GetDepartmentByID', '2': '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetByIDRequest', '3': '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetByIDResponse', '4': const {}},
    const {'1': 'UpdateDepartment', '2': '.github.constantine27k.crnt_data_manager.api.department.DepartmentUpdateRequest', '3': '.github.constantine27k.crnt_data_manager.api.department.DepartmentUpdateResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use departmentRegistryServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> DepartmentRegistryServiceBase$messageJson = const {
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentCreateRequest': DepartmentCreateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.department.Department': Department$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentCreateResponse': DepartmentCreateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentAddProjectRequest': DepartmentAddProjectRequest$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentAddProjectResponse': DepartmentAddProjectResponse$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentRemoveProjectRequest': DepartmentRemoveProjectRequest$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentRemoveProjectResponse': DepartmentRemoveProjectResponse$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetRequest': DepartmentGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetResponse': DepartmentGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetByIDRequest': DepartmentGetByIDRequest$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentGetByIDResponse': DepartmentGetByIDResponse$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentUpdateRequest': DepartmentUpdateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.department.DepartmentUpdateResponse': DepartmentUpdateResponse$json,
};

/// Descriptor for `DepartmentRegistry`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List departmentRegistryServiceDescriptor = $convert.base64Decode('ChJEZXBhcnRtZW50UmVnaXN0cnkS0AEKEENyZWF0ZURlcGFydG1lbnQSTy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudENyZWF0ZVJlcXVlc3QaUC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudENyZWF0ZVJlc3BvbnNlIhmC0+STAhM6ASoiDi92MS9kZXBhcnRtZW50EvMBChREZXBhcnRtZW50QWRkUHJvamVjdBJTLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuZGVwYXJ0bWVudC5EZXBhcnRtZW50QWRkUHJvamVjdFJlcXVlc3QaVC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudEFkZFByb2plY3RSZXNwb25zZSIwgtPkkwIqIigvdjEvZGVwYXJ0bWVudC97aWR9L3Byb2plY3Qve3Byb2plY3RfaWR9EvwBChdEZXBhcnRtZW50UmVtb3ZlUHJvamVjdBJWLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuZGVwYXJ0bWVudC5EZXBhcnRtZW50UmVtb3ZlUHJvamVjdFJlcXVlc3QaVy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudFJlbW92ZVByb2plY3RSZXNwb25zZSIwgtPkkwIqKigvdjEvZGVwYXJ0bWVudC97aWR9L3Byb2plY3Qve3Byb2plY3RfaWR9EsYBCg5HZXREZXBhcnRtZW50cxJMLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuZGVwYXJ0bWVudC5EZXBhcnRtZW50R2V0UmVxdWVzdBpNLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuZGVwYXJ0bWVudC5EZXBhcnRtZW50R2V0UmVzcG9uc2UiF4LT5JMCERIPL3YxL2RlcGFydG1lbnRzEtUBChFHZXREZXBhcnRtZW50QnlJRBJQLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuZGVwYXJ0bWVudC5EZXBhcnRtZW50R2V0QnlJRFJlcXVlc3QaUS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLmRlcGFydG1lbnQuRGVwYXJ0bWVudEdldEJ5SURSZXNwb25zZSIbgtPkkwIVEhMvdjEvZGVwYXJ0bWVudC97aWR9EtUBChBVcGRhdGVEZXBhcnRtZW50Ek8uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5kZXBhcnRtZW50LkRlcGFydG1lbnRVcGRhdGVSZXF1ZXN0GlAuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5kZXBhcnRtZW50LkRlcGFydG1lbnRVcGRhdGVSZXNwb25zZSIegtPkkwIYOgEqGhMvdjEvZGVwYXJ0bWVudC97aWR9');
