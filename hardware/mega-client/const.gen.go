// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package mega

const (
	// TWI_LISTEN_MAX_LENGTH as defined in mega-firmware/protocol.h:11
	TWI_LISTEN_MAX_LENGTH = 36
	// REQUEST_MAX_LENGTH as defined in mega-firmware/protocol.h:13
	REQUEST_MAX_LENGTH = 70
	// RESPONSE_MAX_LENGTH as defined in mega-firmware/protocol.h:24
	RESPONSE_MAX_LENGTH = 70
	// MDB_BLOCK_SIZE as defined in mega-firmware/protocol.h:58
	MDB_BLOCK_SIZE = 36
	// MDB_ACK as defined in mega-firmware/protocol.h:59
	MDB_ACK = 0
	// MDB_RET as defined in mega-firmware/protocol.h:60
	MDB_RET = 170
	// MDB_NAK as defined in mega-firmware/protocol.h:61
	MDB_NAK = 255
)

// COMMAND_STATUS as declared in mega-firmware/protocol.h:15
const COMMAND_STATUS Command_t = 1

// COMMAND_CONFIG as declared in mega-firmware/protocol.h:16
const COMMAND_CONFIG Command_t = 2

// COMMAND_RESET as declared in mega-firmware/protocol.h:17
const COMMAND_RESET Command_t = 3

// COMMAND_DEBUG as declared in mega-firmware/protocol.h:18
const COMMAND_DEBUG Command_t = 4

// COMMAND_FLASH as declared in mega-firmware/protocol.h:19
const COMMAND_FLASH Command_t = 5

// COMMAND_MDB_BUS_RESET as declared in mega-firmware/protocol.h:20
const COMMAND_MDB_BUS_RESET Command_t = 7

// COMMAND_MDB_TRANSACTION_SIMPLE as declared in mega-firmware/protocol.h:21
const COMMAND_MDB_TRANSACTION_SIMPLE Command_t = 8

// COMMAND_MDB_TRANSACTION_CUSTOM as declared in mega-firmware/protocol.h:22
const COMMAND_MDB_TRANSACTION_CUSTOM Command_t = 9

// RESPONSE_OK as declared in mega-firmware/protocol.h:26
const RESPONSE_OK Response_t = 1

// RESPONSE_RESET as declared in mega-firmware/protocol.h:27
const RESPONSE_RESET Response_t = 2

// RESPONSE_TWI_LISTEN as declared in mega-firmware/protocol.h:28
const RESPONSE_TWI_LISTEN Response_t = 3

// RESPONSE_ERROR as declared in mega-firmware/protocol.h:29
const RESPONSE_ERROR Response_t = 128

// ERROR_BAD_PACKET as declared in mega-firmware/protocol.h:32
const ERROR_BAD_PACKET Errcode_t = 1

// ERROR_INVALID_CRC as declared in mega-firmware/protocol.h:33
const ERROR_INVALID_CRC Errcode_t = 2

// ERROR_INVALID_ID as declared in mega-firmware/protocol.h:34
const ERROR_INVALID_ID Errcode_t = 3

// ERROR_UNKNOWN_COMMAND as declared in mega-firmware/protocol.h:35
const ERROR_UNKNOWN_COMMAND Errcode_t = 4

// ERROR_INVALID_DATA as declared in mega-firmware/protocol.h:36
const ERROR_INVALID_DATA Errcode_t = 5

// ERROR_BUFFER_OVERFLOW as declared in mega-firmware/protocol.h:37
const ERROR_BUFFER_OVERFLOW Errcode_t = 6

// ERROR_NOT_IMPLEMENTED as declared in mega-firmware/protocol.h:38
const ERROR_NOT_IMPLEMENTED Errcode_t = 7

// ERROR_RESPONSE_OVERWRITE as declared in mega-firmware/protocol.h:39
const ERROR_RESPONSE_OVERWRITE Errcode_t = 8

// ERROR_RESPONSE_EMPTY as declared in mega-firmware/protocol.h:40
const ERROR_RESPONSE_EMPTY Errcode_t = 9

// FIELD_INVALID as declared in mega-firmware/protocol.h:44
const FIELD_INVALID Field_t = 0

// FIELD_PROTOCOL as declared in mega-firmware/protocol.h:45
const FIELD_PROTOCOL Field_t = 1

// FIELD_FIRMWARE_VERSION as declared in mega-firmware/protocol.h:46
const FIELD_FIRMWARE_VERSION Field_t = 2

// FIELD_ERROR2 as declared in mega-firmware/protocol.h:47
const FIELD_ERROR2 Field_t = 3

// FIELD_ERRORN as declared in mega-firmware/protocol.h:48
const FIELD_ERRORN Field_t = 4

// FIELD_MCUSR as declared in mega-firmware/protocol.h:49
const FIELD_MCUSR Field_t = 5

// FIELD_CLOCK10U as declared in mega-firmware/protocol.h:50
const FIELD_CLOCK10U Field_t = 6

// FIELD_TWI_LENGTH as declared in mega-firmware/protocol.h:51
const FIELD_TWI_LENGTH Field_t = 7

// FIELD_TWI_DATA as declared in mega-firmware/protocol.h:52
const FIELD_TWI_DATA Field_t = 8

// FIELD_MDB_LENGTH as declared in mega-firmware/protocol.h:53
const FIELD_MDB_LENGTH Field_t = 9

// FIELD_MDB_RESULT as declared in mega-firmware/protocol.h:54
const FIELD_MDB_RESULT Field_t = 10

// FIELD_MDB_DATA as declared in mega-firmware/protocol.h:55
const FIELD_MDB_DATA Field_t = 11

// FIELD_MDB_DURATION10U as declared in mega-firmware/protocol.h:56
const FIELD_MDB_DURATION10U Field_t = 12

// MDB_STATE_IDLE as declared in mega-firmware/protocol.h:64
const MDB_STATE_IDLE byte = 0

// MDB_STATE_ERROR as declared in mega-firmware/protocol.h:65
const MDB_STATE_ERROR byte = 1

// MDB_STATE_SEND as declared in mega-firmware/protocol.h:66
const MDB_STATE_SEND byte = 2

// MDB_STATE_RECV as declared in mega-firmware/protocol.h:67
const MDB_STATE_RECV byte = 3

// MDB_STATE_RECV_END as declared in mega-firmware/protocol.h:68
const MDB_STATE_RECV_END byte = 4

// MDB_STATE_BUS_RESET as declared in mega-firmware/protocol.h:69
const MDB_STATE_BUS_RESET byte = 5

// MDB_STATE_DONE as declared in mega-firmware/protocol.h:70
const MDB_STATE_DONE byte = 6

// MDB_RESULT_SUCCESS as declared in mega-firmware/protocol.h:73
const MDB_RESULT_SUCCESS Mdb_result_t = 1

// MDB_RESULT_BUSY as declared in mega-firmware/protocol.h:74
const MDB_RESULT_BUSY Mdb_result_t = 8

// MDB_RESULT_INVALID_CHK as declared in mega-firmware/protocol.h:75
const MDB_RESULT_INVALID_CHK Mdb_result_t = 9

// MDB_RESULT_NAK as declared in mega-firmware/protocol.h:76
const MDB_RESULT_NAK Mdb_result_t = 10

// MDB_RESULT_TIMEOUT as declared in mega-firmware/protocol.h:77
const MDB_RESULT_TIMEOUT Mdb_result_t = 11

// MDB_RESULT_INVALID_END as declared in mega-firmware/protocol.h:78
const MDB_RESULT_INVALID_END Mdb_result_t = 12

// MDB_RESULT_RECEIVE_OVERFLOW as declared in mega-firmware/protocol.h:79
const MDB_RESULT_RECEIVE_OVERFLOW Mdb_result_t = 13

// MDB_RESULT_SEND_OVERFLOW as declared in mega-firmware/protocol.h:80
const MDB_RESULT_SEND_OVERFLOW Mdb_result_t = 14

// MDB_RESULT_CODE_ERROR as declared in mega-firmware/protocol.h:81
const MDB_RESULT_CODE_ERROR Mdb_result_t = 15

// MDB_RESULT_UART_READ_UNEXPECTED as declared in mega-firmware/protocol.h:82
const MDB_RESULT_UART_READ_UNEXPECTED Mdb_result_t = 16

// MDB_RESULT_UART_READ_ERROR as declared in mega-firmware/protocol.h:83
const MDB_RESULT_UART_READ_ERROR Mdb_result_t = 17

// MDB_RESULT_UART_READ_OVERFLOW as declared in mega-firmware/protocol.h:84
const MDB_RESULT_UART_READ_OVERFLOW Mdb_result_t = 18

// MDB_RESULT_UART_READ_PARITY as declared in mega-firmware/protocol.h:85
const MDB_RESULT_UART_READ_PARITY Mdb_result_t = 19

// MDB_RESULT_UART_SEND_BUSY as declared in mega-firmware/protocol.h:86
const MDB_RESULT_UART_SEND_BUSY Mdb_result_t = 20

// MDB_RESULT_UART_TXC_UNEXPECTED as declared in mega-firmware/protocol.h:87
const MDB_RESULT_UART_TXC_UNEXPECTED Mdb_result_t = 21

// MDB_RESULT_TIMER_CODE_ERROR as declared in mega-firmware/protocol.h:88
const MDB_RESULT_TIMER_CODE_ERROR Mdb_result_t = 24
