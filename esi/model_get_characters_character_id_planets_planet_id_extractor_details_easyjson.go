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

func easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails
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
func easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) {
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
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetailsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) {
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
		case "cycle_time":
			out.CycleTime = int32(in.Int32())
		case "head_radius":
			out.HeadRadius = float32(in.Float32())
		case "heads":
			if in.IsNull() {
				in.Skip()
				out.Heads = nil
			} else {
				in.Delim('[')
				if out.Heads == nil {
					if !in.IsDelim(']') {
						out.Heads = make([]GetCharactersCharacterIdPlanetsPlanetIdHead, 0, 5)
					} else {
						out.Heads = []GetCharactersCharacterIdPlanetsPlanetIdHead{}
					}
				} else {
					out.Heads = (out.Heads)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdPlanetsPlanetIdHead
					easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Heads = append(out.Heads, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "product_type_id":
			out.ProductTypeId = int32(in.Int32())
		case "qty_per_cycle":
			out.QtyPerCycle = int32(in.Int32())
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
func easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) {
	out.RawByte('{')
	first := true
	_ = first
	if in.CycleTime != 0 {
		const prefix string = ",\"cycle_time\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.CycleTime))
	}
	if in.HeadRadius != 0 {
		const prefix string = ",\"head_radius\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.HeadRadius))
	}
	if len(in.Heads) != 0 {
		const prefix string = ",\"heads\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Heads {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.ProductTypeId != 0 {
		const prefix string = ",\"product_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ProductTypeId))
	}
	if in.QtyPerCycle != 0 {
		const prefix string = ",\"qty_per_cycle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.QtyPerCycle))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson1bb873dcDecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdPlanetsPlanetIdHead) {
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
		case "head_id":
			out.HeadId = int32(in.Int32())
		case "latitude":
			out.Latitude = float32(in.Float32())
		case "longitude":
			out.Longitude = float32(in.Float32())
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
func easyjson1bb873dcEncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdPlanetsPlanetIdHead) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HeadId != 0 {
		const prefix string = ",\"head_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int32(int32(in.HeadId))
	}
	if in.Latitude != 0 {
		const prefix string = ",\"latitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Latitude))
	}
	if in.Longitude != 0 {
		const prefix string = ",\"longitude\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Float32(float32(in.Longitude))
	}
	out.RawByte('}')
}
