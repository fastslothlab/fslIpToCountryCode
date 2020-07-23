package fslIpToCountryCode

import (
	"net"
	"strconv"
)

type Reader struct{
	Buf []byte `json:",emitempty"`
	Ipv4StartNode uint32 `json:",emitempty"`
	NodeNum uint32 `json:",emitempty"`
	IsoIdCodeMap map[uint32]string `json:",emitempty"`
	IsoIdCodeList []string `json:",emitempty"`
	Ipv4FirstByteNodeCache []uint32 `json:",emitempty"`
	CheckSum string `json:",emitempty"`
}

func (r *Reader) CheckValid() string{
	if r.NodeNum==0{
		return "r.NodeNum==0"
	}
	if r.Ipv4StartNode >= r.NodeNum{
		return "r.Ipv4StartNode >= r.NodeNum "+strconv.Itoa(int(r.Ipv4StartNode))+" "+strconv.Itoa(int(r.NodeNum))
	}
	if len(r.Buf)!=int(r.NodeNum*6)+1{
		return "len(r.Buf)!=r.NodeNum*6 "+strconv.Itoa(int(len(r.Buf)))+" "+strconv.Itoa(int(r.NodeNum*6))
	}
	if r.CheckSum==""{
		return `r.CheckSum=="" `
	}
	return ""
}

func (r *Reader) CompatibleToOldVersion(){
	if len(r.Buf) == int(r.NodeNum*6){
		r.Buf = append(r.Buf,0)
	}
}

func (r *Reader) getIdByIp(ip net.IP) uint32{
	node:=uint32(0)
	i:=uint8(0)
	nodeCount:= r.NodeNum
	if len(ip)==4{
		//bitCount = 32
		if len(r.Ipv4FirstByteNodeCache)>0{
			node = r.Ipv4FirstByteNodeCache[ip[0]]
			if node==nodeCount{
				return 0
			}
			if node>nodeCount{
				return node
			}
			i = uint8(8)
		}else{
			node = r.Ipv4StartNode
		}
	}
	for {
		bit := IpGetBit(ip,i)
		node = r.ReadNode(node, bit)
		if node==nodeCount{
			return 0
		}
		if node>nodeCount{
			return node
		}
		i++
	}
}

func (r *Reader) MustGetCountryIsoCode(ip net.IP) (code string) {
	if ip==nil{
		return ""
	}
	id := r.getIdByIp(ip)
	if id==0{
		return ""
	}
	if len(r.IsoIdCodeList)>0{
		id = id-r.NodeNum-1
		return r.IsoIdCodeList[id]
	}else{
		return r.IsoIdCodeMap[id]
	}
}
func (r *Reader) MustGetCountryIsoCodeByString(ip string) (code string){
	tmpBuf:=[16]byte{}
	ipObj:=ParseIpToBuf(ip,tmpBuf[:])
	code= r.MustGetCountryIsoCode(ipObj)
	return code
}

func (r *Reader) GetAllCountryCode() (output []string){
	for _,code:=range r.IsoIdCodeMap{
		if code==""{
			continue
		}
		output = append(output,code)
	}
	return output
}