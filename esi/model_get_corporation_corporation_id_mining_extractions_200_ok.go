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

import (
	"time"
)

/* A list of GetCorporationCorporationIdMiningExtractions200Ok. */
//easyjson:json
type GetCorporationCorporationIdMiningExtractions200OkList []GetCorporationCorporationIdMiningExtractions200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationCorporationIdMiningExtractions200Ok struct {
	ChunkArrivalTime    time.Time `json:"chunk_arrival_time,omitempty"`    /* The time at which the chunk being extracted will arrive and can be fractured by the moon mining drill.  */
	ExtractionStartTime time.Time `json:"extraction_start_time,omitempty"` /* The time at which the current extraction was initiated.  */
	MoonId              int32     `json:"moon_id,omitempty"`               /* moon_id integer */
	NaturalDecayTime    time.Time `json:"natural_decay_time,omitempty"`    /* The time at which the chunk being extracted will naturally fracture if it is not first fractured by the moon mining drill.  */
	StructureId         int64     `json:"structure_id,omitempty"`          /* structure_id integer */
}
