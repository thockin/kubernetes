#!/usr/bin/env bash

# Copyright 2015 The Kubernetes Authors.
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

# This script checks whether updating of swagger type documentation is needed or
# not. We should run `hack/update-generated-swagger-docs.sh` if swagger type
# documentation is out of date.
# Usage: `hack/verify-generated-swagger-docs.sh`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

kube::util::ensure_clean_working_dir

_tmpdir="$(kube::realpath "$(mktemp -d -t swagger-docs.XXXXXX)")"
git worktree add -f -q "${_tmpdir}" HEAD
kube::util::trap_add "git worktree remove -f ${_tmpdir}" EXIT
ln -s "${KUBE_ROOT}/_output" "${_tmpdir}/_output" # for GOCACHE
cd "${_tmpdir}"

# Update the generated swagger docs
hack/update-generated-swagger-docs.sh

# Test for diffs
diffs=$(git status --porcelain --untracked-files=no | wc -l)
if [[ ${diffs} -gt 0 ]]; then
  echo "Generated swagger type documentation is out of date:"
  git diff
  exit 1
fi

echo "Generated swagger type documentation up to date."
