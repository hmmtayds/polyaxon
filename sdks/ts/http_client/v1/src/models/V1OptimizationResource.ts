// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.9-rc6
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1ResourceType,
    V1ResourceTypeFromJSON,
    V1ResourceTypeFromJSONTyped,
    V1ResourceTypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1OptimizationResource
 */
export interface V1OptimizationResource {
    /**
     * 
     * @type {string}
     * @memberof V1OptimizationResource
     */
    name?: string;
    /**
     * 
     * @type {V1ResourceType}
     * @memberof V1OptimizationResource
     */
    type?: V1ResourceType;
}

export function V1OptimizationResourceFromJSON(json: any): V1OptimizationResource {
    return V1OptimizationResourceFromJSONTyped(json, false);
}

export function V1OptimizationResourceFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1OptimizationResource {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'name': !exists(json, 'name') ? undefined : json['name'],
        'type': !exists(json, 'type') ? undefined : V1ResourceTypeFromJSON(json['type']),
    };
}

export function V1OptimizationResourceToJSON(value?: V1OptimizationResource | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'name': value.name,
        'type': V1ResourceTypeToJSON(value.type),
    };
}


