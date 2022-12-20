/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.15
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetKillmailsKillmailIdKillmailHashVictim. */
//easyjson:json
type GetKillmailsKillmailIdKillmailHashVictimList []GetKillmailsKillmailIdKillmailHashVictim

/* victim object */
//easyjson:json
type GetKillmailsKillmailIdKillmailHashVictim struct {
	AllianceId    int32                                      `json:"alliance_id,omitempty"`    /* alliance_id integer */
	CharacterId   int32                                      `json:"character_id,omitempty"`   /* character_id integer */
	CorporationId int32                                      `json:"corporation_id,omitempty"` /* corporation_id integer */
	DamageTaken   int32                                      `json:"damage_taken,omitempty"`   /* How much total damage was taken by the victim  */
	FactionId     int32                                      `json:"faction_id,omitempty"`     /* faction_id integer */
	Items         []GetKillmailsKillmailIdKillmailHashItem   `json:"items,omitempty"`          /* items array */
	Position      GetKillmailsKillmailIdKillmailHashPosition `json:"position,omitempty"`
	ShipTypeId    int32                                      `json:"ship_type_id,omitempty"` /* The ship that the victim was piloting and was destroyed  */
}
