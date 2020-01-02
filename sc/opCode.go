package sc

//https://github.com/neo-project/neo-vm/blob/master/src/neo-vm/OpCode.cs

type OpCode byte

const (
	// Constants
	PUSHINT8   OpCode = 0x00 // Operand Size = 1
	PUSHINT16  OpCode = 0x01 // Operand Size = 2
	PUSHINT32  OpCode = 0x02 // Operand Size = 4
	PUSHINT64  OpCode = 0x03 // Operand Size = 8
	PUSHINT128 OpCode = 0x04 // Operand Size = 16
	PUSHINT256 OpCode = 0x05 // Operand Size = 32
	PUSHA      OpCode = 0x0A // Convert the next four bytes to an address, and push the address onto the stack.
	PUSHNULL   OpCode = 0x0B // "null" is pushed onto the stack.
	PUSHDATA1  OpCode = 0x0C // Operand SizePrefix = 1. The next byte contains the number of bytes to be pushed onto the stack.
	PUSHDATA2  OpCode = 0x0D // Operand SizePrefix = 2. The next two bytes contains the number of bytes to be pushed onto the stack.
	PUSHDATA4  OpCode = 0x0E // Operand SizePrefix = 4. The next four bytes contains the number of bytes to be pushed onto the stack.
	PUSHM1     OpCode = 0x0F // The number -1 is pushed onto the stack.
	PUSH0      OpCode = 0x10 // The number 0 is pushed onto the stack.
	PUSH1      OpCode = 0x11 // The number 1 is pushed onto the stack.
	PUSH2      OpCode = 0x12 // The number 2 is pushed onto the stack.
	PUSH3      OpCode = 0x13 // The number 3 is pushed onto the stack.
	PUSH4      OpCode = 0x14 // The number 4 is pushed onto the stack.
	PUSH5      OpCode = 0x15 // The number 5 is pushed onto the stack.
	PUSH6      OpCode = 0x16 // The number 6 is pushed onto the stack.
	PUSH7      OpCode = 0x17 // The number 7 is pushed onto the stack.
	PUSH8      OpCode = 0x18 // The number 8 is pushed onto the stack.
	PUSH9      OpCode = 0x19 // The number 9 is pushed onto the stack.
	PUSH10     OpCode = 0x1A // The number 10 is pushed onto the stack.
	PUSH11     OpCode = 0x1B // The number 11 is pushed onto the stack.
	PUSH12     OpCode = 0x1C // The number 12 is pushed onto the stack.
	PUSH13     OpCode = 0x1D // The number 13 is pushed onto the stack.
	PUSH14     OpCode = 0x1E // The number 14 is pushed onto the stack.
	PUSH15     OpCode = 0x1F // The number 15 is pushed onto the stack.
	PUSH16     OpCode = 0x20 // The number 16 is pushed onto the stack.

	// Flow control
	NOP        OpCode = 0x21 // The "NOP" operation does nothing. It is intended to fill in space if opcodes are patched.
	JMP        OpCode = 0x22 // Operand Size = 1. Unconditionally transfers control to a target instruction. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMP_L      OpCode = 0x23 // Operand Size = 4. Unconditionally transfers control to a target instruction. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPIF      OpCode = 0x24 // Operand Size = 1. Transfers control to a target instruction if the value is "true", not "null", or non-zero. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPIF_L    OpCode = 0x25 // Operand Size = 4. Transfers control to a target instruction if the value is "true", not "null", or non-zero. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPIFNOT   OpCode = 0x26 // Operand Size = 1. Transfers control to a target instruction if the value is "false", a "null" reference, or zero. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPIFNOT_L OpCode = 0x27 // Operand Size = 4. Transfers control to a target instruction if the value is "false", a "null" reference, or zero. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPEQ      OpCode = 0x28 // Operand Size = 1. Transfers control to a target instruction if two values are equal. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPEQ_L    OpCode = 0x29 // Operand Size = 4. Transfers control to a target instruction if two values are equal. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPNE      OpCode = 0x2A // Operand Size = 1. Transfers control to a target instruction when two values are not equal. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPNE_L    OpCode = 0x2B // Operand Size = 4. Transfers control to a target instruction when two values are not equal. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPGT      OpCode = 0x2C // Operand Size = 1. Transfers control to a target instruction if the first value is greater than the second value. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPGT_L    OpCode = 0x2D // Operand Size = 4. Transfers control to a target instruction if the first value is greater than the second value. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPGE      OpCode = 0x2E // Operand Size = 1. Transfers control to a target instruction if the first value is greater than or equal to the second value. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPGE_L    OpCode = 0x2F // Operand Size = 4. Transfers control to a target instruction if the first value is greater than or equal to the second value. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPLT      OpCode = 0x30 // Operand Size = 1. Transfers control to a target instruction if the first value is less than the second value. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPLT_L    OpCode = 0x31 // Operand Size = 4. Transfers control to a target instruction if the first value is less than the second value. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	JMPLE      OpCode = 0x32 // Operand Size = 1. Transfers control to a target instruction if the first value is less than or equal to the second value. The target instruction is represented as a 1-byte signed offset from the beginning of the current instruction.
	JMPLE_L    OpCode = 0x33 // Operand Size = 4. Transfers control to a target instruction if the first value is less than or equal to the second value. The target instruction is represented as a 4-bytes signed offset from the beginning of the current instruction.
	CALL       OpCode = 0x34 // Operand Size = 1. Calls the function at the target address which is represented as a 1-byte signed offset from the beginning of the current instruction.
	CALL_L     OpCode = 0x35 // Operand Size = 4. Calls the function at the target address which is represented as a 4-bytes signed offset from the beginning of the current instruction.
	CALLA      OpCode = 0x36 // Pop the address of a function from the stack, and call the function.
	THROW      OpCode = 0x37
	THROWIF    OpCode = 0x38
	THROWIFNOT OpCode = 0x39
	//TRY        OpCode = 0x3B
	//TRY_L      OpCode = 0x3C
	//ENDT       OpCode = 0x3D
	//ENDC       OpCode = 0x3E
	//ENDF       OpCode = 0x3F
	RET     OpCode = 0x40 // Returns from the current method.
	SYSCALL OpCode = 0x41 // Operand Size = 4. Calls to an interop service.

	// Stack
	DEPTH    OpCode = 0x43 // Puts the number of stack items onto the stack.
	DROP     OpCode = 0x45 // Removes the top stack item.
	NIP      OpCode = 0x46 // Removes the second-to-top stack item.
	XDROP    OpCode = 0x48 // The item n back in the main stack is removed.
	CLEAR    OpCode = 0x49 // Clear the stack.
	DUP      OpCode = 0x4A // Duplicates the top stack item.
	OVER     OpCode = 0x4B // Copies the second-to-top stack item to the top.
	PICK     OpCode = 0x4D // The item n back in the stack is copied to the top.
	TUCK     OpCode = 0x4E // The item at the top of the stack is copied and inserted before the second-to-top item.
	SWAP     OpCode = 0x50 // The top two items on the stack are swapped.
	ROT      OpCode = 0x51 // The top three items on the stack are rotated to the left.
	ROLL     OpCode = 0x52 // The item n back in the stack is moved to the top.
	REVERSE3 OpCode = 0x53 // Reverse the order of the top 3 items on the stack.
	REVERSE4 OpCode = 0x54 // Reverse the order of the top 4 items on the stack.
	REVERSEN OpCode = 0x55 // Pop the number N on the stack, and reverse the order of the top N items on the stack.

	// Slot
	INITSSLOT OpCode = 0x56 // Operand Size = 1. Initialize the static field list for the current execution context.
	INITSLOT  OpCode = 0x57 // Operand Size = 2. Initialize the argument slot and the local variable list for the current execution context.
	LDSFLD0   OpCode = 0x58 // Loads the static field at index 0 onto the evaluation stack.
	LDSFLD1   OpCode = 0x59 // Loads the static field at index 1 onto the evaluation stack.
	LDSFLD2   OpCode = 0x5A // Loads the static field at index 2 onto the evaluation stack.
	LDSFLD3   OpCode = 0x5B // Loads the static field at index 3 onto the evaluation stack.
	LDSFLD4   OpCode = 0x5C // Loads the static field at index 4 onto the evaluation stack.
	LDSFLD5   OpCode = 0x5D // Loads the static field at index 5 onto the evaluation stack.
	LDSFLD6   OpCode = 0x5E // Loads the static field at index 6 onto the evaluation stack.
	LDSFLD    OpCode = 0x5F // Operand Size = 1. Loads the static field at a specified index onto the evaluation stack. The index is represented as a 1-byte unsigned integer.
	STSFLD0   OpCode = 0x60 // Stores the value on top of the evaluation stack in the static field list at index 0.
	STSFLD1   OpCode = 0x61 // Stores the value on top of the evaluation stack in the static field list at index 1.
	STSFLD2   OpCode = 0x62 // Stores the value on top of the evaluation stack in the static field list at index 2.
	STSFLD3   OpCode = 0x63 // Stores the value on top of the evaluation stack in the static field list at index 3.
	STSFLD4   OpCode = 0x64 // Stores the value on top of the evaluation stack in the static field list at index 4.
	STSFLD5   OpCode = 0x65 // Stores the value on top of the evaluation stack in the static field list at index 5.
	STSFLD6   OpCode = 0x66 // Stores the value on top of the evaluation stack in the static field list at index 6.
	STSFLD    OpCode = 0x67 // Operand Size = 1. Stores the value on top of the evaluation stack in the static field list at a specified index. The index is represented as a 1-byte unsigned integer.
	LDLOC0    OpCode = 0x68 // Loads the local variable at index 0 onto the evaluation stack.
	LDLOC1    OpCode = 0x69 // Loads the local variable at index 1 onto the evaluation stack.
	LDLOC2    OpCode = 0x6A // Loads the local variable at index 2 onto the evaluation stack.
	LDLOC3    OpCode = 0x6B // Loads the local variable at index 3 onto the evaluation stack.
	LDLOC4    OpCode = 0x6C // Loads the local variable at index 4 onto the evaluation stack.
	LDLOC5    OpCode = 0x6D // Loads the local variable at index 5 onto the evaluation stack.
	LDLOC6    OpCode = 0x6E // Loads the local variable at index 6 onto the evaluation stack.
	LDLOC     OpCode = 0x6F // Operand Size = 1. Loads the local variable at a specified index onto the evaluation stack. The index is represented as a 1-byte unsigned integer.
	STLOC0    OpCode = 0x70 // Stores the value on top of the evaluation stack in the local variable list at index 0.
	STLOC1    OpCode = 0x71 // Stores the value on top of the evaluation stack in the local variable list at index 1.
	STLOC2    OpCode = 0x72 // Stores the value on top of the evaluation stack in the local variable list at index 2.
	STLOC3    OpCode = 0x73 // Stores the value on top of the evaluation stack in the local variable list at index 3.
	STLOC4    OpCode = 0x74 // Stores the value on top of the evaluation stack in the local variable list at index 4.
	STLOC5    OpCode = 0x75 // Stores the value on top of the evaluation stack in the local variable list at index 5.
	STLOC6    OpCode = 0x76 // Stores the value on top of the evaluation stack in the local variable list at index 6.
	STLOC     OpCode = 0x77 // Operand Size = 1. Stores the value on top of the evaluation stack in the local variable list at a specified index. The index is represented as a 1-byte unsigned integer.
	LDARG0    OpCode = 0x78 // Loads the argument at index 0 onto the evaluation stack.
	LDARG1    OpCode = 0x79 // Loads the argument at index 1 onto the evaluation stack.
	LDARG2    OpCode = 0x7A // Loads the argument at index 2 onto the evaluation stack.
	LDARG3    OpCode = 0x7B // Loads the argument at index 3 onto the evaluation stack.
	LDARG4    OpCode = 0x7C // Loads the argument at index 4 onto the evaluation stack.
	LDARG5    OpCode = 0x7D // Loads the argument at index 5 onto the evaluation stack.
	LDARG6    OpCode = 0x7E // Loads the argument at index 6 onto the evaluation stack.
	LDARG     OpCode = 0x7F // Operand Size = 1. Loads the argument at a specified index onto the evaluation stack. The index is represented as a 1-byte unsigned integer.
	STARG0    OpCode = 0x80 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG1    OpCode = 0x81 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG2    OpCode = 0x82 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG3    OpCode = 0x83 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG4    OpCode = 0x84 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG5    OpCode = 0x85 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG6    OpCode = 0x86 // Stores the value on top of the evaluation stack in the argument slot at index 0.
	STARG     OpCode = 0x87 // Operand Size = 1. Stores the value on top of the evaluation stack in the argument slot at a specified index. The index is represented as a 1-byte unsigned integer.

	// Splice
	NEWBUFFER OpCode = 0x88
	MEMCPY    OpCode = 0x89
	CAT       OpCode = 0x8B // Concatenates two strings.
	SUBSTR    OpCode = 0x8C // Returns a section of a string.
	LEFT      OpCode = 0x8D // Keeps only characters left of the specified point in a string.
	RIGHT     OpCode = 0x8E // Keeps only characters right of the specified point in a string.

	// Bitwise logic
	INVERT   OpCode = 0x90 // Flips all of the bits in the input.
	AND      OpCode = 0x91 // Boolean and between each bit in the inputs.
	OR       OpCode = 0x92 // Boolean or between each bit in the inputs.
	XOR      OpCode = 0x93 // Boolean exclusive or between each bit in the inputs.
	EQUAL    OpCode = 0x97 // Returns 1 if the inputs are exactly equal, 0 otherwise.
	NOTEQUAL OpCode = 0x98 // Returns 1 if the inputs are not equal, 0 otherwise.

	// Arithmetic
	SIGN        OpCode = 0x99 // Puts the sign of top stack item on top of the main stack. If value is negative, put -1; if positive, put 1; if value is zero, put 0.
	ABS         OpCode = 0x9A // The input is made positive.
	NEGATE      OpCode = 0x9B // The sign of the input is flipped.
	INC         OpCode = 0x9C // 1 is added to the input.
	DEC         OpCode = 0x9D // 1 is subtracted from the input.
	ADD         OpCode = 0x9E // a is added to b.
	SUB         OpCode = 0x9F // b is subtracted from a.
	MUL         OpCode = 0xA0 // a is multiplied by b.
	DIV         OpCode = 0xA1 // a is divided by b.
	MOD         OpCode = 0xA2 // Returns the remainder after dividing a by b.
	SHL         OpCode = 0xA8 // Shifts a left b bits preserving sign.
	SHR         OpCode = 0xA9 // Shifts a right b bits preserving sign.
	NOT         OpCode = 0xAA // If the input is 0 or 1 it is flipped. Otherwise the output will be 0.
	BOOLAND     OpCode = 0xAB // If both a and b are not 0 the output is 1. Otherwise 0.
	BOOLOR      OpCode = 0xAC // If a or b is not 0 the output is 1. Otherwise 0.
	NZ          OpCode = 0xB1 // Returns 0 if the input is 0. 1 otherwise.
	NUMEQUAL    OpCode = 0xB3 // Returns 1 if the numbers are equal 0 otherwise.
	NUMNOTEQUAL OpCode = 0xB4 // Returns 1 if the numbers are not equal 0 otherwise.
	LT          OpCode = 0xB5 // Returns 1 if a is less than b, 0 otherwise.
	LE          OpCode = 0xB6 // Returns 1 if a is less than or equal to b, 0 otherwise.
	GT          OpCode = 0xB7 // Returns 1 if a is greater than b, 0 otherwise.
	GE          OpCode = 0xB8 // Returns 1 if a is greater than or equal to b, 0 otherwise.
	MIN         OpCode = 0xB9 // Returns the smaller of a and b.
	MAX         OpCode = 0xBA // Returns the larger of a and b.
	WITHIN      OpCode = 0xBB // Returns 1 if x is within the specified range (left-inclusive), 0 otherwise.

	// Compound-type
	PACK         OpCode = 0xC0 // A value n is taken from top of main stack. The next n items on main stack are removed, put inside n-sized array and this array is put on top of the main stack.
	UNPACK       OpCode = 0xC1 // An array is removed from top of the main stack. Its elements are put on top of the main stack (in reverse order) and the array size is also put on main stack.
	NEWARRAY0    OpCode = 0xC2 // An empty array (with size 0) is put on top of the main stack.
	NEWARRAY     OpCode = 0xC3 // A value n is taken from top of main stack. A null-filled array with size n is put on top of the main stack.
	NEWARRAY_T   OpCode = 0xC4 // Operand Size = 1. A value n is taken from top of main stack. An array of type T with size n is put on top of the main stack.
	NEWSTRUCT0   OpCode = 0xC5 // An empty struct (with size 0) is put on top of the main stack.
	NEWSTRUCT    OpCode = 0xC6 // A value n is taken from top of main stack. A zero-filled struct type with size n is put on top of the main stack.
	NEWMAP       OpCode = 0xC8 // A Map is created and put on top of the main stack.
	SIZE         OpCode = 0xCA // An array is removed from top of the main stack. Its size is put on top of the main stack.
	HASKEY       OpCode = 0xCB // An input index n (or key) and an array (or map) are removed from the top of the main stack. Puts True on top of main stack if array[n] (or map[n]) exist, and False otherwise.
	KEYS         OpCode = 0xCC // A map is taken from top of the main stack. The keys of this map are put on top of the main stack.
	VALUES       OpCode = 0xCD // A map is taken from top of the main stack. The values of this map are put on top of the main stack.
	PICKITEM     OpCode = 0xCE // An input index n (or key) and an array (or map) are taken from main stack. Element array[n] (or map[n]) is put on top of the main stack.
	APPEND       OpCode = 0xCF // The item on top of main stack is removed and appended to the second item on top of the main stack.
	SETITEM      OpCode = 0xD0 // A value v, index n (or key) and an array (or map) are taken from main stack. Attribution array[n]=v (or map[n]=v) is performed.
	REVERSEITEMS OpCode = 0xD1 // An array is removed from the top of the main stack and its elements are reversed.
	REMOVE       OpCode = 0xD2 // An input index n (or key) and an array (or map) are removed from the top of the main stack. Element array[n] (or map[n]) is removed.
	CLEARITEMS   OpCode = 0xD3 // Remove all the items from the compound-type.

	// Types
	ISNULL  OpCode = 0xD8 // Returns true if the input is null. Returns false otherwise.
	ISTYPE  OpCode = 0xD9 // Operand Size = 1. Returns true if the top item is of the specified type.
	CONVERT OpCode = 0xDB // Operand Size = 1. Converts the top item to the specified type.
)
