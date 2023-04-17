///
//  Generated code. Do not modify.
//  source: api/state/status/status.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Common extends $pb.ProtobufEnum {
  static const Common STATUS_COMMON_UNKNOWN = Common._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_UNKNOWN');
  static const Common STATUS_COMMON_BACKLOG = Common._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_BACKLOG');
  static const Common STATUS_COMMON_IN_PROGRESS = Common._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_IN_PROGRESS');
  static const Common STATUS_COMMON_DONE = Common._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_DONE');
  static const Common STATUS_COMMON_READY_FOR_REVIEW = Common._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_READY_FOR_REVIEW');
  static const Common STATUS_COMMON_IN_REVIEW = Common._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_IN_REVIEW');
  static const Common STATUS_COMMON_REVIEW_FAILED = Common._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_REVIEW_FAILED');
  static const Common STATUS_COMMON_REVIEW_PASSED = Common._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_REVIEW_PASSED');
  static const Common STATUS_COMMON_GIVEN_TO_CUSTOMER = Common._(8, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_GIVEN_TO_CUSTOMER');
  static const Common STATUS_COMMON_CLOSED = Common._(9, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_CLOSED');
  static const Common STATUS_COMMON_ON_HOLD = Common._(10, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_COMMON_ON_HOLD');

  static const $core.List<Common> values = <Common> [
    STATUS_COMMON_UNKNOWN,
    STATUS_COMMON_BACKLOG,
    STATUS_COMMON_IN_PROGRESS,
    STATUS_COMMON_DONE,
    STATUS_COMMON_READY_FOR_REVIEW,
    STATUS_COMMON_IN_REVIEW,
    STATUS_COMMON_REVIEW_FAILED,
    STATUS_COMMON_REVIEW_PASSED,
    STATUS_COMMON_GIVEN_TO_CUSTOMER,
    STATUS_COMMON_CLOSED,
    STATUS_COMMON_ON_HOLD,
  ];

  static final $core.Map<$core.int, Common> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Common? valueOf($core.int value) => _byValue[value];

  const Common._($core.int v, $core.String n) : super(v, n);
}

class Development extends $pb.ProtobufEnum {
  static const Development STATUS_DEVELOPMENT_UNKNOWN = Development._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_UNKNOWN');
  static const Development STATUS_DEVELOPMENT_BACKLOG = Development._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_BACKLOG');
  static const Development STATUS_DEVELOPMENT_IN_PROGRESS = Development._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_IN_PROGRESS');
  static const Development STATUS_DEVELOPMENT_IN_REVIEW = Development._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_IN_REVIEW');
  static const Development STATUS_DEVELOPMENT_READY_FOR_TESTING = Development._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_READY_FOR_TESTING');
  static const Development STATUS_DEVELOPMENT_TESTING = Development._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_TESTING');
  static const Development STATUS_DEVELOPMENT_TESTING_PASSED = Development._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_TESTING_PASSED');
  static const Development STATUS_DEVELOPMENT_DONE = Development._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_DONE');
  static const Development STATUS_DEVELOPMENT_READY_TO_DEPLOY = Development._(8, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_READY_TO_DEPLOY');
  static const Development STATUS_DEVELOPMENT_CLOSED = Development._(9, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_CLOSED');
  static const Development STATUS_DEVELOPMENT_ON_HOLD = Development._(10, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_DEVELOPMENT_ON_HOLD');

  static const $core.List<Development> values = <Development> [
    STATUS_DEVELOPMENT_UNKNOWN,
    STATUS_DEVELOPMENT_BACKLOG,
    STATUS_DEVELOPMENT_IN_PROGRESS,
    STATUS_DEVELOPMENT_IN_REVIEW,
    STATUS_DEVELOPMENT_READY_FOR_TESTING,
    STATUS_DEVELOPMENT_TESTING,
    STATUS_DEVELOPMENT_TESTING_PASSED,
    STATUS_DEVELOPMENT_DONE,
    STATUS_DEVELOPMENT_READY_TO_DEPLOY,
    STATUS_DEVELOPMENT_CLOSED,
    STATUS_DEVELOPMENT_ON_HOLD,
  ];

  static final $core.Map<$core.int, Development> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Development? valueOf($core.int value) => _byValue[value];

  const Development._($core.int v, $core.String n) : super(v, n);
}

class Epic extends $pb.ProtobufEnum {
  static const Epic STATUS_EPIC_UNKNOWN = Epic._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_UNKNOWN');
  static const Epic STATUS_EPIC_BACKLOG = Epic._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_BACKLOG');
  static const Epic STATUS_EPIC_PLANNING = Epic._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_PLANNING');
  static const Epic STATUS_EPIC_DESIGNING = Epic._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_DESIGNING');
  static const Epic STATUS_EPIC_IT = Epic._(4, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_IT');
  static const Epic STATUS_EPIC_DONE = Epic._(5, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_DONE');
  static const Epic STATUS_EPIC_CLOSED = Epic._(6, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_CLOSED');
  static const Epic STATUS_EPIC_ON_HOLD = Epic._(7, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_EPIC_ON_HOLD');

  static const $core.List<Epic> values = <Epic> [
    STATUS_EPIC_UNKNOWN,
    STATUS_EPIC_BACKLOG,
    STATUS_EPIC_PLANNING,
    STATUS_EPIC_DESIGNING,
    STATUS_EPIC_IT,
    STATUS_EPIC_DONE,
    STATUS_EPIC_CLOSED,
    STATUS_EPIC_ON_HOLD,
  ];

  static final $core.Map<$core.int, Epic> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Epic? valueOf($core.int value) => _byValue[value];

  const Epic._($core.int v, $core.String n) : super(v, n);
}

class Sprint extends $pb.ProtobufEnum {
  static const Sprint STATUS_SPRINT_UNKNOWN = Sprint._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_SPRINT_UNKNOWN');
  static const Sprint STATUS_SPRINT_BACKLOG = Sprint._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_SPRINT_BACKLOG');
  static const Sprint STATUS_SPRINT_ACTIVE = Sprint._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_SPRINT_ACTIVE');
  static const Sprint STATUS_SPRINT_FINISHED = Sprint._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'STATUS_SPRINT_FINISHED');

  static const $core.List<Sprint> values = <Sprint> [
    STATUS_SPRINT_UNKNOWN,
    STATUS_SPRINT_BACKLOG,
    STATUS_SPRINT_ACTIVE,
    STATUS_SPRINT_FINISHED,
  ];

  static final $core.Map<$core.int, Sprint> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Sprint? valueOf($core.int value) => _byValue[value];

  const Sprint._($core.int v, $core.String n) : super(v, n);
}

