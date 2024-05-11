// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for google.golang.org/protobuf/reflect/protoreflect, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: google.golang.org/protobuf/reflect/protoreflect (exports: EnumDescriptor,EnumType,EnumNumber,Message,FileDescriptor; functions: )

// Package protoreflect is a stub of google.golang.org/protobuf/reflect/protoreflect, generated by depstubber.
package protoreflect

import ()

type Cardinality int8

func (_ Cardinality) GoString() string {
	return ""
}

func (_ Cardinality) IsValid() bool {
	return false
}

func (_ Cardinality) String() string {
	return ""
}

type Descriptor interface {
	FullName() FullName
	Index() int
	IsPlaceholder() bool
	Name() Name
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	Syntax() Syntax
}

type Enum interface {
	Descriptor() EnumDescriptor
	Number() EnumNumber
	Type() EnumType
}

type EnumDescriptor interface {
	FullName() FullName
	Index() int
	IsPlaceholder() bool
	Name() Name
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ EnumDescriptor)
	ReservedNames() Names
	ReservedRanges() EnumRanges
	Syntax() Syntax
	Values() EnumValueDescriptors
}

type EnumDescriptors interface {
	ByName(_ Name) EnumDescriptor
	Get(_ int) EnumDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type EnumNumber int32

type EnumRanges interface {
	Get(_ int) [2]EnumNumber
	Has(_ EnumNumber) bool
	Len() int
	ProtoInternal(_ interface{})
}

type EnumType interface {
	Descriptor() EnumDescriptor
	New(_ EnumNumber) Enum
}

type EnumValueDescriptor interface {
	FullName() FullName
	Index() int
	IsPlaceholder() bool
	Name() Name
	Number() EnumNumber
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ EnumValueDescriptor)
	Syntax() Syntax
}

type EnumValueDescriptors interface {
	ByName(_ Name) EnumValueDescriptor
	ByNumber(_ EnumNumber) EnumValueDescriptor
	Get(_ int) EnumValueDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type ExtensionDescriptors interface {
	ByName(_ Name) FieldDescriptor
	Get(_ int) FieldDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type ExtensionType interface {
	InterfaceOf(_ Value) interface{}
	IsValidInterface(_ interface{}) bool
	IsValidValue(_ Value) bool
	New() Value
	TypeDescriptor() ExtensionTypeDescriptor
	ValueOf(_ interface{}) Value
	Zero() Value
}

type ExtensionTypeDescriptor interface {
	Cardinality() Cardinality
	ContainingMessage() MessageDescriptor
	ContainingOneof() OneofDescriptor
	Default() Value
	DefaultEnumValue() EnumValueDescriptor
	Descriptor() FieldDescriptor
	Enum() EnumDescriptor
	FullName() FullName
	HasDefault() bool
	HasJSONName() bool
	HasOptionalKeyword() bool
	HasPresence() bool
	Index() int
	IsExtension() bool
	IsList() bool
	IsMap() bool
	IsPacked() bool
	IsPlaceholder() bool
	IsWeak() bool
	JSONName() string
	Kind() Kind
	MapKey() FieldDescriptor
	MapValue() FieldDescriptor
	Message() MessageDescriptor
	Name() Name
	Number() interface{}
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ FieldDescriptor)
	Syntax() Syntax
	Type() ExtensionType
}

type FieldDescriptor interface {
	Cardinality() Cardinality
	ContainingMessage() MessageDescriptor
	ContainingOneof() OneofDescriptor
	Default() Value
	DefaultEnumValue() EnumValueDescriptor
	Enum() EnumDescriptor
	FullName() FullName
	HasDefault() bool
	HasJSONName() bool
	HasOptionalKeyword() bool
	HasPresence() bool
	Index() int
	IsExtension() bool
	IsList() bool
	IsMap() bool
	IsPacked() bool
	IsPlaceholder() bool
	IsWeak() bool
	JSONName() string
	Kind() Kind
	MapKey() FieldDescriptor
	MapValue() FieldDescriptor
	Message() MessageDescriptor
	Name() Name
	Number() interface{}
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ FieldDescriptor)
	Syntax() Syntax
}

type FieldDescriptors interface {
	ByJSONName(_ string) FieldDescriptor
	ByName(_ Name) FieldDescriptor
	ByNumber(_ interface{}) FieldDescriptor
	Get(_ int) FieldDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type FieldNumber int32

type FieldNumbers interface {
	Get(_ int) interface{}
	Has(_ interface{}) bool
	Len() int
	ProtoInternal(_ interface{})
}

type FieldRanges interface {
	Get(_ int) [2]interface{}
	Has(_ interface{}) bool
	Len() int
	ProtoInternal(_ interface{})
}

type FileDescriptor interface {
	Enums() EnumDescriptors
	Extensions() ExtensionDescriptors
	FullName() FullName
	Imports() FileImports
	Index() int
	IsPlaceholder() bool
	Messages() MessageDescriptors
	Name() Name
	Options() ProtoMessage
	Package() FullName
	Parent() Descriptor
	ParentFile() FileDescriptor
	Path() string
	ProtoInternal(_ interface{})
	ProtoType(_ FileDescriptor)
	Services() ServiceDescriptors
	SourceLocations() SourceLocations
	Syntax() Syntax
}

type FileImport struct {
	FileDescriptor FileDescriptor
	IsPublic       bool
	IsWeak         bool
}

func (_ FileImport) Enums() EnumDescriptors {
	return nil
}

func (_ FileImport) Extensions() ExtensionDescriptors {
	return nil
}

func (_ FileImport) FullName() FullName {
	return ""
}

func (_ FileImport) Imports() FileImports {
	return nil
}

func (_ FileImport) Index() int {
	return 0
}

func (_ FileImport) IsPlaceholder() bool {
	return false
}

func (_ FileImport) Messages() MessageDescriptors {
	return nil
}

func (_ FileImport) Name() Name {
	return ""
}

func (_ FileImport) Options() ProtoMessage {
	return nil
}

func (_ FileImport) Package() FullName {
	return ""
}

func (_ FileImport) Parent() Descriptor {
	return nil
}

func (_ FileImport) ParentFile() FileDescriptor {
	return nil
}

func (_ FileImport) Path() string {
	return ""
}

func (_ FileImport) ProtoInternal(_ interface{}) {}

func (_ FileImport) ProtoType(_ FileDescriptor) {}

func (_ FileImport) Services() ServiceDescriptors {
	return nil
}

func (_ FileImport) SourceLocations() SourceLocations {
	return nil
}

func (_ FileImport) Syntax() Syntax {
	return 0
}

type FileImports interface {
	Get(_ int) FileImport
	Len() int
	ProtoInternal(_ interface{})
}

type FullName string

func (_ FullName) Append(_ Name) FullName {
	return ""
}

func (_ FullName) IsValid() bool {
	return false
}

func (_ FullName) Name() Name {
	return ""
}

func (_ FullName) Parent() FullName {
	return ""
}

type Kind int8

func (_ Kind) GoString() string {
	return ""
}

func (_ Kind) IsValid() bool {
	return false
}

func (_ Kind) String() string {
	return ""
}

type List interface {
	Append(_ Value)
	AppendMutable() Value
	Get(_ int) Value
	IsValid() bool
	Len() int
	NewElement() Value
	Set(_ int, _ Value)
	Truncate(_ int)
}

type Map interface {
	Clear(_ MapKey)
	Get(_ MapKey) Value
	Has(_ MapKey) bool
	IsValid() bool
	Len() int
	Mutable(_ MapKey) Value
	NewValue() Value
	Range(_ func(MapKey, Value) bool)
	Set(_ MapKey, _ Value)
}

type MapKey struct {
	DoNotCompare interface{}
}

func (_ MapKey) Bool() bool {
	return false
}

func (_ MapKey) Int() int64 {
	return 0
}

func (_ MapKey) Interface() interface{} {
	return nil
}

func (_ MapKey) IsValid() bool {
	return false
}

func (_ MapKey) String() string {
	return ""
}

func (_ MapKey) Uint() uint64 {
	return 0
}

func (_ MapKey) Value() Value {
	return Value{}
}

type Message interface {
	Clear(_ FieldDescriptor)
	Descriptor() MessageDescriptor
	Get(_ FieldDescriptor) Value
	GetUnknown() RawFields
	Has(_ FieldDescriptor) bool
	Interface() ProtoMessage
	IsValid() bool
	Mutable(_ FieldDescriptor) Value
	New() Message
	NewField(_ FieldDescriptor) Value
	ProtoMethods() *struct {
		NoUnkeyedLiterals interface{}
		Flags             uint64
		Size              func(struct {
			NoUnkeyedLiterals interface{}
			Message           Message
			Flags             byte
		}) struct {
			NoUnkeyedLiterals interface{}
			Size              int
		}
		Marshal func(struct {
			NoUnkeyedLiterals interface{}
			Message           Message
			Buf               []byte
			Flags             byte
		}) (struct {
			NoUnkeyedLiterals interface{}
			Buf               []byte
		}, error)
		Unmarshal func(struct {
			NoUnkeyedLiterals interface{}
			Message           Message
			Buf               []byte
			Flags             byte
			Resolver          interface {
				FindExtensionByName(_ FullName) (ExtensionType, error)
				FindExtensionByNumber(_ FullName, _ interface{}) (ExtensionType, error)
			}
		}) (struct {
			NoUnkeyedLiterals interface{}
			Flags             byte
		}, error)
		Merge func(struct {
			NoUnkeyedLiterals interface{}
			Source            Message
			Destination       Message
		}) struct {
			NoUnkeyedLiterals interface{}
			Flags             byte
		}
		CheckInitialized func(struct {
			NoUnkeyedLiterals interface{}
			Message           Message
		}) (struct {
			NoUnkeyedLiterals interface{}
		}, error)
	}
	Range(_ func(FieldDescriptor, Value) bool)
	Set(_ FieldDescriptor, _ Value)
	SetUnknown(_ RawFields)
	Type() MessageType
	WhichOneof(_ OneofDescriptor) FieldDescriptor
}

type MessageDescriptor interface {
	Enums() EnumDescriptors
	ExtensionRangeOptions(_ int) ProtoMessage
	ExtensionRanges() FieldRanges
	Extensions() ExtensionDescriptors
	Fields() FieldDescriptors
	FullName() FullName
	Index() int
	IsMapEntry() bool
	IsPlaceholder() bool
	Messages() MessageDescriptors
	Name() Name
	Oneofs() OneofDescriptors
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ MessageDescriptor)
	RequiredNumbers() FieldNumbers
	ReservedNames() Names
	ReservedRanges() FieldRanges
	Syntax() Syntax
}

type MessageDescriptors interface {
	ByName(_ Name) MessageDescriptor
	Get(_ int) MessageDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type MessageType interface {
	Descriptor() MessageDescriptor
	New() Message
	Zero() Message
}

type MethodDescriptor interface {
	FullName() FullName
	Index() int
	Input() MessageDescriptor
	IsPlaceholder() bool
	IsStreamingClient() bool
	IsStreamingServer() bool
	Name() Name
	Options() ProtoMessage
	Output() MessageDescriptor
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ MethodDescriptor)
	Syntax() Syntax
}

type MethodDescriptors interface {
	ByName(_ Name) MethodDescriptor
	Get(_ int) MethodDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type Name string

func (_ Name) IsValid() bool {
	return false
}

type Names interface {
	Get(_ int) Name
	Has(_ Name) bool
	Len() int
	ProtoInternal(_ interface{})
}

type OneofDescriptor interface {
	Fields() FieldDescriptors
	FullName() FullName
	Index() int
	IsPlaceholder() bool
	IsSynthetic() bool
	Name() Name
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ OneofDescriptor)
	Syntax() Syntax
}

type OneofDescriptors interface {
	ByName(_ Name) OneofDescriptor
	Get(_ int) OneofDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type ProtoMessage interface {
	ProtoReflect() Message
}

type RawFields []byte

func (_ RawFields) IsValid() bool {
	return false
}

type ServiceDescriptor interface {
	FullName() FullName
	Index() int
	IsPlaceholder() bool
	Methods() MethodDescriptors
	Name() Name
	Options() ProtoMessage
	Parent() Descriptor
	ParentFile() FileDescriptor
	ProtoInternal(_ interface{})
	ProtoType(_ ServiceDescriptor)
	Syntax() Syntax
}

type ServiceDescriptors interface {
	ByName(_ Name) ServiceDescriptor
	Get(_ int) ServiceDescriptor
	Len() int
	ProtoInternal(_ interface{})
}

type SourceLocation struct {
	Path                    SourcePath
	StartLine               int
	StartColumn             int
	EndLine                 int
	EndColumn               int
	LeadingDetachedComments []string
	LeadingComments         string
	TrailingComments        string
}

type SourceLocations interface {
	Get(_ int) SourceLocation
	Len() int
	ProtoInternal(_ interface{})
}

type SourcePath []int32

type Syntax int8

func (_ Syntax) GoString() string {
	return ""
}

func (_ Syntax) IsValid() bool {
	return false
}

func (_ Syntax) String() string {
	return ""
}

type Value struct {
	DoNotCompare interface{}
}

func (_ Value) Bool() bool {
	return false
}

func (_ Value) Bytes() []byte {
	return nil
}

func (_ Value) Enum() EnumNumber {
	return 0
}

func (_ Value) Float() float64 {
	return 0
}

func (_ Value) Int() int64 {
	return 0
}

func (_ Value) Interface() interface{} {
	return nil
}

func (_ Value) IsValid() bool {
	return false
}

func (_ Value) List() List {
	return nil
}

func (_ Value) Map() Map {
	return nil
}

func (_ Value) MapKey() MapKey {
	return MapKey{}
}

func (_ Value) Message() Message {
	return nil
}

func (_ Value) String() string {
	return ""
}

func (_ Value) Uint() uint64 {
	return 0
}
