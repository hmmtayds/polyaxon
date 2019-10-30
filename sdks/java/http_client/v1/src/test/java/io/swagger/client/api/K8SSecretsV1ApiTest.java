// Copyright 2019 Polyaxon, Inc.
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

/*
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.api;

import io.swagger.client.ApiException;
import io.swagger.client.model.V1K8SResource;
import io.swagger.client.model.V1ListK8SResourcesResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for K8SSecretsV1Api
 */
@Ignore
public class K8SSecretsV1ApiTest {

    private final K8SSecretsV1Api api = new K8SSecretsV1Api();

    
    /**
     * List runs
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createK8SSecretsTest() throws ApiException {
        String owner = null;
        V1K8SResource body = null;
        V1K8SResource response = api.createK8SSecrets(owner, body);

        // TODO: test validations
    }
    
    /**
     * Patch run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteK8SSecretTest() throws ApiException {
        String owner = null;
        String uuid = null;
        api.deleteK8SSecret(owner, uuid);

        // TODO: test validations
    }
    
    /**
     * Create new run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getK8SSecretTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1K8SResource response = api.getK8SSecret(owner, uuid);

        // TODO: test validations
    }
    
    /**
     * List bookmarked runs for user
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listK8SSecretNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListK8SResourcesResponse response = api.listK8SSecretNames(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * List archived runs for user
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listK8SSecretsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListK8SResourcesResponse response = api.listK8SSecrets(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * Update run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchK8SSecretTest() throws ApiException {
        String owner = null;
        String k8sResourceUuid = null;
        V1K8SResource body = null;
        V1K8SResource response = api.patchK8SSecret(owner, k8sResourceUuid, body);

        // TODO: test validations
    }
    
    /**
     * Get run
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateK8SSecretTest() throws ApiException {
        String owner = null;
        String k8sResourceUuid = null;
        V1K8SResource body = null;
        V1K8SResource response = api.updateK8SSecret(owner, k8sResourceUuid, body);

        // TODO: test validations
    }
    
}
