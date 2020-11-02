# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: tendermint/types/params.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from gogoproto import gogo_pb2 as gogoproto_dot_gogo__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
    name="tendermint/types/params.proto",
    package="tendermint.types",
    syntax="proto3",
    serialized_options=b"Z7github.com/tendermint/tendermint/proto/tendermint/types\250\342\036\001",
    create_key=_descriptor._internal_create_key,
    serialized_pb=b'\n\x1dtendermint/types/params.proto\x12\x10tendermint.types\x1a\x14gogoproto/gogo.proto\x1a\x1egoogle/protobuf/duration.proto"\xf3\x01\n\x0f\x43onsensusParams\x12\x32\n\x05\x62lock\x18\x01 \x01(\x0b\x32\x1d.tendermint.types.BlockParamsB\x04\xc8\xde\x1f\x00\x12\x38\n\x08\x65vidence\x18\x02 \x01(\x0b\x32 .tendermint.types.EvidenceParamsB\x04\xc8\xde\x1f\x00\x12:\n\tvalidator\x18\x03 \x01(\x0b\x32!.tendermint.types.ValidatorParamsB\x04\xc8\xde\x1f\x00\x12\x36\n\x07version\x18\x04 \x01(\x0b\x32\x1f.tendermint.types.VersionParamsB\x04\xc8\xde\x1f\x00"G\n\x0b\x42lockParams\x12\x11\n\tmax_bytes\x18\x01 \x01(\x03\x12\x0f\n\x07max_gas\x18\x02 \x01(\x03\x12\x14\n\x0ctime_iota_ms\x18\x03 \x01(\x03"~\n\x0e\x45videnceParams\x12\x1a\n\x12max_age_num_blocks\x18\x01 \x01(\x03\x12=\n\x10max_age_duration\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationB\x08\xc8\xde\x1f\x00\x98\xdf\x1f\x01\x12\x11\n\tmax_bytes\x18\x03 \x01(\x03"2\n\x0fValidatorParams\x12\x15\n\rpub_key_types\x18\x01 \x03(\t:\x08\xb8\xa0\x1f\x01\xe8\xa0\x1f\x01".\n\rVersionParams\x12\x13\n\x0b\x61pp_version\x18\x01 \x01(\x04:\x08\xb8\xa0\x1f\x01\xe8\xa0\x1f\x01">\n\x0cHashedParams\x12\x17\n\x0f\x62lock_max_bytes\x18\x01 \x01(\x03\x12\x15\n\rblock_max_gas\x18\x02 \x01(\x03\x42=Z7github.com/tendermint/tendermint/proto/tendermint/types\xa8\xe2\x1e\x01\x62\x06proto3',
    dependencies=[
        gogoproto_dot_gogo__pb2.DESCRIPTOR,
        google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,
    ],
)


_CONSENSUSPARAMS = _descriptor.Descriptor(
    name="ConsensusParams",
    full_name="tendermint.types.ConsensusParams",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="block",
            full_name="tendermint.types.ConsensusParams.block",
            index=0,
            number=1,
            type=11,
            cpp_type=10,
            label=1,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="evidence",
            full_name="tendermint.types.ConsensusParams.evidence",
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=1,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="validator",
            full_name="tendermint.types.ConsensusParams.validator",
            index=2,
            number=3,
            type=11,
            cpp_type=10,
            label=1,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="version",
            full_name="tendermint.types.ConsensusParams.version",
            index=3,
            number=4,
            type=11,
            cpp_type=10,
            label=1,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=106,
    serialized_end=349,
)


_BLOCKPARAMS = _descriptor.Descriptor(
    name="BlockParams",
    full_name="tendermint.types.BlockParams",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="max_bytes",
            full_name="tendermint.types.BlockParams.max_bytes",
            index=0,
            number=1,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="max_gas",
            full_name="tendermint.types.BlockParams.max_gas",
            index=1,
            number=2,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="time_iota_ms",
            full_name="tendermint.types.BlockParams.time_iota_ms",
            index=2,
            number=3,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=351,
    serialized_end=422,
)


_EVIDENCEPARAMS = _descriptor.Descriptor(
    name="EvidenceParams",
    full_name="tendermint.types.EvidenceParams",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="max_age_num_blocks",
            full_name="tendermint.types.EvidenceParams.max_age_num_blocks",
            index=0,
            number=1,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="max_age_duration",
            full_name="tendermint.types.EvidenceParams.max_age_duration",
            index=1,
            number=2,
            type=11,
            cpp_type=10,
            label=1,
            has_default_value=False,
            default_value=None,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=b"\310\336\037\000\230\337\037\001",
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="max_bytes",
            full_name="tendermint.types.EvidenceParams.max_bytes",
            index=2,
            number=3,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=424,
    serialized_end=550,
)


_VALIDATORPARAMS = _descriptor.Descriptor(
    name="ValidatorParams",
    full_name="tendermint.types.ValidatorParams",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="pub_key_types",
            full_name="tendermint.types.ValidatorParams.pub_key_types",
            index=0,
            number=1,
            type=9,
            cpp_type=9,
            label=3,
            has_default_value=False,
            default_value=[],
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\270\240\037\001\350\240\037\001",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=552,
    serialized_end=602,
)


_VERSIONPARAMS = _descriptor.Descriptor(
    name="VersionParams",
    full_name="tendermint.types.VersionParams",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="app_version",
            full_name="tendermint.types.VersionParams.app_version",
            index=0,
            number=1,
            type=4,
            cpp_type=4,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=b"\270\240\037\001\350\240\037\001",
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=604,
    serialized_end=650,
)


_HASHEDPARAMS = _descriptor.Descriptor(
    name="HashedParams",
    full_name="tendermint.types.HashedParams",
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    create_key=_descriptor._internal_create_key,
    fields=[
        _descriptor.FieldDescriptor(
            name="block_max_bytes",
            full_name="tendermint.types.HashedParams.block_max_bytes",
            index=0,
            number=1,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
        _descriptor.FieldDescriptor(
            name="block_max_gas",
            full_name="tendermint.types.HashedParams.block_max_gas",
            index=1,
            number=2,
            type=3,
            cpp_type=2,
            label=1,
            has_default_value=False,
            default_value=0,
            message_type=None,
            enum_type=None,
            containing_type=None,
            is_extension=False,
            extension_scope=None,
            serialized_options=None,
            file=DESCRIPTOR,
            create_key=_descriptor._internal_create_key,
        ),
    ],
    extensions=[],
    nested_types=[],
    enum_types=[],
    serialized_options=None,
    is_extendable=False,
    syntax="proto3",
    extension_ranges=[],
    oneofs=[],
    serialized_start=652,
    serialized_end=714,
)

_CONSENSUSPARAMS.fields_by_name["block"].message_type = _BLOCKPARAMS
_CONSENSUSPARAMS.fields_by_name["evidence"].message_type = _EVIDENCEPARAMS
_CONSENSUSPARAMS.fields_by_name["validator"].message_type = _VALIDATORPARAMS
_CONSENSUSPARAMS.fields_by_name["version"].message_type = _VERSIONPARAMS
_EVIDENCEPARAMS.fields_by_name[
    "max_age_duration"
].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
DESCRIPTOR.message_types_by_name["ConsensusParams"] = _CONSENSUSPARAMS
DESCRIPTOR.message_types_by_name["BlockParams"] = _BLOCKPARAMS
DESCRIPTOR.message_types_by_name["EvidenceParams"] = _EVIDENCEPARAMS
DESCRIPTOR.message_types_by_name["ValidatorParams"] = _VALIDATORPARAMS
DESCRIPTOR.message_types_by_name["VersionParams"] = _VERSIONPARAMS
DESCRIPTOR.message_types_by_name["HashedParams"] = _HASHEDPARAMS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ConsensusParams = _reflection.GeneratedProtocolMessageType(
    "ConsensusParams",
    (_message.Message,),
    {
        "DESCRIPTOR": _CONSENSUSPARAMS,
        "__module__": "tendermint.types.params_pb2"
        # @@protoc_insertion_point(class_scope:tendermint.types.ConsensusParams)
    },
)
_sym_db.RegisterMessage(ConsensusParams)

BlockParams = _reflection.GeneratedProtocolMessageType(
    "BlockParams",
    (_message.Message,),
    {
        "DESCRIPTOR": _BLOCKPARAMS,
        "__module__": "tendermint.types.params_pb2"
        # @@protoc_insertion_point(class_scope:tendermint.types.BlockParams)
    },
)
_sym_db.RegisterMessage(BlockParams)

EvidenceParams = _reflection.GeneratedProtocolMessageType(
    "EvidenceParams",
    (_message.Message,),
    {
        "DESCRIPTOR": _EVIDENCEPARAMS,
        "__module__": "tendermint.types.params_pb2"
        # @@protoc_insertion_point(class_scope:tendermint.types.EvidenceParams)
    },
)
_sym_db.RegisterMessage(EvidenceParams)

ValidatorParams = _reflection.GeneratedProtocolMessageType(
    "ValidatorParams",
    (_message.Message,),
    {
        "DESCRIPTOR": _VALIDATORPARAMS,
        "__module__": "tendermint.types.params_pb2"
        # @@protoc_insertion_point(class_scope:tendermint.types.ValidatorParams)
    },
)
_sym_db.RegisterMessage(ValidatorParams)

VersionParams = _reflection.GeneratedProtocolMessageType(
    "VersionParams",
    (_message.Message,),
    {
        "DESCRIPTOR": _VERSIONPARAMS,
        "__module__": "tendermint.types.params_pb2"
        # @@protoc_insertion_point(class_scope:tendermint.types.VersionParams)
    },
)
_sym_db.RegisterMessage(VersionParams)

HashedParams = _reflection.GeneratedProtocolMessageType(
    "HashedParams",
    (_message.Message,),
    {
        "DESCRIPTOR": _HASHEDPARAMS,
        "__module__": "tendermint.types.params_pb2"
        # @@protoc_insertion_point(class_scope:tendermint.types.HashedParams)
    },
)
_sym_db.RegisterMessage(HashedParams)


DESCRIPTOR._options = None
_CONSENSUSPARAMS.fields_by_name["block"]._options = None
_CONSENSUSPARAMS.fields_by_name["evidence"]._options = None
_CONSENSUSPARAMS.fields_by_name["validator"]._options = None
_CONSENSUSPARAMS.fields_by_name["version"]._options = None
_EVIDENCEPARAMS.fields_by_name["max_age_duration"]._options = None
_VALIDATORPARAMS._options = None
_VERSIONPARAMS._options = None
# @@protoc_insertion_point(module_scope)
