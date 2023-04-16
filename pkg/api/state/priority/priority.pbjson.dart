///
//  Generated code. Do not modify.
//  source: api/state/priority/priority.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use priorityDescriptor instead')
const Priority$json = const {
  '1': 'Priority',
  '2': const [
    const {'1': 'PRIORITY_UNKNOWN', '2': 0},
    const {'1': 'PRIORITY_CRITICAL', '2': 1},
    const {'1': 'PRIORITY_HIGH', '2': 2},
    const {'1': 'PRIORITY_MEDIUM', '2': 3},
    const {'1': 'PRIORITY_LOW', '2': 4},
  ],
};

/// Descriptor for `Priority`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List priorityDescriptor = $convert.base64Decode('CghQcmlvcml0eRIUChBQUklPUklUWV9VTktOT1dOEAASFQoRUFJJT1JJVFlfQ1JJVElDQUwQARIRCg1QUklPUklUWV9ISUdIEAISEwoPUFJJT1JJVFlfTUVESVVNEAMSEAoMUFJJT1JJVFlfTE9XEAQ=');
