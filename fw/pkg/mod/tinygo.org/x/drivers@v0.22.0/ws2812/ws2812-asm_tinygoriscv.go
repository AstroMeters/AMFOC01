//go:build tinygo.riscv32
// +build tinygo.riscv32

package ws2812

// Warning: autogenerated file. Instead of modifying this file, change
// gen-ws2812.go and run "go generate".

import "runtime/interrupt"
import "unsafe"

/*
#include <stdint.h>

__attribute__((always_inline))
void ws2812_writeByte160(char c, uint32_t *portSet, uint32_t *portClear, uint32_t maskSet, uint32_t maskClear) {
	// Timings:
	// T0H: 56 - 58 cycles or 350.0ns - 362.5ns
	// T1H: 168 - 170 cycles or 1050.0ns - 1062.5ns
	// TLD: 184 -    cycles or 1150.0ns -
	uint32_t value = (uint32_t)c << 23;
	char i = 8;
	__asm__ __volatile__(
		"1: // send_bit\n"
		"\t  sw    %[maskSet], %[portSet]     // [1]   T0H and T0L start here\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  slli  %[value], %[value], 1      // [1]   shift value left by 1\n"
		"\t  bltz  %[value], 2f               // [1/3] skip_store\n"
		"\t  sw    %[maskClear], %[portClear] // [1]   T0H -> T0L transition\n"
		"\t2: // skip_store\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  sw    %[maskClear], %[portClear] // [1]   T1H -> T1L transition\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  addi  %[i], %[i], -1             // [1]\n"
		"\t  bnez  %[i], 1b                   // [1/3] send_bit\n"
	: [value]"+r"(value),
	  [i]"+r"(i)
	: [maskSet]"r"(maskSet),
	  [portSet]"m"(*portSet),
	  [maskClear]"r"(maskClear),
	  [portClear]"m"(*portClear));
}

__attribute__((always_inline))
void ws2812_writeByte320(char c, uint32_t *portSet, uint32_t *portClear, uint32_t maskSet, uint32_t maskClear) {
	// Timings:
	// T0H: 112 - 114 cycles or 350.0ns - 356.2ns
	// T1H: 336 - 338 cycles or 1050.0ns - 1056.2ns
	// TLD: 368 -    cycles or 1150.0ns -
	uint32_t value = (uint32_t)c << 23;
	char i = 8;
	__asm__ __volatile__(
		"1: // send_bit\n"
		"\t  sw    %[maskSet], %[portSet]     // [1]   T0H and T0L start here\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  slli  %[value], %[value], 1      // [1]   shift value left by 1\n"
		"\t  bltz  %[value], 2f               // [1/3] skip_store\n"
		"\t  sw    %[maskClear], %[portClear] // [1]   T0H -> T0L transition\n"
		"\t2: // skip_store\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  sw    %[maskClear], %[portClear] // [1]   T1H -> T1L transition\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  nop\n"
		"\t  addi  %[i], %[i], -1             // [1]\n"
		"\t  bnez  %[i], 1b                   // [1/3] send_bit\n"
	: [value]"+r"(value),
	  [i]"+r"(i)
	: [maskSet]"r"(maskSet),
	  [portSet]"m"(*portSet),
	  [maskClear]"r"(maskClear),
	  [portClear]"m"(*portClear));
}
*/
import "C"

func (d Device) writeByte160(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	mask := interrupt.Disable()
	C.ws2812_writeByte160(C.char(c), (*uint32)(unsafe.Pointer(portSet)), (*uint32)(unsafe.Pointer(portClear)), maskSet, maskClear)

	interrupt.Restore(mask)
}

func (d Device) writeByte320(c byte) {
	portSet, maskSet := d.Pin.PortMaskSet()
	portClear, maskClear := d.Pin.PortMaskClear()

	mask := interrupt.Disable()
	C.ws2812_writeByte320(C.char(c), (*uint32)(unsafe.Pointer(portSet)), (*uint32)(unsafe.Pointer(portClear)), maskSet, maskClear)

	interrupt.Restore(mask)
}
