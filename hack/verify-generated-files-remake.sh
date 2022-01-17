#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
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

# This script verifies that the expected results are obtained when creating each
# type of file(e.g. codegen tool itself, a file in a package that needs codegen,
# and etc.) for verification and then generating the code(executes
# `make generated_files`).
# Usage: `hack/verify-generated-files-remake.sh`.

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"
export GO111MODULE=on # TODO(thockin): remove this when init.sh stops disabling it

kube::util::ensure_clean_working_dir

_tmpdir="$(kube::realpath "$(mktemp -d -t verify-generated-files.XXXXXX)")"
git worktree add -f -q "${_tmpdir}" HEAD
kube::util::trap_add "git worktree remove -f ${_tmpdir}" EXIT
ln -s "${KUBE_ROOT}/_output" "${_tmpdir}/_output" # for GOCACHE
cd "${_tmpdir}"

# $1 = filename pattern as in "zz_generated.$1.go"
function find_genfiles() {
    find .                         \
        \(                         \
          -not \(                  \
            \(                     \
                -path ./_\* -o     \
                -path ./.\*        \
            \) -prune              \
          \)                       \
        \) -name "zz_generated.$1.go"
}

# $1 = filename pattern as in "zz_generated.$1.go"
# $2 timestamp file
function newer() {
    find_genfiles "$1" | while read -r F; do
        if [[ "${F}" -nt "$2" ]]; then
            echo "${F}"
        fi
    done | LC_ALL=C sort
}

# $1 = filename pattern as in "zz_generated.$1.go"
# $2 timestamp file
function older() {
    find_genfiles "$1" | while read -r F; do
        if [[ "$2" -nt "${F}" ]]; then
            echo "${F}"
        fi
    done | LC_ALL=C sort
}

# Pipe commands through this to indent their output.
function indent() {
     sed 's/^/    /'
}

STAMP=/tmp/stamp.$RANDOM

#
# Test back-to-back builds.
#

echo "CASE: back-to-back builds"
make generated_files 2>&1 | indent
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer deepcopy "${STAMP}")"
if [[ -n "${X}" ]]; then
    echo "Generated files changed on back-to-back 'make' runs:"
    echo "  ${X}" | tr '\n' ' '
    echo ""
    exit 1
fi

#
# Test when we touch a file in a package that needs codegen.
#

echo
echo "CASE: touch a file in a package that needs codegen"
DIR=staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1
touch "${DIR}/types.go"
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer deepcopy "${STAMP}")"
if [[ -z "${X}" || ${X} != "./${DIR}/zz_generated.deepcopy.go" ]]; then
    echo "Wrong generated deepcopy files changed after touching src file:"
    echo "  ${X:-(none)}" | tr '\n' ' '
    echo ""
    exit 1
fi
X="$(newer defaults "${STAMP}")"
if [[ -z "${X}" || ${X} != "./${DIR}/zz_generated.defaults.go" ]]; then
    echo "Wrong generated defaults files changed after touching src file:"
    echo "  ${X:-(none)}" | tr '\n' ' '
    echo ""
    exit 1
fi
X="$(newer conversion "${STAMP}")"
if [[ -z "${X}" || ${X} != "./${DIR}/zz_generated.conversion.go" ]]; then
    echo "Wrong generated conversion files changed after touching src file:"
    echo "  ${X:-(none)}" | tr '\n' ' '
    echo ""
    exit 1
fi

#
# Test when we do unrelated things in a package that needs codegen.
#

echo
echo "CASE: touch a non-go file in a package that needs codegen"
DIR=staging/src/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1
touch "${DIR}/foo"
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer deepcopy "${STAMP}")"
if [[ -n "${X}" ]]; then
    echo "Generated files changed when an unrelated file was added"
    rm -f "${DIR}/foo"
    return 1
fi

echo
echo "CASE: remove a non-go file in a package that needs codegen"
rm "${DIR}/foo"
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer deepcopy "${STAMP}")"
if [[ -n "${X}" ]]; then
    echo "Generated files changed when an unrelated file was removed"
    rm -f "${DIR}/foo"
    return 1
fi

echo
echo "CASE: touch the directory of a package that needs codegen"
touch "${DIR}"
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer deepcopy "${STAMP}")"
if [[ -n "${X}" ]]; then
    echo "Generated files changed when a dir was touched"
    rm -f "${DIR}/foo"
    return 1
fi

#
# Test when the codegen tool itself changes: deepcopy
#

echo
echo "CASE: touch a file in the codegen tool"
touch staging/src/k8s.io/code-generator/cmd/deepcopy-gen/main.go
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(older deepcopy "${STAMP}")"
if [[ -n "${X}" ]]; then
    echo "Generated deepcopy files did not change after touching code-generator file:"
    echo "  ${X}" | tr '\n' ' '
    echo ""
    exit 1
fi

echo
echo "CASE: touch a file in a dep of the codegen tool"
#FIXME: touch vendor/k8s.io/gengo/examples/deepcopy-gen/generators/deepcopy.go
touch gengo2/v2/examples/deepcopy-gen/generators/deepcopy.go
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(older deepcopy "${STAMP}")"
if [[ -n "${X}" ]]; then
    echo "Generated deepcopy files did not change after touching code-generator dep file:"
    echo "  ${X}" | tr '\n' ' '
    echo ""
    exit 1
fi

#
# Test when we touch a file in a package that needs codegen for all openapi specs.
#

echo
echo "CASE: touch a file in a package that needs openapi codegen (all)"
touch "staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go"
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer openapi "${STAMP}")"
if [[ -z "${X}" || ${X} != "./pkg/generated/openapi/zz_generated.openapi.go
./staging/src/k8s.io/apiextensions-apiserver/pkg/generated/openapi/zz_generated.openapi.go
./staging/src/k8s.io/code-generator/examples/apiserver/openapi/zz_generated.openapi.go
./staging/src/k8s.io/kube-aggregator/pkg/generated/openapi/zz_generated.openapi.go
./staging/src/k8s.io/sample-apiserver/pkg/generated/openapi/zz_generated.openapi.go" ]]; then
    echo "Wrong generated openapi files changed after touching src file:"
    echo "  ${X:-(none)}" | tr '\n' ' '
    echo ""
    exit 1
fi

#
# Test when we touch a file in a package that needs codegen for only the main openapi spec.
#

echo
echo "CASE: touch a file in a package that needs openapi codegen (main only)"
touch "staging/src/k8s.io/api/apps/v1/types.go"
touch "${STAMP}"
make generated_files 2>&1 | indent
X="$(newer openapi "${STAMP}")"
if [[ -z "${X}" || ${X} != "./pkg/generated/openapi/zz_generated.openapi.go" ]]; then
    echo "Wrong generated openapi files changed after touching src file:"
    echo "  ${X:-(none)}" | tr '\n' ' '
    echo ""
    exit 1
fi

#
# Test when we touch a file, modify the violation file it should fail, and UPDATE_API_KNOWN_VIOLATIONS=true updates it.
#

echo
echo "CASE: touch a file in a dep of the openapi tool with violations, then fix"
touch "staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/types.go"
sed -i '$d' api/api-rules/violation_exceptions.list # remove an error
sed -i '$d' api/api-rules/codegen_violation_exceptions.list #remove an error
if make generated_files 2>&1 | indent; then
    echo "Expected make generated_files to fail with API violations."
    echo ""
    exit 1
fi
touch "${STAMP}"
make generated_files UPDATE_API_KNOWN_VIOLATIONS=true 2>&1 | indent
X="$(newer openapi "${STAMP}")"
if [[ -z "${X}" || ${X} != "./pkg/generated/openapi/zz_generated.openapi.go
./staging/src/k8s.io/apiextensions-apiserver/pkg/generated/openapi/zz_generated.openapi.go
./staging/src/k8s.io/code-generator/examples/apiserver/openapi/zz_generated.openapi.go
./staging/src/k8s.io/kube-aggregator/pkg/generated/openapi/zz_generated.openapi.go
./staging/src/k8s.io/sample-apiserver/pkg/generated/openapi/zz_generated.openapi.go" ]]; then
    echo "Wrong generated openapi files changed after updating violation files:"
    echo "  ${X:-(none)}" | tr '\n' ' '
    echo ""
    exit 1
fi
for f in api/api-rules/violation_exceptions.list api/api-rules/codegen_violation_exceptions.list; do
    if ! git diff --quiet "$f"; then
        echo "Violation file \"$f\" was not updated with UPDATE_API_KNOWN_VIOLATIONS=true."
        echo ""
        exit 1
    fi
done

echo
echo "PASS"
