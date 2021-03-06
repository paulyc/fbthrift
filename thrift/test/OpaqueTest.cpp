/*
 * Copyright 2013-present Facebook, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

#include <string>
#include <unordered_map>
#include <folly/portability/GTest.h>
#include <thrift/test/gen-cpp2/OpaqueTest_types.h>
#include <thrift/lib/cpp2/protocol/BinaryProtocol.h>

using namespace apache::thrift::test;

template<class T>
T getTestStruct() {
  T a;
  a.d1 = OpaqueDouble1{123.0};
  a.d2 = OpaqueDouble2{234.0};
  for (int i = 1; i <= 5; ++i) {
    a.dmap[i] = OpaqueDouble1{9000.0 + i};
  }
  for (int i = 1; i <= 5; ++i) {
    a.ids.push_back(NonConvertibleId{111 * i});
  }
  return a;
}

OpaqueTestStruct getTestStruct() {
  return getTestStruct<OpaqueTestStruct>();
}

template<typename T>
void checkTypedefs() {
  EXPECT_FALSE((std::is_same<double, decltype(T::d1)>::value));
  EXPECT_FALSE((std::is_same<double, decltype(T::d2)>::value));
  EXPECT_FALSE((std::is_same<double,
                             decltype(T::dmap.begin()->second)>::value));
  EXPECT_FALSE((std::is_same<int64_t, decltype(T::ids[0])>::value));
  EXPECT_TRUE((std::is_same<decltype(T::d1),
                            decltype(T::dmap.begin()->second)>::value));
}

TEST(Opaque, Typedefs) {
  checkTypedefs<OpaqueTestStruct>();
}

TEST(Opaque, Serialize) {
  using namespace apache::thrift;

  OpaqueTestStruct a = getTestStruct();
  BinaryProtocolWriter protWriter;
  size_t bufSize = a.serializedSize(&protWriter);
  folly::IOBufQueue queue;
  protWriter.setOutput(&queue, bufSize);
  a.write(&protWriter);

  auto buf = queue.move();
  BinaryProtocolReader protReader;
  protReader.setInput(buf.get());
  OpaqueTestStruct a2;
  a2.read(&protReader);
  EXPECT_EQ(getTestStruct(), a2);
}
