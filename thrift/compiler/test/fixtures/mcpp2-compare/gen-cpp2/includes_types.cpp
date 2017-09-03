/**
 * Autogenerated by Thrift
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "src/gen-cpp2/includes_types.h"
#include "src/gen-cpp2/includes_types.tcc"

#include <algorithm>
#include <folly/Indestructible.h>

#include "src/gen-cpp2/includes_data.h"

namespace a { namespace different { namespace ns {

const _AnEnum_EnumMapFactory::ValuesToNamesMapType _AnEnum_VALUES_TO_NAMES = _AnEnum_EnumMapFactory::makeValuesToNamesMap();
const _AnEnum_EnumMapFactory::NamesToValuesMapType _AnEnum_NAMES_TO_VALUES = _AnEnum_EnumMapFactory::makeNamesToValuesMap();

}}} // a::different::ns
namespace std {

} // std
namespace apache { namespace thrift {

template <> const std::size_t TEnumTraits< ::a::different::ns::AnEnum>::size = 2;
template <> const folly::Range<const  ::a::different::ns::AnEnum*> TEnumTraits< ::a::different::ns::AnEnum>::values = folly::range( ::a::different::ns::_AnEnumEnumDataStorage::values);
template <> const folly::Range<const folly::StringPiece*> TEnumTraits< ::a::different::ns::AnEnum>::names = folly::range( ::a::different::ns::_AnEnumEnumDataStorage::names);
template <> const char* TEnumTraits< ::a::different::ns::AnEnum>::findName( ::a::different::ns::AnEnum value) {
  static auto const map = folly::Indestructible< ::a::different::ns::_AnEnum_EnumMapFactory::ValuesToNamesMapType>{ ::a::different::ns::_AnEnum_EnumMapFactory::makeValuesToNamesMap()};
  return findName(*map, value);
}

template <> bool TEnumTraits< ::a::different::ns::AnEnum>::findValue(const char* name,  ::a::different::ns::AnEnum* outValue) {
  static auto const map = folly::Indestructible< ::a::different::ns::_AnEnum_EnumMapFactory::NamesToValuesMapType>{ ::a::different::ns::_AnEnum_EnumMapFactory::makeNamesToValuesMap()};
  return findValue(*map, name, outValue);
}

}} // apache::thrift
namespace a { namespace different { namespace ns {

void AStruct::__clear() {
  // clear all fields
  FieldA = 0;
  __isset.__clear();
}

bool AStruct::operator==(const AStruct& rhs) const {
  if (!((FieldA == rhs.FieldA))) {
    return false;
  }
  return true;
}

void AStruct::translateFieldName(FOLLY_MAYBE_UNUSED folly::StringPiece _fname, FOLLY_MAYBE_UNUSED int16_t& fid, FOLLY_MAYBE_UNUSED apache::thrift::protocol::TType& _ftype) {
  if (false) {}
  else if (_fname == "FieldA") {
    fid = 1;
    _ftype = apache::thrift::protocol::T_I32;
  }
}

void swap(AStruct& a, AStruct& b) {
  using ::std::swap;
  swap(a.FieldA, b.FieldA);
  swap(a.__isset, b.__isset);
}

template uint32_t AStruct::read<>(apache::thrift::BinaryProtocolReader*);
template uint32_t AStruct::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t AStruct::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t AStruct::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t AStruct::read<>(apache::thrift::CompactProtocolReader*);
template uint32_t AStruct::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t AStruct::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t AStruct::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;

}}} // a::different::ns
namespace apache { namespace thrift {

}} // apache::thrift
namespace a { namespace different { namespace ns {

void AStructB::__clear() {
  // clear all fields
  FieldA.reset();
  __isset.__clear();
}

bool AStructB::operator==(const AStructB& rhs) const {
  if (!(((FieldA && rhs.FieldA && *FieldA == *rhs.FieldA) ||(!FieldA && !rhs.FieldA)))) {
    return false;
  }
  return true;
}

void AStructB::translateFieldName(FOLLY_MAYBE_UNUSED folly::StringPiece _fname, FOLLY_MAYBE_UNUSED int16_t& fid, FOLLY_MAYBE_UNUSED apache::thrift::protocol::TType& _ftype) {
  if (false) {}
  else if (_fname == "FieldA") {
    fid = 1;
    _ftype = apache::thrift::protocol::T_STRUCT;
  }
}

void swap(AStructB& a, AStructB& b) {
  using ::std::swap;
  swap(a.FieldA, b.FieldA);
  swap(a.__isset, b.__isset);
}

template uint32_t AStructB::read<>(apache::thrift::BinaryProtocolReader*);
template uint32_t AStructB::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t AStructB::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t AStructB::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t AStructB::read<>(apache::thrift::CompactProtocolReader*);
template uint32_t AStructB::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t AStructB::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t AStructB::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;

}}} // a::different::ns
namespace apache { namespace thrift {

}} // apache::thrift
namespace a { namespace different { namespace ns {

}}} // a::different::ns
