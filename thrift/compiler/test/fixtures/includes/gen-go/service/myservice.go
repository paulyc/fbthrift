// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package service

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift-go"
	module0 "module"
	includes1 "includes"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

var _ = module0.GoUnusedProtection__
var _ = includes1.GoUnusedProtection__
type MyService interface {  //This is a service-level docblock

  // This is a function-level docblock
  // 
  // Parameters:
  //  - S
  //  - I
  Query(s *module0.MyStruct, i *includes1.Included) (err error)
  // Parameters:
  //  - S
  //  - I: arg doc
  HasArgDocs(s *module0.MyStruct, i *includes1.Included) (err error)
}

//This is a service-level docblock
type MyServiceClient struct {
  CC thrift.ClientConn
}

func (client *MyServiceClient) Close() error {
  return client.CC.Close()
}

func (client *MyServiceClient) Open() error {
  return client.CC.Open()
}

func (client *MyServiceClient) IsOpen() bool {
  return client.CC.IsOpen()
}

func NewMyServiceClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *MyServiceClient {
  return &MyServiceClient{ CC: thrift.NewClientConn(t, f) }
}

func NewMyServiceClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *MyServiceClient {
  return &MyServiceClient{ CC: thrift.NewClientConnWithProtocols(t, iprot, oprot) }
}

// This is a function-level docblock
// 
// Parameters:
//  - S
//  - I
func (p *MyServiceClient) Query(s *module0.MyStruct, i *includes1.Included) (err error) {
  args := MyServiceQueryArgs{
    S : s,
    I : i,
  }
  err = p.CC.SendMsg("query", &args, thrift.CALL)
  if err != nil { return }
  return p.recvQuery()
}


func (p *MyServiceClient) recvQuery() (err error) {
  var result MyServiceQueryResult
  return p.CC.RecvMsg("query", &result)
}

// Parameters:
//  - S
//  - I: arg doc
func (p *MyServiceClient) HasArgDocs(s *module0.MyStruct, i *includes1.Included) (err error) {
  args := MyServiceHasArgDocsArgs{
    S : s,
    I : i,
  }
  err = p.CC.SendMsg("has_arg_docs", &args, thrift.CALL)
  if err != nil { return }
  return p.recvHasArgDocs()
}


func (p *MyServiceClient) recvHasArgDocs() (err error) {
  var result MyServiceHasArgDocsResult
  return p.CC.RecvMsg("has_arg_docs", &result)
}


//This is a service-level docblock
type MyServiceThreadsafeClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
  Mu sync.Mutex
}

func NewMyServiceThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *MyServiceThreadsafeClient {
  return &MyServiceThreadsafeClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewMyServiceThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *MyServiceThreadsafeClient {
  return &MyServiceThreadsafeClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *MyServiceThreadsafeClient) Threadsafe() {}

// This is a function-level docblock
// 
// Parameters:
//  - S
//  - I
func (p *MyServiceThreadsafeClient) Query(s *module0.MyStruct, i *includes1.Included) (err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendQuery(s, i); err != nil { return }
  return p.recvQuery()
}

func (p *MyServiceThreadsafeClient) sendQuery(s *module0.MyStruct, i *includes1.Included)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("query", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := MyServiceQueryArgs{
  S : s,
  I : i,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *MyServiceThreadsafeClient) recvQuery() (err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "query" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "query failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "query failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error2 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error3 error
    error3, err = error2.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error3
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "query failed: invalid message type")
    return
  }
  result := MyServiceQueryResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  return
}

// Parameters:
//  - S
//  - I: arg doc
func (p *MyServiceThreadsafeClient) HasArgDocs(s *module0.MyStruct, i *includes1.Included) (err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendHasArgDocs(s, i); err != nil { return }
  return p.recvHasArgDocs()
}

func (p *MyServiceThreadsafeClient) sendHasArgDocs(s *module0.MyStruct, i *includes1.Included)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("has_arg_docs", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := MyServiceHasArgDocsArgs{
  S : s,
  I : i,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *MyServiceThreadsafeClient) recvHasArgDocs() (err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "has_arg_docs" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "has_arg_docs failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "has_arg_docs failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error4 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error5 error
    error5, err = error4.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error5
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "has_arg_docs failed: invalid message type")
    return
  }
  result := MyServiceHasArgDocsResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  return
}


type MyServiceProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  handler MyService
}

func (p *MyServiceProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *MyServiceProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *MyServiceProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func NewMyServiceProcessor(handler MyService) *MyServiceProcessor {
  self6 := &MyServiceProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction)}
  self6.processorMap["query"] = &myServiceProcessorQuery{handler:handler}
  self6.processorMap["has_arg_docs"] = &myServiceProcessorHasArgDocs{handler:handler}
  return self6
}

type myServiceProcessorQuery struct {
  handler MyService
}

func (p *myServiceProcessorQuery) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := MyServiceQueryArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *myServiceProcessorQuery) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("query", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *myServiceProcessorQuery) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*MyServiceQueryArgs)
  var result MyServiceQueryResult
  if err := p.handler.Query(args.S, args.I); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing query: " + err.Error())
      return x, x
    }
  }
  return &result, nil
}

type myServiceProcessorHasArgDocs struct {
  handler MyService
}

func (p *myServiceProcessorHasArgDocs) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := MyServiceHasArgDocsArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *myServiceProcessorHasArgDocs) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("has_arg_docs", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *myServiceProcessorHasArgDocs) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*MyServiceHasArgDocsArgs)
  var result MyServiceHasArgDocsResult
  if err := p.handler.HasArgDocs(args.S, args.I); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing has_arg_docs: " + err.Error())
      return x, x
    }
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - S
//  - I
type MyServiceQueryArgs struct {
  thrift.IRequest
  S *module0.MyStruct `thrift:"s,1" db:"s" json:"s"`
  I *includes1.Included `thrift:"i,2" db:"i" json:"i"`
}

func NewMyServiceQueryArgs() *MyServiceQueryArgs {
  return &MyServiceQueryArgs{
S: module0.NewMyStruct(),

I: includes1.NewIncluded(),
}
}

var MyServiceQueryArgs_S_DEFAULT *module0.MyStruct
func (p *MyServiceQueryArgs) GetS() *module0.MyStruct {
  if !p.IsSetS() {
    return MyServiceQueryArgs_S_DEFAULT
  }
return p.S
}
var MyServiceQueryArgs_I_DEFAULT *includes1.Included
func (p *MyServiceQueryArgs) GetI() *includes1.Included {
  if !p.IsSetI() {
    return MyServiceQueryArgs_I_DEFAULT
  }
return p.I
}
func (p *MyServiceQueryArgs) IsSetS() bool {
  return p.S != nil
}

func (p *MyServiceQueryArgs) IsSetI() bool {
  return p.I != nil
}

func (p *MyServiceQueryArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServiceQueryArgs)  ReadField1(iprot thrift.Protocol) error {
  p.S = module0.NewMyStruct()
  if err := p.S.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.S), err)
  }
  return nil
}

func (p *MyServiceQueryArgs)  ReadField2(iprot thrift.Protocol) error {
  p.I = includes1.NewIncluded()
  if err := p.I.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.I), err)
  }
  return nil
}

func (p *MyServiceQueryArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("query_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServiceQueryArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("s", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:s: ", p), err) }
  if err := p.S.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.S), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:s: ", p), err) }
  return err
}

func (p *MyServiceQueryArgs) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("i", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:i: ", p), err) }
  if err := p.I.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.I), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:i: ", p), err) }
  return err
}

func (p *MyServiceQueryArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServiceQueryArgs(%+v)", *p)
}

type MyServiceQueryResult struct {
  thrift.IResponse
}

func NewMyServiceQueryResult() *MyServiceQueryResult {
  return &MyServiceQueryResult{}
}

func (p *MyServiceQueryResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServiceQueryResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("query_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServiceQueryResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServiceQueryResult(%+v)", *p)
}

// Attributes:
//  - S
//  - I: arg doc
type MyServiceHasArgDocsArgs struct {
  thrift.IRequest
  S *module0.MyStruct `thrift:"s,1" db:"s" json:"s"`
  I *includes1.Included `thrift:"i,2" db:"i" json:"i"`
}

func NewMyServiceHasArgDocsArgs() *MyServiceHasArgDocsArgs {
  return &MyServiceHasArgDocsArgs{
S: module0.NewMyStruct(),

I: includes1.NewIncluded(),
}
}

var MyServiceHasArgDocsArgs_S_DEFAULT *module0.MyStruct
func (p *MyServiceHasArgDocsArgs) GetS() *module0.MyStruct {
  if !p.IsSetS() {
    return MyServiceHasArgDocsArgs_S_DEFAULT
  }
return p.S
}
var MyServiceHasArgDocsArgs_I_DEFAULT *includes1.Included
func (p *MyServiceHasArgDocsArgs) GetI() *includes1.Included {
  if !p.IsSetI() {
    return MyServiceHasArgDocsArgs_I_DEFAULT
  }
return p.I
}
func (p *MyServiceHasArgDocsArgs) IsSetS() bool {
  return p.S != nil
}

func (p *MyServiceHasArgDocsArgs) IsSetI() bool {
  return p.I != nil
}

func (p *MyServiceHasArgDocsArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServiceHasArgDocsArgs)  ReadField1(iprot thrift.Protocol) error {
  p.S = module0.NewMyStruct()
  if err := p.S.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.S), err)
  }
  return nil
}

func (p *MyServiceHasArgDocsArgs)  ReadField2(iprot thrift.Protocol) error {
  p.I = includes1.NewIncluded()
  if err := p.I.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.I), err)
  }
  return nil
}

func (p *MyServiceHasArgDocsArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("has_arg_docs_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServiceHasArgDocsArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("s", thrift.STRUCT, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:s: ", p), err) }
  if err := p.S.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.S), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:s: ", p), err) }
  return err
}

func (p *MyServiceHasArgDocsArgs) writeField2(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("i", thrift.STRUCT, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:i: ", p), err) }
  if err := p.I.Write(oprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.I), err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:i: ", p), err) }
  return err
}

func (p *MyServiceHasArgDocsArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServiceHasArgDocsArgs(%+v)", *p)
}

type MyServiceHasArgDocsResult struct {
  thrift.IResponse
}

func NewMyServiceHasArgDocsResult() *MyServiceHasArgDocsResult {
  return &MyServiceHasArgDocsResult{}
}

func (p *MyServiceHasArgDocsResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *MyServiceHasArgDocsResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("has_arg_docs_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *MyServiceHasArgDocsResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("MyServiceHasArgDocsResult(%+v)", *p)
}


