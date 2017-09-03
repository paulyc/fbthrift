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

namespace cpp2 {

void Included::__clear() {
  // clear all fields
  MyIntField = 0LL;
  __isset.__clear();
}

bool Included::operator==(const Included& rhs) const {
  if (!((MyIntField == rhs.MyIntField))) {
    return false;
  }
  return true;
}

void Included::translateFieldName(FOLLY_MAYBE_UNUSED folly::StringPiece _fname, FOLLY_MAYBE_UNUSED int16_t& fid, FOLLY_MAYBE_UNUSED apache::thrift::protocol::TType& _ftype) {
  if (false) {}
  else if (_fname == "MyIntField") {
    fid = 1;
    _ftype = apache::thrift::protocol::T_I64;
  }
}

void swap(Included& a, Included& b) {
  using ::std::swap;
  swap(a.MyIntField, b.MyIntField);
  swap(a.__isset, b.__isset);
}

template uint32_t Included::read<>(apache::thrift::BinaryProtocolReader*);
template uint32_t Included::write<>(apache::thrift::BinaryProtocolWriter*) const;
template uint32_t Included::serializedSize<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t Included::serializedSizeZC<>(apache::thrift::BinaryProtocolWriter const*) const;
template uint32_t Included::read<>(apache::thrift::CompactProtocolReader*);
template uint32_t Included::write<>(apache::thrift::CompactProtocolWriter*) const;
template uint32_t Included::serializedSize<>(apache::thrift::CompactProtocolWriter const*) const;
template uint32_t Included::serializedSizeZC<>(apache::thrift::CompactProtocolWriter const*) const;

} // cpp2
namespace apache { namespace thrift {

}} // apache::thrift
namespace cpp2 {

} // cpp2
