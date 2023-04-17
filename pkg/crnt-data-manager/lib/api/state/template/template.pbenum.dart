///
//  Generated code. Do not modify.
//  source: api/state/template/template.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Template extends $pb.ProtobufEnum {
  static const Template TEMPLATE_UNKNOWN = Template._(0, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TEMPLATE_UNKNOWN');
  static const Template TEMPLATE_COMMON = Template._(1, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TEMPLATE_COMMON');
  static const Template TEMPLATE_DEVELOPMENT = Template._(2, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TEMPLATE_DEVELOPMENT');
  static const Template TEMPLATE_EPIC = Template._(3, const $core.bool.fromEnvironment('protobuf.omit_enum_names') ? '' : 'TEMPLATE_EPIC');

  static const $core.List<Template> values = <Template> [
    TEMPLATE_UNKNOWN,
    TEMPLATE_COMMON,
    TEMPLATE_DEVELOPMENT,
    TEMPLATE_EPIC,
  ];

  static final $core.Map<$core.int, Template> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Template? valueOf($core.int value) => _byValue[value];

  const Template._($core.int v, $core.String n) : super(v, n);
}

