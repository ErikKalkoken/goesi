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

func easyjson63aac689DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsOrbitalList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdStatsOrbitalList, 0, 2)
			} else {
				*out = GetCharactersCharacterIdStatsOrbitalList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdStatsOrbital
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
func easyjson63aac689EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdStatsOrbitalList) {
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
func (v GetCharactersCharacterIdStatsOrbitalList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson63aac689EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsOrbitalList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson63aac689EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsOrbitalList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson63aac689DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsOrbitalList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson63aac689DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson63aac689DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdStatsOrbital) {
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
		case "strike_characters_killed":
			out.StrikeCharactersKilled = int64(in.Int64())
		case "strike_damage_to_players_armor_amount":
			out.StrikeDamageToPlayersArmorAmount = int64(in.Int64())
		case "strike_damage_to_players_shield_amount":
			out.StrikeDamageToPlayersShieldAmount = int64(in.Int64())
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
func easyjson63aac689EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdStatsOrbital) {
	out.RawByte('{')
	first := true
	_ = first
	if in.StrikeCharactersKilled != 0 {
		const prefix string = ",\"strike_characters_killed\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.StrikeCharactersKilled))
	}
	if in.StrikeDamageToPlayersArmorAmount != 0 {
		const prefix string = ",\"strike_damage_to_players_armor_amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.StrikeDamageToPlayersArmorAmount))
	}
	if in.StrikeDamageToPlayersShieldAmount != 0 {
		const prefix string = ",\"strike_damage_to_players_shield_amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.StrikeDamageToPlayersShieldAmount))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdStatsOrbital) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson63aac689EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdStatsOrbital) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson63aac689EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsOrbital) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson63aac689DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdStatsOrbital) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson63aac689DecodeGithubComAntihaxGoesiEsi1(l, v)
}
