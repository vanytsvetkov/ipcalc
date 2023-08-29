package functions

import (
	"fmt"
	"github.com/fatih/color"
	"net"
)

func IPcalc(prefix string) (string, error) {
	var result string

	ip, ipNet, err := net.ParseCIDR(prefix)
	if err != nil {
		return result, err
	}

	ipv4 := Byte(ip.To4())
	netmask := Byte(ipNet.Mask)
	maskSize, _ := ipNet.Mask.Size()
	netmaskStr := fmt.Sprintf("%s = %d", netmask.String(), maskSize)
	/*
		In hexadecimal notation, 0xFFFFFFFF is a 32-bit number equal to 4294967295 in decimal.
		It's the largest 32-bit unsigned integer, that is 2**32-1 in decimal.
		The >> operator is used for bitwise right shifting.
		In this case, "maskSize" specifies that the bits will be shifted maskSize-th positions to the right.
	*/
	wildcard := ByteFromUint32(0xFFFFFFFF >> maskSize)

	network := Byte(ipNet.IP)
	networkCIDR := fmt.Sprintf("%v/%d", network.String(), maskSize)

	var hostMin, hostMax, broadcast Byte
	var hosts uint32

	switch maskSize {
	case 31:
		hosts = 2
		hostMin = ipv4
		hostMax = ByteFromUint32(ipv4.ToUint32() + 1)
		broadcast = NAByte()
	case 32:
		hosts = 1
		hostMin = NAByte()
		hostMax = NAByte()
		broadcast = NAByte()
	default:
		//	"|" operation performs a bitwise OR operation on unsigned integer types.
		broadcast = ByteFromUint32(ipv4.ToUint32() | wildcard.ToUint32())
		//	1<<0 is 00000000.00000000.00000000.00000001
		hostMin = ByteFromUint32(network.ToUint32() | 1<<0)
		// "&^" operator performs a bitwise AND NOT on unsigned integer types,
		hostMax = ByteFromUint32(broadcast.ToUint32() &^ 1 << 0) // clear last bit on broadcast

		hosts = (broadcast.ToUint32() - network.ToUint32()) - 1
	}

	result += fmt.Sprintf("Address:    %-29s %s\n", color.BlueString(ipv4.String()), color.YellowString(ipv4.BinaryString()))
	result += fmt.Sprintf("Netmask:    %-29s %s\n", color.BlueString(netmaskStr), color.RedString(netmask.BinaryString()))
	result += fmt.Sprintf("Wildcard:   %-29s %s\n", color.BlueString(wildcard.String()), color.YellowString(wildcard.BinaryString()))
	result += fmt.Sprintln("=>")
	if hostMin != nil && hostMax != nil {
		result += fmt.Sprintf("Network:    %-29s %s\n", color.BlueString(networkCIDR), color.YellowString(network.BinaryString()))
		result += fmt.Sprintf("HostMin:    %-29s %s\n", color.BlueString(hostMin.String()), color.YellowString(hostMin.BinaryString()))
		result += fmt.Sprintf("HostMax:    %-29s %s\n", color.BlueString(hostMax.String()), color.YellowString(hostMax.BinaryString()))
	} else {
		result += fmt.Sprintf("Hostroute:  %-29s %s\n", color.BlueString(networkCIDR), color.YellowString(network.BinaryString()))
	}
	if broadcast != nil {
		result += fmt.Sprintf("Broadcast:  %-29s %s\n", color.BlueString(broadcast.String()), color.YellowString(broadcast.BinaryString()))
	}
	result += fmt.Sprintf("Hosts/Net:  %s\n", color.BlueString(fmt.Sprintf("%d", hosts)))

	return result, nil
}
