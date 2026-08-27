package cpugo

func (cpu *Cpu) Execute(cycles *U32, memory *Mem) {
	for *cycles > 0 {
		ins := cpu.FetchBytePc(cycles, memory)
		switch ins {
		case INS_LDA_IM:
			var value Byte = cpu.FetchBytePc(cycles, memory)
			cpu.A = value
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ZP:
			var location Byte = cpu.FetchBytePc(cycles, memory)
			cpu.A = memory.AccessZp(location, cycles)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ZPX:
			var location Byte = cpu.FetchBytePc(cycles, memory)
			cpu.A = memory.AccessZp(location+cpu.X, cycles)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ABS:
			var locationll Byte = cpu.FetchBytePc(cycles, memory)
			var locationhh Byte = cpu.FetchBytePc(cycles, memory)
			var actuallocation Tbyte = (Tbyte(locationhh)<<8 | Tbyte(locationll+cpu.Y))
			cpu.A = memory.Access(actuallocation, cycles)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ABSX:
			var locationll Byte = cpu.FetchBytePc(cycles, memory)
			var locationhh Byte = cpu.FetchBytePc(cycles, memory)
			var actuallocation Tbyte = (Tbyte(locationhh)<<8 | Tbyte(locationll+cpu.Y))
			cpu.A = memory.Access(actuallocation, cycles)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_ABSY:
			var locationll Byte = cpu.FetchBytePc(cycles, memory)
			var locationhh Byte = cpu.FetchBytePc(cycles, memory)
			var actuallocation Tbyte = (Tbyte(locationhh)<<8 | Tbyte(locationll+cpu.Y))
			cpu.A = memory.Access(actuallocation, cycles)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_INDX:
			var zpaddress Byte = cpu.FetchBytePc(cycles, memory)
			var address Tbyte = Tbyte(Tbyte(memory.AccessZp((zpaddress+cpu.X+1), cycles))<<8 | Tbyte(memory.AccessZp((zpaddress+cpu.X), cycles)))
			cpu.PC = address
			(*cycles)--
			cpu.A = cpu.FetchBytePc(cycles, memory)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		case INS_LDA_INDY:
			var zpaddress Byte = cpu.FetchBytePc(cycles, memory)
			var address Tbyte = Tbyte(Tbyte(memory.AccessZp(zpaddress+1, cycles))<<8 | (Tbyte(memory.AccessZp(zpaddress, cycles)) + Tbyte(cpu.Y)))
			cpu.PC = address
			cpu.A = cpu.FetchBytePc(cycles, memory)
			cpu.PRSR.Z = (cpu.A == 0)
			cpu.PRSR.N = (cpu.A & 0b10000000) > 0
		}
	}
}
