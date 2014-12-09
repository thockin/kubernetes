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

# The repositories are really slow and there are GCE mirrors
sed -i -e "\|^deb.*http://http.debian.net/debian| s/^/#/" /etc/apt/sources.list
sed -i -e "\|^deb.*http://ftp.debian.org/debian| s/^/#/" /etc/apt/sources.list.d/backports.list

# Prepopulate the name of the Master
mkdir -p /etc/salt/minion.d
echo "master: $MASTER_NAME" > /etc/salt/minion.d/master.conf

cat <<EOF >/etc/salt/minion.d/log-level-debug.conf
log_level: debug
log_level_logfile: debug
EOF

# Our minions will have a pool role to distinguish them from the master.
cat <<EOF >/etc/salt/minion.d/grains.conf
grains:
  roles:
    - kubernetes-pool
  cbr-cidr: $MINION_IP_RANGE
  cloud: gce
EOF

# Decide if enable the cache
if [[ "${ENABLE_DOCKER_REGISTRY_CACHE}" == "true" ]]; then 
    REGION=$(echo "${ZONE}" | cut -f 1,2 -d -)
    echo "Enable docker registry cache at region: " $REGION
    DOCKER_OPTS="--registry-mirror=\"https://${REGION}.docker-cache.clustermaster.net\""

    cat <<EOF >>/etc/salt/minion.d/grains.conf
  docker_opts: $DOCKER_OPTS
EOF
fi

install-salt

# Wait a few minutes and trigger another Salt run to better recover from
# any transient errors.
echo "Sleeping 180"
sleep 180
salt-call state.highstate || true
