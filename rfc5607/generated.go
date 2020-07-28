// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc5607

import (
	"strconv"

	"layeh.com/radius"

	. "layeh.com/radius/rfc2865"
)

const (
	FramedManagement_Type              radius.Type = 133
	ManagementTransportProtection_Type radius.Type = 134
	ManagementPolicyID_Type            radius.Type = 135
	ManagementPrivilegeLevel_Type      radius.Type = 136
)

func init() {
	ServiceType_Strings[ServiceType_Value_FramedManagement] = "Framed-Management"
}

const (
	ServiceType_Value_FramedManagement ServiceType = 18
)

type FramedManagement uint32

const (
	FramedManagement_Value_SNMP     FramedManagement = 1
	FramedManagement_Value_WebBased FramedManagement = 2
	FramedManagement_Value_Netconf  FramedManagement = 3
	FramedManagement_Value_FTP      FramedManagement = 4
	FramedManagement_Value_TFTP     FramedManagement = 5
	FramedManagement_Value_SFTP     FramedManagement = 6
	FramedManagement_Value_RCP      FramedManagement = 7
	FramedManagement_Value_SCP      FramedManagement = 8
)

var FramedManagement_Strings = map[FramedManagement]string{
	FramedManagement_Value_SNMP:     "SNMP",
	FramedManagement_Value_WebBased: "Web-Based",
	FramedManagement_Value_Netconf:  "Netconf",
	FramedManagement_Value_FTP:      "FTP",
	FramedManagement_Value_TFTP:     "TFTP",
	FramedManagement_Value_SFTP:     "SFTP",
	FramedManagement_Value_RCP:      "RCP",
	FramedManagement_Value_SCP:      "SCP",
}

func (a FramedManagement) String() string {
	if str, ok := FramedManagement_Strings[a]; ok {
		return str
	}
	return "FramedManagement(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func FramedManagement_Add(p *radius.Packet, value FramedManagement) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(FramedManagement_Type, a)
	return
}

func FramedManagement_Get(p *radius.Packet) (value FramedManagement) {
	value, _ = FramedManagement_Lookup(p)
	return
}

func FramedManagement_Gets(p *radius.Packet) (values []FramedManagement, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != FramedManagement_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, FramedManagement(i))
	}
	return
}

func FramedManagement_Lookup(p *radius.Packet) (value FramedManagement, err error) {
	a, ok := p.Lookup(FramedManagement_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = FramedManagement(i)
	return
}

func FramedManagement_Set(p *radius.Packet, value FramedManagement) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(FramedManagement_Type, a)
	return
}

func FramedManagement_Del(p *radius.Packet) {
	p.Attributes.Del(FramedManagement_Type)
}

type ManagementTransportProtection uint32

const (
	ManagementTransportProtection_Value_NoProtection                       ManagementTransportProtection = 1
	ManagementTransportProtection_Value_IntegrityProtection                ManagementTransportProtection = 2
	ManagementTransportProtection_Value_IntegrityConfidentialityProtection ManagementTransportProtection = 3
)

var ManagementTransportProtection_Strings = map[ManagementTransportProtection]string{
	ManagementTransportProtection_Value_NoProtection:                       "No-Protection",
	ManagementTransportProtection_Value_IntegrityProtection:                "Integrity-Protection",
	ManagementTransportProtection_Value_IntegrityConfidentialityProtection: "Integrity-Confidentiality-Protection",
}

func (a ManagementTransportProtection) String() string {
	if str, ok := ManagementTransportProtection_Strings[a]; ok {
		return str
	}
	return "ManagementTransportProtection(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ManagementTransportProtection_Add(p *radius.Packet, value ManagementTransportProtection) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(ManagementTransportProtection_Type, a)
	return
}

func ManagementTransportProtection_Get(p *radius.Packet) (value ManagementTransportProtection) {
	value, _ = ManagementTransportProtection_Lookup(p)
	return
}

func ManagementTransportProtection_Gets(p *radius.Packet) (values []ManagementTransportProtection, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != ManagementTransportProtection_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ManagementTransportProtection(i))
	}
	return
}

func ManagementTransportProtection_Lookup(p *radius.Packet) (value ManagementTransportProtection, err error) {
	a, ok := p.Lookup(ManagementTransportProtection_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ManagementTransportProtection(i)
	return
}

func ManagementTransportProtection_Set(p *radius.Packet, value ManagementTransportProtection) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(ManagementTransportProtection_Type, a)
	return
}

func ManagementTransportProtection_Del(p *radius.Packet) {
	p.Attributes.Del(ManagementTransportProtection_Type)
}

func ManagementPolicyID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(ManagementPolicyID_Type, a)
	return
}

func ManagementPolicyID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(ManagementPolicyID_Type, a)
	return
}

func ManagementPolicyID_Get(p *radius.Packet) (value []byte) {
	value, _ = ManagementPolicyID_Lookup(p)
	return
}

func ManagementPolicyID_GetString(p *radius.Packet) (value string) {
	value, _ = ManagementPolicyID_LookupString(p)
	return
}

func ManagementPolicyID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, avp := range p.Attributes {
		if avp.Type != ManagementPolicyID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManagementPolicyID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, avp := range p.Attributes {
		if avp.Type != ManagementPolicyID_Type {
			continue
		}
		attr := avp.Attribute
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ManagementPolicyID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(ManagementPolicyID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ManagementPolicyID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(ManagementPolicyID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ManagementPolicyID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(ManagementPolicyID_Type, a)
	return
}

func ManagementPolicyID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(ManagementPolicyID_Type, a)
	return
}

func ManagementPolicyID_Del(p *radius.Packet) {
	p.Attributes.Del(ManagementPolicyID_Type)
}

type ManagementPrivilegeLevel uint32

var ManagementPrivilegeLevel_Strings = map[ManagementPrivilegeLevel]string{}

func (a ManagementPrivilegeLevel) String() string {
	if str, ok := ManagementPrivilegeLevel_Strings[a]; ok {
		return str
	}
	return "ManagementPrivilegeLevel(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ManagementPrivilegeLevel_Add(p *radius.Packet, value ManagementPrivilegeLevel) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(ManagementPrivilegeLevel_Type, a)
	return
}

func ManagementPrivilegeLevel_Get(p *radius.Packet) (value ManagementPrivilegeLevel) {
	value, _ = ManagementPrivilegeLevel_Lookup(p)
	return
}

func ManagementPrivilegeLevel_Gets(p *radius.Packet) (values []ManagementPrivilegeLevel, err error) {
	var i uint32
	for _, avp := range p.Attributes {
		if avp.Type != ManagementPrivilegeLevel_Type {
			continue
		}
		attr := avp.Attribute
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ManagementPrivilegeLevel(i))
	}
	return
}

func ManagementPrivilegeLevel_Lookup(p *radius.Packet) (value ManagementPrivilegeLevel, err error) {
	a, ok := p.Lookup(ManagementPrivilegeLevel_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ManagementPrivilegeLevel(i)
	return
}

func ManagementPrivilegeLevel_Set(p *radius.Packet, value ManagementPrivilegeLevel) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(ManagementPrivilegeLevel_Type, a)
	return
}

func ManagementPrivilegeLevel_Del(p *radius.Packet) {
	p.Attributes.Del(ManagementPrivilegeLevel_Type)
}