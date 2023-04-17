///
//  Generated code. Do not modify.
//  source: api/tasks/issue/issue.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class IssueType extends $pb.ProtobufEnum {
  static const IssueType ISSUE_TYPE_UNKNOWN = IssueType._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ISSUE_TYPE_UNKNOWN');
  static const IssueType ISSUE_TYPE_EPIC = IssueType._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ISSUE_TYPE_EPIC');
  static const IssueType ISSUE_TYPE_STORY = IssueType._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ISSUE_TYPE_STORY');
  static const IssueType ISSUE_TYPE_TASK = IssueType._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ISSUE_TYPE_TASK');
  static const IssueType ISSUE_TYPE_SUBTASK = IssueType._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'ISSUE_TYPE_SUBTASK');

  static const $core.List<IssueType> values = <IssueType> [
    ISSUE_TYPE_UNKNOWN,
    ISSUE_TYPE_EPIC,
    ISSUE_TYPE_STORY,
    ISSUE_TYPE_TASK,
    ISSUE_TYPE_SUBTASK,
  ];

  static final $core.Map<$core.int, IssueType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static IssueType? valueOf($core.int value) => _byValue[value];

  const IssueType._($core.int v, $core.String n) : super(v, n);
}

