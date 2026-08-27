package cpugo

import "fmt"

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
	for i := Utt(0); i < MAX_MEM; i++ {
		m.Mem_map[i] = 0
	}
}

// whilist Fetchbyte uses the PC in the cpu struct, mem.Access uses an external address
func (m *Mem) Access(address Tbyte, cycles *Utt) Byte {
	data := m.Mem_map[address]
	*cycles--
	return data
}

// whilist Fetchbyte uses the PC in the cpu struct, mem.Access uses an external address
func (m *Mem) AccessZp(address Byte, cycles *Utt) Byte {
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
func (cpu *Cpu) FetchBytePc(cycles *Utt, memory *Mem) Byte {
	var data Byte = memory.Mem_map[cpu.PC]
	cpu.PC++
	*cycles--
	return data
}

// fetches the byte @zpaddress without
func (cpu *Cpu) FetchBytezp(cycles *Utt, memory *Mem, zpaddress byte) Byte {
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
