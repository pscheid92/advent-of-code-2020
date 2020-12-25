package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type InstructionCode string

const (
	NoOp       InstructionCode = "nop"
	Accumulate                 = "acc"
	Jump                       = "jmp"
)

type Instruction struct {
	command  InstructionCode
	argument int
}

type Machine struct {
	Accumulator    int
	ProgramCounter int
	Program        []Instruction
	Seen           []bool
}

var (
	EndOfProgram = errors.New("EndOfProgram")
	AlreadySeen  = errors.New("Instruction already seen")
)

func NewMachine() *Machine {
	return &Machine{
		Accumulator:    0,
		ProgramCounter: 0,
		Program:        []Instruction{},
		Seen:           []bool{},
	}
}

func (m *Machine) LoadProgram(program []Instruction) {
	m.Accumulator = 0
	m.ProgramCounter = 0
	m.Program = program
	m.Seen = make([]bool, len(program))
}

func (m *Machine) RunProgram() error {
	var err error

	for err = nil; err != EndOfProgram; err = m.Step() {
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Machine) Step() error {
	// PC behind program means we are done
	if m.ProgramCounter >= len(m.Program) {
		return EndOfProgram
	}

	// Stop if this operation was already seen
	if m.Seen[m.ProgramCounter] {
		return AlreadySeen
	} else {
		m.Seen[m.ProgramCounter] = true
	}

	instruction := m.Program[m.ProgramCounter]
	switch instruction.command {
	case NoOp:
		// intentionally do nothing
	case Accumulate:
		m.Accumulator += instruction.argument
	case Jump:
		// we need to jump one operation less to compensate for the PC increment later
		m.ProgramCounter += instruction.argument - 1
	default:
		// handle unknown operations as nop
		// this should be impossible due to verification during program load
	}

	m.ProgramCounter++
	return nil
}

func ParseCodeToProgram(code []string) ([]Instruction, error) {
	program := make([]Instruction, len(code))
	for i, c := range code {
		instruction, err := parseCodeLineToInstruction(c)
		if err != nil {
			return nil, fmt.Errorf("error reading program: %w", err)
		}
		program[i] = instruction
	}
	return program, nil
}

func parseCodeLineToInstruction(codeLine string) (Instruction, error) {
	parts := strings.Split(codeLine, " ")
	if len(parts) != 2 {
		return Instruction{}, fmt.Errorf("unexpected number of command parts (awaited 2, found %d)", len(parts))
	}

	var command InstructionCode
	switch parts[0] {
	case "nop":
		command = NoOp
	case "acc":
		command = Accumulate
	case "jmp":
		command = Jump
	default:
		return Instruction{}, fmt.Errorf("encountered unknown instruction code '%s'", parts[0])
	}

	argument, err := strconv.Atoi(parts[1])
	if err != nil {
		return Instruction{}, fmt.Errorf("error parsing instruction argument '%s': %w", parts[1], err)
	}

	return Instruction{command: command, argument: argument}, nil
}
