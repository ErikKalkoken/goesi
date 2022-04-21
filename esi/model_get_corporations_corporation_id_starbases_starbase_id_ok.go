/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.11
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

/* A list of GetCorporationsCorporationIdStarbasesStarbaseIdOk. */
//easyjson:json
type GetCorporationsCorporationIdStarbasesStarbaseIdOkList []GetCorporationsCorporationIdStarbasesStarbaseIdOk

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdStarbasesStarbaseIdOk struct {
	AllowAllianceMembers                bool                                                  `json:"allow_alliance_members,omitempty"`                   /* allow_alliance_members boolean */
	AllowCorporationMembers             bool                                                  `json:"allow_corporation_members,omitempty"`                /* allow_corporation_members boolean */
	Anchor                              string                                                `json:"anchor,omitempty"`                                   /* Who can anchor starbase (POS) and its structures */
	AttackIfAtWar                       bool                                                  `json:"attack_if_at_war,omitempty"`                         /* attack_if_at_war boolean */
	AttackIfOtherSecurityStatusDropping bool                                                  `json:"attack_if_other_security_status_dropping,omitempty"` /* attack_if_other_security_status_dropping boolean */
	AttackSecurityStatusThreshold       float32                                               `json:"attack_security_status_threshold,omitempty"`         /* Starbase (POS) will attack if target's security standing is lower than this value */
	AttackStandingThreshold             float32                                               `json:"attack_standing_threshold,omitempty"`                /* Starbase (POS) will attack if target's standing is lower than this value */
	FuelBayTake                         string                                                `json:"fuel_bay_take,omitempty"`                            /* Who can take fuel blocks out of the starbase (POS)'s fuel bay */
	FuelBayView                         string                                                `json:"fuel_bay_view,omitempty"`                            /* Who can view the starbase (POS)'s fule bay. Characters either need to have required role or belong to the starbase (POS) owner's corporation or alliance, as described by the enum, all other access settings follows the same scheme */
	Fuels                               []GetCorporationsCorporationIdStarbasesStarbaseIdFuel `json:"fuels,omitempty"`                                    /* Fuel blocks and other things that will be consumed when operating a starbase (POS) */
	Offline                             string                                                `json:"offline,omitempty"`                                  /* Who can offline starbase (POS) and its structures */
	Online                              string                                                `json:"online,omitempty"`                                   /* Who can online starbase (POS) and its structures */
	Unanchor                            string                                                `json:"unanchor,omitempty"`                                 /* Who can unanchor starbase (POS) and its structures */
	UseAllianceStandings                bool                                                  `json:"use_alliance_standings,omitempty"`                   /* True if the starbase (POS) is using alliance standings, otherwise using corporation's */
}
