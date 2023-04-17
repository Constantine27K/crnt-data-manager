///
//  Generated code. Do not modify.
//  source: api/comments/comments.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use commentDescriptor instead')
const Comment$json = const {
  '1': 'Comment',
  '2': const [
    const {'1': 'author', '3': 1, '4': 1, '5': 9, '10': 'author'},
    const {'1': 'written_at', '3': 2, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'writtenAt'},
    const {'1': 'updated_at', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.Timestamp', '10': 'updatedAt'},
    const {'1': 'updated_by', '3': 4, '4': 1, '5': 9, '10': 'updatedBy'},
    const {'1': 'text', '3': 5, '4': 1, '5': 9, '10': 'text'},
  ],
};

/// Descriptor for `Comment`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List commentDescriptor = $convert.base64Decode('CgdDb21tZW50EhYKBmF1dGhvchgBIAEoCVIGYXV0aG9yEjkKCndyaXR0ZW5fYXQYAiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wUgl3cml0dGVuQXQSOQoKdXBkYXRlZF9hdBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBSCXVwZGF0ZWRBdBIdCgp1cGRhdGVkX2J5GAQgASgJUgl1cGRhdGVkQnkSEgoEdGV4dBgFIAEoCVIEdGV4dA==');
@$core.Deprecated('Use commentsDescriptor instead')
const Comments$json = const {
  '1': 'Comments',
  '2': const [
    const {'1': 'items', '3': 1, '4': 3, '5': 11, '6': '.github.constantine27k.crnt_data_manager.api.comments.Comment', '10': 'items'},
  ],
};

/// Descriptor for `Comments`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List commentsDescriptor = $convert.base64Decode('CghDb21tZW50cxJTCgVpdGVtcxgBIAMoCzI9LmdpdGh1Yi5jb25zdGFudGluZTI3ay5jcm50X2RhdGFfbWFuYWdlci5hcGkuY29tbWVudHMuQ29tbWVudFIFaXRlbXM=');
