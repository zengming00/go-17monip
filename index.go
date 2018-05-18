package go17monip

import (
	"errors"
	"net"
)

var mCCMAP = map[byte]string{
	105: "HN", 36: "RO", 12: "BR", 4: "TW", 94: "JO", 120: "KG", 226: "IO", 67: "SR", 197: "GQ", 219: "PW", 191: "GL", 135: "NW",
	33: "UA", 69: "PA", 202: "SS", 37: "IR", 137: "NC", 121: "JM", 189: "TG", 166: "SL", 169: "AG", 81: "BO", 158: "BJ", 186: "LR",
	78: "GE", 20: "SE", 15: "NL", 184: "SX", 165: "DJ", 162: "YK", 156: "KY", 47: "MA", 70: "KW", 119: "AN", 175: "NE", 212: "TD",
	179: "GF", 98: "AM", 148: "GP", 76: "MD", 49: "GR", 230: "YT", 139: "CD", 214: "NR", 77: "SD", 142: "BZ", 128: "BN", 130: "MG",
	31: "BE", 170: "SZ", 91: "QA", 45: "MY", 223: "AQ", 196: "KN", 133: "BB", 93: "AZ", 50: "HU", 138: "ZW", 29: "VN", 75: "EE",
	51: "UY", 80: "SC", 140: "FJ", 21: "ZA", 6: "JP", 27: "TR", 3: "HK", 161: "BF", 233: "NU", 109: "NA",
	192: "BI", 1: "US", 107: "NP", 83: "BD", 23: "PL", 194: "CV", 28: "NO", 231: "TK", 204: "VC", 188: "VG", 146: "ET", 229: "MS",
	108: "BH", 58: "TN", 26: "CO", 48: "PT", 122: "UG", 150: "LI", 157: "LA", 174: "FO", 56: "AE", 65: "LT", 7: "GB", 236: "CX",
	185: "GN", 208: "ST", 39: "TH", 206: "TC", 217: "FM", 154: "YE", 32: "DK", 134: "HT", 127: "MN", 13: "IT", 163: "ML", 118: "SN",
	85: "SY", 228: "MH", 132: "ME", 18: "ES", 171: "GY", 123: "CU", 152: "GI", 5: "MO", 89: "IS", 106: "LB", 221: "CF", 54: "BG", 74: "LU",
	38: "CL", 16: "RU", 14: "AU", 201: "AS", 10: "FR", 24: "AR", 149: "JE", 0: "--", 207: "GW", 199: "DM", 42: "SG", 210: "TL", 2: "CN",
	129: "CI", 62: "EC", 200: "GD", 86: "OM", 178: "SM", 22: "CH", 176: "MM", 35: "AT", 64: "CR", 57: "DZ",
	114: "LY", 116: "ZM", 203: "VA", 90: "CY", 60: "KZ", 100: "MT", 218: "MF", 82: "PR", 209: "TM", 43: "NZ", 101: "MU", 53: "PK",
	19: "MX", 59: "PE", 102: "LK", 131: "RE", 172: "GG", 104: "TT", 84: "PY", 9: "KR", 144: "VI", 25: "ID", 99: "PS", 61: "SK", 8: "DE",
	159: "PF", 30: "FI", 181: "BT", 34: "EG", 97: "IQ", 124: "CM", 155: "TJ", 52: "PH", 190: "SO", 46: "VE", 73: "DO", 113: "AL", 87: "MK",
	183: "GM", 151: "MQ", 11: "CA", 143: "BW", 136: "AF", 234: "KP", 153: "LS", 198: "MP", 168: "AW", 141: "BS", 103: "GH", 68: "NG", 160: "PG",
	164: "CG", 211: "TO", 193: "VU", 195: "WS", 177: "AD", 205: "SB", 220: "KM", 187: "BQ", 63: "SI", 17: "IN", 112: "RW", 167: "MC", 72: "LV",
	173: "MR", 213: "AI", 71: "BY", 66: "HR", 216: "CK", 55: "KE", 79: "AO", 145: "BM",
	117: "KH", 224: "PM", 147: "IM", 215: "TV", 40: "CZ", 44: "SA", 222: "ER", 115: "GA", 95: "SV", 182: "LC", 88: "TZ", 180: "AX",
	235: "CC", 92: "BA", 111: "NI", 126: "GU", 41: "IE", 225: "KI", 110: "MZ", 96: "GT", 227: "WF", 237: "FK", 232: "NF", 125: "UZ",
}

const (
	IPv4a = 12
	IPv4b = 13
	IPv4c = 14
	IPv4d = 15
)

// func panicif(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

var ipBuffer []byte
var indexBuffer []byte
var ipBufferLength uint32
var indexBufferLength int

func init() {
	// 由json生成ccmap
	// bts, err := ioutil.ReadFile("./cc.json")
	// panicif(err)

	// var cc map[string]int
	// err = json.Unmarshal(bts, &cc)
	// panicif(err)

	// var ccmap = make(map[int]string)
	// for k, v := range cc {
	// 	// fmt.Println(k, v)
	// 	ccmap[v] = k
	// }
	// // fmt.Println(ccmap)
	// fmt.Printf("%#v", ccmap)

	//////////////////////////////
	// ipBuffer, err := ioutil.ReadFile("./ip.dat")
	// panicif(err)
	// indexBuffer, err := ioutil.ReadFile("./index.dat")
	// panicif(err)

	ipBuffer = _escFSMustByte(false, "/ip.dat")
	indexBuffer = _escFSMustByte(false, "/index.dat")

	ipBufferLength = uint32(len(ipBuffer))
	indexBufferLength = len(indexBuffer)
}

const mIP_FIND_DEFAULT = "--"

func IpFind(ip net.IP) string {
	var ipInt = Ip2long(ip)
	var preIP = ((uint32(ip[IPv4a]) << 8) + uint32(ip[IPv4b])) * 4
	start, err := ReadUint32BE(indexBuffer, preIP)
	if err != nil {
		return mIP_FIND_DEFAULT
	}
	for i := start * 5; i < ipBufferLength; i = i + 5 {
		v, err := ReadUint32BE(ipBuffer, i)
		if err != nil {
			return mIP_FIND_DEFAULT
		}
		if v >= ipInt && (i+4) < ipBufferLength {
			if val, ok := mCCMAP[ipBuffer[i+4]]; ok {
				return val
			}
		}
	}
	return mIP_FIND_DEFAULT
}

func ReadUint32BE(data []byte, idx uint32) (uint32, error) {
	dataLen := len(data)
	if dataLen < 4 || (uint32(dataLen)-idx) < 4 {
		return 0, errors.New("data.len < 4")
	}
	return uint32(data[idx])<<24 | uint32(data[idx+1])<<16 | uint32(data[idx+2])<<8 | uint32(data[idx+3]), nil
}

func ToUint32BE(data []byte) (uint32, error) {
	if len(data) < 4 {
		return 0, errors.New("data.len < 4")
	}
	return uint32(data[0])<<24 | uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3]), nil
}

func Ip2long(ip net.IP) uint32 {
	v, _ := ReadUint32BE(ip, IPv4a)
	return v
}

func Long2ip(ip uint32) net.IP {
	a := byte((ip >> 24) & 0xFF)
	b := byte((ip >> 16) & 0xFF)
	c := byte((ip >> 8) & 0xFF)
	d := byte(ip & 0xFF)
	return net.IPv4(a, b, c, d)
}
