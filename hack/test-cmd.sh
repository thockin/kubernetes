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

# This command checks that the built commands can function together for
# simple scenarios.  It does not require Docker so it can run in travis.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

function cleanup()
{
    [[ -n ${APISERVER_PID-} ]] && kill ${APISERVER_PID} 1>&2 2>/dev/null
    [[ -n ${CTLRMGR_PID-} ]] && kill ${CTLRMGR_PID} 1>&2 2>/dev/null
    [[ -n ${KUBELET_PID-} ]] && kill ${KUBELET_PID} 1>&2 2>/dev/null
    [[ -n ${PROXY_PID-} ]] && kill ${PROXY_PID} 1>&2 2>/dev/null

    kube::etcd::cleanup

    kube::log::status "Clean up complete"
}

trap cleanup EXIT SIGINT

kube::etcd::start

ETCD_HOST=${ETCD_HOST:-127.0.0.1}
ETCD_PORT=${ETCD_PORT:-4001}
API_PORT=${API_PORT:-8080}
API_HOST=${API_HOST:-127.0.0.1}
KUBELET_PORT=${KUBELET_PORT:-10250}
CTLRMGR_PORT=${CTLRMGR_PORT:-10252}

# Check kubectl
kube::log::status "Running kubectl with no options"
"${KUBE_OUTPUT_HOSTBIN}/kubectl"

# Start kubelet
kube::log::status "Starting kubelet"
"${KUBE_OUTPUT_HOSTBIN}/kubelet" \
  --root_dir=/tmp/kubelet.$$ \
  --etcd_servers="http://${ETCD_HOST}:${ETCD_PORT}" \
  --hostname_override="127.0.0.1" \
  --address="127.0.0.1" \
  --port="$KUBELET_PORT" 1>&2 &
KUBELET_PID=$!

kube::util::wait_for_url "http://127.0.0.1:${KUBELET_PORT}/healthz" "kubelet: "

# Start kube-apiserver
kube::log::status "Starting kube-apiserver"
"${KUBE_OUTPUT_HOSTBIN}/kube-apiserver" \
  --address="127.0.0.1" \
  --public_address_override="127.0.0.1" \
  --port="${API_PORT}" \
  --etcd_servers="http://${ETCD_HOST}:${ETCD_PORT}" \
  --public_address_override="127.0.0.1" \
  --kubelet_port=${KUBELET_PORT} \
  --portal_net="10.0.0.0/24" 1>&2 &
APISERVER_PID=$!

kube::util::wait_for_url "http://127.0.0.1:${API_PORT}/healthz" "apiserver: "

kube_cmd=(
  "${KUBE_OUTPUT_HOSTBIN}/kubectl"
)

kube_flags=(
  -s "http://127.0.0.1:${API_PORT}"
  --match-server-version
)

# Start controller manager
kube::log::status "Starting CONTROLLER-MANAGER"
"${KUBE_OUTPUT_HOSTBIN}/kube-controller-manager" \
  --machines="127.0.0.1" \
  --master="127.0.0.1:${API_PORT}" 1>&2 &
CTLRMGR_PID=$!

kube::util::wait_for_url "http://127.0.0.1:${CTLRMGR_PORT}/healthz" "controller-manager: "

kube::log::status "Testing kubectl(pods)"
"${kube_cmd[@]}" get pods "${kube_flags[@]}"
"${kube_cmd[@]}" create -f examples/guestbook/redis-master.json "${kube_flags[@]}"
"${kube_cmd[@]}" get pods "${kube_flags[@]}"
"${kube_cmd[@]}" get pod redis-master "${kube_flags[@]}"
[[ "$("${kube_cmd[@]}" get pod redis-master -o template --output-version=v1beta1 -t '{{ .id }}' "${kube_flags[@]}")" == "redis-master" ]]
output_pod=$("${kube_cmd[@]}" get pod redis-master -o json --output-version=v1beta1 "${kube_flags[@]}")
"${kube_cmd[@]}" delete pod redis-master "${kube_flags[@]}"
[[ $("${kube_cmd[@]}" get pods -o template -t '{{ len .items }}' "${kube_flags[@]}") -eq 0 ]]
echo $output_pod | "${kube_cmd[@]}" create -f - "${kube_flags[@]}"
[[ $("${kube_cmd[@]}" get pods -o template -t '{{ len .items }}' "${kube_flags[@]}") -eq 1 ]]
"${kube_cmd[@]}" get pods -o yaml "${kube_flags[@]}" | grep -q "id: redis-master"
"${kube_cmd[@]}" describe pod redis-master "${kube_flags[@]}" | grep -q 'Name:.*redis-master'
"${kube_cmd[@]}" delete -f examples/guestbook/redis-master.json "${kube_flags[@]}"

kube::log::status "Testing kubectl(services)"
"${kube_cmd[@]}" get services "${kube_flags[@]}"
"${kube_cmd[@]}" create -f examples/guestbook/frontend-service.json "${kube_flags[@]}"
"${kube_cmd[@]}" get services "${kube_flags[@]}"
"${kube_cmd[@]}" delete service frontend "${kube_flags[@]}"

kube::log::status "Testing kubectl(replicationcontrollers)"
"${kube_cmd[@]}" get replicationcontrollers "${kube_flags[@]}"
"${kube_cmd[@]}" create -f examples/guestbook/frontend-controller.json "${kube_flags[@]}"
"${kube_cmd[@]}" get replicationcontrollers "${kube_flags[@]}"
"${kube_cmd[@]}" describe replicationcontroller frontendController "${kube_flags[@]}" | grep -q 'Replicas:.*3 desired'

kube::log::status "Testing kubectl(minions)"
"${kube_cmd[@]}" get minions "${kube_flags[@]}"
"${kube_cmd[@]}" get minions 127.0.0.1 "${kube_flags[@]}"

kube::log::status "TEST PASSED"

# Start proxy
#PROXY_LOG=/tmp/kube-proxy.log
#${KUBE_OUTPUT_HOSTBIN}/kube-proxy \
#  --etcd_servers="http://127.0.0.1:${ETCD_PORT}" 1>&2 &
#PROXY_PID=$!
