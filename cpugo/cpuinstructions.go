package cpugo

func (cpu *Cpu) Execute(cycles u32, memory *Mem) {
	for cycles > 0 {
		ins := cpu.FetchByte(&cycles, memory)
		switch ins {
		case INS_LDA_IM:
			var value Byte = cpu.FetchByte(&cycles, memory)
			cpu.A = value
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ZP:
			var location Byte = cpu.FetchByte(&cycles, memory)
			cpu.A = memory.Access(location, &cycles)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ZPX:
			var location Byte = cpu.FetchByte(&cycles, memory)
			cpu.A = memory.Mem_map[location+cpu.X]
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		}
	}
}
