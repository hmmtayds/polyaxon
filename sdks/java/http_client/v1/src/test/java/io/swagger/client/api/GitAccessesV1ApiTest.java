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
import io.swagger.client.model.V1HostAccess;
import io.swagger.client.model.V1ListHostAccessesResponse;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for GitAccessesV1Api
 */
@Ignore
public class GitAccessesV1ApiTest {

    private final GitAccessesV1Api api = new GitAccessesV1Api();

    
    /**
     * List runs
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createGitAccessTest() throws ApiException {
        String owner = null;
        V1HostAccess body = null;
        V1HostAccess response = api.createGitAccess(owner, body);

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
    public void deleteGitAccessTest() throws ApiException {
        String owner = null;
        String uuid = null;
        api.deleteGitAccess(owner, uuid);

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
    public void getGitAccessTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1HostAccess response = api.getGitAccess(owner, uuid);

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
    public void listGitAccessNamesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListHostAccessesResponse response = api.listGitAccessNames(owner, offset, limit, sort, query);

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
    public void listGitAccessesTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListHostAccessesResponse response = api.listGitAccesses(owner, offset, limit, sort, query);

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
    public void patchGitAccessTest() throws ApiException {
        String owner = null;
        String hostAccessUuid = null;
        V1HostAccess body = null;
        V1HostAccess response = api.patchGitAccess(owner, hostAccessUuid, body);

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
    public void updateGitAccessTest() throws ApiException {
        String owner = null;
        String hostAccessUuid = null;
        V1HostAccess body = null;
        V1HostAccess response = api.updateGitAccess(owner, hostAccessUuid, body);

        // TODO: test validations
    }
    
}
