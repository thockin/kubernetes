#!/bin/bash

# Copyright 2014 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

## Contains configuration values for interacting with the Vagrant cluster

# NUMBER OF MINIONS IN THE CLUSTER
NUM_MINIONS=${KUBERNETES_NUM_MINIONS-"3"}

# IP LOCATIONS FOR INTERACTING WITH THE MASTER
export KUBE_MASTER_IP="10.245.1.2"

INSTANCE_PREFIX=kubernetes
MASTER_NAME="${INSTANCE_PREFIX}-master"
MASTER_TAG="${INSTANCE_PREFIX}-master"
MINION_TAG="${INSTANCE_PREFIX}-minion"
# Unable to use hostnames yet because DNS is not in cluster, so we revert external look-up name to use the minion IP
#MINION_NAMES=($(eval echo ${INSTANCE_PREFIX}-minion-{1..${NUM_MINIONS}}))

# IP LOCATIONS FOR INTERACTING WITH THE MINIONS
MINION_IP_BASE="10.245.2."
declare -A VAGRANT_MINION_NAMES_BY_IP

for (( i=0; i <${NUM_MINIONS}; i++)) do
  KUBE_MINION_IP_ADDRESSES[$i]="${MINION_IP_BASE}$[$i+2]"
  MINION_IP[$i]="${MINION_IP_BASE}$[$i+2]"
  MINION_NAMES[$i]="${MINION_IP[$i]}"
  VAGRANT_MINION_NAMES[$i]="minion-$[$i+1]"

  VAGRANT_MINION_NAMES_BY_IP["${MINION_IP[$i]}"]="${VAGRANT_MINION_NAMES[$i]}"
done

