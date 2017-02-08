/* 
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.3.10.dev19
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

package goesiv1

/* logo_urls object */
type GetCorporationsCorporationIdLogoUrls struct {

	/* 128x128 string */
	Var128x128 string `json:"128x128,omitempty"`

	/* 256x256 string */
	Var256x256 string `json:"256x256,omitempty"`

	/* 64x64 string */
	Var64x64 string `json:"64x64,omitempty"`
}
