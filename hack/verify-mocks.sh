#!/usr/bin/env bash

# Copyright 2021 The Kubernetes Authors.
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

# This script checks whether updating of Mock files generated from Interfaces
# is needed or not. We should run `hack/update-mocks.sh` 
# if Mock files are out of date.
# Usage: `hack/verify-mocks.sh`.


set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
export KUBE_ROOT
source "${KUBE_ROOT}/hack/lib/init.sh"

# Explicitly opt into go modules
export GO111MODULE=on

kube::util::ensure_clean_working_dir

_tmpdir="$(kube::realpath "$(mktemp -d -t verify-mocks.XXXXXX)")"
git worktree add -f -q "${_tmpdir}" HEAD
kube::util::trap_add "git worktree remove -f ${_tmpdir}" EXIT
ln -s "${KUBE_ROOT}/_output" "${_tmpdir}/_output" # for GOCACHE
cd "${_tmpdir}"

# Update the mocks in ${_tmpdir}
hack/update-mocks.sh

# Test for diffs
diffs=$(git status --porcelain --untracked-files=no | wc -l)
if [[ ${diffs} -gt 0 ]]; then
  echo "Mock files are out of date. Please run hack/update-mocks.sh" >&2
  git diff
  exit 1
fi

echo "up to date"
