package osquerylogs

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/panther-labs/panther/internal/log_analysis/log_processor/parsers"
	"github.com/panther-labs/panther/internal/log_analysis/log_processor/parsers/timestamp"
)

var SnapshotDesc = `Snapshot contains all the data included in OsQuery differential logs
Reference: https://osquery.readthedocs.io/en/stable/deployment/logging/`

// nolint:lll
type Snapshot struct {
	Action         *string                `json:"action,omitempty" validate:"required,eq=snapshot" description:""`
	CalendarTime   *timestamp.ANSICwithTZ `json:"calendarTime,omitempty" validate:"required" description:"The time of the event (UTC)."`
	Counter        *int                   `json:"counter,omitempty,string" validate:"required" description:""`
	Decorations    map[string]string      `json:"decorations,omitempty" description:""`
	Epoch          *int                   `json:"epoch,omitempty,string" validate:"required" description:""`
	HostIdentifier *string                `json:"hostIdentifier,omitempty" validate:"required" description:""`
	Name           *string                `json:"name,omitempty" validate:"required" description:""`
	Snapshot       []map[string]string    `json:"snapshot,omitempty" validate:"required" description:""`
	UnixTime       *int                   `json:"unixTime,omitempty,string" validate:"required" description:""`

	// NOTE: added to end of struct to allow expansion later
	parsers.PantherLog
}

// SnapshotParser parses OsQuery snapshot logs
type SnapshotParser struct{}

// Parse returns the parsed events or nil if parsing failed
func (p *SnapshotParser) Parse(log string) []interface{} {
	event := &Snapshot{}
	err := jsoniter.UnmarshalFromString(log, event)
	if err != nil {
		zap.L().Debug("failed to unmarshal log", zap.Error(err))
		return nil
	}

	event.updatePantherFields(p)

	if err := parsers.Validator.Struct(event); err != nil {
		zap.L().Debug("failed to validate log", zap.Error(err))
		return nil
	}
	return []interface{}{event}
}

// LogType returns the log type supported by this parser
func (p *SnapshotParser) LogType() string {
	return "Osquery.Snapshot"
}

func (event *Snapshot) updatePantherFields(p *SnapshotParser) {
	// panther fields
	if event.CalendarTime != nil {
		event.SetRequired(p.LogType(), timestamp.RFC3339(*event.CalendarTime))
	}
	if event.HostIdentifier != nil {
		event.AppendAnyDomainNames(*event.HostIdentifier)
	}
}
