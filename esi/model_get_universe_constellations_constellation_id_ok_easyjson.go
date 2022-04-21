// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7d9c79f0DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetUniverseConstellationsConstellationIdOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseConstellationsConstellationIdOkList, 0, 0)
			} else {
				*out = GetUniverseConstellationsConstellationIdOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseConstellationsConstellationIdOk
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7d9c79f0EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetUniverseConstellationsConstellationIdOkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseConstellationsConstellationIdOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d9c79f0EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseConstellationsConstellationIdOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d9c79f0EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d9c79f0DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d9c79f0DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson7d9c79f0DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetUniverseConstellationsConstellationIdOk) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "constellation_id":
			out.ConstellationId = int32(in.Int32())
		case "name":
			out.Name = string(in.String())
		case "position":
			(out.Position).UnmarshalEasyJSON(in)
		case "region_id":
			out.RegionId = int32(in.Int32())
		case "systems":
			if in.IsNull() {
				in.Skip()
				out.Systems = nil
			} else {
				in.Delim('[')
				if out.Systems == nil {
					if !in.IsDelim(']') {
						out.Systems = make([]int32, 0, 16)
					} else {
						out.Systems = []int32{}
					}
				} else {
					out.Systems = (out.Systems)[:0]
				}
				for !in.IsDelim(']') {
					var v4 int32
					v4 = int32(in.Int32())
					out.Systems = append(out.Systems, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7d9c79f0EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetUniverseConstellationsConstellationIdOk) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ConstellationId != 0 {
		const prefix string = ",\"constellation_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.ConstellationId))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if true {
		const prefix string = ",\"position\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Position).MarshalEasyJSON(out)
	}
	if in.RegionId != 0 {
		const prefix string = ",\"region_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.RegionId))
	}
	if len(in.Systems) != 0 {
		const prefix string = ",\"systems\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Systems {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int32(int32(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseConstellationsConstellationIdOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d9c79f0EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseConstellationsConstellationIdOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d9c79f0EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d9c79f0DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseConstellationsConstellationIdOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d9c79f0DecodeGithubComAntihaxGoesiEsi1(l, v)
}
