package go_vm

type Instruction struct {
	opCode OpCode   // 操作码
	args   []*Value // 参数
}

type OpCode int

func (opCode OpCode) String() string {
	return OpCodeToString[opCode]
}

// 所有操作码
const (
	_     OpCode = iota
	CONST        // CONST 123  局部变量入栈帧
	ADD          // ADD  pop出两个值，相加，存入栈帧
	SUB          // SUB  pop出两个值，前-后，存入栈帧
	MUL          // MUL  pop出两个值，前*后，存入栈帧
	DIV          // DIV  pop出两个值，前/后，存入栈帧
	JMP          // JMP 6  跳到6的位置执行指令
	JEQ          // 相等跳转
	JNE          // 不相等跳转
	CALL         // CALL 6 2  跳到6位置执行，上一个栈帧中pop出2个参数放到新栈帧。返回值入栈，返回值个数入栈
	RET          // RET 2  跳转到返回地址执行，2表示返回两个参数
	// 下面是预设函数
	PRINT // PRINT(d interface{}) 打印参数
	HALT  // HALT() 退出进程
)

var OpCodeToString = map[OpCode]string{
	CONST: "CONST",
	ADD:   "ADD",
	SUB:   "SUB",
	MUL:   "MUL",
	DIV:   "DIV",
	JMP:   "JMP",
	JEQ:   "JEQ",
	JNE:   "JNE",
	CALL:  "CALL",
	RET:   "RET",
	PRINT: "PRINT",
	HALT:  "HALT",
}

var OpCodeToTokenType = map[OpCode]TokenType{
	CONST: TokenType_CONST,
	ADD:   TokenType_ADD,
	SUB:   TokenType_SUB,
	MUL:   TokenType_MUL,
	DIV:   TokenType_DIV,
	JMP:   TokenType_JMP,
	JEQ:   TokenType_JEQ,
	JNE:   TokenType_JNE,
	CALL:  TokenType_CALL,
	RET:   TokenType_RET,
	PRINT: TokenType_PRINT,
	HALT:  TokenType_HALT,
}
