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

import React from 'react';
import { Alert, Button, Card, Flex, Icon } from 'pouncejs';
import { INTEGRATION_TYPES, RESOURCE_TYPES } from 'Source/constants';
import GenerateFiltersGroup from 'Components/utils/generate-filters-group';

import { useQuery, gql } from '@apollo/client';
import { ComplianceStatusEnum, ListResourcesInput, Integration } from 'Generated/schema';
import { capitalize } from 'Helpers/utils';
import FormikTextInput from 'Components/fields/text-input';
import FormikCombobox from 'Components/fields/combobox';
import FormikMultiCombobox from 'Components/fields/multi-combobox';
import ErrorBoundary from 'Components/error-boundary';
import pick from 'lodash-es/pick';
import useRequestParamsWithPagination from 'Hooks/useRequestParamsWithPagination';
import isEmpty from 'lodash-es/isEmpty';

const statusOptions = Object.values(ComplianceStatusEnum);

export const filters = {
  types: {
    component: FormikMultiCombobox,
    props: {
      items: RESOURCE_TYPES,
      label: 'Types',
      searchable: true,
      inputProps: {
        placeholder: 'Start typing resources...',
      },
    },
  },
  integrationId: {
    component: FormikCombobox,
    props: {
      label: 'Source',
      searchable: true,
      items: [] as Integration[],
      itemToString: (integration: Integration) => integration.integrationLabel,
      inputProps: {
        placeholder: 'Choose a source...',
      },
    },
  },
  complianceStatus: {
    component: FormikCombobox,
    props: {
      label: 'Status',
      itemToString: (status: ComplianceStatusEnum) => capitalize(status.toLowerCase()),
      items: statusOptions,
      inputProps: {
        placeholder: 'Choose a status...',
      },
    },
  },
  idContains: {
    component: FormikTextInput,
    props: {
      label: 'ID / Name',
      placeholder: 'Enter part of an id or a name...',
    },
  },
};

const LIST_ACCOUNT_IDS = gql`
  query ListAccountIds {
    integrations(input: { integrationType: "${INTEGRATION_TYPES.AWS_INFRA}"}) {
        integrationLabel
        integrationId
    }
  }
`;

// The values of the filters that the resources page will show
export type ListResourcesFiltersValues = Pick<
  ListResourcesInput,
  'types' | 'complianceStatus' | 'idContains' | 'integrationId'
>;

// we mutate the shape of the integrationID here since we want it to have a different shape, that's
// dependant on the response of another API
type MutatedListResourcesFiltersValues = Omit<ListResourcesFiltersValues, 'integrationId'> & {
  integrationId: Pick<Integration, 'integrationId' | 'integrationLabel'>;
};

interface ListResourcesFiltersProps {
  onCancel: () => void;
  onSubmit: (values: ListResourcesFiltersValues) => void;
  initialValues: ListResourcesFiltersValues;
}

const ListResourcesActions: React.FC = () => {
  const [areFiltersVisible, setFiltersVisibility] = React.useState(false);
  const { requestParams, updateRequestParamsAndResetPaging } = useRequestParamsWithPagination<
    ListResourcesInput
  >();

  const { error, data } = useQuery<{ integrations: Integration[] }>(LIST_ACCOUNT_IDS, {
    fetchPolicy: 'cache-first',
  });

  if (data) {
    filters.integrationId.props.items = data.integrations;
  }

  // Just because the `integrationId` field has objects as items, when a value is selected we have
  // an object of the shape {integrationId,integrationLabel} selected in our form. We need to
  // convert that back to integrationId so that the `onSubmit` given as prop to this component can
  // get the value it should expect
  const handleFiltersSubmit = React.useCallback(
    ({ integrationId: integrationObj, ...values }: MutatedListResourcesFiltersValues) => {
      updateRequestParamsAndResetPaging({
        ...values,
        integrationId: integrationObj ? integrationObj.integrationId : null,
      });
    },
    []
  );

  // Mutate initial values since the initial values provide an `integrationId` and we want to map
  // that to an `Integration` object, since that is the kind of items that the MultiCombobox has
  const filterKeys = Object.keys(filters) as (keyof ListResourcesFiltersValues)[];
  const filtersCount = filterKeys.filter(key => !isEmpty(requestParams[key])).length;
  const mutatedInitialValues = React.useMemo(
    () => ({
      ...(pick(requestParams, filterKeys) as ListResourcesFiltersValues),
      integrationId:
        data?.integrations.find(i => i.integrationId === requestParams.integrationId) || null,
    }),
    [requestParams, data]
  );

  return (
    <React.Fragment>
      {error && <Alert variant="error" title="Failed to fetch available sources" discardable />}
      <Flex justifyContent="flex-end" mb={6} position="relative">
        <Button
          size="large"
          variant="default"
          onClick={() => setFiltersVisibility(!areFiltersVisible)}
        >
          <Flex>
            <Icon type="filter" size="small" mr={3} />
            Filter Options {filtersCount ? `(${filtersCount})` : ''}
          </Flex>
        </Button>
      </Flex>
      {areFiltersVisible && (
        <ErrorBoundary>
          <Card p={6} mb={6}>
            <GenerateFiltersGroup<MutatedListResourcesFiltersValues>
              filters={filters}
              onCancel={() => setFiltersVisibility(false)}
              onSubmit={handleFiltersSubmit}
              initialValues={mutatedInitialValues}
            />
          </Card>
        </ErrorBoundary>
      )}
    </React.Fragment>
  );
};

export default React.memo(ListResourcesActions);
