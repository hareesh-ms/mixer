// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Model objects related to the SQL database.
package sqldb

import (
	"fmt"
	"strings"
)

// Observation struct to represent a row in the observations table
type Observation struct {
	Entity            string  `db:"entity"`
	Variable          string  `db:"variable"`
	Date              string  `db:"date"`
	Value             float64 `db:"value"`
	Provenance        string  `db:"provenance"`
	Unit              string  `db:"unit"`
	ScalingFactor     string  `db:"scaling_factor"`
	MeasurementMethod string  `db:"measurement_method"`
	ObservationPeriod string  `db:"observation_period"`
	Properties        string  `db:"properties"`
}

// SVSummary represents a SV summary row.
type SVSummary struct {
	Variable        string      `db:"variable"`
	EntityType      string      `db:"entity_type"`
	EntityCount     int32       `db:"entity_count"`
	MinValue        float64     `db:"min_value"`
	MaxValue        float64     `db:"max_value"`
	SampleEntityIds StringSlice `db:"sample_entity_ids"`
}

// StringSlice is a custom scanner for comma-separated strings.
type StringSlice []string

// Scan implements the sql.Scanner interface and decodes a comma-separated string field into a StringSlice ([]string) value.
func (s *StringSlice) Scan(src interface{}) error {
	if src == nil {
		*s = []string{}
		return nil
	}

	val, ok := src.(string)
	if !ok {
		return fmt.Errorf("failed to decode []string: (%v)", src)
	}
	*s = strings.Split(val, ",")
	return nil
}
