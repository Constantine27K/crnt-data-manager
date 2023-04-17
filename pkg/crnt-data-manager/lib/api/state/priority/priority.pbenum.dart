///
//  Generated code. Do not modify.
//  source: api/state/priority/priority.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Priority extends $pb.ProtobufEnum {
  static const Priority PRIORITY_UNKNOWN = Priority._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PRIORITY_UNKNOWN');
  static const Priority PRIORITY_CRITICAL = Priority._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PRIORITY_CRITICAL');
  static const Priority PRIORITY_HIGH = Priority._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PRIORITY_HIGH');
  static const Priority PRIORITY_MEDIUM = Priority._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PRIORITY_MEDIUM');
  static const Priority PRIORITY_LOW = Priority._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'PRIORITY_LOW');

  static const $core.List<Priority> values = <Priority> [
    PRIORITY_UNKNOWN,
    PRIORITY_CRITICAL,
    PRIORITY_HIGH,
    PRIORITY_MEDIUM,
    PRIORITY_LOW,
  ];

  static final $core.Map<$core.int, Priority> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Priority? valueOf($core.int value) => _byValue[value];

  const Priority._($core.int v, $core.String n) : super(v, n);
}

