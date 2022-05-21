#!/usr/bin/env bash

# Copyright 2014 The Kubernetes Authors.
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

set -o errexit
set -o nounset
set -o pipefail

# Unset CDPATH so that path interpolation can work correctly
# https://github.com/kubernetes/kubernetes/issues/52255
unset CDPATH

#FIXME: Remove this
# Until all GOPATH references are removed from all build scripts as well,
# explicitly disable module mode to avoid picking up user-set GO111MODULE preferences.
# As individual scripts (like hack/update-vendor.sh) make use of go modules,
# they can explicitly set GO111MODULE=on
export GO111MODULE=off

# The root of the build/dist directory
KUBE_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

# Where output goes.  We should avoid redefining these anywhere else.
#
# KUBE_OUTPUT: the root directory (absolute) where this build should drop any
#     files (subdirs are encouraged).
# KUBE_OUTPUT_BIN: the directory in which compiled binaries will be placed,
#     under OS/ARCH specific subdirs
# THIS_PLATFORM_BIN: a symlink to the output directory for binaries built for
#     the current host platform (e.g. build/test tools).
#
# Compat: The KUBE_OUTPUT_SUBPATH variable is sometimes passed in by callers.
# If it is specified, we'll use it as we always have, and fall back on
# compatible behavior.  Perhaps one day we can remove this.
if [[ -n "${KUBE_OUTPUT_SUBPATH:-}" ]]; then
    if [[ -n "${KUBE_BUILDNAME:-}" ]]; then
        echo "ERROR: \$KUBE_OUTPUT_SUBPATH and \$KUBE_BUILDNAME are mutually exclusive"
        exit 1
    fi
    export KUBE_OUTPUT="${KUBE_ROOT}/${KUBE_OUTPUT_SUBPATH}"
    export THIS_PLATFORM_BIN="${KUBE_ROOT}/_output/bin"
else
    # The "name" of this build variety.
    KUBE_BUILDNAME="${KUBE_BUILDNAME:-"local"}"
    KUBE_OUTPUT_ROOT="${KUBE_ROOT}/_output"
    export KUBE_OUTPUT="${KUBE_OUTPUT_ROOT}/${KUBE_BUILDNAME}"
    export THIS_PLATFORM_BIN="${KUBE_OUTPUT_ROOT}/bin"
fi
export KUBE_OUTPUT_BIN="${KUBE_OUTPUT}/bin"

# This controls rsync compression. Set to a value > 0 to enable rsync
# compression for build container
KUBE_RSYNC_COMPRESS="${KUBE_RSYNC_COMPRESS:-0}"

# Set no_proxy for localhost if behind a proxy, otherwise,
# the connections to localhost in scripts will time out
export no_proxy="127.0.0.1,localhost${no_proxy:+,${no_proxy}}"

source "${KUBE_ROOT}/hack/lib/util.sh"
source "${KUBE_ROOT}/hack/lib/logging.sh"

kube::log::install_errexit
kube::util::ensure-bash-version

source "${KUBE_ROOT}/hack/lib/version.sh"
source "${KUBE_ROOT}/hack/lib/golang.sh"
source "${KUBE_ROOT}/hack/lib/etcd.sh"

# list of all available group versions.  This should be used when generated code
# or when starting an API server that you want to have everything.
# most preferred version for a group should appear first
KUBE_AVAILABLE_GROUP_VERSIONS="${KUBE_AVAILABLE_GROUP_VERSIONS:-\
v1 \
admissionregistration.k8s.io/v1 \
admissionregistration.k8s.io/v1beta1 \
admission.k8s.io/v1 \
admission.k8s.io/v1beta1 \
apps/v1 \
apps/v1beta1 \
apps/v1beta2 \
authentication.k8s.io/v1 \
authentication.k8s.io/v1alpha1 \
authentication.k8s.io/v1beta1 \
authorization.k8s.io/v1 \
authorization.k8s.io/v1beta1 \
autoscaling/v1 \
autoscaling/v2 \
autoscaling/v2beta1 \
autoscaling/v2beta2 \
batch/v1 \
batch/v1beta1 \
certificates.k8s.io/v1 \
certificates.k8s.io/v1beta1 \
coordination.k8s.io/v1beta1 \
coordination.k8s.io/v1 \
discovery.k8s.io/v1 \
discovery.k8s.io/v1beta1 \
extensions/v1beta1 \
events.k8s.io/v1 \
events.k8s.io/v1beta1 \
imagepolicy.k8s.io/v1alpha1 \
networking.k8s.io/v1 \
networking.k8s.io/v1alpha1 \
networking.k8s.io/v1beta1 \
node.k8s.io/v1 \
node.k8s.io/v1alpha1 \
node.k8s.io/v1beta1 \
policy/v1 \
policy/v1beta1 \
rbac.authorization.k8s.io/v1 \
rbac.authorization.k8s.io/v1beta1 \
rbac.authorization.k8s.io/v1alpha1 \
scheduling.k8s.io/v1alpha1 \
scheduling.k8s.io/v1beta1 \
scheduling.k8s.io/v1 \
storage.k8s.io/v1beta1 \
storage.k8s.io/v1 \
storage.k8s.io/v1alpha1 \
flowcontrol.apiserver.k8s.io/v1alpha1 \
flowcontrol.apiserver.k8s.io/v1beta1 \
flowcontrol.apiserver.k8s.io/v1beta2 \
flowcontrol.apiserver.k8s.io/v1beta3 \
internal.apiserver.k8s.io/v1alpha1 \
}"

# not all group versions are exposed by the server.  This list contains those
# which are not available so we don't generate clients or swagger for them
KUBE_NONSERVER_GROUP_VERSIONS="
 abac.authorization.kubernetes.io/v0 \
 abac.authorization.kubernetes.io/v1beta1 \
 apidiscovery.k8s.io/v2beta1 \
 componentconfig/v1alpha1 \
 imagepolicy.k8s.io/v1alpha1\
 admission.k8s.io/v1\
 admission.k8s.io/v1beta1\
"
export KUBE_NONSERVER_GROUP_VERSIONS

# This emulates "readlink -f" which is not available on MacOS X.
# Test:
# T=/tmp/$$.$RANDOM
# mkdir $T
# touch $T/file
# mkdir $T/dir
# ln -s $T/file $T/linkfile
# ln -s $T/dir $T/linkdir
# function testone() {
#   X=$(readlink -f $1 2>&1)
#   Y=$(kube::readlinkdashf $1 2>&1)
#   if [ "$X" != "$Y" ]; then
#     echo readlinkdashf $1: expected "$X", got "$Y"
#   fi
# }
# testone /
# testone /tmp
# testone $T
# testone $T/file
# testone $T/dir
# testone $T/linkfile
# testone $T/linkdir
# testone $T/nonexistant
# testone $T/linkdir/file
# testone $T/linkdir/dir
# testone $T/linkdir/linkfile
# testone $T/linkdir/linkdir
function kube::readlinkdashf {
  # run in a subshell for simpler 'cd'
  (
    if [[ -d "${1}" ]]; then # This also catch symlinks to dirs.
      cd "${1}"
      pwd -P
    else
      cd "$(dirname "${1}")"
      local f
      f=$(basename "${1}")
      if [[ -L "${f}" ]]; then
        readlink "${f}"
      else
        echo "$(pwd -P)/${f}"
      fi
    fi
  )
}

# This emulates "realpath" which is not available on MacOS X
# Test:
# T=/tmp/$$.$RANDOM
# mkdir $T
# touch $T/file
# mkdir $T/dir
# ln -s $T/file $T/linkfile
# ln -s $T/dir $T/linkdir
# function testone() {
#   X=$(realpath $1 2>&1)
#   Y=$(kube::realpath $1 2>&1)
#   if [ "$X" != "$Y" ]; then
#     echo realpath $1: expected "$X", got "$Y"
#   fi
# }
# testone /
# testone /tmp
# testone $T
# testone $T/file
# testone $T/dir
# testone $T/linkfile
# testone $T/linkdir
# testone $T/nonexistant
# testone $T/linkdir/file
# testone $T/linkdir/dir
# testone $T/linkdir/linkfile
# testone $T/linkdir/linkdir
kube::realpath() {
  if [[ ! -e "${1}" ]]; then
    echo "${1}: No such file or directory" >&2
    return 1
  fi
  kube::readlinkdashf "${1}"
}
