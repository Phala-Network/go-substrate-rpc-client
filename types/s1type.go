package types

import (
	"fmt"
	"math/big"

	"github.com/Phala-Network/go-substrate-rpc-client/v3/scale"
)

type CompactU32 U32

// Encode implements encoding for CompactU32, which just unwraps the bytes of CompactU32
func (cu32 CompactU32) Encode(encoder scale.Encoder) error {
	return encoder.EncodeUintCompact(*big.NewInt(0).SetUint64(uint64(cu32)))
}

// Decode implements decoding for CompactU32, which just wraps the bytes in CompactU32
func (cu32 *CompactU32) Decode(decoder scale.Decoder) error {
	u, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}
	*cu32 = CompactU32(u.Uint64())
	return nil
}

type Si1LookupTypeId big.Int

func NewSi1LookupTypeID(value *big.Int) Si1LookupTypeId {
	return Si1LookupTypeId(*value)
}

func NewSi1LookupTypeIDFromUInt(value uint64) Si1LookupTypeId {
	return NewSi1LookupTypeID(new(big.Int).SetUint64(value))
}

func (d *Si1LookupTypeId) Int64() int64 {
	i := big.Int(*d)
	return i.Int64()
}

func (d *Si1LookupTypeId) Decode(decoder scale.Decoder) error {
	ui, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}

	*d = Si1LookupTypeId(*ui)
	return nil
}

func (d Si1LookupTypeId) Encode(encoder scale.Encoder) error {
	err := encoder.EncodeUintCompact(big.Int(d))
	if err != nil {
		return err
	}
	return nil
}

type Si1Path []Text

type Si1TypeParameter struct {
	Name    Text
	HasType bool
	Type    Si1LookupTypeId
}

func (m Si1TypeParameter) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.Name)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.HasType)
	if err != nil {
		return err
	}

	if m.HasType {
		err = encoder.Encode(m.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Si1TypeParameter) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.HasType)
	if err != nil {
		return err
	}

	if m.HasType {
		err = decoder.Decode(&m.Type)
		if err != nil {
			return err
		}
	}

	return nil
}

type Si1Field struct {
	HasName     bool
	Name        Text
	Type        Si1LookupTypeId
	HasTypeName bool
	TypeName    Text
	Docs        []Text
}

func (m Si1Field) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.HasName)
	if err != nil {
		return err
	}

	if m.HasName {
		err = encoder.Encode(m.Name)
		if err != nil {
			return err
		}
	}

	err = encoder.Encode(m.Type)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.HasTypeName)
	if err != nil {
		return err
	}

	if m.HasTypeName {
		err = encoder.Encode(m.TypeName)
		if err != nil {
			return err
		}
	}

	err = encoder.Encode(m.Docs)
	if err != nil {
		return err
	}

	return nil
}

func (m *Si1Field) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.HasName)
	if err != nil {
		return err
	}

	if m.HasName {
		err = decoder.Decode(&m.Name)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.Type)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.HasTypeName)
	if err != nil {
		return err
	}

	if m.HasTypeName {
		err = decoder.Decode(&m.TypeName)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.Docs)
	if err != nil {
		return err
	}

	return nil
}

type Si1Variant struct {
	Name   Text
	Fields []Si1Field
	Index  U8
	Docs   []Text
}

func (m Si1Variant) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.Name)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.Fields)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.Index)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.Docs)
	if err != nil {
		return err
	}

	return nil
}

func (m *Si1Variant) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Fields)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Index)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Docs)
	if err != nil {
		return err
	}

	return nil
}

type Si1TypeDefComposite struct {
	Fields []Si1Field
}

type Si1TypeDefVariant struct {
	Variants []Si1Variant
}

type Si1TypeDefSequence struct {
	Type Si1LookupTypeId
}

type Si1TypeDefArray struct {
	Len  U32
	Type Si1LookupTypeId
}

type Si1TypeDefTuple []Si1LookupTypeId

type Si1TypeDefPrimitive struct {
	Value string
}

func (d *Si1TypeDefPrimitive) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}
	switch b {
	case 0:
		d.Value = "Bool"
	case 1:
		d.Value = "Char"
	case 2:
		d.Value = "Str"
	case 3:
		d.Value = "U8"
	case 4:
		d.Value = "U16"
	case 5:
		d.Value = "U32"
	case 6:
		d.Value = "U64"
	case 7:
		d.Value = "U128"
	case 8:
		d.Value = "U256"
	case 9:
		d.Value = "I8"
	case 10:
		d.Value = "I16"
	case 11:
		d.Value = "I32"
	case 12:
		d.Value = "I64"
	case 13:
		d.Value = "I128"
	case 14:
		d.Value = "I256"
	default:
		return fmt.Errorf("Si1TypeDefPrimitive do not support this type: %d", b)
	}
	return nil
}

func (d Si1TypeDefPrimitive) Encode(encoder scale.Encoder) error {
	switch d.Value {
	case "Bool":
		return encoder.PushByte(0)
	case "Char":
		return encoder.PushByte(1)
	case "Str":
		return encoder.PushByte(2)
	case "U8":
		return encoder.PushByte(3)
	case "U16":
		return encoder.PushByte(4)
	case "U32":
		return encoder.PushByte(5)
	case "U64":
		return encoder.PushByte(6)
	case "U128":
		return encoder.PushByte(7)
	case "U256":
		return encoder.PushByte(8)
	case "I8":
		return encoder.PushByte(9)
	case "I16":
		return encoder.PushByte(10)
	case "I32":
		return encoder.PushByte(11)
	case "I64":
		return encoder.PushByte(12)
	case "I128":
		return encoder.PushByte(13)
	case "I256":
		return encoder.PushByte(14)
	default:
		//TODO(nuno): Not sure what to do
		return nil
	}
}

type Si1TypeDefCompact struct {
	Type Si1LookupTypeId
}

type Si1TypeDefBitSequence struct {
	BitStoreType Si1LookupTypeId
	BitOrderType Si1LookupTypeId
}

type Si1TypeDef struct {
	IsComposite          bool
	AsComposite          Si1TypeDefComposite
	IsVariant            bool
	AsVariant            Si1TypeDefVariant
	IsSequence           bool
	AsSequence           Si1TypeDefSequence
	IsArray              bool
	AsArray              Si1TypeDefArray
	IsTuple              bool
	AsTuple              Si1TypeDefTuple
	IsPrimitive          bool
	AsPrimitive          Si1TypeDefPrimitive
	IsCompact            bool
	AsCompact            Si1TypeDefCompact
	IsBitSequence        bool
	AsBitSequence        Si1TypeDefBitSequence
	IsHistoricMetaCompat bool
	AsHistoricMetaCompat Type
}

func (m Si1TypeDef) Encode(encoder scale.Encoder) error {
	var err error
	switch {
	case m.IsComposite:
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsComposite)
		if err != nil {
			return err
		}
	case m.IsVariant:
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsVariant)
		if err != nil {
			return err
		}
	case m.IsSequence:
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsSequence)
		if err != nil {
			return err
		}
	case m.IsTuple:
		err = encoder.PushByte(3)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsTuple)
		if err != nil {
			return err
		}
	case m.IsArray:
		err = encoder.PushByte(4)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsArray)
		if err != nil {
			return err
		}
	case m.IsPrimitive:
		err = encoder.PushByte(5)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsPrimitive)
		if err != nil {
			return err
		}
	case m.IsCompact:
		err = encoder.PushByte(6)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsCompact)
		if err != nil {
			return err
		}
	case m.IsBitSequence:
		err = encoder.PushByte(7)
		if err != nil {
			return err
		}
		err = encoder.Encode(m.AsBitSequence)
		if err != nil {
			return err
		}
	case m.IsHistoricMetaCompat:
		err := encoder.PushByte(8)
		if err != nil {
			return err
		}
		m.IsHistoricMetaCompat = true
		return encoder.Encode(&m.AsHistoricMetaCompat)
	default:
		return fmt.Errorf("invalid variant for Si1TypeDef")
	}

	return nil
}

func (m *Si1TypeDef) Decode(decoder scale.Decoder) error {
	tag, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch tag {
	case 0:
		m.IsComposite = true
		err = decoder.Decode(&m.AsComposite)
	case 1:
		m.IsVariant = true
		err = decoder.Decode(&m.AsVariant)
	case 2:
		m.IsSequence = true
		err = decoder.Decode(&m.AsSequence)
	case 3:
		m.IsArray = true
		err = decoder.Decode(&m.AsArray)
	case 4:
		m.IsTuple = true
		err = decoder.Decode(&m.AsTuple)
	case 5:
		m.IsPrimitive = true
		err = decoder.Decode(&m.AsPrimitive)
	case 6:
		m.IsCompact = true
		err = decoder.Decode(&m.AsCompact)
	case 7:
		m.IsBitSequence = true
		err = decoder.Decode(&m.AsBitSequence)
	case 8:
		m.IsHistoricMetaCompat = true
		return decoder.Decode(&m.AsHistoricMetaCompat)
	default:
		err = fmt.Errorf("invalid variant for Si1TypeDef: %v", tag)
	}

	return err
}

type Si1Type struct {
	Path   Si1Path
	Params []Si1TypeParameter
	Def    Si1TypeDef
	Docs   []Text
}

func (d *Si1Type) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Path)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Params)
	if err != nil {
		return err
	}
	err = decoder.Decode(&d.Def)
	if err != nil {
		return err
	}
	return decoder.Decode(&d.Docs)
}
