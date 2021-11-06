package types

import (
	"fmt"
	"math/big"

	"github.com/centrifuge/go-substrate-rpc-client/v3/scale"
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

type Si1LookupTypeId CompactU32

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
	IsBool bool
	IsChar bool
	IsStr  bool
	IsU8   bool
	IsU16  bool
	IsU32  bool
	IsU64  bool
	IsU128 bool
	IsU256 bool
	IsI8   bool
	IsI16  bool
	IsI32  bool
	IsI64  bool
	IsI128 bool
	IsI256 bool
}

func (m Si1TypeDefPrimitive) Encode(encoder scale.Encoder) error {
	var err error
	switch {
	case m.IsBool:
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
	case m.IsChar:
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
	case m.IsStr:
		err = encoder.PushByte(2)
		if err != nil {
			return err
		}
	case m.IsU8:
		err = encoder.PushByte(3)
		if err != nil {
			return err
		}
	case m.IsU16:
		err = encoder.PushByte(4)
		if err != nil {
			return err
		}
	case m.IsU32:
		err = encoder.PushByte(5)
		if err != nil {
			return err
		}
	case m.IsU64:
		err = encoder.PushByte(6)
		if err != nil {
			return err
		}
	case m.IsU128:
		err = encoder.PushByte(7)
		if err != nil {
			return err
		}
	case m.IsU256:
		err = encoder.PushByte(8)
		if err != nil {
			return err
		}
	case m.IsI8:
		err = encoder.PushByte(9)
		if err != nil {
			return err
		}
	case m.IsI16:
		err = encoder.PushByte(10)
		if err != nil {
			return err
		}
	case m.IsI32:
		err = encoder.PushByte(11)
		if err != nil {
			return err
		}
	case m.IsI64:
		err = encoder.PushByte(12)
		if err != nil {
			return err
		}
	case m.IsI128:
		err = encoder.PushByte(13)
		if err != nil {
			return err
		}
	case m.IsI256:
		err = encoder.PushByte(14)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid variant for Si1TypeDefPrimitive")
	}

	return nil
}

func (m *Si1TypeDefPrimitive) Decode(decoder scale.Decoder) error {
	tag, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch tag {
	case 0:
		m.IsBool = true
	case 1:
		m.IsChar = true
	case 2:
		m.IsStr = true
	case 3:
		m.IsU8 = true
	case 4:
		m.IsU16 = true
	case 5:
		m.IsU32 = true
	case 6:
		m.IsU64 = true
	case 7:
		m.IsU128 = true
	case 8:
		m.IsU256 = true
	case 9:
		m.IsI8 = true
	case 10:
		m.IsI16 = true
	case 11:
		m.IsI32 = true
	case 12:
		m.IsI64 = true
	case 13:
		m.IsI128 = true
	case 14:
		m.IsI256 = true
	default:
		err = fmt.Errorf("invalid variant for Si1TypeDef: %v", tag)
	}

	return err
}

type Si1TypeDefCompact struct {
	Type Si1LookupTypeId
}

type Si1TypeDefBitSequence struct {
	BitStoreType Si1LookupTypeId
	BitOrderType Si1LookupTypeId
}

type Si1TypeDef struct {
	IsComposite   bool
	AsComposite   Si1TypeDefComposite
	IsVariant     bool
	AsVariant     Si1TypeDefVariant
	IsSequence    bool
	AsSequence    Si1TypeDefSequence
	IsArray       bool
	AsArray       Si1TypeDefArray
	IsTuple       bool
	AsTuple       Si1TypeDefTuple
	IsPrimitive   bool
	AsPrimitive   Si1TypeDefPrimitive
	IsCompact     bool
	AsCompact     Si1TypeDefCompact
	IsBitSequence bool
	AsBitSequence Si1TypeDefBitSequence
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
