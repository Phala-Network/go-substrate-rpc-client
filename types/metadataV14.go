package types

import (
	"fmt"
	"hash"
	"strings"

	"github.com/Phala-Network/go-substrate-rpc-client/v3/scale"
	"github.com/Phala-Network/go-substrate-rpc-client/v3/xxhash"
)

type PortableType struct {
	Id   Si1LookupTypeId
	Type Si1Type
}

func (d *PortableType) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&d.Id)
	if err != nil {
		return fmt.Errorf("decode Si1LookupTypeID error: %v", err)
	}

	return decoder.Decode(&d.Type)
}

type PortableRegistryV14 struct {
	Types []PortableType
}

type StorageEntryModifierV14 struct {
	IsOptional bool
	IsDefault  bool
}

func (m StorageEntryModifierV14) Encode(encoder scale.Encoder) error {
	var err error
	switch {
	case m.IsOptional:
		err = encoder.PushByte(0)
		if err != nil {
			return err
		}
	case m.IsDefault:
		err = encoder.PushByte(1)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid variant for Si1TypeDefPrimitive")
	}

	return nil
}

func (m *StorageEntryModifierV14) Decode(decoder scale.Decoder) error {
	tag, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch tag {
	case 0:
		m.IsOptional = true
	case 1:
		m.IsDefault = true
	default:
		err = fmt.Errorf("invalid variant for Si1TypeDef: %v", tag)
	}

	return err
}

type StorageEntryTypeMap struct {
	Hashers []StorageHasherV10
	Key     Si1LookupTypeId
	Value   Si1LookupTypeId
}

type StorageEntryTypeV14 struct {
	IsPlain bool
	AsPlain Si1LookupTypeId
	IsMap   bool
	AsMap   StorageEntryTypeMap
}

func (m StorageEntryTypeV14) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.IsPlain)
	if err != nil {
		return err
	}

	if m.IsPlain {
		err = encoder.Encode(m.AsPlain)
		if err != nil {
			return err
		}
	}

	err = encoder.Encode(m.IsMap)
	if err != nil {
		return err
	}

	if m.IsMap {
		err = encoder.Encode(m.AsMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *StorageEntryTypeV14) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.IsPlain)
	if err != nil {
		return err
	}

	if m.IsPlain {
		err = decoder.Decode(&m.AsPlain)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.IsMap)
	if err != nil {
		return err
	}

	if m.IsMap {
		err = decoder.Decode(&m.AsMap)
		if err != nil {
			return err
		}
	}

	return err
}

type StorageEntryMetadataV14 struct {
	Name     Text
	Modifier StorageEntryModifierV14
	Type     StorageEntryTypeV14
	Fallback Bytes
	Docs     []Text
}

func (storage *StorageEntryMetadataV14) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&storage.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&storage.Modifier)
	if err != nil {
		return err
	}

	err = decoder.Decode(&storage.Type)
	if err != nil {
		return err
	}

	err = decoder.Decode(&storage.Fallback)
	if err != nil {
		return err
	}

	return decoder.Decode(&storage.Docs)
}

func (s StorageEntryMetadataV14) IsPlain() bool {
	return s.Type.IsPlain
}

func (s StorageEntryMetadataV14) IsMap() bool {
	return s.Type.IsMap
}

func (s StorageEntryMetadataV14) IsDoubleMap() bool {
	return false
}

func (s StorageEntryMetadataV14) IsNMap() bool {
	return false
}

func (s StorageEntryMetadataV14) Hasher() (hash.Hash, error) {
	return xxhash.New128(nil), nil
}

func (s StorageEntryMetadataV14) Hasher2() (hash.Hash, error) {
	return nil, fmt.Errorf("StorageEntryMetadataV14 does not implement Hasher2()")
}

func (s StorageEntryMetadataV14) Hashers() ([]hash.Hash, error) {
	var hashes []hash.Hash
	if s.Type.IsMap {
		for _, hasher := range s.Type.AsMap.Hashers {
			h, err := hasher.HashFunc()
			if err != nil {
				return nil, err
			}
			hashes = append(hashes, h)
		}
	}
	return hashes, nil
}

type StorageMetadataV14 struct {
	Prefix  Text
	Entries []StorageEntryMetadataV14
}

func (storage *StorageMetadataV14) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&storage.Prefix)
	if err != nil {
		return err
	}
	return decoder.Decode(&storage.Entries)
}

type PalletCallMetadataV14 struct {
	Type Si1LookupTypeId
}

type EventMetadataV14 struct {
	Type Si1LookupTypeId
}

type PalletConstantMetadataV14 struct {
	Name  Text
	Type  Si1LookupTypeId
	Value Bytes
	Docs  []Text
}

type PalletErrorMetadataV14 struct {
	Type Si1LookupTypeId
}

type PalletMetadataV14 struct {
	Name       Text
	HasStorage bool
	Storage    StorageMetadataV14
	HasCalls   bool
	Calls      PalletCallMetadataV14
	HasEvents  bool
	Events     EventMetadataV14
	Constants  []PalletConstantMetadataV14
	HasErrors  bool
	Errors     []PalletErrorMetadataV14
	Index      uint8
}

func (m *PalletMetadataV14) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Name)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.HasStorage)
	if err != nil {
		return err
	}

	if m.HasStorage {
		err = decoder.Decode(&m.Storage)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.HasCalls)
	if err != nil {
		return err
	}

	if m.HasCalls {
		err = decoder.Decode(&m.Calls)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.HasEvents)
	if err != nil {
		return err
	}

	if m.HasEvents {
		err = decoder.Decode(&m.Events)
		if err != nil {
			return err
		}
	}

	err = decoder.Decode(&m.Constants)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.HasErrors)
	if err != nil {
		return err
	}

	if m.HasErrors {
		err = decoder.Decode(&m.Errors)
		if err != nil {
			return err
		}
	}

	return decoder.Decode(&m.Index)
}

func (m PalletMetadataV14) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.Name)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.HasStorage)
	if err != nil {
		return err
	}

	if m.HasStorage {
		err = encoder.Encode(m.Storage)
		if err != nil {
			return err
		}
	}

	err = encoder.Encode(m.HasCalls)
	if err != nil {
		return err
	}

	if m.HasCalls {
		err = encoder.Encode(m.Calls)
		if err != nil {
			return err
		}
	}

	err = encoder.Encode(m.HasEvents)
	if err != nil {
		return err
	}

	if m.HasEvents {
		err = encoder.Encode(m.Events)
		if err != nil {
			return err
		}
	}

	err = encoder.Encode(m.Constants)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.HasErrors)
	if err != nil {
		return err
	}

	if m.HasErrors {
		err = encoder.Encode(m.Errors)
		if err != nil {
			return err
		}
	}

	return encoder.Encode(m.Index)
}

func (m *PalletMetadataV14) FindConstantValue(constant Text) ([]byte, error) {
	for _, cons := range m.Constants {
		if cons.Name == constant {
			return cons.Value, nil
		}
	}
	return nil, fmt.Errorf("could not find constant %s", constant)
}

type SignedExtensionMetadataV14 struct {
	Identifier       Text
	Type             Si1LookupTypeId
	AdditionalSigned Si1LookupTypeId
}

type ExtrinsicMetadataV14 struct {
	Type             Si1LookupTypeId
	Version          U8
	SignedExtensions []SignedExtensionMetadataV14
}

type MetadataV14 struct {
	Lookup    PortableRegistryV14
	Pallets   []PalletMetadataV14
	Extrinsic ExtrinsicMetadataV14
	Type      Si1LookupTypeId
}

func (m *MetadataV14) Decode(decoder scale.Decoder) error {
	err := decoder.Decode(&m.Lookup)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Pallets)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Extrinsic)
	if err != nil {
		return err
	}

	err = decoder.Decode(&m.Type)
	if err != nil {
		return err
	}

	return nil
}

func (m MetadataV14) Encode(encoder scale.Encoder) error {
	err := encoder.Encode(m.Lookup)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.Pallets)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.Extrinsic)
	if err != nil {
		return err
	}

	err = encoder.Encode(m.Type)
	if err != nil {
		return err
	}

	return nil
}

func (m *MetadataV14) FindCallIndex(call string) (CallIndex, error) {
	s := strings.Split(call, ".")
	for _, mod := range m.Pallets {
		if !mod.HasCalls {
			continue
		}
		if string(mod.Name) != s[0] {
			continue
		}
		callType := mod.Calls.Type
		for _, lookUp := range m.Lookup.Types {
			if lookUp.Id.Int64() == callType.Int64() {
				if len(lookUp.Type.Def.AsVariant.Variants) > 0 {
					for _, vars := range lookUp.Type.Def.AsVariant.Variants {
						if string(vars.Name) == s[1] {
							return CallIndex{uint8(mod.Index), uint8(vars.Index)}, nil
						}
					}
				}
			}
		}
	}
	return CallIndex{}, fmt.Errorf("module %v not found in metadata for call %v", s[0], call)
}

func (m *MetadataV14) FindEventNamesForEventID(eventID EventID) (Text, Text, error) {
	// for _, mod := range m.Pallets {
	// 	if !mod.HasEvents {
	// 		continue
	// 	}
	// 	if mod.Index != eventID[0] {
	// 		continue
	// 	}
	// 	if int(eventID[1]) >= len(mod.Events) {
	// 		return "", "", fmt.Errorf("event index %v for module %v out of range", eventID[1], mod.Name)
	// 	}
	// 	return mod.Name, mod.Events[eventID[1]].Name, nil
	// }
	return "", "", fmt.Errorf("module index %v out of range", eventID[0])
}

func (m *MetadataV14) FindStorageEntryMetadata(module string, fn string) (StorageEntryMetadata, error) {
	for _, mod := range m.Pallets {
		if !mod.HasStorage {
			continue
		}
		if string(mod.Storage.Prefix) != module {
			continue
		}
		for _, s := range mod.Storage.Entries {
			if string(s.Name) != fn {
				continue
			}

			return s, nil
		}
		return nil, fmt.Errorf("storage %v not found within module %v", fn, module)
	}
	return nil, fmt.Errorf("module %v not found in metadata", module)
}

func (m *MetadataV14) FindConstantValue(module Text, constant Text) ([]byte, error) {
	// for _, mod := range m.Pallets {
	// 	if mod.Name == module {
	// 		value, err := mod.FindConstantValue(constant)
	// 		if err == nil {
	// 			return value, nil
	// 		}
	// 	}
	// }
	return nil, fmt.Errorf("could not find constant %s.%s", module, constant)
}

func (m *MetadataV14) ExistsModuleMetadata(module string) bool {
	for _, mod := range m.Pallets {
		if string(mod.Name) == module {
			return true
		}
	}
	return false
}
