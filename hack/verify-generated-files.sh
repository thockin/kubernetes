#!/usr/bin/env bash

# Copyright 2018 The Kubernetes Authors.
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

# This script checks whether updating of generated code is needed or not. We
# should run `make generated_files` if generated code is out of date.
# Usage: `hack/verify-generated-files.sh`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
export KUBE_ROOT
source "${KUBE_ROOT}/hack/lib/init.sh"
export GO111MODULE=on # TODO(thockin): remove this when init.sh stops disabling it

kube::util::ensure_clean_working_dir

_tmpdir="$(kube::realpath "$(mktemp -d -t verify-generated-files.XXXXXX)")"
git worktree add -f -q "${_tmpdir}" HEAD
kube::util::trap_add "git worktree remove -f ${_tmpdir}" EXIT
ln -s "${KUBE_ROOT}/_output" "${_tmpdir}/_output" # for GOCACHE
cd "${_tmpdir}"

# regenerate any generated code
make generated_files

changed_files=$(git status --porcelain)

if [[ -n "${changed_files}" ]]; then
  echo "!!! Generated code is out of date:" >&2
  echo "${changed_files}" >&2
  echo >&2
  echo "Please run make generated_files." >&2
  exit 1
fi
