/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.0
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

/* A list of GetWarsWarIdOk. */
//easyjson:json
type GetWarsWarIdOkList []GetWarsWarIdOk

/* 200 ok object */
//easyjson:json
type GetWarsWarIdOk struct {
	Id            int32                 `json:"id,omitempty"`              /* ID of the specified war */
	Declared      time.Time             `json:"declared,omitempty"`        /* Time that the war was declared */
	Started       time.Time             `json:"started,omitempty"`         /* Time when the war started and both sides could shoot each other */
	Retracted     time.Time             `json:"retracted,omitempty"`       /* Time the war was retracted but both sides could still shoot each other */
	Finished      time.Time             `json:"finished,omitempty"`        /* Time the war ended and shooting was no longer allowed */
	Mutual        bool                  `json:"mutual,omitempty"`          /* Was the war declared mutual by both parties */
	OpenForAllies bool                  `json:"open_for_allies,omitempty"` /* Is the war currently open for allies or not */
	Aggressor     GetWarsWarIdAggressor `json:"aggressor,omitempty"`
	Defender      GetWarsWarIdDefender  `json:"defender,omitempty"`
	Allies        []GetWarsWarIdAlly    `json:"allies,omitempty"` /* allied corporations or alliances, each object contains either corporation_id or alliance_id */
}