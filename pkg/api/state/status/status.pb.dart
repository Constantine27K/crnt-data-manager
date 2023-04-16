///
//  Generated code. Do not modify.
//  source: api/state/status/status.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'status.pbenum.dart';

export 'status.pbenum.dart';

enum IssueStatus_Status {
  common, 
  development, 
  epic, 
  notSet
}

class IssueStatus extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, IssueStatus_Status> _IssueStatus_StatusByTag = {
    1 : IssueStatus_Status.common,
    2 : IssueStatus_Status.development,
    3 : IssueStatus_Status.epic,
    0 : IssueStatus_Status.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueStatus', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.state.status'), createEmptyInstance: create)
    ..oo(0, [1, 2, 3])
    ..aOM<IssueStatusCommon>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'common', subBuilder: IssueStatusCommon.create)
    ..aOM<IssueStatusDevelopment>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'development', subBuilder: IssueStatusDevelopment.create)
    ..aOM<IssueStatusEpic>(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'epic', subBuilder: IssueStatusEpic.create)
    ..hasRequiredFields = false
  ;

  IssueStatus._() : super();
  factory IssueStatus({
    IssueStatusCommon? common,
    IssueStatusDevelopment? development,
    IssueStatusEpic? epic,
  }) {
    final _result = create();
    if (common != null) {
      _result.common = common;
    }
    if (development != null) {
      _result.development = development;
    }
    if (epic != null) {
      _result.epic = epic;
    }
    return _result;
  }
  factory IssueStatus.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueStatus.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueStatus clone() => IssueStatus()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueStatus copyWith(void Function(IssueStatus) updates) => super.copyWith((message) => updates(message as IssueStatus)) as IssueStatus; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueStatus create() => IssueStatus._();
  IssueStatus createEmptyInstance() => create();
  static $pb.PbList<IssueStatus> createRepeated() => $pb.PbList<IssueStatus>();
  @$core.pragma('dart2js:noInline')
  static IssueStatus getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueStatus>(create);
  static IssueStatus? _defaultInstance;

  IssueStatus_Status whichStatus() => _IssueStatus_StatusByTag[$_whichOneof(0)]!;
  void clearStatus() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  IssueStatusCommon get common => $_getN(0);
  @$pb.TagNumber(1)
  set common(IssueStatusCommon v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCommon() => $_has(0);
  @$pb.TagNumber(1)
  void clearCommon() => clearField(1);
  @$pb.TagNumber(1)
  IssueStatusCommon ensureCommon() => $_ensure(0);

  @$pb.TagNumber(2)
  IssueStatusDevelopment get development => $_getN(1);
  @$pb.TagNumber(2)
  set development(IssueStatusDevelopment v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasDevelopment() => $_has(1);
  @$pb.TagNumber(2)
  void clearDevelopment() => clearField(2);
  @$pb.TagNumber(2)
  IssueStatusDevelopment ensureDevelopment() => $_ensure(1);

  @$pb.TagNumber(3)
  IssueStatusEpic get epic => $_getN(2);
  @$pb.TagNumber(3)
  set epic(IssueStatusEpic v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasEpic() => $_has(2);
  @$pb.TagNumber(3)
  void clearEpic() => clearField(3);
  @$pb.TagNumber(3)
  IssueStatusEpic ensureEpic() => $_ensure(2);
}

class IssueStatusCommon extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueStatusCommon', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.state.status'), createEmptyInstance: create)
    ..e<Common>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: Common.STATUS_COMMON_UNKNOWN, valueOf: Common.valueOf, enumValues: Common.values)
    ..hasRequiredFields = false
  ;

  IssueStatusCommon._() : super();
  factory IssueStatusCommon({
    Common? status,
  }) {
    final _result = create();
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory IssueStatusCommon.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueStatusCommon.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueStatusCommon clone() => IssueStatusCommon()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueStatusCommon copyWith(void Function(IssueStatusCommon) updates) => super.copyWith((message) => updates(message as IssueStatusCommon)) as IssueStatusCommon; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueStatusCommon create() => IssueStatusCommon._();
  IssueStatusCommon createEmptyInstance() => create();
  static $pb.PbList<IssueStatusCommon> createRepeated() => $pb.PbList<IssueStatusCommon>();
  @$core.pragma('dart2js:noInline')
  static IssueStatusCommon getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueStatusCommon>(create);
  static IssueStatusCommon? _defaultInstance;

  @$pb.TagNumber(1)
  Common get status => $_getN(0);
  @$pb.TagNumber(1)
  set status(Common v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(1)
  void clearStatus() => clearField(1);
}

class IssueStatusDevelopment extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueStatusDevelopment', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.state.status'), createEmptyInstance: create)
    ..e<Development>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: Development.STATUS_DEVELOPMENT_UNKNOWN, valueOf: Development.valueOf, enumValues: Development.values)
    ..hasRequiredFields = false
  ;

  IssueStatusDevelopment._() : super();
  factory IssueStatusDevelopment({
    Development? status,
  }) {
    final _result = create();
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory IssueStatusDevelopment.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueStatusDevelopment.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueStatusDevelopment clone() => IssueStatusDevelopment()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueStatusDevelopment copyWith(void Function(IssueStatusDevelopment) updates) => super.copyWith((message) => updates(message as IssueStatusDevelopment)) as IssueStatusDevelopment; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueStatusDevelopment create() => IssueStatusDevelopment._();
  IssueStatusDevelopment createEmptyInstance() => create();
  static $pb.PbList<IssueStatusDevelopment> createRepeated() => $pb.PbList<IssueStatusDevelopment>();
  @$core.pragma('dart2js:noInline')
  static IssueStatusDevelopment getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueStatusDevelopment>(create);
  static IssueStatusDevelopment? _defaultInstance;

  @$pb.TagNumber(2)
  Development get status => $_getN(0);
  @$pb.TagNumber(2)
  set status(Development v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(2)
  void clearStatus() => clearField(2);
}

class IssueStatusEpic extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'IssueStatusEpic', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'github.constantine27k.crnt_data_manager.api.state.status'), createEmptyInstance: create)
    ..e<Epic>(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'status', $pb.PbFieldType.OE, defaultOrMaker: Epic.STATUS_EPIC_UNKNOWN, valueOf: Epic.valueOf, enumValues: Epic.values)
    ..hasRequiredFields = false
  ;

  IssueStatusEpic._() : super();
  factory IssueStatusEpic({
    Epic? status,
  }) {
    final _result = create();
    if (status != null) {
      _result.status = status;
    }
    return _result;
  }
  factory IssueStatusEpic.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IssueStatusEpic.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  IssueStatusEpic clone() => IssueStatusEpic()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  IssueStatusEpic copyWith(void Function(IssueStatusEpic) updates) => super.copyWith((message) => updates(message as IssueStatusEpic)) as IssueStatusEpic; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IssueStatusEpic create() => IssueStatusEpic._();
  IssueStatusEpic createEmptyInstance() => create();
  static $pb.PbList<IssueStatusEpic> createRepeated() => $pb.PbList<IssueStatusEpic>();
  @$core.pragma('dart2js:noInline')
  static IssueStatusEpic getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IssueStatusEpic>(create);
  static IssueStatusEpic? _defaultInstance;

  @$pb.TagNumber(2)
  Epic get status => $_getN(0);
  @$pb.TagNumber(2)
  set status(Epic v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasStatus() => $_has(0);
  @$pb.TagNumber(2)
  void clearStatus() => clearField(2);
}

