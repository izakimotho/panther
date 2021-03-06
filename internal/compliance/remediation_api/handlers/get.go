package apihandlers

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
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/panther-labs/panther/pkg/gatewayapi"
	"github.com/panther-labs/panther/pkg/genericapi"
)

// GetRemediations returns the list of remediations available for an organization
func GetRemediations(_ *events.APIGatewayProxyRequest) *events.APIGatewayProxyResponse {
	zap.L().Info("getting list of remediations")
	// TODO - differentiate between different error types
	remediations, err := invoker.GetRemediations()
	if err != nil {
		if _, ok := err.(*genericapi.DoesNotExistError); ok {
			return gatewayapi.MarshalResponse(RemediationLambdaNotFound, http.StatusNotFound)
		}
		zap.L().Warn("failed to fetch available remediations", zap.Error(err))
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}
	}

	body, err := jsoniter.MarshalToString(remediations)
	if err != nil {
		zap.L().Error("failed to marshal remediations", zap.Error(err))
		return &events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError}
	}

	return &events.APIGatewayProxyResponse{StatusCode: http.StatusOK, Body: body}
}
