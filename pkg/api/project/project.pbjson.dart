///
//  Generated code. Do not modify.
//  source: api/project/project.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use projectDescriptor instead')
const Project$json = const {
  '1': 'Project',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'short_name', '3': 3, '4': 1, '5': 9, '10': 'shortName'},
    const {'1': 'is_archived', '3': 4, '4': 1, '5': 8, '10': 'isArchived'},
    const {'1': 'responsible_team_ids', '3': 5, '4': 3, '5': 3, '10': 'responsibleTeamIds'},
    const {'1': 'description', '3': 6, '4': 1, '5': 9, '10': 'description'},
    const {'1': 'department', '3': 7, '4': 1, '5': 3, '10': 'department'},
    const {'1': 'responsible', '3': 8, '4': 1, '5': 9, '10': 'responsible'},
  ],
};

/// Descriptor for `Project`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectDescriptor = $convert.base64Decode('CgdQcm9qZWN0Eg4KAmlkGAEgASgDUgJpZBISCgRuYW1lGAIgASgJUgRuYW1lEh0KCnNob3J0X25hbWUYAyABKAlSCXNob3J0TmFtZRIfCgtpc19hcmNoaXZlZBgEIAEoCFIKaXNBcmNoaXZlZBIwChRyZXNwb25zaWJsZV90ZWFtX2lkcxgFIAMoA1IScmVzcG9uc2libGVUZWFtSWRzEiAKC2Rlc2NyaXB0aW9uGAYgASgJUgtkZXNjcmlwdGlvbhIeCgpkZXBhcnRtZW50GAcgASgDUgpkZXBhcnRtZW50EiAKC3Jlc3BvbnNpYmxlGAggASgJUgtyZXNwb25zaWJsZQ==');
@$core.Deprecated('Use projectCreateRequestDescriptor instead')
const ProjectCreateRequest$json = const {
  '1': 'ProjectCreateRequest',
  '2': const [
    const {'1': 'project', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.project.Project', '10': 'project'},
  ],
};

/// Descriptor for `ProjectCreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectCreateRequestDescriptor = $convert.base64Decode('ChRQcm9qZWN0Q3JlYXRlUmVxdWVzdBJWCgdwcm9qZWN0GAEgASgLMjwuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5wcm9qZWN0LlByb2plY3RSB3Byb2plY3Q=');
@$core.Deprecated('Use projectCreateResponseDescriptor instead')
const ProjectCreateResponse$json = const {
  '1': 'ProjectCreateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `ProjectCreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectCreateResponseDescriptor = $convert.base64Decode('ChVQcm9qZWN0Q3JlYXRlUmVzcG9uc2USDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use projectAddTeamRequestDescriptor instead')
const ProjectAddTeamRequest$json = const {
  '1': 'ProjectAddTeamRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'team_id', '3': 2, '4': 1, '5': 3, '10': 'teamId'},
  ],
};

/// Descriptor for `ProjectAddTeamRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectAddTeamRequestDescriptor = $convert.base64Decode('ChVQcm9qZWN0QWRkVGVhbVJlcXVlc3QSDgoCaWQYASABKANSAmlkEhcKB3RlYW1faWQYAiABKANSBnRlYW1JZA==');
@$core.Deprecated('Use projectAddTeamResponseDescriptor instead')
const ProjectAddTeamResponse$json = const {
  '1': 'ProjectAddTeamResponse',
  '2': const [
    const {'1': 'project_id', '3': 1, '4': 1, '5': 3, '10': 'projectId'},
  ],
};

/// Descriptor for `ProjectAddTeamResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectAddTeamResponseDescriptor = $convert.base64Decode('ChZQcm9qZWN0QWRkVGVhbVJlc3BvbnNlEh0KCnByb2plY3RfaWQYASABKANSCXByb2plY3RJZA==');
@$core.Deprecated('Use projectRemoveTeamRequestDescriptor instead')
const ProjectRemoveTeamRequest$json = const {
  '1': 'ProjectRemoveTeamRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'team_id', '3': 2, '4': 1, '5': 3, '10': 'teamId'},
  ],
};

/// Descriptor for `ProjectRemoveTeamRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectRemoveTeamRequestDescriptor = $convert.base64Decode('ChhQcm9qZWN0UmVtb3ZlVGVhbVJlcXVlc3QSDgoCaWQYASABKANSAmlkEhcKB3RlYW1faWQYAiABKANSBnRlYW1JZA==');
@$core.Deprecated('Use projectRemoveTeamResponseDescriptor instead')
const ProjectRemoveTeamResponse$json = const {
  '1': 'ProjectRemoveTeamResponse',
  '2': const [
    const {'1': 'project_id', '3': 1, '4': 1, '5': 3, '10': 'projectId'},
  ],
};

/// Descriptor for `ProjectRemoveTeamResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectRemoveTeamResponseDescriptor = $convert.base64Decode('ChlQcm9qZWN0UmVtb3ZlVGVhbVJlc3BvbnNlEh0KCnByb2plY3RfaWQYASABKANSCXByb2plY3RJZA==');
@$core.Deprecated('Use projectUpdateRequestDescriptor instead')
const ProjectUpdateRequest$json = const {
  '1': 'ProjectUpdateRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'project', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.project.Project', '10': 'project'},
  ],
};

/// Descriptor for `ProjectUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectUpdateRequestDescriptor = $convert.base64Decode('ChRQcm9qZWN0VXBkYXRlUmVxdWVzdBIOCgJpZBgBIAEoA1ICaWQSVgoHcHJvamVjdBgCIAEoCzI8LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkucHJvamVjdC5Qcm9qZWN0Ugdwcm9qZWN0');
@$core.Deprecated('Use projectUpdateResponseDescriptor instead')
const ProjectUpdateResponse$json = const {
  '1': 'ProjectUpdateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `ProjectUpdateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectUpdateResponseDescriptor = $convert.base64Decode('ChVQcm9qZWN0VXBkYXRlUmVzcG9uc2USDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use projectGetRequestDescriptor instead')
const ProjectGetRequest$json = const {
  '1': 'ProjectGetRequest',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 3, '10': 'ids'},
    const {'1': 'names', '3': 2, '4': 3, '5': 9, '10': 'names'},
    const {'1': 'short_names', '3': 3, '4': 3, '5': 9, '10': 'shortNames'},
    const {'1': 'is_archived', '3': 4, '4': 1, '5': 8, '10': 'isArchived'},
    const {'1': 'responsible_team_ids', '3': 5, '4': 3, '5': 3, '10': 'responsibleTeamIds'},
  ],
};

/// Descriptor for `ProjectGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectGetRequestDescriptor = $convert.base64Decode('ChFQcm9qZWN0R2V0UmVxdWVzdBIQCgNpZHMYASADKANSA2lkcxIUCgVuYW1lcxgCIAMoCVIFbmFtZXMSHwoLc2hvcnRfbmFtZXMYAyADKAlSCnNob3J0TmFtZXMSHwoLaXNfYXJjaGl2ZWQYBCABKAhSCmlzQXJjaGl2ZWQSMAoUcmVzcG9uc2libGVfdGVhbV9pZHMYBSADKANSEnJlc3BvbnNpYmxlVGVhbUlkcw==');
@$core.Deprecated('Use projectGetResponseDescriptor instead')
const ProjectGetResponse$json = const {
  '1': 'ProjectGetResponse',
  '2': const [
    const {'1': 'projects', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.project.Project', '10': 'projects'},
  ],
};

/// Descriptor for `ProjectGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectGetResponseDescriptor = $convert.base64Decode('ChJQcm9qZWN0R2V0UmVzcG9uc2USWAoIcHJvamVjdHMYASADKAsyPC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdFIIcHJvamVjdHM=');
@$core.Deprecated('Use projectGetByIDRequestDescriptor instead')
const ProjectGetByIDRequest$json = const {
  '1': 'ProjectGetByIDRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `ProjectGetByIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectGetByIDRequestDescriptor = $convert.base64Decode('ChVQcm9qZWN0R2V0QnlJRFJlcXVlc3QSDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use projectGetByIDResponseDescriptor instead')
const ProjectGetByIDResponse$json = const {
  '1': 'ProjectGetByIDResponse',
  '2': const [
    const {'1': 'project', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.project.Project', '10': 'project'},
  ],
};

/// Descriptor for `ProjectGetByIDResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List projectGetByIDResponseDescriptor = $convert.base64Decode('ChZQcm9qZWN0R2V0QnlJRFJlc3BvbnNlElYKB3Byb2plY3QYASABKAsyPC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdFIHcHJvamVjdA==');
const $core.Map<$core.String, $core.dynamic> ProjectRegistryServiceBase$json = const {
  '1': 'ProjectRegistry',
  '2': const [
    const {'1': 'CreateProject', '2': '.github.constantine27k.crnt_data_manager.api.project.ProjectCreateRequest', '3': '.github.constantine27k.crnt_data_manager.api.project.ProjectCreateResponse', '4': const {}},
    const {'1': 'AddResponsibleTeam', '2': '.github.constantine27k.crnt_data_manager.api.project.ProjectAddTeamRequest', '3': '.github.constantine27k.crnt_data_manager.api.project.ProjectAddTeamResponse', '4': const {}},
    const {'1': 'RemoveResponsibleTeam', '2': '.github.constantine27k.crnt_data_manager.api.project.ProjectRemoveTeamRequest', '3': '.github.constantine27k.crnt_data_manager.api.project.ProjectRemoveTeamResponse', '4': const {}},
    const {'1': 'UpdateProject', '2': '.github.constantine27k.crnt_data_manager.api.project.ProjectUpdateRequest', '3': '.github.constantine27k.crnt_data_manager.api.project.ProjectUpdateResponse', '4': const {}},
    const {'1': 'GetProjects', '2': '.github.constantine27k.crnt_data_manager.api.project.ProjectGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.project.ProjectGetResponse', '4': const {}},
    const {'1': 'GetProjectByID', '2': '.github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDRequest', '3': '.github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use projectRegistryServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> ProjectRegistryServiceBase$messageJson = const {
  '.github.constantine27k.crnt_data_manager.api.project.ProjectCreateRequest': ProjectCreateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.project.Project': Project$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectCreateResponse': ProjectCreateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectAddTeamRequest': ProjectAddTeamRequest$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectAddTeamResponse': ProjectAddTeamResponse$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectRemoveTeamRequest': ProjectRemoveTeamRequest$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectRemoveTeamResponse': ProjectRemoveTeamResponse$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectUpdateRequest': ProjectUpdateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectUpdateResponse': ProjectUpdateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectGetRequest': ProjectGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectGetResponse': ProjectGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDRequest': ProjectGetByIDRequest$json,
  '.github.constantine27k.crnt_data_manager.api.project.ProjectGetByIDResponse': ProjectGetByIDResponse$json,
};

/// Descriptor for `ProjectRegistry`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List projectRegistryServiceDescriptor = $convert.base64Decode('Cg9Qcm9qZWN0UmVnaXN0cnkSvgEKDUNyZWF0ZVByb2plY3QSSS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdENyZWF0ZVJlcXVlc3QaSi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdENyZWF0ZVJlc3BvbnNlIhaC0+STAhA6ASoiCy92MS9wcm9qZWN0EtYBChJBZGRSZXNwb25zaWJsZVRlYW0SSi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdEFkZFRlYW1SZXF1ZXN0GksuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5wcm9qZWN0LlByb2plY3RBZGRUZWFtUmVzcG9uc2UiJ4LT5JMCISIfL3YxL3Byb2plY3Qve2lkfS90ZWFtL3t0ZWFtX2lkfRLfAQoVUmVtb3ZlUmVzcG9uc2libGVUZWFtEk0uZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5wcm9qZWN0LlByb2plY3RSZW1vdmVUZWFtUmVxdWVzdBpOLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkucHJvamVjdC5Qcm9qZWN0UmVtb3ZlVGVhbVJlc3BvbnNlIieC0+STAiEqHy92MS9wcm9qZWN0L3tpZH0vdGVhbS97dGVhbV9pZH0SwwEKDVVwZGF0ZVByb2plY3QSSS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdFVwZGF0ZVJlcXVlc3QaSi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnByb2plY3QuUHJvamVjdFVwZGF0ZVJlc3BvbnNlIhuC0+STAhU6ASoaEC92MS9wcm9qZWN0L3tpZH0StAEKC0dldFByb2plY3RzEkYuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5wcm9qZWN0LlByb2plY3RHZXRSZXF1ZXN0GkcuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5wcm9qZWN0LlByb2plY3RHZXRSZXNwb25zZSIUgtPkkwIOEgwvdjEvcHJvamVjdHMSwwEKDkdldFByb2plY3RCeUlEEkouZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS5wcm9qZWN0LlByb2plY3RHZXRCeUlEUmVxdWVzdBpLLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkucHJvamVjdC5Qcm9qZWN0R2V0QnlJRFJlc3BvbnNlIhiC0+STAhISEC92MS9wcm9qZWN0L3tpZH0=');
