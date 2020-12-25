package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMachineWithCode(t *testing.T) {
	code := []string{
		"nop +4",
		"acc +6",
		"jmp -4",
	}

	expectedSeen := []bool{false, false, false}
	expectedProgram := []Instruction{
		{NoOp, 4},
		{Accumulate, 6},
		{Jump, -4},
	}

	program, err := ParseCodeToProgram(code)
	assert.NoError(t, err)

	machine := NewMachine()
	machine.LoadProgram(program)

	assert.Equal(t, 0, machine.ProgramCounter)
	assert.Equal(t, 0, machine.Accumulator)
	assert.Equal(t, expectedProgram, machine.Program)
	assert.Equal(t, expectedSeen, machine.Seen)
}

func TestParseCodeToProgram2(t *testing.T) {
	code := []string{
		"nop +4",
		"acc +6",
		"jmp -4",
	}

	expectedProgram := []Instruction{
		{NoOp, 4},
		{Accumulate, 6},
		{Jump, -4},
	}

	program, err := ParseCodeToProgram(code)
	assert.NoError(t, err)
	assert.Equal(t, expectedProgram, program)
}

func TestParseCodeLineToInstruction(t *testing.T) {
	successCases := []struct {
		codeLine string
		command  InstructionCode
		argument int
	}{
		{"nop +4", NoOp, 4},
		{"acc +6", Accumulate, 6},
		{"jmp -4", Jump, -4},
	}

	for _, c := range successCases {
		op, err := parseCodeLineToInstruction(c.codeLine)
		assert.NoError(t, err)
		assert.Equal(t, c.command, op.command)
		assert.Equal(t, c.argument, op.argument)
	}

	errorCases := []string{
		"nop +4 -1",
		"nop -2 10",
		"unknown -3",
		"nop foobar",
	}

	for _, c := range errorCases {
		_, err := parseCodeLineToInstruction(c)
		assert.Error(t, err)
	}
}

func TestMachine_Step(t *testing.T) {
	code := []string{
		"nop 0",
		"acc 1",
		"jmp +3",
		"acc 666",
		"acc 666",
		"acc -1",
	}

	cases := []struct {
		pc  int
		acc int
		err error
	}{
		{1, 0, nil},
		{2, 1, nil},
		{5, 1, nil},
		{6, 0, nil},
		{6, 0, EndOfProgram},
	}

	program, err := ParseCodeToProgram(code)
	assert.NoError(t, err)

	machine := NewMachine()
	machine.LoadProgram(program)

	for _, c := range cases {
		err = machine.Step()
		assert.Equal(t, c.pc, machine.ProgramCounter)
		assert.Equal(t, c.acc, machine.Accumulator)
		assert.Equal(t, c.err, err)
	}
}

func TestMachine_RunProgram(t *testing.T) {
	code := []string{
		"nop 0",
		"acc 1",
		"jmp +3",
		"acc 666",
		"acc 666",
		"acc -1",
	}

	program, err := ParseCodeToProgram(code)
	assert.NoError(t, err)

	machine := NewMachine()
	machine.LoadProgram(program)

	err = machine.RunProgram()
	assert.NoError(t, err)
	assert.Equal(t, 0, machine.Accumulator)
}

func TestMachine_RunProgram_AlreadySeen(t *testing.T) {
	code := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	program, err := ParseCodeToProgram(code)
	assert.NoError(t, err)

	machine := NewMachine()
	machine.LoadProgram(program)

	err = machine.RunProgram()
	assert.Error(t, err)
	assert.Equal(t, AlreadySeen, err)
	assert.Equal(t, 5, machine.Accumulator)
}
