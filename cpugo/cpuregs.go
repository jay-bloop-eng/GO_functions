package cpugo

import "fmt"

const MAX_MEM u32 = 1024 * 64
const (
	INS_LDA_IM  = 0xA9
	INS_LDA_ZP  = 0xA5
	INS_LDA_ZPX = 0xB5
)

type (
	Byte  uint8
	Tbyte uint16
	u32   uint32
)

type PSR struct { // PSR = Processor status register
	C bool
	Z bool
	I bool
	D bool
	B bool
	V bool
	N bool
}

type Mem struct {
	Mem_map [MAX_MEM]Byte
}

func (m *Mem) Initialize() {
	for i := u32(0); i < MAX_MEM; i++ {
		m.Mem_map[i] = 0
	}
}

// whilist Fetchbyte uses the PC in the cpu struct, mem.Access uses an external address
func (m *Mem) Access(address Byte, cycles *u32) Byte {
	data := m.Mem_map[address]
	*cycles--
	return data
}

type Cpu struct {
	A, X, Y, SP Byte  //A = acumulator register, X and Y = index registers, SP = stack pointer
	PC          Tbyte //PC = program counter
	PRSR        PSR   //PRSR = is the CPU structures PSR

}

func (cpu *Cpu) Reset(memory *Mem) {
	cpu.PC = 0xFFFC
	cpu.SP = 0x00FF
	cpu.PRSR = PSR{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}
	cpu.A, cpu.X, cpu.Y = 0, 0, 0
	memory.Initialize()
}

// fetches the byte in PC++
func (cpu *Cpu) FetchByte(cycles *u32, memory *Mem) Byte {
	var data Byte = memory.Mem_map[cpu.PC]
	cpu.PC++
	*cycles--
	return data
}

// fetches the byte @zpaddress without
func (cpu *Cpu) FetchBytezp(cycles *u32, memory *Mem, zpaddress byte) Byte {
	var data Byte
	if zpaddress > 0b00001111 {
		fmt.Printf("invalid memory location/ NOT in 0p: %x", zpaddress)
	} else {
		data = memory.Mem_map[zpaddress]
		*cycles--
		return data
	}
	return data
}
