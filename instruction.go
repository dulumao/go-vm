package go_vm

type Instruction struct {
	opCode OpCode        // 操作码
	args   []*Value // 参数
}

type OpCode int

// 所有操作码
// 8位指令后缀是B、16位是S、32位是L、64位是Q
const (
	_      OpCode = iota
	CONSTQ        // CONSTQ 123  局部变量入栈帧
	ADDQ          // ADDQ  pop出两个值，相加，存入栈帧
	SUBQ          // SUBQ  pop出两个值，前-后，存入栈帧
	MULQ          // MULQ  pop出两个值，前*后，存入栈帧
	DIVQ          // DIVQ  pop出两个值，前/后，存入栈帧
	JMP           // JMP 6  跳到6的位置执行指令
	CALL          // CALL 6 2  跳到6位置执行，pop出2个参数。返回值入栈，返回值个数入栈
	RET           // RET  跳转到返回地址执行
	// 下面是预设函数
	PRINT // PRINT(d interface{}) 打印参数
	HALT  // HALT() 退出进程
)

var OpCodeToTokenType = map[OpCode]TokenType{
	CONSTQ: TokenType_CONSTQ,
	ADDQ:   TokenType_ADDQ,
	SUBQ:   TokenType_SUBQ,
	MULQ:   TokenType_MULQ,
	DIVQ:   TokenType_DIVQ,
	JMP:    TokenType_JMP,
	CALL:   TokenType_CALL,
	RET:    TokenType_RET,
	PRINT:  TokenType_PRINT,
	HALT:   TokenType_HALT,
}