// source: google/api/serviceusage/v1/resources.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_api_auth_pb = require('../../../../google/api/auth_pb.js');
goog.object.extend(proto, google_api_auth_pb);
var google_api_documentation_pb = require('../../../../google/api/documentation_pb.js');
goog.object.extend(proto, google_api_documentation_pb);
var google_api_endpoint_pb = require('../../../../google/api/endpoint_pb.js');
goog.object.extend(proto, google_api_endpoint_pb);
var google_api_monitored_resource_pb = require('../../../../google/api/monitored_resource_pb.js');
goog.object.extend(proto, google_api_monitored_resource_pb);
var google_api_monitoring_pb = require('../../../../google/api/monitoring_pb.js');
goog.object.extend(proto, google_api_monitoring_pb);
var google_api_quota_pb = require('../../../../google/api/quota_pb.js');
goog.object.extend(proto, google_api_quota_pb);
var google_api_usage_pb = require('../../../../google/api/usage_pb.js');
goog.object.extend(proto, google_api_usage_pb);
var google_protobuf_api_pb = require('google-protobuf/google/protobuf/api_pb.js');
goog.object.extend(proto, google_protobuf_api_pb);
var google_api_annotations_pb = require('../../../../google/api/annotations_pb.js');
goog.object.extend(proto, google_api_annotations_pb);
goog.exportSymbol('proto.google.api.serviceusage.v1.OperationMetadata', null, global);
goog.exportSymbol('proto.google.api.serviceusage.v1.Service', null, global);
goog.exportSymbol('proto.google.api.serviceusage.v1.ServiceConfig', null, global);
goog.exportSymbol('proto.google.api.serviceusage.v1.State', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.google.api.serviceusage.v1.Service = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.google.api.serviceusage.v1.Service, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.google.api.serviceusage.v1.Service.displayName = 'proto.google.api.serviceusage.v1.Service';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.google.api.serviceusage.v1.ServiceConfig = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.google.api.serviceusage.v1.ServiceConfig.repeatedFields_, null);
};
goog.inherits(proto.google.api.serviceusage.v1.ServiceConfig, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.google.api.serviceusage.v1.ServiceConfig.displayName = 'proto.google.api.serviceusage.v1.ServiceConfig';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.google.api.serviceusage.v1.OperationMetadata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.google.api.serviceusage.v1.OperationMetadata.repeatedFields_, null);
};
goog.inherits(proto.google.api.serviceusage.v1.OperationMetadata, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.google.api.serviceusage.v1.OperationMetadata.displayName = 'proto.google.api.serviceusage.v1.OperationMetadata';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.google.api.serviceusage.v1.Service.prototype.toObject = function(opt_includeInstance) {
  return proto.google.api.serviceusage.v1.Service.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.google.api.serviceusage.v1.Service} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.serviceusage.v1.Service.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    parent: jspb.Message.getFieldWithDefault(msg, 5, ""),
    config: (f = msg.getConfig()) && proto.google.api.serviceusage.v1.ServiceConfig.toObject(includeInstance, f),
    state: jspb.Message.getFieldWithDefault(msg, 4, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.google.api.serviceusage.v1.Service}
 */
proto.google.api.serviceusage.v1.Service.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.google.api.serviceusage.v1.Service;
  return proto.google.api.serviceusage.v1.Service.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.google.api.serviceusage.v1.Service} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.google.api.serviceusage.v1.Service}
 */
proto.google.api.serviceusage.v1.Service.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setParent(value);
      break;
    case 2:
      var value = new proto.google.api.serviceusage.v1.ServiceConfig;
      reader.readMessage(value,proto.google.api.serviceusage.v1.ServiceConfig.deserializeBinaryFromReader);
      msg.setConfig(value);
      break;
    case 4:
      var value = /** @type {!proto.google.api.serviceusage.v1.State} */ (reader.readEnum());
      msg.setState(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.google.api.serviceusage.v1.Service.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.google.api.serviceusage.v1.Service.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.google.api.serviceusage.v1.Service} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.serviceusage.v1.Service.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getParent();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getConfig();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.google.api.serviceusage.v1.ServiceConfig.serializeBinaryToWriter
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.google.api.serviceusage.v1.Service.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.google.api.serviceusage.v1.Service} returns this
 */
proto.google.api.serviceusage.v1.Service.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string parent = 5;
 * @return {string}
 */
proto.google.api.serviceusage.v1.Service.prototype.getParent = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.google.api.serviceusage.v1.Service} returns this
 */
proto.google.api.serviceusage.v1.Service.prototype.setParent = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional ServiceConfig config = 2;
 * @return {?proto.google.api.serviceusage.v1.ServiceConfig}
 */
proto.google.api.serviceusage.v1.Service.prototype.getConfig = function() {
  return /** @type{?proto.google.api.serviceusage.v1.ServiceConfig} */ (
    jspb.Message.getWrapperField(this, proto.google.api.serviceusage.v1.ServiceConfig, 2));
};


/**
 * @param {?proto.google.api.serviceusage.v1.ServiceConfig|undefined} value
 * @return {!proto.google.api.serviceusage.v1.Service} returns this
*/
proto.google.api.serviceusage.v1.Service.prototype.setConfig = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.google.api.serviceusage.v1.Service} returns this
 */
proto.google.api.serviceusage.v1.Service.prototype.clearConfig = function() {
  return this.setConfig(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.google.api.serviceusage.v1.Service.prototype.hasConfig = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional State state = 4;
 * @return {!proto.google.api.serviceusage.v1.State}
 */
proto.google.api.serviceusage.v1.Service.prototype.getState = function() {
  return /** @type {!proto.google.api.serviceusage.v1.State} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.google.api.serviceusage.v1.State} value
 * @return {!proto.google.api.serviceusage.v1.Service} returns this
 */
proto.google.api.serviceusage.v1.Service.prototype.setState = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.google.api.serviceusage.v1.ServiceConfig.repeatedFields_ = [3,18,25];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.toObject = function(opt_includeInstance) {
  return proto.google.api.serviceusage.v1.ServiceConfig.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.google.api.serviceusage.v1.ServiceConfig} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.serviceusage.v1.ServiceConfig.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    title: jspb.Message.getFieldWithDefault(msg, 2, ""),
    apisList: jspb.Message.toObjectList(msg.getApisList(),
    google_protobuf_api_pb.Api.toObject, includeInstance),
    documentation: (f = msg.getDocumentation()) && google_api_documentation_pb.Documentation.toObject(includeInstance, f),
    quota: (f = msg.getQuota()) && google_api_quota_pb.Quota.toObject(includeInstance, f),
    authentication: (f = msg.getAuthentication()) && google_api_auth_pb.Authentication.toObject(includeInstance, f),
    usage: (f = msg.getUsage()) && google_api_usage_pb.Usage.toObject(includeInstance, f),
    endpointsList: jspb.Message.toObjectList(msg.getEndpointsList(),
    google_api_endpoint_pb.Endpoint.toObject, includeInstance),
    monitoredResourcesList: jspb.Message.toObjectList(msg.getMonitoredResourcesList(),
    google_api_monitored_resource_pb.MonitoredResourceDescriptor.toObject, includeInstance),
    monitoring: (f = msg.getMonitoring()) && google_api_monitoring_pb.Monitoring.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig}
 */
proto.google.api.serviceusage.v1.ServiceConfig.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.google.api.serviceusage.v1.ServiceConfig;
  return proto.google.api.serviceusage.v1.ServiceConfig.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.google.api.serviceusage.v1.ServiceConfig} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig}
 */
proto.google.api.serviceusage.v1.ServiceConfig.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setTitle(value);
      break;
    case 3:
      var value = new google_protobuf_api_pb.Api;
      reader.readMessage(value,google_protobuf_api_pb.Api.deserializeBinaryFromReader);
      msg.addApis(value);
      break;
    case 6:
      var value = new google_api_documentation_pb.Documentation;
      reader.readMessage(value,google_api_documentation_pb.Documentation.deserializeBinaryFromReader);
      msg.setDocumentation(value);
      break;
    case 10:
      var value = new google_api_quota_pb.Quota;
      reader.readMessage(value,google_api_quota_pb.Quota.deserializeBinaryFromReader);
      msg.setQuota(value);
      break;
    case 11:
      var value = new google_api_auth_pb.Authentication;
      reader.readMessage(value,google_api_auth_pb.Authentication.deserializeBinaryFromReader);
      msg.setAuthentication(value);
      break;
    case 15:
      var value = new google_api_usage_pb.Usage;
      reader.readMessage(value,google_api_usage_pb.Usage.deserializeBinaryFromReader);
      msg.setUsage(value);
      break;
    case 18:
      var value = new google_api_endpoint_pb.Endpoint;
      reader.readMessage(value,google_api_endpoint_pb.Endpoint.deserializeBinaryFromReader);
      msg.addEndpoints(value);
      break;
    case 25:
      var value = new google_api_monitored_resource_pb.MonitoredResourceDescriptor;
      reader.readMessage(value,google_api_monitored_resource_pb.MonitoredResourceDescriptor.deserializeBinaryFromReader);
      msg.addMonitoredResources(value);
      break;
    case 28:
      var value = new google_api_monitoring_pb.Monitoring;
      reader.readMessage(value,google_api_monitoring_pb.Monitoring.deserializeBinaryFromReader);
      msg.setMonitoring(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.google.api.serviceusage.v1.ServiceConfig.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.google.api.serviceusage.v1.ServiceConfig} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.serviceusage.v1.ServiceConfig.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getTitle();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getApisList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      google_protobuf_api_pb.Api.serializeBinaryToWriter
    );
  }
  f = message.getDocumentation();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_api_documentation_pb.Documentation.serializeBinaryToWriter
    );
  }
  f = message.getQuota();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      google_api_quota_pb.Quota.serializeBinaryToWriter
    );
  }
  f = message.getAuthentication();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      google_api_auth_pb.Authentication.serializeBinaryToWriter
    );
  }
  f = message.getUsage();
  if (f != null) {
    writer.writeMessage(
      15,
      f,
      google_api_usage_pb.Usage.serializeBinaryToWriter
    );
  }
  f = message.getEndpointsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      18,
      f,
      google_api_endpoint_pb.Endpoint.serializeBinaryToWriter
    );
  }
  f = message.getMonitoredResourcesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      25,
      f,
      google_api_monitored_resource_pb.MonitoredResourceDescriptor.serializeBinaryToWriter
    );
  }
  f = message.getMonitoring();
  if (f != null) {
    writer.writeMessage(
      28,
      f,
      google_api_monitoring_pb.Monitoring.serializeBinaryToWriter
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string title = 2;
 * @return {string}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getTitle = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setTitle = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * repeated google.protobuf.Api apis = 3;
 * @return {!Array<!proto.google.protobuf.Api>}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getApisList = function() {
  return /** @type{!Array<!proto.google.protobuf.Api>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_protobuf_api_pb.Api, 3));
};


/**
 * @param {!Array<!proto.google.protobuf.Api>} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setApisList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.google.protobuf.Api=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.protobuf.Api}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.addApis = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.google.protobuf.Api, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearApisList = function() {
  return this.setApisList([]);
};


/**
 * optional google.api.Documentation documentation = 6;
 * @return {?proto.google.api.Documentation}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getDocumentation = function() {
  return /** @type{?proto.google.api.Documentation} */ (
    jspb.Message.getWrapperField(this, google_api_documentation_pb.Documentation, 6));
};


/**
 * @param {?proto.google.api.Documentation|undefined} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setDocumentation = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearDocumentation = function() {
  return this.setDocumentation(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.hasDocumentation = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional google.api.Quota quota = 10;
 * @return {?proto.google.api.Quota}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getQuota = function() {
  return /** @type{?proto.google.api.Quota} */ (
    jspb.Message.getWrapperField(this, google_api_quota_pb.Quota, 10));
};


/**
 * @param {?proto.google.api.Quota|undefined} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setQuota = function(value) {
  return jspb.Message.setWrapperField(this, 10, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearQuota = function() {
  return this.setQuota(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.hasQuota = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional google.api.Authentication authentication = 11;
 * @return {?proto.google.api.Authentication}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getAuthentication = function() {
  return /** @type{?proto.google.api.Authentication} */ (
    jspb.Message.getWrapperField(this, google_api_auth_pb.Authentication, 11));
};


/**
 * @param {?proto.google.api.Authentication|undefined} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setAuthentication = function(value) {
  return jspb.Message.setWrapperField(this, 11, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearAuthentication = function() {
  return this.setAuthentication(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.hasAuthentication = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional google.api.Usage usage = 15;
 * @return {?proto.google.api.Usage}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getUsage = function() {
  return /** @type{?proto.google.api.Usage} */ (
    jspb.Message.getWrapperField(this, google_api_usage_pb.Usage, 15));
};


/**
 * @param {?proto.google.api.Usage|undefined} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setUsage = function(value) {
  return jspb.Message.setWrapperField(this, 15, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearUsage = function() {
  return this.setUsage(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.hasUsage = function() {
  return jspb.Message.getField(this, 15) != null;
};


/**
 * repeated google.api.Endpoint endpoints = 18;
 * @return {!Array<!proto.google.api.Endpoint>}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getEndpointsList = function() {
  return /** @type{!Array<!proto.google.api.Endpoint>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_api_endpoint_pb.Endpoint, 18));
};


/**
 * @param {!Array<!proto.google.api.Endpoint>} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setEndpointsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 18, value);
};


/**
 * @param {!proto.google.api.Endpoint=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.api.Endpoint}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.addEndpoints = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 18, opt_value, proto.google.api.Endpoint, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearEndpointsList = function() {
  return this.setEndpointsList([]);
};


/**
 * repeated google.api.MonitoredResourceDescriptor monitored_resources = 25;
 * @return {!Array<!proto.google.api.MonitoredResourceDescriptor>}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getMonitoredResourcesList = function() {
  return /** @type{!Array<!proto.google.api.MonitoredResourceDescriptor>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_api_monitored_resource_pb.MonitoredResourceDescriptor, 25));
};


/**
 * @param {!Array<!proto.google.api.MonitoredResourceDescriptor>} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setMonitoredResourcesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 25, value);
};


/**
 * @param {!proto.google.api.MonitoredResourceDescriptor=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.api.MonitoredResourceDescriptor}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.addMonitoredResources = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 25, opt_value, proto.google.api.MonitoredResourceDescriptor, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearMonitoredResourcesList = function() {
  return this.setMonitoredResourcesList([]);
};


/**
 * optional google.api.Monitoring monitoring = 28;
 * @return {?proto.google.api.Monitoring}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.getMonitoring = function() {
  return /** @type{?proto.google.api.Monitoring} */ (
    jspb.Message.getWrapperField(this, google_api_monitoring_pb.Monitoring, 28));
};


/**
 * @param {?proto.google.api.Monitoring|undefined} value
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
*/
proto.google.api.serviceusage.v1.ServiceConfig.prototype.setMonitoring = function(value) {
  return jspb.Message.setWrapperField(this, 28, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.google.api.serviceusage.v1.ServiceConfig} returns this
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.clearMonitoring = function() {
  return this.setMonitoring(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.google.api.serviceusage.v1.ServiceConfig.prototype.hasMonitoring = function() {
  return jspb.Message.getField(this, 28) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.google.api.serviceusage.v1.OperationMetadata.repeatedFields_ = [2];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.google.api.serviceusage.v1.OperationMetadata.prototype.toObject = function(opt_includeInstance) {
  return proto.google.api.serviceusage.v1.OperationMetadata.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.google.api.serviceusage.v1.OperationMetadata} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.serviceusage.v1.OperationMetadata.toObject = function(includeInstance, msg) {
  var f, obj = {
    resourceNamesList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.google.api.serviceusage.v1.OperationMetadata}
 */
proto.google.api.serviceusage.v1.OperationMetadata.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.google.api.serviceusage.v1.OperationMetadata;
  return proto.google.api.serviceusage.v1.OperationMetadata.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.google.api.serviceusage.v1.OperationMetadata} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.google.api.serviceusage.v1.OperationMetadata}
 */
proto.google.api.serviceusage.v1.OperationMetadata.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addResourceNames(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.google.api.serviceusage.v1.OperationMetadata.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.google.api.serviceusage.v1.OperationMetadata.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.google.api.serviceusage.v1.OperationMetadata} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.google.api.serviceusage.v1.OperationMetadata.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getResourceNamesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
};


/**
 * repeated string resource_names = 2;
 * @return {!Array<string>}
 */
proto.google.api.serviceusage.v1.OperationMetadata.prototype.getResourceNamesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.google.api.serviceusage.v1.OperationMetadata} returns this
 */
proto.google.api.serviceusage.v1.OperationMetadata.prototype.setResourceNamesList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.google.api.serviceusage.v1.OperationMetadata} returns this
 */
proto.google.api.serviceusage.v1.OperationMetadata.prototype.addResourceNames = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.google.api.serviceusage.v1.OperationMetadata} returns this
 */
proto.google.api.serviceusage.v1.OperationMetadata.prototype.clearResourceNamesList = function() {
  return this.setResourceNamesList([]);
};


/**
 * @enum {number}
 */
proto.google.api.serviceusage.v1.State = {
  STATE_UNSPECIFIED: 0,
  DISABLED: 1,
  ENABLED: 2
};

goog.object.extend(exports, proto.google.api.serviceusage.v1);
