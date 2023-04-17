///
//  Generated code. Do not modify.
//  source: api/team/team.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use teamDescriptor instead')
const Team$json = const {
  '1': 'Team',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'members', '3': 3, '4': 3, '5': 9, '10': 'members'},
    const {'1': 'projects', '3': 4, '4': 3, '5': 3, '10': 'projects'},
    const {'1': 'tech_owner', '3': 5, '4': 1, '5': 9, '10': 'techOwner'},
    const {'1': 'business_owner', '3': 6, '4': 1, '5': 9, '10': 'businessOwner'},
    const {'1': 'department', '3': 7, '4': 1, '5': 9, '10': 'department'},
  ],
};

/// Descriptor for `Team`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamDescriptor = $convert.base64Decode('CgRUZWFtEg4KAmlkGAEgASgDUgJpZBISCgRuYW1lGAIgASgJUgRuYW1lEhgKB21lbWJlcnMYAyADKAlSB21lbWJlcnMSGgoIcHJvamVjdHMYBCADKANSCHByb2plY3RzEh0KCnRlY2hfb3duZXIYBSABKAlSCXRlY2hPd25lchIlCg5idXNpbmVzc19vd25lchgGIAEoCVINYnVzaW5lc3NPd25lchIeCgpkZXBhcnRtZW50GAcgASgJUgpkZXBhcnRtZW50');
@$core.Deprecated('Use teamCreateRequestDescriptor instead')
const TeamCreateRequest$json = const {
  '1': 'TeamCreateRequest',
  '2': const [
    const {'1': 'team', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.team.Team', '10': 'team'},
  ],
};

/// Descriptor for `TeamCreateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamCreateRequestDescriptor = $convert.base64Decode('ChFUZWFtQ3JlYXRlUmVxdWVzdBJKCgR0ZWFtGAEgASgLMjYuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50ZWFtLlRlYW1SBHRlYW0=');
@$core.Deprecated('Use teamCreateResponseDescriptor instead')
const TeamCreateResponse$json = const {
  '1': 'TeamCreateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `TeamCreateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamCreateResponseDescriptor = $convert.base64Decode('ChJUZWFtQ3JlYXRlUmVzcG9uc2USDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use teamUpdateRequestDescriptor instead')
const TeamUpdateRequest$json = const {
  '1': 'TeamUpdateRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
    const {'1': 'team', '3': 2, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.team.Team', '10': 'team'},
  ],
};

/// Descriptor for `TeamUpdateRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamUpdateRequestDescriptor = $convert.base64Decode('ChFUZWFtVXBkYXRlUmVxdWVzdBIOCgJpZBgBIAEoA1ICaWQSSgoEdGVhbRgCIAEoCzI2LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGVhbS5UZWFtUgR0ZWFt');
@$core.Deprecated('Use teamUpdateResponseDescriptor instead')
const TeamUpdateResponse$json = const {
  '1': 'TeamUpdateResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `TeamUpdateResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamUpdateResponseDescriptor = $convert.base64Decode('ChJUZWFtVXBkYXRlUmVzcG9uc2USDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use teamGetRequestDescriptor instead')
const TeamGetRequest$json = const {
  '1': 'TeamGetRequest',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 3, '10': 'ids'},
    const {'1': 'names', '3': 2, '4': 3, '5': 9, '10': 'names'},
    const {'1': 'departments', '3': 3, '4': 3, '5': 9, '10': 'departments'},
  ],
};

/// Descriptor for `TeamGetRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamGetRequestDescriptor = $convert.base64Decode('Cg5UZWFtR2V0UmVxdWVzdBIQCgNpZHMYASADKANSA2lkcxIUCgVuYW1lcxgCIAMoCVIFbmFtZXMSIAoLZGVwYXJ0bWVudHMYAyADKAlSC2RlcGFydG1lbnRz');
@$core.Deprecated('Use teamGetResponseDescriptor instead')
const TeamGetResponse$json = const {
  '1': 'TeamGetResponse',
  '2': const [
    const {'1': 'teams', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.team.Team', '10': 'teams'},
  ],
};

/// Descriptor for `TeamGetResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamGetResponseDescriptor = $convert.base64Decode('Cg9UZWFtR2V0UmVzcG9uc2USTAoFdGVhbXMYASADKAsyNi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRlYW0uVGVhbVIFdGVhbXM=');
@$core.Deprecated('Use teamGetByIDRequestDescriptor instead')
const TeamGetByIDRequest$json = const {
  '1': 'TeamGetByIDRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 3, '10': 'id'},
  ],
};

/// Descriptor for `TeamGetByIDRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamGetByIDRequestDescriptor = $convert.base64Decode('ChJUZWFtR2V0QnlJRFJlcXVlc3QSDgoCaWQYASABKANSAmlk');
@$core.Deprecated('Use teamGetByIDResponseDescriptor instead')
const TeamGetByIDResponse$json = const {
  '1': 'TeamGetByIDResponse',
  '2': const [
    const {'1': 'team', '3': 1, '4': 1, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.team.Team', '10': 'team'},
  ],
};

/// Descriptor for `TeamGetByIDResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List teamGetByIDResponseDescriptor = $convert.base64Decode('ChNUZWFtR2V0QnlJRFJlc3BvbnNlEkoKBHRlYW0YASABKAsyNi5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRlYW0uVGVhbVIEdGVhbQ==');
const $core.Map<$core.String, $core.dynamic> TeamRegistryServiceBase$json = const {
  '1': 'TeamRegistry',
  '2': const [
    const {'1': 'CreateTeam', '2': '.github.constantine27k.crnt_data_manager.api.team.TeamCreateRequest', '3': '.github.constantine27k.crnt_data_manager.api.team.TeamCreateResponse', '4': const {}},
    const {'1': 'UpdateTeam', '2': '.github.constantine27k.crnt_data_manager.api.team.TeamUpdateRequest', '3': '.github.constantine27k.crnt_data_manager.api.team.TeamUpdateResponse', '4': const {}},
    const {'1': 'GetTeams', '2': '.github.constantine27k.crnt_data_manager.api.team.TeamGetRequest', '3': '.github.constantine27k.crnt_data_manager.api.team.TeamGetResponse', '4': const {}},
    const {'1': 'GetTeamByID', '2': '.github.constantine27k.crnt_data_manager.api.team.TeamGetByIDRequest', '3': '.github.constantine27k.crnt_data_manager.api.team.TeamGetByIDResponse', '4': const {}},
  ],
};

@$core.Deprecated('Use teamRegistryServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> TeamRegistryServiceBase$messageJson = const {
  '.github.constantine27k.crnt_data_manager.api.team.TeamCreateRequest': TeamCreateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.team.Team': Team$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamCreateResponse': TeamCreateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamUpdateRequest': TeamUpdateRequest$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamUpdateResponse': TeamUpdateResponse$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamGetRequest': TeamGetRequest$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamGetResponse': TeamGetResponse$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamGetByIDRequest': TeamGetByIDRequest$json,
  '.github.constantine27k.crnt_data_manager.api.team.TeamGetByIDResponse': TeamGetByIDResponse$json,
};

/// Descriptor for `TeamRegistry`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List teamRegistryServiceDescriptor = $convert.base64Decode('CgxUZWFtUmVnaXN0cnkSrAEKCkNyZWF0ZVRlYW0SQy5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRlYW0uVGVhbUNyZWF0ZVJlcXVlc3QaRC5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRlYW0uVGVhbUNyZWF0ZVJlc3BvbnNlIhOC0+STAg06ASoiCC92MS90ZWFtErEBCgpVcGRhdGVUZWFtEkMuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50ZWFtLlRlYW1VcGRhdGVSZXF1ZXN0GkQuZ2l0aHViLmNvbnN0YW50aW5lMjdrLmNybnRfZGF0YV9tYW5hZ2VyLmFwaS50ZWFtLlRlYW1VcGRhdGVSZXNwb25zZSIYgtPkkwISOgEqGg0vdjEvdGVhbS97aWR9EqIBCghHZXRUZWFtcxJALmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGVhbS5UZWFtR2V0UmVxdWVzdBpBLmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGVhbS5UZWFtR2V0UmVzcG9uc2UiEYLT5JMCCxIJL3YxL3RlYW1zErEBCgtHZXRUZWFtQnlJRBJELmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkudGVhbS5UZWFtR2V0QnlJRFJlcXVlc3QaRS5naXRodWIuY29uc3RhbnRpbmUyN2suY3JudF9kYXRhX21hbmFnZXIuYXBpLnRlYW0uVGVhbUdldEJ5SURSZXNwb25zZSIVgtPkkwIPEg0vdjEvdGVhbS97aWR9');
