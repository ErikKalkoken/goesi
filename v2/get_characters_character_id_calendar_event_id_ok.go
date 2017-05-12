/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.4.5
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

package goesiv2

import (
	"time"
)

/* Full details of a specific event */
type GetCharactersCharacterIdCalendarEventIdOk struct {
	DurationInMinutes int32     `json:"duration_in_minutes,omitempty"` /* duration_in_minutes integer */
	EventDate         time.Time `json:"event_date,omitempty"`          /* event_date string */
	EventId           int32     `json:"event_id,omitempty"`            /* event_id integer */
	EventResponse     string    `json:"event_response,omitempty"`      /* event_response string */
	EventText         string    `json:"event_text,omitempty"`          /* event_text string */
	Importance        int32     `json:"importance,omitempty"`          /* importance integer */
	OwnerId           int32     `json:"owner_id,omitempty"`            /* owner_id integer */
	OwnerName         string    `json:"owner_name,omitempty"`          /* owner_name string */
	OwnerTypeId       int32     `json:"owner_type_id,omitempty"`       /* owner_type_id integer */
	Title             string    `json:"title,omitempty"`               /* title string */
}
