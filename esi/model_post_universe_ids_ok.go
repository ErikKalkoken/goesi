/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.7.15
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

/* A list of PostUniverseIdsOk. */
//easyjson:json
type PostUniverseIdsOkList []PostUniverseIdsOk

/* 200 ok object */
//easyjson:json
type PostUniverseIdsOk struct {
	Agents         []PostUniverseIdsAgent         `json:"agents,omitempty"`          /* agents array */
	Alliances      []PostUniverseIdsAlliance      `json:"alliances,omitempty"`       /* alliances array */
	Characters     []PostUniverseIdsCharacter     `json:"characters,omitempty"`      /* characters array */
	Constellations []PostUniverseIdsConstellation `json:"constellations,omitempty"`  /* constellations array */
	Corporations   []PostUniverseIdsCorporation   `json:"corporations,omitempty"`    /* corporations array */
	Factions       []PostUniverseIdsFaction       `json:"factions,omitempty"`        /* factions array */
	InventoryTypes []PostUniverseIdsInventoryType `json:"inventory_types,omitempty"` /* inventory_types array */
	Regions        []PostUniverseIdsRegion        `json:"regions,omitempty"`         /* regions array */
	Stations       []PostUniverseIdsStation       `json:"stations,omitempty"`        /* stations array */
	Systems        []PostUniverseIdsSystem        `json:"systems,omitempty"`         /* systems array */
}
