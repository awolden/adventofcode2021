package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/awolden/adventofcode2021/helpers"
)

type Packet struct {
	version      int
	packetType   int
	lengthTypeId int
	value        int
	subPackets   []Packet
}

// lazy parse
var BITS = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func main() {
	input := readInput()
	packets := parseToPackets(&input)
	fmt.Println("===================")
	fmt.Println("version count", countVersion(packets))
	fmt.Println("process count", processPacket(packets[0]))
}

func countVersion(packets []Packet) int {
	sum := 0
	for _, packet := range packets {
		sum += packet.version
		if len(packet.subPackets) > 0 {
			sum += countVersion(packet.subPackets)
		}
	}
	return sum
}

func processPacket(packet Packet) int {
	vals := []int{}
	for _, packet := range packet.subPackets {
		vals = append(vals, processPacket(packet))
	}

	switch packet.packetType {
	case 0:
		return helpers.FindArraySum(vals)
	case 1:
		return helpers.FindArrayProduct(vals)
	case 2:
		return helpers.Min(vals)
	case 3:
		return helpers.Max(vals)
	case 4:
		return packet.value
	case 5:
		if vals[0] > vals[1] {
			return 1
		}
		return 0
	case 6:
		if vals[0] < vals[1] {
			return 1
		}
		return 0
	case 7:
		if vals[0] == vals[1] {
			return 1
		}
		return 0
	case 8:
	case 9:
	}
	return 0
}

func parseToPackets(input *[]int) []Packet {
	packets := []Packet{}

	if isPadding(*input) {
		*input = []int{}
		return []Packet{}
	}

	version, _ := strconv.ParseInt(join((*input)[:3]), 2, 64)
	*input = (*input)[3:] //cursor += 3
	packetType, _ := strconv.ParseInt(join((*input)[:3]), 2, 64)
	*input = (*input)[3:] //cursor += 3
	value := 0

	// literal
	if packetType == 4 {
		// build groups for literal
		done := false
		fullLiteral := []int{}
		for !done {
			group := (*input)[:5]
			*input = (*input)[5:] //cursor += 5
			if group[0] == 0 {
				done = true
			}
			fullLiteral = append(fullLiteral, group[1:]...)
			parsedVal, _ := strconv.ParseInt(join(fullLiteral), 2, 64)
			value = int(parsedVal)
		}
		return []Packet{{
			packetType: int(packetType),
			version:    int(version),
			value:      value,
		}}
	} else {
		//operator
		lengthTypeId, _ := strconv.ParseInt(join((*input)[:1]), 2, 64)
		*input = (*input)[1:] //cursor += 1

		switch lengthTypeId {
		case 1:
			numOfPackets, _ := strconv.ParseInt(join((*input)[:11]), 2, 64)
			*input = (*input)[11:] //cursor += 11
			for i := 0; i < int(numOfPackets); i++ {
				packets = append(packets, parseToPackets(input)...)
			}
		case 0:
			packetLength, _ := strconv.ParseInt(join((*input)[:15]), 2, 64)
			*input = (*input)[15:] // cursor += 15
			subBits := (*input)[:packetLength]
			(*input) = (*input)[packetLength:]
			for len(subBits) > 0 {
				packets = append(packets, parseToPackets(&subBits)...)
			}
		}
		packet := Packet{
			packetType:   int(packetType),
			version:      int(version),
			lengthTypeId: int(lengthTypeId),
			value:        value,
			subPackets:   packets,
		}
		return []Packet{packet}
	}
}

func readInput() []int {
	rawInput := helpers.GetFileArray("./input")
	str := ""

	for _, c := range rawInput[0] {
		str += BITS[string(c)]
	}

	bits := asBits(str)
	for len(bits)%4 != 0 {
		bits = append([]int{0}, bits...)
	}
	return bits
}

func isPadding(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}
func join(nums []int) string {
	var str string
	for i := range nums {
		str += strconv.Itoa(int(nums[i]))
	}
	return str
}
func asBits(val string) []int {
	bits := []int{}
	cArr := strings.Split(val, "")
	for _, c := range cArr {
		i, _ := strconv.Atoi(c)
		bits = append(bits, i)
	}
	return bits
}
