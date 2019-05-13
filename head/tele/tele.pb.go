// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tele.proto

package tele

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type State int32

const (
	State_Invalid      State = 0
	State_Boot         State = 1
	State_Work         State = 2
	State_Disconnected State = 3
	State_Problem      State = 4
)

var State_name = map[int32]string{
	0: "Invalid",
	1: "Boot",
	2: "Work",
	3: "Disconnected",
	4: "Problem",
}
var State_value = map[string]int32{
	"Invalid":      0,
	"Boot":         1,
	"Work":         2,
	"Disconnected": 3,
	"Problem":      4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}
func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0}
}

type PaymentMethod int32

const (
	PaymentMethod_Nothing  PaymentMethod = 0
	PaymentMethod_Cash     PaymentMethod = 1
	PaymentMethod_Cashless PaymentMethod = 2
)

var PaymentMethod_name = map[int32]string{
	0: "Nothing",
	1: "Cash",
	2: "Cashless",
}
var PaymentMethod_value = map[string]int32{
	"Nothing":  0,
	"Cash":     1,
	"Cashless": 2,
}

func (x PaymentMethod) String() string {
	return proto.EnumName(PaymentMethod_name, int32(x))
}
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{1}
}

// Optimising for rare, bulk delivery on cell network.
// "Touching network" is expensive, while 10 or 900 bytes is about same cost.
type Telemetry struct {
	VmId                 int32                  `protobuf:"varint,1,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	Error                *Telemetry_Error       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Inventory            *Telemetry_Inventory   `protobuf:"bytes,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Money                *Telemetry_Money       `protobuf:"bytes,4,opt,name=money,proto3" json:"money,omitempty"`
	Transaction          *Telemetry_Transaction `protobuf:"bytes,5,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Stat                 *Telemetry_Stat        `protobuf:"bytes,6,opt,name=stat,proto3" json:"stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Telemetry) Reset()         { *m = Telemetry{} }
func (m *Telemetry) String() string { return proto.CompactTextString(m) }
func (*Telemetry) ProtoMessage()    {}
func (*Telemetry) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0}
}
func (m *Telemetry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry.Unmarshal(m, b)
}
func (m *Telemetry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry.Marshal(b, m, deterministic)
}
func (dst *Telemetry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry.Merge(dst, src)
}
func (m *Telemetry) XXX_Size() int {
	return xxx_messageInfo_Telemetry.Size(m)
}
func (m *Telemetry) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry proto.InternalMessageInfo

func (m *Telemetry) GetVmId() int32 {
	if m != nil {
		return m.VmId
	}
	return 0
}

func (m *Telemetry) GetError() *Telemetry_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Telemetry) GetInventory() *Telemetry_Inventory {
	if m != nil {
		return m.Inventory
	}
	return nil
}

func (m *Telemetry) GetMoney() *Telemetry_Money {
	if m != nil {
		return m.Money
	}
	return nil
}

func (m *Telemetry) GetTransaction() *Telemetry_Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *Telemetry) GetStat() *Telemetry_Stat {
	if m != nil {
		return m.Stat
	}
	return nil
}

type Telemetry_Error struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry_Error) Reset()         { *m = Telemetry_Error{} }
func (m *Telemetry_Error) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Error) ProtoMessage()    {}
func (*Telemetry_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0, 0}
}
func (m *Telemetry_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Error.Unmarshal(m, b)
}
func (m *Telemetry_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Error.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Error.Merge(dst, src)
}
func (m *Telemetry_Error) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Error.Size(m)
}
func (m *Telemetry_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Error proto.InternalMessageInfo

func (m *Telemetry_Error) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Telemetry_Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Telemetry_Error) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Telemetry_Inventory struct {
	Water                int32    `protobuf:"varint,1,opt,name=water,proto3" json:"water,omitempty"`
	Coffee               int32    `protobuf:"varint,2,opt,name=coffee,proto3" json:"coffee,omitempty"`
	Cup                  int32    `protobuf:"varint,3,opt,name=cup,proto3" json:"cup,omitempty"`
	Hoppers              []int32  `protobuf:"varint,4,rep,packed,name=hoppers,proto3" json:"hoppers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry_Inventory) Reset()         { *m = Telemetry_Inventory{} }
func (m *Telemetry_Inventory) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Inventory) ProtoMessage()    {}
func (*Telemetry_Inventory) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0, 1}
}
func (m *Telemetry_Inventory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Inventory.Unmarshal(m, b)
}
func (m *Telemetry_Inventory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Inventory.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Inventory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Inventory.Merge(dst, src)
}
func (m *Telemetry_Inventory) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Inventory.Size(m)
}
func (m *Telemetry_Inventory) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Inventory.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Inventory proto.InternalMessageInfo

func (m *Telemetry_Inventory) GetWater() int32 {
	if m != nil {
		return m.Water
	}
	return 0
}

func (m *Telemetry_Inventory) GetCoffee() int32 {
	if m != nil {
		return m.Coffee
	}
	return 0
}

func (m *Telemetry_Inventory) GetCup() int32 {
	if m != nil {
		return m.Cup
	}
	return 0
}

func (m *Telemetry_Inventory) GetHoppers() []int32 {
	if m != nil {
		return m.Hoppers
	}
	return nil
}

type Telemetry_Money struct {
	TotalBills           uint32   `protobuf:"varint,1,opt,name=total_bills,json=totalBills,proto3" json:"total_bills,omitempty"`
	TotalCoins           uint32   `protobuf:"varint,2,opt,name=total_coins,json=totalCoins,proto3" json:"total_coins,omitempty"`
	Bills                []uint32 `protobuf:"varint,3,rep,packed,name=bills,proto3" json:"bills,omitempty"`
	Coins                []uint32 `protobuf:"varint,4,rep,packed,name=coins,proto3" json:"coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Telemetry_Money) Reset()         { *m = Telemetry_Money{} }
func (m *Telemetry_Money) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Money) ProtoMessage()    {}
func (*Telemetry_Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0, 2}
}
func (m *Telemetry_Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Money.Unmarshal(m, b)
}
func (m *Telemetry_Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Money.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Money.Merge(dst, src)
}
func (m *Telemetry_Money) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Money.Size(m)
}
func (m *Telemetry_Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Money proto.InternalMessageInfo

func (m *Telemetry_Money) GetTotalBills() uint32 {
	if m != nil {
		return m.TotalBills
	}
	return 0
}

func (m *Telemetry_Money) GetTotalCoins() uint32 {
	if m != nil {
		return m.TotalCoins
	}
	return 0
}

func (m *Telemetry_Money) GetBills() []uint32 {
	if m != nil {
		return m.Bills
	}
	return nil
}

func (m *Telemetry_Money) GetCoins() []uint32 {
	if m != nil {
		return m.Coins
	}
	return nil
}

type Telemetry_Transaction struct {
	Code                 int32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Options              []int32       `protobuf:"varint,2,rep,packed,name=options,proto3" json:"options,omitempty"`
	Price                uint32        `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	PaymentMethod        PaymentMethod `protobuf:"varint,4,opt,name=payment_method,json=paymentMethod,proto3,enum=tele.PaymentMethod" json:"payment_method,omitempty"`
	CreditBills          uint32        `protobuf:"varint,5,opt,name=credit_bills,json=creditBills,proto3" json:"credit_bills,omitempty"`
	CreditCoins          uint32        `protobuf:"varint,6,opt,name=credit_coins,json=creditCoins,proto3" json:"credit_coins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Telemetry_Transaction) Reset()         { *m = Telemetry_Transaction{} }
func (m *Telemetry_Transaction) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Transaction) ProtoMessage()    {}
func (*Telemetry_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0, 3}
}
func (m *Telemetry_Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Transaction.Unmarshal(m, b)
}
func (m *Telemetry_Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Transaction.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Transaction.Merge(dst, src)
}
func (m *Telemetry_Transaction) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Transaction.Size(m)
}
func (m *Telemetry_Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Transaction proto.InternalMessageInfo

func (m *Telemetry_Transaction) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Telemetry_Transaction) GetOptions() []int32 {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Telemetry_Transaction) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Telemetry_Transaction) GetPaymentMethod() PaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return PaymentMethod_Nothing
}

func (m *Telemetry_Transaction) GetCreditBills() uint32 {
	if m != nil {
		return m.CreditBills
	}
	return 0
}

func (m *Telemetry_Transaction) GetCreditCoins() uint32 {
	if m != nil {
		return m.CreditCoins
	}
	return 0
}

type Telemetry_Stat struct {
	BillRejected         map[uint32]uint32 `protobuf:"bytes,16,rep,name=bill_rejected,json=billRejected,proto3" json:"bill_rejected,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	CoinRejected         map[uint32]uint32 `protobuf:"bytes,17,rep,name=coin_rejected,json=coinRejected,proto3" json:"coin_rejected,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Telemetry_Stat) Reset()         { *m = Telemetry_Stat{} }
func (m *Telemetry_Stat) String() string { return proto.CompactTextString(m) }
func (*Telemetry_Stat) ProtoMessage()    {}
func (*Telemetry_Stat) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{0, 4}
}
func (m *Telemetry_Stat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Telemetry_Stat.Unmarshal(m, b)
}
func (m *Telemetry_Stat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Telemetry_Stat.Marshal(b, m, deterministic)
}
func (dst *Telemetry_Stat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Telemetry_Stat.Merge(dst, src)
}
func (m *Telemetry_Stat) XXX_Size() int {
	return xxx_messageInfo_Telemetry_Stat.Size(m)
}
func (m *Telemetry_Stat) XXX_DiscardUnknown() {
	xxx_messageInfo_Telemetry_Stat.DiscardUnknown(m)
}

var xxx_messageInfo_Telemetry_Stat proto.InternalMessageInfo

func (m *Telemetry_Stat) GetBillRejected() map[uint32]uint32 {
	if m != nil {
		return m.BillRejected
	}
	return nil
}

func (m *Telemetry_Stat) GetCoinRejected() map[uint32]uint32 {
	if m != nil {
		return m.CoinRejected
	}
	return nil
}

type Command struct {
	Id         uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ReplyTopic string `protobuf:"bytes,2,opt,name=reply_topic,json=replyTopic,proto3" json:"reply_topic,omitempty"`
	// Types that are valid to be assigned to Task:
	//	*Command_Report
	//	*Command_Abort
	//	*Command_Dispense
	//	*Command_SetGiftCredit
	Task                 isCommand_Task `protobuf_oneof:"task"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{1}
}
func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (dst *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(dst, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Command) GetReplyTopic() string {
	if m != nil {
		return m.ReplyTopic
	}
	return ""
}

type isCommand_Task interface {
	isCommand_Task()
}

type Command_Report struct {
	Report *Command_ArgReport `protobuf:"bytes,3,opt,name=report,proto3,oneof"`
}

type Command_Abort struct {
	Abort *Command_ArgAbort `protobuf:"bytes,4,opt,name=abort,proto3,oneof"`
}

type Command_Dispense struct {
	Dispense *Command_ArgDispense `protobuf:"bytes,5,opt,name=dispense,proto3,oneof"`
}

type Command_SetGiftCredit struct {
	SetGiftCredit *Command_ArgSetGiftCredit `protobuf:"bytes,6,opt,name=set_gift_credit,json=setGiftCredit,proto3,oneof"`
}

func (*Command_Report) isCommand_Task() {}

func (*Command_Abort) isCommand_Task() {}

func (*Command_Dispense) isCommand_Task() {}

func (*Command_SetGiftCredit) isCommand_Task() {}

func (m *Command) GetTask() isCommand_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *Command) GetReport() *Command_ArgReport {
	if x, ok := m.GetTask().(*Command_Report); ok {
		return x.Report
	}
	return nil
}

func (m *Command) GetAbort() *Command_ArgAbort {
	if x, ok := m.GetTask().(*Command_Abort); ok {
		return x.Abort
	}
	return nil
}

func (m *Command) GetDispense() *Command_ArgDispense {
	if x, ok := m.GetTask().(*Command_Dispense); ok {
		return x.Dispense
	}
	return nil
}

func (m *Command) GetSetGiftCredit() *Command_ArgSetGiftCredit {
	if x, ok := m.GetTask().(*Command_SetGiftCredit); ok {
		return x.SetGiftCredit
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Command) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Command_OneofMarshaler, _Command_OneofUnmarshaler, _Command_OneofSizer, []interface{}{
		(*Command_Report)(nil),
		(*Command_Abort)(nil),
		(*Command_Dispense)(nil),
		(*Command_SetGiftCredit)(nil),
	}
}

func _Command_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Command)
	// task
	switch x := m.Task.(type) {
	case *Command_Report:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Report); err != nil {
			return err
		}
	case *Command_Abort:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Abort); err != nil {
			return err
		}
	case *Command_Dispense:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dispense); err != nil {
			return err
		}
	case *Command_SetGiftCredit:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SetGiftCredit); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Command.Task has unexpected type %T", x)
	}
	return nil
}

func _Command_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Command)
	switch tag {
	case 3: // task.report
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Command_ArgReport)
		err := b.DecodeMessage(msg)
		m.Task = &Command_Report{msg}
		return true, err
	case 4: // task.abort
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Command_ArgAbort)
		err := b.DecodeMessage(msg)
		m.Task = &Command_Abort{msg}
		return true, err
	case 5: // task.dispense
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Command_ArgDispense)
		err := b.DecodeMessage(msg)
		m.Task = &Command_Dispense{msg}
		return true, err
	case 6: // task.set_gift_credit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Command_ArgSetGiftCredit)
		err := b.DecodeMessage(msg)
		m.Task = &Command_SetGiftCredit{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Command_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Command)
	// task
	switch x := m.Task.(type) {
	case *Command_Report:
		s := proto.Size(x.Report)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_Abort:
		s := proto.Size(x.Abort)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_Dispense:
		s := proto.Size(x.Dispense)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Command_SetGiftCredit:
		s := proto.Size(x.SetGiftCredit)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Command_ArgReport struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command_ArgReport) Reset()         { *m = Command_ArgReport{} }
func (m *Command_ArgReport) String() string { return proto.CompactTextString(m) }
func (*Command_ArgReport) ProtoMessage()    {}
func (*Command_ArgReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{1, 0}
}
func (m *Command_ArgReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command_ArgReport.Unmarshal(m, b)
}
func (m *Command_ArgReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command_ArgReport.Marshal(b, m, deterministic)
}
func (dst *Command_ArgReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command_ArgReport.Merge(dst, src)
}
func (m *Command_ArgReport) XXX_Size() int {
	return xxx_messageInfo_Command_ArgReport.Size(m)
}
func (m *Command_ArgReport) XXX_DiscardUnknown() {
	xxx_messageInfo_Command_ArgReport.DiscardUnknown(m)
}

var xxx_messageInfo_Command_ArgReport proto.InternalMessageInfo

type Command_ArgAbort struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command_ArgAbort) Reset()         { *m = Command_ArgAbort{} }
func (m *Command_ArgAbort) String() string { return proto.CompactTextString(m) }
func (*Command_ArgAbort) ProtoMessage()    {}
func (*Command_ArgAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{1, 1}
}
func (m *Command_ArgAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command_ArgAbort.Unmarshal(m, b)
}
func (m *Command_ArgAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command_ArgAbort.Marshal(b, m, deterministic)
}
func (dst *Command_ArgAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command_ArgAbort.Merge(dst, src)
}
func (m *Command_ArgAbort) XXX_Size() int {
	return xxx_messageInfo_Command_ArgAbort.Size(m)
}
func (m *Command_ArgAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_Command_ArgAbort.DiscardUnknown(m)
}

var xxx_messageInfo_Command_ArgAbort proto.InternalMessageInfo

type Command_ArgDispense struct {
	Amount               uint32   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command_ArgDispense) Reset()         { *m = Command_ArgDispense{} }
func (m *Command_ArgDispense) String() string { return proto.CompactTextString(m) }
func (*Command_ArgDispense) ProtoMessage()    {}
func (*Command_ArgDispense) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{1, 2}
}
func (m *Command_ArgDispense) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command_ArgDispense.Unmarshal(m, b)
}
func (m *Command_ArgDispense) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command_ArgDispense.Marshal(b, m, deterministic)
}
func (dst *Command_ArgDispense) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command_ArgDispense.Merge(dst, src)
}
func (m *Command_ArgDispense) XXX_Size() int {
	return xxx_messageInfo_Command_ArgDispense.Size(m)
}
func (m *Command_ArgDispense) XXX_DiscardUnknown() {
	xxx_messageInfo_Command_ArgDispense.DiscardUnknown(m)
}

var xxx_messageInfo_Command_ArgDispense proto.InternalMessageInfo

func (m *Command_ArgDispense) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Command_ArgSetGiftCredit struct {
	Amount               uint32   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command_ArgSetGiftCredit) Reset()         { *m = Command_ArgSetGiftCredit{} }
func (m *Command_ArgSetGiftCredit) String() string { return proto.CompactTextString(m) }
func (*Command_ArgSetGiftCredit) ProtoMessage()    {}
func (*Command_ArgSetGiftCredit) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{1, 3}
}
func (m *Command_ArgSetGiftCredit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command_ArgSetGiftCredit.Unmarshal(m, b)
}
func (m *Command_ArgSetGiftCredit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command_ArgSetGiftCredit.Marshal(b, m, deterministic)
}
func (dst *Command_ArgSetGiftCredit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command_ArgSetGiftCredit.Merge(dst, src)
}
func (m *Command_ArgSetGiftCredit) XXX_Size() int {
	return xxx_messageInfo_Command_ArgSetGiftCredit.Size(m)
}
func (m *Command_ArgSetGiftCredit) XXX_DiscardUnknown() {
	xxx_messageInfo_Command_ArgSetGiftCredit.DiscardUnknown(m)
}

var xxx_messageInfo_Command_ArgSetGiftCredit proto.InternalMessageInfo

func (m *Command_ArgSetGiftCredit) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Response struct {
	CommandId            uint32   `protobuf:"varint,1,opt,name=command_id,json=commandId,proto3" json:"command_id,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_tele_a7e94f26f60f621e, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCommandId() uint32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Telemetry)(nil), "tele.Telemetry")
	proto.RegisterType((*Telemetry_Error)(nil), "tele.Telemetry.Error")
	proto.RegisterType((*Telemetry_Inventory)(nil), "tele.Telemetry.Inventory")
	proto.RegisterType((*Telemetry_Money)(nil), "tele.Telemetry.Money")
	proto.RegisterType((*Telemetry_Transaction)(nil), "tele.Telemetry.Transaction")
	proto.RegisterType((*Telemetry_Stat)(nil), "tele.Telemetry.Stat")
	proto.RegisterMapType((map[uint32]uint32)(nil), "tele.Telemetry.Stat.BillRejectedEntry")
	proto.RegisterMapType((map[uint32]uint32)(nil), "tele.Telemetry.Stat.CoinRejectedEntry")
	proto.RegisterType((*Command)(nil), "tele.Command")
	proto.RegisterType((*Command_ArgReport)(nil), "tele.Command.ArgReport")
	proto.RegisterType((*Command_ArgAbort)(nil), "tele.Command.ArgAbort")
	proto.RegisterType((*Command_ArgDispense)(nil), "tele.Command.ArgDispense")
	proto.RegisterType((*Command_ArgSetGiftCredit)(nil), "tele.Command.ArgSetGiftCredit")
	proto.RegisterType((*Response)(nil), "tele.Response")
	proto.RegisterEnum("tele.State", State_name, State_value)
	proto.RegisterEnum("tele.PaymentMethod", PaymentMethod_name, PaymentMethod_value)
}

func init() { proto.RegisterFile("tele.proto", fileDescriptor_tele_a7e94f26f60f621e) }

var fileDescriptor_tele_a7e94f26f60f621e = []byte{
	// 829 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xe1, 0x6a, 0x1b, 0x47,
	0x10, 0xb6, 0xa4, 0x3b, 0x59, 0x9a, 0xb3, 0xdc, 0xcb, 0x26, 0x4d, 0x2f, 0x57, 0xda, 0xaa, 0x86,
	0x16, 0xe1, 0x82, 0xa0, 0x6e, 0x21, 0x25, 0x50, 0x4a, 0xec, 0x84, 0xd8, 0x84, 0x94, 0xb0, 0x36,
	0xf4, 0xa7, 0x58, 0xdd, 0x8d, 0xed, 0xab, 0xef, 0x6e, 0x8f, 0xdd, 0xb5, 0x8a, 0xde, 0xa3, 0x4f,
	0xd0, 0x47, 0xe9, 0x2b, 0xf4, 0x85, 0xca, 0xec, 0xae, 0xa4, 0x8b, 0xe4, 0x16, 0xf2, 0x6f, 0x67,
	0xe6, 0xfb, 0xbe, 0xd9, 0x99, 0x9d, 0x9b, 0x03, 0x30, 0x58, 0xe2, 0xb4, 0x51, 0xd2, 0x48, 0x16,
	0xd0, 0xf9, 0xe8, 0xef, 0x01, 0x0c, 0xaf, 0xb0, 0xc4, 0x0a, 0x8d, 0x5a, 0xb2, 0xc7, 0x10, 0x2e,
	0xaa, 0x59, 0x91, 0x27, 0x9d, 0x71, 0x67, 0x12, 0xf2, 0x60, 0x51, 0x5d, 0xe4, 0xec, 0x3b, 0x08,
	0x51, 0x29, 0xa9, 0x92, 0xee, 0xb8, 0x33, 0x89, 0x4e, 0x3e, 0x9d, 0x5a, 0x91, 0x35, 0x69, 0xfa,
	0x9a, 0x82, 0xdc, 0x61, 0xd8, 0x73, 0x18, 0x16, 0xf5, 0x02, 0x6b, 0x23, 0xd5, 0x32, 0xe9, 0x59,
	0xc2, 0xb3, 0x6d, 0xc2, 0xc5, 0x0a, 0xc0, 0x37, 0x58, 0xca, 0x52, 0xc9, 0x1a, 0x97, 0x49, 0xf0,
	0x70, 0x96, 0x77, 0x14, 0xe4, 0x0e, 0xc3, 0x7e, 0x86, 0xc8, 0x28, 0x51, 0x6b, 0x91, 0x99, 0x42,
	0xd6, 0x49, 0x68, 0x29, 0x9f, 0x6f, 0x53, 0xae, 0x36, 0x10, 0xde, 0xc6, 0xb3, 0x09, 0x04, 0xda,
	0x08, 0x93, 0xf4, 0x2d, 0xef, 0xc9, 0x36, 0xef, 0xd2, 0x08, 0xc3, 0x2d, 0x22, 0x7d, 0x0b, 0xa1,
	0x2d, 0x8f, 0x31, 0x08, 0x32, 0x99, 0xa3, 0x6d, 0xcc, 0x88, 0xdb, 0x33, 0x4b, 0x60, 0xbf, 0x42,
	0xad, 0xc5, 0x0d, 0xda, 0xd6, 0x0c, 0xf9, 0xca, 0x64, 0x4f, 0x20, 0xcc, 0xe4, 0x7d, 0x6d, 0x6c,
	0x07, 0x46, 0xdc, 0x19, 0x29, 0xc2, 0x70, 0x5d, 0x3a, 0x41, 0xfe, 0x10, 0x06, 0x95, 0x6f, 0xb5,
	0x33, 0xd8, 0x53, 0xe8, 0x67, 0xf2, 0xfa, 0x1a, 0x9d, 0x62, 0xc8, 0xbd, 0xc5, 0x62, 0xe8, 0x65,
	0xf7, 0x8d, 0x95, 0x0b, 0x39, 0x1d, 0x29, 0xf9, 0xad, 0x6c, 0x1a, 0x54, 0x3a, 0x09, 0xc6, 0xbd,
	0x49, 0xc8, 0x57, 0x66, 0x7a, 0x0f, 0xa1, 0x6d, 0x16, 0xfb, 0x0a, 0x22, 0x23, 0x8d, 0x28, 0x67,
	0xf3, 0xa2, 0x2c, 0xb5, 0xbf, 0x3a, 0x58, 0xd7, 0x29, 0x79, 0x36, 0x80, 0x4c, 0x16, 0xb5, 0xb6,
	0x29, 0x57, 0x80, 0x33, 0xf2, 0xd0, 0x25, 0x1d, 0xb7, 0x37, 0xee, 0x51, 0x1d, 0xd6, 0x70, 0xd5,
	0x11, 0x21, 0x70, 0x5e, 0x6b, 0xa4, 0xff, 0x74, 0x20, 0x6a, 0x75, 0xfc, 0x83, 0x8e, 0x85, 0x9b,
	0x8e, 0xc9, 0x86, 0xa2, 0x94, 0xcc, 0x5e, 0xda, 0x9b, 0xa4, 0xd9, 0xa8, 0x22, 0xc3, 0x55, 0xc7,
	0xac, 0xc1, 0x5e, 0xc0, 0x61, 0x23, 0x96, 0x15, 0xd6, 0x66, 0x56, 0xa1, 0xb9, 0x95, 0xb9, 0x9d,
	0x8e, 0xc3, 0x93, 0xc7, 0xee, 0xc9, 0xde, 0xbb, 0xd8, 0x3b, 0x1b, 0xe2, 0xa3, 0xa6, 0x6d, 0xb2,
	0xaf, 0xe1, 0x20, 0x53, 0x98, 0x17, 0xc6, 0x97, 0x1f, 0x5a, 0xe1, 0xc8, 0xf9, 0x5c, 0xfd, 0x1b,
	0x88, 0xab, 0xa7, 0xdf, 0x86, 0xd8, 0x0e, 0xa4, 0x7f, 0x75, 0x21, 0xa0, 0x79, 0x60, 0x6f, 0x61,
	0x44, 0x3a, 0x33, 0x85, 0xbf, 0x63, 0x66, 0x30, 0x4f, 0xe2, 0x71, 0x6f, 0x12, 0x9d, 0x7c, 0xfb,
	0xd0, 0xf0, 0x4c, 0x49, 0x9e, 0x7b, 0xe0, 0xeb, 0xda, 0xa8, 0x25, 0x3f, 0x98, 0xb7, 0x5c, 0x24,
	0x46, 0x19, 0x37, 0x62, 0x8f, 0xfe, 0x47, 0x8c, 0x2e, 0xb2, 0x25, 0x96, 0xb5, 0x5c, 0xe9, 0x2f,
	0xf0, 0x68, 0x27, 0x1f, 0x0d, 0xcc, 0x1d, 0x2e, 0xfd, 0x9b, 0xd3, 0x91, 0x3a, 0xbc, 0x10, 0xe5,
	0x3d, 0xfa, 0x67, 0x76, 0xc6, 0x8b, 0xee, 0x4f, 0x1d, 0x12, 0xd8, 0xc9, 0xf1, 0x31, 0x02, 0x47,
	0x7f, 0xf6, 0x60, 0xff, 0x4c, 0x56, 0x95, 0xa8, 0x73, 0x76, 0x08, 0x5d, 0xbf, 0x3f, 0x46, 0xbc,
	0x5b, 0xe4, 0x34, 0x63, 0x0a, 0x9b, 0x72, 0x39, 0x33, 0xb2, 0x29, 0x32, 0xff, 0xa1, 0x80, 0x75,
	0x5d, 0x91, 0x87, 0x7d, 0x0f, 0x7d, 0x85, 0x8d, 0x54, 0xc6, 0xaf, 0x8b, 0xcf, 0x5c, 0x13, 0xbc,
	0xde, 0xf4, 0xa5, 0xba, 0xe1, 0x36, 0x7c, 0xbe, 0xc7, 0x3d, 0x90, 0x4d, 0x21, 0x14, 0x73, 0x62,
	0xb8, 0x5d, 0xf1, 0x74, 0x87, 0xf1, 0x72, 0xee, 0x08, 0x0e, 0xc6, 0x9e, 0xc3, 0x20, 0x2f, 0x74,
	0x83, 0xb5, 0x46, 0xbf, 0x2b, 0x9e, 0xed, 0x50, 0x5e, 0x79, 0xc0, 0xf9, 0x1e, 0x5f, 0x83, 0xd9,
	0x39, 0x7c, 0xa2, 0xd1, 0xcc, 0x6e, 0x8a, 0x6b, 0x33, 0x73, 0x53, 0xe1, 0x77, 0xc6, 0x97, 0x3b,
	0xfc, 0x4b, 0x34, 0x6f, 0x8a, 0x6b, 0x73, 0x66, 0x51, 0xe7, 0x7b, 0x7c, 0xa4, 0xdb, 0x8e, 0x34,
	0x82, 0xe1, 0xba, 0x92, 0x14, 0x60, 0xb0, 0xba, 0x64, 0xfa, 0x0d, 0x44, 0xad, 0xec, 0xb4, 0x00,
	0x44, 0x65, 0x57, 0x87, 0x6b, 0xa1, 0xb7, 0xd2, 0x63, 0x88, 0xb7, 0x93, 0xfc, 0x17, 0xf6, 0xb4,
	0x0f, 0x81, 0x11, 0xfa, 0xee, 0xe8, 0x12, 0x06, 0x1c, 0x75, 0x23, 0x49, 0xf7, 0x0b, 0x80, 0xcc,
	0x5d, 0x76, 0xb6, 0x7e, 0x9e, 0xa1, 0xf7, 0x5c, 0xe4, 0xf4, 0xb6, 0x9b, 0x1d, 0x3f, 0x5c, 0x2d,
	0x73, 0x06, 0x41, 0x2e, 0x8c, 0xb0, 0x0f, 0x33, 0xe4, 0xf6, 0x7c, 0xfc, 0x06, 0x42, 0x9a, 0x4a,
	0x64, 0x11, 0xec, 0x5f, 0xd4, 0x0b, 0x51, 0x16, 0x79, 0xbc, 0xc7, 0x06, 0x10, 0x9c, 0x4a, 0x69,
	0xe2, 0x0e, 0x9d, 0x7e, 0x93, 0xea, 0x2e, 0xee, 0xb2, 0x18, 0x0e, 0x5e, 0x15, 0x3a, 0x93, 0x75,
	0x6d, 0xc7, 0x2a, 0xee, 0x11, 0xe5, 0xbd, 0x92, 0xf3, 0x12, 0xab, 0x38, 0x38, 0xfe, 0x11, 0x46,
	0x1f, 0x7c, 0xbf, 0x14, 0xfd, 0x55, 0x9a, 0xdb, 0xa2, 0xbe, 0x71, 0x82, 0x67, 0x42, 0xdf, 0xc6,
	0x1d, 0x76, 0x00, 0x03, 0x3a, 0x95, 0xa8, 0x75, 0xdc, 0x9d, 0xf7, 0xed, 0xcf, 0xeb, 0x87, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xba, 0x9e, 0x75, 0x17, 0xca, 0x06, 0x00, 0x00,
}
