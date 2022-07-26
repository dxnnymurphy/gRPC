# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: anomaly.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ranomaly.proto\x12\x02pb\x1a\x1fgoogle/protobuf/timestamp.proto\"d\n\x06Metric\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\r\n\x05topic\x18\x02 \x01(\t\x12\x13\n\x0brefreshbool\x18\x03 \x01(\t\x12\x0c\n\x04stop\x18\x04 \x01(\t\"e\n\x07Metric1\x12-\n\tstartTime\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12+\n\x07\x65ndTime\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"-\n\x0e\x41nomalyRequest\x12\x1b\n\x07metrics\x18\x01 \x03(\x0b\x32\n.pb.Metric\"#\n\x0f\x41nomalyResponse\x12\x10\n\x08response\x18\x01 \x03(\t\"1\n\x11ModelTrainRequest\x12\x1c\n\x07metrics\x18\x01 \x03(\x0b\x32\x0b.pb.Metric1\"&\n\x12ModelTrainResponse\x12\x10\n\x08response\x18\x01 \x03(\t2\x82\x01\n\x10\x41nomalyDetection\x12\x34\n\x07Predict\x12\x12.pb.AnomalyRequest\x1a\x13.pb.AnomalyResponse\"\x00\x12\x38\n\x05Train\x12\x15.pb.ModelTrainRequest\x1a\x16.pb.ModelTrainResponse\"\x00\x42 Z\x1egithub.com/dxnnymurphy/gRPC/pbb\x06proto3')



_METRIC = DESCRIPTOR.message_types_by_name['Metric']
_METRIC1 = DESCRIPTOR.message_types_by_name['Metric1']
_ANOMALYREQUEST = DESCRIPTOR.message_types_by_name['AnomalyRequest']
_ANOMALYRESPONSE = DESCRIPTOR.message_types_by_name['AnomalyResponse']
_MODELTRAINREQUEST = DESCRIPTOR.message_types_by_name['ModelTrainRequest']
_MODELTRAINRESPONSE = DESCRIPTOR.message_types_by_name['ModelTrainResponse']
Metric = _reflection.GeneratedProtocolMessageType('Metric', (_message.Message,), {
  'DESCRIPTOR' : _METRIC,
  '__module__' : 'anomaly_pb2'
  # @@protoc_insertion_point(class_scope:pb.Metric)
  })
_sym_db.RegisterMessage(Metric)

Metric1 = _reflection.GeneratedProtocolMessageType('Metric1', (_message.Message,), {
  'DESCRIPTOR' : _METRIC1,
  '__module__' : 'anomaly_pb2'
  # @@protoc_insertion_point(class_scope:pb.Metric1)
  })
_sym_db.RegisterMessage(Metric1)

AnomalyRequest = _reflection.GeneratedProtocolMessageType('AnomalyRequest', (_message.Message,), {
  'DESCRIPTOR' : _ANOMALYREQUEST,
  '__module__' : 'anomaly_pb2'
  # @@protoc_insertion_point(class_scope:pb.AnomalyRequest)
  })
_sym_db.RegisterMessage(AnomalyRequest)

AnomalyResponse = _reflection.GeneratedProtocolMessageType('AnomalyResponse', (_message.Message,), {
  'DESCRIPTOR' : _ANOMALYRESPONSE,
  '__module__' : 'anomaly_pb2'
  # @@protoc_insertion_point(class_scope:pb.AnomalyResponse)
  })
_sym_db.RegisterMessage(AnomalyResponse)

ModelTrainRequest = _reflection.GeneratedProtocolMessageType('ModelTrainRequest', (_message.Message,), {
  'DESCRIPTOR' : _MODELTRAINREQUEST,
  '__module__' : 'anomaly_pb2'
  # @@protoc_insertion_point(class_scope:pb.ModelTrainRequest)
  })
_sym_db.RegisterMessage(ModelTrainRequest)

ModelTrainResponse = _reflection.GeneratedProtocolMessageType('ModelTrainResponse', (_message.Message,), {
  'DESCRIPTOR' : _MODELTRAINRESPONSE,
  '__module__' : 'anomaly_pb2'
  # @@protoc_insertion_point(class_scope:pb.ModelTrainResponse)
  })
_sym_db.RegisterMessage(ModelTrainResponse)

_ANOMALYDETECTION = DESCRIPTOR.services_by_name['AnomalyDetection']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\036github.com/dxnnymurphy/gRPC/pb'
  _METRIC._serialized_start=54
  _METRIC._serialized_end=154
  _METRIC1._serialized_start=156
  _METRIC1._serialized_end=257
  _ANOMALYREQUEST._serialized_start=259
  _ANOMALYREQUEST._serialized_end=304
  _ANOMALYRESPONSE._serialized_start=306
  _ANOMALYRESPONSE._serialized_end=341
  _MODELTRAINREQUEST._serialized_start=343
  _MODELTRAINREQUEST._serialized_end=392
  _MODELTRAINRESPONSE._serialized_start=394
  _MODELTRAINRESPONSE._serialized_end=432
  _ANOMALYDETECTION._serialized_start=435
  _ANOMALYDETECTION._serialized_end=565
# @@protoc_insertion_point(module_scope)
