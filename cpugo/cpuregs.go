package cpugo

const MAX_MEM u32 = 1024 * 64
const (
	INS_LDA_IM = 0xA9
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
func (cpu *Cpu) FetchByte(cycles *u32, memory *Mem) Byte {
	var data Byte = memory.Mem_map[cpu.PC]
	cpu.PC++
	*cycles--
	return data
}
func (cpu *Cpu) Execute(cycles u32, memory *Mem) {
	for cycles > 0 {
		ins := cpu.FetchByte(&cycles, memory)
		switch ins {
		case INS_LDA_IM:
			var value Byte = cpu.FetchByte(&cycles, memory)
			cpu.A = value
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		}
	}
}
