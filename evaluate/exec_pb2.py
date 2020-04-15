# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: exec.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='exec.proto',
  package='evaluate',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=b'\n\nexec.proto\x12\x08\x65valuate\"&\n\x04\x43ode\x12\x0c\n\x04\x63ode\x18\x01 \x01(\t\x12\x10\n\x08\x66unction\x18\x02 \x01(\t\"*\n\rCompileResult\x12\n\n\x02id\x18\x01 \x01(\r\x12\r\n\x05\x65rror\x18\x02 \x01(\t\"#\n\x05Input\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t\x12\x0c\n\x04\x63ode\x18\x02 \x01(\r\"%\n\x06Result\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t\x12\r\n\x05\x65rror\x18\x02 \x01(\t2m\n\x08\x45xecutor\x12\x34\n\x07\x43ompile\x12\x0e.evaluate.Code\x1a\x17.evaluate.CompileResult\"\x00\x12+\n\x04\x43\x61ll\x12\x0f.evaluate.Input\x1a\x10.evaluate.Result\"\x00\x62\x06proto3'
)




_CODE = _descriptor.Descriptor(
  name='Code',
  full_name='evaluate.Code',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='evaluate.Code.code', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='function', full_name='evaluate.Code.function', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=62,
)


_COMPILERESULT = _descriptor.Descriptor(
  name='CompileResult',
  full_name='evaluate.CompileResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='evaluate.CompileResult.id', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='evaluate.CompileResult.error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=64,
  serialized_end=106,
)


_INPUT = _descriptor.Descriptor(
  name='Input',
  full_name='evaluate.Input',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='evaluate.Input.data', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='code', full_name='evaluate.Input.code', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=108,
  serialized_end=143,
)


_RESULT = _descriptor.Descriptor(
  name='Result',
  full_name='evaluate.Result',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='evaluate.Result.data', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='evaluate.Result.error', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=145,
  serialized_end=182,
)

DESCRIPTOR.message_types_by_name['Code'] = _CODE
DESCRIPTOR.message_types_by_name['CompileResult'] = _COMPILERESULT
DESCRIPTOR.message_types_by_name['Input'] = _INPUT
DESCRIPTOR.message_types_by_name['Result'] = _RESULT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Code = _reflection.GeneratedProtocolMessageType('Code', (_message.Message,), {
  'DESCRIPTOR' : _CODE,
  '__module__' : 'exec_pb2'
  # @@protoc_insertion_point(class_scope:evaluate.Code)
  })
_sym_db.RegisterMessage(Code)

CompileResult = _reflection.GeneratedProtocolMessageType('CompileResult', (_message.Message,), {
  'DESCRIPTOR' : _COMPILERESULT,
  '__module__' : 'exec_pb2'
  # @@protoc_insertion_point(class_scope:evaluate.CompileResult)
  })
_sym_db.RegisterMessage(CompileResult)

Input = _reflection.GeneratedProtocolMessageType('Input', (_message.Message,), {
  'DESCRIPTOR' : _INPUT,
  '__module__' : 'exec_pb2'
  # @@protoc_insertion_point(class_scope:evaluate.Input)
  })
_sym_db.RegisterMessage(Input)

Result = _reflection.GeneratedProtocolMessageType('Result', (_message.Message,), {
  'DESCRIPTOR' : _RESULT,
  '__module__' : 'exec_pb2'
  # @@protoc_insertion_point(class_scope:evaluate.Result)
  })
_sym_db.RegisterMessage(Result)



_EXECUTOR = _descriptor.ServiceDescriptor(
  name='Executor',
  full_name='evaluate.Executor',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=184,
  serialized_end=293,
  methods=[
  _descriptor.MethodDescriptor(
    name='Compile',
    full_name='evaluate.Executor.Compile',
    index=0,
    containing_service=None,
    input_type=_CODE,
    output_type=_COMPILERESULT,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Call',
    full_name='evaluate.Executor.Call',
    index=1,
    containing_service=None,
    input_type=_INPUT,
    output_type=_RESULT,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_EXECUTOR)

DESCRIPTOR.services_by_name['Executor'] = _EXECUTOR

# @@protoc_insertion_point(module_scope)
