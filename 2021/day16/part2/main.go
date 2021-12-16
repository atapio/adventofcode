package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

type Bits struct {
	data     []byte
	pos      int
	BitsRead int
}

func (b *Bits) Init(input string) error {
	d, err := hex.DecodeString(input)
	if err != nil {
		return err
	}
	b.data = d
	b.pos = 0
	b.BitsRead = 0

	return nil
}

func (b *Bits) readBit() byte {
	v := b.data[0] & (1 << (7 - b.pos))
	b.pos++
	b.BitsRead++
	if b.pos > 7 {
		b.data = b.data[1:]
		b.pos = 0
	}

	bit := byte(0)
	if v > 0 {
		bit = 1
	}

	//fmt.Printf("readbit: %d\n", bit)

	return bit
}

func (b *Bits) ReadBits(count int) int {
	v := int(b.readBit())
	for i := 1; i < count; i++ {
		v = v << 1
		v += int(b.readBit())
	}

	return v
}

type Packet struct {
	Version int
	TypeID  int
	Literal int

	SubPackets []*Packet
}

func (p *Packet) Parse(b *Bits) error {
	p.SubPackets = []*Packet{}
	p.Version = b.ReadBits(3)
	p.TypeID = b.ReadBits(3)

	fmt.Printf("start packet v: %d t: %d\n", p.Version, p.TypeID)
	switch p.TypeID {
	case 4:
		fmt.Println("literal:")
		p.ReadLiteral(b)
	default:
		fmt.Println("operator packet:")
		err := p.ParseOperatorPacket(b)
		if err != nil {
			return err
		}
		fmt.Printf("typeID: %d\n", p.TypeID)
		switch p.TypeID {
		case 0:
			for _, subPacket := range p.SubPackets {
				p.Literal += subPacket.Literal
			}
		case 1:
			p.Literal = 1
			for _, subPacket := range p.SubPackets {
				p.Literal *= subPacket.Literal
			}
		case 2:
			first, rest := p.SubPackets[0], p.SubPackets[1:]
			p.Literal = first.Literal

			for _, subPacket := range rest {
				if subPacket.Literal < p.Literal {
					p.Literal = subPacket.Literal
				}
			}
		case 3:
			first, rest := p.SubPackets[0], p.SubPackets[1:]
			p.Literal = first.Literal

			for _, subPacket := range rest {
				if subPacket.Literal > p.Literal {
					p.Literal = subPacket.Literal
				}
			}
		case 5:
			p.Literal = 0
			if p.SubPackets[0].Literal > p.SubPackets[1].Literal {
				p.Literal = 1
			}
		case 6:
			p.Literal = 0
			if p.SubPackets[0].Literal < p.SubPackets[1].Literal {
				p.Literal = 1
			}
		case 7:
			p.Literal = 0
			if p.SubPackets[0].Literal == p.SubPackets[1].Literal {
				p.Literal = 1
			}
		default:
			fmt.Printf("TypeID %d not implemented\n", p.TypeID)
		}
	}

	fmt.Printf("end packet v: %d t: %d l: %d s:%d\n", p.Version, p.TypeID, p.Literal, len(p.SubPackets))
	return nil
}

func (p *Packet) ReadLiteral(b *Bits) {
	for b.ReadBits(1) > 0 {
		p.Literal = p.Literal << 4
		p.Literal += b.ReadBits(4)
	}
	p.Literal = p.Literal << 4
	p.Literal += b.ReadBits(4)
}

func (p *Packet) ParseOperatorPacket(b *Bits) error {
	switch b.ReadBits(1) {
	case 0:
		length := b.ReadBits(15)
		endPos := b.BitsRead + length
		fmt.Printf("length type 0 %d bits\n", length)
		for b.BitsRead < endPos {
			fmt.Printf("pos %d endPos %d\n", b.pos, endPos)
			packet := &Packet{}
			err := packet.Parse(b)
			if err != nil {
				return err
			}
			p.SubPackets = append(p.SubPackets, packet)
		}
	case 1:
		packets := b.ReadBits(11)
		fmt.Printf("length type 0 %d packets\n", packets)
		for i := 0; i < packets; i++ {
			packet := &Packet{}
			err := packet.Parse(b)
			if err != nil {
				return err
			}
			p.SubPackets = append(p.SubPackets, packet)
		}
	}
	return nil
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	fmt.Println("start")
	bits := &Bits{}
	err := bits.Init(lines[0])
	if err != nil {
		return 0, err
	}

	packet := Packet{}
	packet.Parse(bits)
	answer = packet.Literal

	return answer, nil
}

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return input, nil
}

func main() {
	input, err := parseFile("input")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
