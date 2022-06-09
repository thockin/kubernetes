#!/usr/bin/env bash

# Copyright 2022 The Kubernetes Authors.
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

# This script pre-populates the local GOMODCACHE with source code from our
# "vendor" directory. That source code must be complete enough for "make all",

set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source "${KUBE_ROOT}/hack/lib/init.sh"

# This sets GOMODCACHE
kube::golang::setup_env

cd "${KUBE_ROOT}"

# See https://go.dev/ref/mod#module-cache for an explanation of the
# expected content.
if [ -d "${GOMODCACHE}" ]; then
    chmod -R u+rw "${GOMODCACHE}"
    rm -r "${GOMODCACHE}"
fi
mkdir -p "${GOMODCACHE}"

# Capital letters in module paths and versions are escaped using exclamation
# points (Azure is escaped as !azure) to avoid conflicts on case-insensitive
# file systems.
escape_path () {
    echo "$1" | sed -e 's/\([A-Z]\)/!\1/g' | tr A-Z a-z
}

grep '#.*=>' vendor/modules.txt | grep -v '=> ./' | sed -e 's/.*=> //' | sort -u | while read -r module version; do
    from="${KUBE_ROOT}/vendor/${module}"

    # Some modules are listed in modules.txt although they are not needed.
    if ! [ -d "${from}" ]; then
        continue
    fi

    # Go always wants cache/download/$module and will hang
    # when only the source is present. Therefore we need
    # to populate the download directory.
    moddir="${GOMODCACHE}/cache/download/$(escape_path "${module}/@v")"
    basename="${moddir}/$(escape_path "$version")"
    mkdir -p "${moddir}"
    echo "$version" >"${moddir}/list"
    echo "{\"Version\": \"$version\"}" >"${basename}.info"
    if [ -f "${from}/go.mod" ]; then
        # Use the original go.mod. It got placed in the vendor directory
        # by hack/update-vendor.sh. "go mod vendor" normally doesn't
        # do that, but we need the original one to match the checksum.
        cp "${from}/go.mod" "${basename}.mod"
    else
        # Make up a simple go.mod file.
        # TODO: this might not be needed and if it is needed,
        # it might not work - check this.
        echo "module ${module}" >"${basename}.mod"
    fi

    # Create a .zip file (see https://go.dev/ref/mod#zip-files) and the
    # corresponding hash. Placing the source ourselves into GOMODCACHE didn't
    # work.
    #
    # "Our" zip file content does not match the "official" content because only
    # relevant files got vendored. Therefore the comparison against the
    # checksum value recorded in the various go.sum files will fail.  However,
    # we can skip that check by creating a .ziphash file that has the right
    # checksum.
    #
    # This is a HACK! It relies on implementation details in the Go
    # 1.18 module handling which could change. It would be better if Go
    # supported building from "vendor" in workspace mode.
    #
    # Instead of vendoring, we could add the full content of each module
    # zip file to our repo. But that would make the repo larger and make
    # it harder to review dependency changes because it would not be obvious
    # anymore which of the files are actually needed.
    if grep "^${module} ${version} " "${KUBE_ROOT}/go.sum" | sed -e 's/.* //' > "${basename}.ziphash"; then
        tmp="${GOMODCACHE}/${module}@${version}"
        mkdir -p "$(dirname "${tmp}")"
        cp -r --link "${from}" "${tmp}"

        # Sub-modules must not be included.
        find "${tmp}" -name go.mod | while read -r modfile; do
            modulepath="$(dirname "${modfile}")"
            if [ "${modulepath}" != "${tmp}" ]; then
                rm -rf "${modulepath}"
            fi
        done

        # We could skip compression here via `--suffixes .go`, but
        # it didn't save much time (12s vs 14s on a fast machine) and
        # required more space (101M vs. 29M).
        #
        # TODO: can we rely on zip being installed?
        (cd ${GOMODCACHE} && zip --quiet --recurse-paths "${basename}.zip" "${module}@${version}")
        rm -rf "${tmp}"
    else
        # For some modules we just have the go.mod file, but not the actual source.
        # We don't need the zip archive for those.
        rm "${basename}.ziphash"
    fi
done
