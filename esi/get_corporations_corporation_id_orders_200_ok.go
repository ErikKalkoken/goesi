/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.7.3
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

/* A list of GetCorporationsCorporationIdOrders200Ok. */
//easyjson:json
type GetCorporationsCorporationIdOrders200OkList []GetCorporationsCorporationIdOrders200Ok

/* 200 ok object */
//easyjson:json
type GetCorporationsCorporationIdOrders200Ok struct {
	OrderId        int64     `json:"order_id,omitempty"`        /* Unique order ID */
	TypeId         int32     `json:"type_id,omitempty"`         /* The type ID of the item transacted in this order */
	RegionId       int32     `json:"region_id,omitempty"`       /* ID of the region where order was placed */
	LocationId     int64     `json:"location_id,omitempty"`     /* ID of the location where order was placed */
	Range_         string    `json:"range,omitempty"`           /* Valid order range, numbers are ranges in jumps */
	IsBuyOrder     bool      `json:"is_buy_order,omitempty"`    /* True for a bid (buy) order. False for an offer (sell) order */
	Price          float32   `json:"price,omitempty"`           /* Cost per unit for this order */
	VolumeTotal    int32     `json:"volume_total,omitempty"`    /* Quantity of items required or offered at time order was placed */
	VolumeRemain   int32     `json:"volume_remain,omitempty"`   /* Quantity of items still required or offered */
	Issued         time.Time `json:"issued,omitempty"`          /* Date and time when this order was issued */
	State          string    `json:"state,omitempty"`           /* Current order state */
	MinVolume      int32     `json:"min_volume,omitempty"`      /* For bids (buy orders), the minimum quantity that will be accepted in a matching offer (sell order) */
	WalletDivision int32     `json:"wallet_division,omitempty"` /* Wallet division of which this order used */
	Duration       int32     `json:"duration,omitempty"`        /* Numer of days for which order is valid (starting from the issued date). An order expires at time issued + duration */
	Escrow         float32   `json:"escrow,omitempty"`          /* For buy orders, the amount of ISK in escrow */
}
