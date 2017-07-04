// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package goesiv3

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

func easyjsonF9365fb2DecodeGithubComAntihaxGoesiV3(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaEffectList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetUniverseTypesTypeIdDogmaEffectList, 0, 8)
			} else {
				*out = GetUniverseTypesTypeIdDogmaEffectList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetUniverseTypesTypeIdDogmaEffect
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
func easyjsonF9365fb2EncodeGithubComAntihaxGoesiV3(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaEffectList) {
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
func (v GetUniverseTypesTypeIdDogmaEffectList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF9365fb2EncodeGithubComAntihaxGoesiV3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaEffectList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF9365fb2EncodeGithubComAntihaxGoesiV3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffectList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF9365fb2DecodeGithubComAntihaxGoesiV3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffectList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF9365fb2DecodeGithubComAntihaxGoesiV3(l, v)
}
func easyjsonF9365fb2DecodeGithubComAntihaxGoesiV31(in *jlexer.Lexer, out *GetUniverseTypesTypeIdDogmaEffect) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "effect_id":
			out.EffectId = int32(in.Int32())
		case "is_default":
			out.IsDefault = bool(in.Bool())
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
func easyjsonF9365fb2EncodeGithubComAntihaxGoesiV31(out *jwriter.Writer, in GetUniverseTypesTypeIdDogmaEffect) {
	out.RawByte('{')
	first := true
	_ = first
	if in.EffectId != 0 {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"effect_id\":")
		out.Int32(int32(in.EffectId))
	}
	if in.IsDefault {
		if !first {
			out.RawByte(',')
		}
		first = false
		out.RawString("\"is_default\":")
		out.Bool(bool(in.IsDefault))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaEffect) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF9365fb2EncodeGithubComAntihaxGoesiV31(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetUniverseTypesTypeIdDogmaEffect) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF9365fb2EncodeGithubComAntihaxGoesiV31(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffect) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF9365fb2DecodeGithubComAntihaxGoesiV31(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetUniverseTypesTypeIdDogmaEffect) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF9365fb2DecodeGithubComAntihaxGoesiV31(l, v)
}