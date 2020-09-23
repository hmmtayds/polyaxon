#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    The version of the OpenAPI document: 1.1.9-rc6
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1OperationCond(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {"io_conidtion": "V1IoCond", "status_condition": "V1StatusCond"}

    attribute_map = {
        "io_conidtion": "io_conidtion",
        "status_condition": "status_condition",
    }

    def __init__(
        self, io_conidtion=None, status_condition=None, local_vars_configuration=None
    ):  # noqa: E501
        """V1OperationCond - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._io_conidtion = None
        self._status_condition = None
        self.discriminator = None

        if io_conidtion is not None:
            self.io_conidtion = io_conidtion
        if status_condition is not None:
            self.status_condition = status_condition

    @property
    def io_conidtion(self):
        """Gets the io_conidtion of this V1OperationCond.  # noqa: E501


        :return: The io_conidtion of this V1OperationCond.  # noqa: E501
        :rtype: V1IoCond
        """
        return self._io_conidtion

    @io_conidtion.setter
    def io_conidtion(self, io_conidtion):
        """Sets the io_conidtion of this V1OperationCond.


        :param io_conidtion: The io_conidtion of this V1OperationCond.  # noqa: E501
        :type: V1IoCond
        """

        self._io_conidtion = io_conidtion

    @property
    def status_condition(self):
        """Gets the status_condition of this V1OperationCond.  # noqa: E501


        :return: The status_condition of this V1OperationCond.  # noqa: E501
        :rtype: V1StatusCond
        """
        return self._status_condition

    @status_condition.setter
    def status_condition(self, status_condition):
        """Sets the status_condition of this V1OperationCond.


        :param status_condition: The status_condition of this V1OperationCond.  # noqa: E501
        :type: V1StatusCond
        """

        self._status_condition = status_condition

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(
                    map(lambda x: x.to_dict() if hasattr(x, "to_dict") else x, value)
                )
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(
                    map(
                        lambda item: (item[0], item[1].to_dict())
                        if hasattr(item[1], "to_dict")
                        else item,
                        value.items(),
                    )
                )
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1OperationCond):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1OperationCond):
            return True

        return self.to_dict() != other.to_dict()
