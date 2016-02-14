// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package example

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type Example interface {
	// Parameters:
	//  - Num1
	//  - Num2
	Add(num1 int32, num2 int32) (r int32, err error)
	// Parameters:
	//  - Num1
	//  - Num2
	//  - ClientTimeoutMs
	AddTimeout(num1 int32, num2 int32, client_timeout_ms int32) (r int32, err error)
}

type ExampleClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewExampleClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ExampleClient {
	return &ExampleClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewExampleClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ExampleClient {
	return &ExampleClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Num1
//  - Num2
func (p *ExampleClient) Add(num1 int32, num2 int32) (r int32, err error) {
	if err = p.sendAdd(num1, num2); err != nil {
		return
	}
	return p.recvAdd()
}

func (p *ExampleClient) sendAdd(num1 int32, num2 int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("add", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AddArgs{
		Num1: num1,
		Num2: num2,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ExampleClient) recvAdd() (value int32, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "add failed: out of sequence response")
		return
	}
	result := AddResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Num1
//  - Num2
//  - ClientTimeoutMs
func (p *ExampleClient) AddTimeout(num1 int32, num2 int32, client_timeout_ms int32) (r int32, err error) {
	if err = p.sendAddTimeout(num1, num2, client_timeout_ms); err != nil {
		return
	}
	return p.recvAddTimeout()
}

func (p *ExampleClient) sendAddTimeout(num1 int32, num2 int32, client_timeout_ms int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("add_timeout", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := AddTimeoutArgs{
		Num1:            num1,
		Num2:            num2,
		ClientTimeoutMs: client_timeout_ms,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ExampleClient) recvAddTimeout() (value int32, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
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
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "add_timeout failed: out of sequence response")
		return
	}
	result := AddTimeoutResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ExampleProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Example
}

func (p *ExampleProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ExampleProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ExampleProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewExampleProcessor(handler Example) *ExampleProcessor {

	self4 := &ExampleProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["add"] = &exampleProcessorAdd{handler: handler}
	self4.processorMap["add_timeout"] = &exampleProcessorAddTimeout{handler: handler}
	return self4
}

func (p *ExampleProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type exampleProcessorAdd struct {
	handler Example
}

func (p *exampleProcessorAdd) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AddArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AddResult{}
	var retval int32
	var err2 error
	if retval, err2 = p.handler.Add(args.Num1, args.Num2); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add: "+err2.Error())
		oprot.WriteMessageBegin("add", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("add", thrift.REPLY, seqId); err2 != nil {
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
	if err != nil {
		return
	}
	return true, err
}

type exampleProcessorAddTimeout struct {
	handler Example
}

func (p *exampleProcessorAddTimeout) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AddTimeoutArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add_timeout", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := AddTimeoutResult{}
	var retval int32
	var err2 error
	if retval, err2 = p.handler.AddTimeout(args.Num1, args.Num2, args.ClientTimeoutMs); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add_timeout: "+err2.Error())
		oprot.WriteMessageBegin("add_timeout", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("add_timeout", thrift.REPLY, seqId); err2 != nil {
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
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type AddArgs struct {
	Num1 int32 `thrift:"num1,1" json:"num1"`
	Num2 int32 `thrift:"num2,2" json:"num2"`
}

func NewAddArgs() *AddArgs {
	return &AddArgs{}
}

func (p *AddArgs) GetNum1() int32 {
	return p.Num1
}

func (p *AddArgs) GetNum2() int32 {
	return p.Num2
}
func (p *AddArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Num1 = v
	}
	return nil
}

func (p *AddArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Num2 = v
	}
	return nil
}

func (p *AddArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *AddArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num1", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:num1: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num1)); err != nil {
		return fmt.Errorf("%T.num1 (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:num1: %s", p, err)
	}
	return err
}

func (p *AddArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num2", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:num2: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num2)); err != nil {
		return fmt.Errorf("%T.num2 (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:num2: %s", p, err)
	}
	return err
}

func (p *AddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddArgs(%+v)", *p)
}

type AddResult struct {
	Success *int32 `thrift:"success,0" json:"success"`
}

func NewAddResult() *AddResult {
	return &AddResult{}
}

var AddResult_Success_DEFAULT int32

func (p *AddResult) GetSuccess() int32 {
	if !p.IsSetSuccess() {
		return AddResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *AddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *AddResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *AddResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *AddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddResult(%+v)", *p)
}

type AddTimeoutArgs struct {
	Num1            int32 `thrift:"num1,1" json:"num1"`
	Num2            int32 `thrift:"num2,2" json:"num2"`
	ClientTimeoutMs int32 `thrift:"client_timeout_ms,3" json:"client_timeout_ms"`
}

func NewAddTimeoutArgs() *AddTimeoutArgs {
	return &AddTimeoutArgs{}
}

func (p *AddTimeoutArgs) GetNum1() int32 {
	return p.Num1
}

func (p *AddTimeoutArgs) GetNum2() int32 {
	return p.Num2
}

func (p *AddTimeoutArgs) GetClientTimeoutMs() int32 {
	return p.ClientTimeoutMs
}
func (p *AddTimeoutArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddTimeoutArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Num1 = v
	}
	return nil
}

func (p *AddTimeoutArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Num2 = v
	}
	return nil
}

func (p *AddTimeoutArgs) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.ClientTimeoutMs = v
	}
	return nil
}

func (p *AddTimeoutArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_timeout_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *AddTimeoutArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num1", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:num1: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num1)); err != nil {
		return fmt.Errorf("%T.num1 (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:num1: %s", p, err)
	}
	return err
}

func (p *AddTimeoutArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("num2", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:num2: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Num2)); err != nil {
		return fmt.Errorf("%T.num2 (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:num2: %s", p, err)
	}
	return err
}

func (p *AddTimeoutArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("client_timeout_ms", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:client_timeout_ms: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.ClientTimeoutMs)); err != nil {
		return fmt.Errorf("%T.client_timeout_ms (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:client_timeout_ms: %s", p, err)
	}
	return err
}

func (p *AddTimeoutArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddTimeoutArgs(%+v)", *p)
}

type AddTimeoutResult struct {
	Success *int32 `thrift:"success,0" json:"success"`
}

func NewAddTimeoutResult() *AddTimeoutResult {
	return &AddTimeoutResult{}
}

var AddTimeoutResult_Success_DEFAULT int32

func (p *AddTimeoutResult) GetSuccess() int32 {
	if !p.IsSetSuccess() {
		return AddTimeoutResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *AddTimeoutResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddTimeoutResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
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
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *AddTimeoutResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 0: %s", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *AddTimeoutResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("add_timeout_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *AddTimeoutResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *AddTimeoutResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddTimeoutResult(%+v)", *p)
}