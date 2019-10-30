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


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import org.threeten.bp.OffsetDateTime;

/**
 * V1K8SResource
 */

public class V1K8SResource {
  @SerializedName("uuid")
  private String uuid = null;

  @SerializedName("name")
  private String name = null;

  @SerializedName("description")
  private String description = null;

  @SerializedName("readme")
  private String readme = null;

  @SerializedName("tags")
  private List<String> tags = null;

  @SerializedName("created_at")
  private OffsetDateTime createdAt = null;

  @SerializedName("updated_at")
  private OffsetDateTime updatedAt = null;

  @SerializedName("k8s_ref")
  private String k8sRef = null;

  @SerializedName("frozen")
  private Boolean frozen = null;

  @SerializedName("disabled")
  private Boolean disabled = null;

  @SerializedName("deleted")
  private Boolean deleted = null;

  @SerializedName("mount_path")
  private String mountPath = null;

  @SerializedName("items")
  private List<String> items = null;

  public V1K8SResource uuid(String uuid) {
    this.uuid = uuid;
    return this;
  }

   /**
   * Get uuid
   * @return uuid
  **/
  @ApiModelProperty(value = "")
  public String getUuid() {
    return uuid;
  }

  public void setUuid(String uuid) {
    this.uuid = uuid;
  }

  public V1K8SResource name(String name) {
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @ApiModelProperty(value = "")
  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public V1K8SResource description(String description) {
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @ApiModelProperty(value = "")
  public String getDescription() {
    return description;
  }

  public void setDescription(String description) {
    this.description = description;
  }

  public V1K8SResource readme(String readme) {
    this.readme = readme;
    return this;
  }

   /**
   * Get readme
   * @return readme
  **/
  @ApiModelProperty(value = "")
  public String getReadme() {
    return readme;
  }

  public void setReadme(String readme) {
    this.readme = readme;
  }

  public V1K8SResource tags(List<String> tags) {
    this.tags = tags;
    return this;
  }

  public V1K8SResource addTagsItem(String tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<String>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Get tags
   * @return tags
  **/
  @ApiModelProperty(value = "")
  public List<String> getTags() {
    return tags;
  }

  public void setTags(List<String> tags) {
    this.tags = tags;
  }

  public V1K8SResource createdAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
    return this;
  }

   /**
   * Get createdAt
   * @return createdAt
  **/
  @ApiModelProperty(value = "")
  public OffsetDateTime getCreatedAt() {
    return createdAt;
  }

  public void setCreatedAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
  }

  public V1K8SResource updatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @ApiModelProperty(value = "")
  public OffsetDateTime getUpdatedAt() {
    return updatedAt;
  }

  public void setUpdatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
  }

  public V1K8SResource k8sRef(String k8sRef) {
    this.k8sRef = k8sRef;
    return this;
  }

   /**
   * Get k8sRef
   * @return k8sRef
  **/
  @ApiModelProperty(value = "")
  public String getK8sRef() {
    return k8sRef;
  }

  public void setK8sRef(String k8sRef) {
    this.k8sRef = k8sRef;
  }

  public V1K8SResource frozen(Boolean frozen) {
    this.frozen = frozen;
    return this;
  }

   /**
   * Get frozen
   * @return frozen
  **/
  @ApiModelProperty(value = "")
  public Boolean isFrozen() {
    return frozen;
  }

  public void setFrozen(Boolean frozen) {
    this.frozen = frozen;
  }

  public V1K8SResource disabled(Boolean disabled) {
    this.disabled = disabled;
    return this;
  }

   /**
   * Get disabled
   * @return disabled
  **/
  @ApiModelProperty(value = "")
  public Boolean isDisabled() {
    return disabled;
  }

  public void setDisabled(Boolean disabled) {
    this.disabled = disabled;
  }

  public V1K8SResource deleted(Boolean deleted) {
    this.deleted = deleted;
    return this;
  }

   /**
   * Get deleted
   * @return deleted
  **/
  @ApiModelProperty(value = "")
  public Boolean isDeleted() {
    return deleted;
  }

  public void setDeleted(Boolean deleted) {
    this.deleted = deleted;
  }

  public V1K8SResource mountPath(String mountPath) {
    this.mountPath = mountPath;
    return this;
  }

   /**
   * Get mountPath
   * @return mountPath
  **/
  @ApiModelProperty(value = "")
  public String getMountPath() {
    return mountPath;
  }

  public void setMountPath(String mountPath) {
    this.mountPath = mountPath;
  }

  public V1K8SResource items(List<String> items) {
    this.items = items;
    return this;
  }

  public V1K8SResource addItemsItem(String itemsItem) {
    if (this.items == null) {
      this.items = new ArrayList<String>();
    }
    this.items.add(itemsItem);
    return this;
  }

   /**
   * Get items
   * @return items
  **/
  @ApiModelProperty(value = "")
  public List<String> getItems() {
    return items;
  }

  public void setItems(List<String> items) {
    this.items = items;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1K8SResource v1K8SResource = (V1K8SResource) o;
    return Objects.equals(this.uuid, v1K8SResource.uuid) &&
        Objects.equals(this.name, v1K8SResource.name) &&
        Objects.equals(this.description, v1K8SResource.description) &&
        Objects.equals(this.readme, v1K8SResource.readme) &&
        Objects.equals(this.tags, v1K8SResource.tags) &&
        Objects.equals(this.createdAt, v1K8SResource.createdAt) &&
        Objects.equals(this.updatedAt, v1K8SResource.updatedAt) &&
        Objects.equals(this.k8sRef, v1K8SResource.k8sRef) &&
        Objects.equals(this.frozen, v1K8SResource.frozen) &&
        Objects.equals(this.disabled, v1K8SResource.disabled) &&
        Objects.equals(this.deleted, v1K8SResource.deleted) &&
        Objects.equals(this.mountPath, v1K8SResource.mountPath) &&
        Objects.equals(this.items, v1K8SResource.items);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, name, description, readme, tags, createdAt, updatedAt, k8sRef, frozen, disabled, deleted, mountPath, items);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1K8SResource {\n");
    
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    readme: ").append(toIndentedString(readme)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
    sb.append("    k8sRef: ").append(toIndentedString(k8sRef)).append("\n");
    sb.append("    frozen: ").append(toIndentedString(frozen)).append("\n");
    sb.append("    disabled: ").append(toIndentedString(disabled)).append("\n");
    sb.append("    deleted: ").append(toIndentedString(deleted)).append("\n");
    sb.append("    mountPath: ").append(toIndentedString(mountPath)).append("\n");
    sb.append("    items: ").append(toIndentedString(items)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

