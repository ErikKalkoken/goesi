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

func easyjson635b4880DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetWarsWarIdKillmails200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetWarsWarIdKillmails200OkList, 0, 2)
			} else {
				*out = GetWarsWarIdKillmails200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetWarsWarIdKillmails200Ok
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
func easyjson635b4880EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetWarsWarIdKillmails200OkList) {
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
func (v GetWarsWarIdKillmails200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson635b4880EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdKillmails200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson635b4880EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdKillmails200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson635b4880DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdKillmails200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson635b4880DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson635b4880DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetWarsWarIdKillmails200Ok) {
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
		case "killmail_id":
			out.KillmailId = int32(in.Int32())
		case "killmail_hash":
			out.KillmailHash = string(in.String())
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
func easyjson635b4880EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetWarsWarIdKillmails200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.KillmailId != 0 {
		const prefix string = ",\"killmail_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.KillmailId))
	}
	if in.KillmailHash != "" {
		const prefix string = ",\"killmail_hash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.KillmailHash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetWarsWarIdKillmails200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson635b4880EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetWarsWarIdKillmails200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson635b4880EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetWarsWarIdKillmails200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson635b4880DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetWarsWarIdKillmails200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson635b4880DecodeGithubComAntihaxGoesiEsi1(l, v)
}