// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	job.proto

It has these top-level messages:
	Job
	JobResult
	Worker
	Log
	Empty
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type JobResult_Status int32

const (
	JobResult_Sucess JobResult_Status = 0
	JobResult_Failed JobResult_Status = 1
)

var JobResult_Status_name = map[int32]string{
	0: "Sucess",
	1: "Failed",
}
var JobResult_Status_value = map[string]int32{
	"Sucess": 0,
	"Failed": 1,
}

func (x JobResult_Status) String() string {
	return proto.EnumName(JobResult_Status_name, int32(x))
}
func (JobResult_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Job struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	WorkerId  string `protobuf:"bytes,2,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	StartTime int64  `protobuf:"varint,3,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime   int64  `protobuf:"varint,4,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *Job) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Job) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

type JobResult struct {
	WorkerId  string `protobuf:"bytes,1,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	JobId     string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	StartTime int64  `protobuf:"varint,3,opt,name=startTime" json:"startTime,omitempty"`
}

func (m *JobResult) Reset()                    { *m = JobResult{} }
func (m *JobResult) String() string            { return proto.CompactTextString(m) }
func (*JobResult) ProtoMessage()               {}
func (*JobResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobResult) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *JobResult) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *JobResult) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

type Worker struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Worker) Reset()                    { *m = Worker{} }
func (m *Worker) String() string            { return proto.CompactTextString(m) }
func (*Worker) ProtoMessage()               {}
func (*Worker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Worker) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Log struct {
	WorkerId string `protobuf:"bytes,1,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	JobId    string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	Message  []byte `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Log) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *Log) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Log) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Job)(nil), "gen.Job")
	proto.RegisterType((*JobResult)(nil), "gen.JobResult")
	proto.RegisterType((*Worker)(nil), "gen.Worker")
	proto.RegisterType((*Log)(nil), "gen.Log")
	proto.RegisterType((*Empty)(nil), "gen.Empty")
	proto.RegisterEnum("gen.JobResult_Status", JobResult_Status_name, JobResult_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for JobService service

type JobServiceClient interface {
	// RegisterJob registers a job given a Worker.ID message
	RegisterJob(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Job, error)
	// RecordLog records
	RecordLog(ctx context.Context, opts ...grpc.CallOption) (JobService_RecordLogClient, error)
	RegisterJobResult(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*Job, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) RegisterJob(ctx context.Context, in *Worker, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/gen.JobService/RegisterJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) RecordLog(ctx context.Context, opts ...grpc.CallOption) (JobService_RecordLogClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_JobService_serviceDesc.Streams[0], c.cc, "/gen.JobService/RecordLog", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobServiceRecordLogClient{stream}
	return x, nil
}

type JobService_RecordLogClient interface {
	Send(*Log) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type jobServiceRecordLogClient struct {
	grpc.ClientStream
}

func (x *jobServiceRecordLogClient) Send(m *Log) error {
	return x.ClientStream.SendMsg(m)
}

func (x *jobServiceRecordLogClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *jobServiceClient) RegisterJobResult(ctx context.Context, in *JobResult, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/gen.JobService/RegisterJobResult", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for JobService service

type JobServiceServer interface {
	// RegisterJob registers a job given a Worker.ID message
	RegisterJob(context.Context, *Worker) (*Job, error)
	// RecordLog records
	RecordLog(JobService_RecordLogServer) error
	RegisterJobResult(context.Context, *JobResult) (*Job, error)
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_RegisterJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Worker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).RegisterJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.JobService/RegisterJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).RegisterJob(ctx, req.(*Worker))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_RecordLog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(JobServiceServer).RecordLog(&jobServiceRecordLogServer{stream})
}

type JobService_RecordLogServer interface {
	SendAndClose(*Empty) error
	Recv() (*Log, error)
	grpc.ServerStream
}

type jobServiceRecordLogServer struct {
	grpc.ServerStream
}

func (x *jobServiceRecordLogServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *jobServiceRecordLogServer) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _JobService_RegisterJobResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).RegisterJobResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gen.JobService/RegisterJobResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).RegisterJobResult(ctx, req.(*JobResult))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gen.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterJob",
			Handler:    _JobService_RegisterJob_Handler,
		},
		{
			MethodName: "RegisterJobResult",
			Handler:    _JobService_RegisterJobResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordLog",
			Handler:       _JobService_RecordLog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "job.proto",
}

func init() { proto.RegisterFile("job.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0x97, 0xd5, 0x75, 0xeb, 0x9d, 0x8c, 0x19, 0x10, 0xea, 0x54, 0x18, 0x01, 0x65, 0x4f,
	0x13, 0xf4, 0x37, 0x28, 0xac, 0xec, 0xa9, 0x15, 0x7c, 0x1c, 0xcd, 0x72, 0x29, 0x99, 0x6b, 0xef,
	0x48, 0x32, 0xc5, 0x27, 0xff, 0x80, 0x3f, 0x5a, 0x9a, 0x3a, 0x36, 0xe7, 0x9b, 0x6f, 0xf7, 0x9e,
	0x03, 0xe7, 0x3b, 0xb9, 0x04, 0xa2, 0x15, 0xc9, 0xe9, 0xc6, 0x90, 0x23, 0x1e, 0x14, 0x58, 0x89,
	0x35, 0x04, 0x09, 0x49, 0x3e, 0x80, 0xb6, 0x56, 0x31, 0x1b, 0xb3, 0x49, 0x94, 0xb6, 0xb5, 0xe2,
	0x97, 0x10, 0xbd, 0x93, 0x79, 0x45, 0xb3, 0xd0, 0x2a, 0x6e, 0x7b, 0xb9, 0xd7, 0x08, 0x33, 0xc5,
	0xaf, 0x01, 0xac, 0xcb, 0x8d, 0x5b, 0x38, 0x5d, 0x62, 0x1c, 0x8c, 0xd9, 0x24, 0x48, 0x23, 0xaf,
	0x3c, 0xeb, 0x12, 0xf9, 0x05, 0xf4, 0xb0, 0x52, 0x8d, 0x79, 0xe2, 0xcd, 0x2e, 0x56, 0xaa, 0xb6,
	0xc4, 0x27, 0x44, 0x09, 0xc9, 0x14, 0xed, 0x76, 0xed, 0x7e, 0x33, 0xd8, 0x11, 0xe3, 0x1c, 0xc2,
	0x15, 0xc9, 0x3d, 0xbd, 0xb3, 0x22, 0x39, 0x53, 0xfc, 0x0a, 0xf6, 0xa0, 0x3f, 0x64, 0x31, 0x86,
	0x30, 0x73, 0xb9, 0xdb, 0x5a, 0x0e, 0x10, 0x66, 0xdb, 0x25, 0x5a, 0x3b, 0x6c, 0xd5, 0xf3, 0x53,
	0xae, 0xd7, 0xa8, 0x86, 0x4c, 0xc4, 0x10, 0xbe, 0x78, 0xc4, 0xf1, 0x8b, 0x45, 0x06, 0xc1, 0x9c,
	0x8a, 0x7f, 0x95, 0x8a, 0xa1, 0x5b, 0xa2, 0xb5, 0x79, 0xd1, 0x54, 0x3a, 0x4d, 0x77, 0xab, 0xe8,
	0x42, 0xe7, 0xb1, 0xdc, 0xb8, 0x8f, 0xfb, 0x2f, 0x06, 0x90, 0x90, 0xcc, 0xd0, 0xbc, 0xe9, 0x25,
	0xf2, 0x5b, 0xe8, 0xa7, 0x58, 0x68, 0xeb, 0xd0, 0xd4, 0xd7, 0xef, 0x4f, 0x0b, 0xac, 0xa6, 0x4d,
	0xb1, 0x51, 0xcf, 0x2f, 0x09, 0x49, 0xd1, 0xe2, 0x37, 0x10, 0xa5, 0xb8, 0x24, 0xa3, 0xea, 0x6a,
	0x8d, 0x31, 0xa7, 0x62, 0x04, 0x7e, 0xf2, 0xc9, 0xa2, 0x35, 0x61, 0xfc, 0x0e, 0xce, 0x0e, 0xe2,
	0x7e, 0xce, 0x3b, 0xd8, 0xe5, 0x34, 0xfb, 0x61, 0xae, 0x0c, 0xfd, 0x0f, 0x78, 0xf8, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0x15, 0x4c, 0x2a, 0x0e, 0x02, 0x00, 0x00,
}