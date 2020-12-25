package main

import "fmt"

type Ticket []int

type FieldSpecification struct {
	Name   string
	ranges [2][2]int
}

func (fs FieldSpecification) IsValid(x int) bool {
	first := fs.ranges[0][0] <= x && x <= fs.ranges[0][1]
	second := fs.ranges[1][0] <= x && x <= fs.ranges[1][1]
	return first || second
}

func RemoveInvalidTickets(tickets []Ticket, specifications []FieldSpecification) []Ticket {
	validTickets := make([]Ticket, 0, len(tickets))
	for _, ticket := range tickets {
		if CheckTicket(ticket, specifications) {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func CheckTicket(ticket Ticket, specifications []FieldSpecification) bool {
	for _, column := range ticket {
		// if any spec matches, this column is not totally invalid
		anySpecMatches := false
		for _, s := range specifications {
			if s.IsValid(column) {
				anySpecMatches = true
				break
			}
		}

		// a column without any matching spec means this ticket is totally invalid
		// we can abort the check and remember the value that made it invalid
		if !anySpecMatches {
			return false
		}
	}
	return true
}

type Matrix struct {
	Specifications []FieldSpecification
	Columns        []int

	data [][]int
	N    int
	M    int
}

func InitMatrix(specifications []FieldSpecification, tickets []Ticket) Matrix {
	// derive dimensions
	N := len(specifications)
	M := len(tickets[0])

	// build columns vector
	columns := make([]int, M)
	for i := 0; i < M; i++ {
		columns[i] = i
	}

	// pre-init matrix
	m := Matrix{
		Specifications: specifications,
		Columns:        columns,
		N:              N,
		M:              M,
		data:           make([][]int, N),
	}

	// calculate matrix
	for i := 0; i < m.N; i++ {
		m.data[i] = make([]int, m.M)
		for j := 0; j < m.M; j++ {
			if columnFulfillsSpec(j, specifications[i], tickets) {
				m.data[i][j] = 1
			}
		}
	}

	return m
}

func (m Matrix) Row(i int) []int {
	return m.data[i]
}

func (m Matrix) Col(j int) []int {
	result := make([]int, m.N)
	for i := 0; i < m.N; i++ {
		result[i] = m.data[i][j]
	}
	return result
}

func (m Matrix) RowSum(i int) (int, int) {
	sum := 0
	last := -1

	for k, x := range m.Row(i) {
		if x > 0 {
			sum += x
			last = k
		}
	}
	return sum, last
}

func (m Matrix) ColSum(j int) (int, int) {
	sum := 0
	last := -1

	cols := m.Col(j)
	for k, x := range cols {
		if x > 0 {
			sum += x
			last = k
		}
	}
	return sum, last
}

func (m *Matrix) RemoveRow(i int) {
	m.N--
	m.Specifications = append(m.Specifications[:i], m.Specifications[i+1:]...)
	m.data = append(m.data[:i], m.data[i+1:]...)
}

func (m *Matrix) RemoveCol(j int) {
	m.M--
	m.Columns = append(m.Columns[:j], m.Columns[j+1:]...)
	for i, row := range m.data {
		m.data[i] = append(row[:j], row[j+1:]...)
	}
}

func columnFulfillsSpec(column int, specification FieldSpecification, tickets []Ticket) bool {
	for _, ticket := range tickets {
		if !specification.IsValid(ticket[column]) {
			return false
		}
	}
	return true
}

func (m Matrix) PrintString() {
	fmt.Printf("%30s ", " ")
	for _, c := range m.Columns {
		fmt.Printf("%3d ", c)
	}
	fmt.Printf("%3s\n", "Σ")

	for i := 0; i < m.N; i++ {
		fmt.Printf("%30s ", m.Specifications[i].Name)
		for j := 0; j < m.M; j++ {
			fmt.Printf("%3d ", m.data[i][j])
		}

		sum, _ := m.RowSum(i)
		fmt.Printf("%3d\n", sum)
	}

	fmt.Printf("%30s ", "Σ")
	for i, _ := range m.Columns {
		sum, _ := m.ColSum(i)
		fmt.Printf("%3d ", sum)
	}
	fmt.Println()
}
