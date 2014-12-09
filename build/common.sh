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

# Common utilities, variables and checks for all build scripts.
set -o errexit
set -o nounset
set -o pipefail

KUBE_ROOT=$(dirname "${BASH_SOURCE}")/..
cd "${KUBE_ROOT}"

# This'll canonicalize the path
KUBE_ROOT=$PWD

source hack/lib/init.sh

# Incoming options
#
readonly KUBE_SKIP_CONFIRMATIONS="${KUBE_SKIP_CONFIRMATIONS:-n}"
readonly KUBE_BUILD_RUN_IMAGES="${KUBE_BUILD_RUN_IMAGES:-n}"
readonly KUBE_GCS_UPLOAD_RELEASE="${KUBE_GCS_UPLOAD_RELEASE:-n}"
readonly KUBE_GCS_NO_CACHING="${KUBE_GCS_NO_CACHING:-y}"
readonly KUBE_GCS_MAKE_PUBLIC="${KUBE_GCS_MAKE_PUBLIC:-y}"
# KUBE_GCS_RELEASE_BUCKET default: kubernetes-releases-${project_hash}
readonly KUBE_GCS_RELEASE_PREFIX=${KUBE_GCS_RELEASE_PREFIX-devel}/
readonly KUBE_GCS_DOCKER_REG_PREFIX=${KUBE_GCS_DOCKER_REG_PREFIX-docker-reg}/


# Constants
readonly KUBE_BUILD_IMAGE_REPO=kube-build
# These get set in verify_prereqs with a unique hash based on KUBE_ROOT
# KUBE_BUILD_IMAGE_TAG=<hash>
# KUBE_BUILD_IMAGE="${KUBE_BUILD_IMAGE_REPO}:${KUBE_BUILD_IMAGE_TAG}"
# KUBE_BUILD_CONTAINER_NAME=kube-build-<hash>
readonly KUBE_BUILD_IMAGE_CROSS_TAG=cross
readonly KUBE_BUILD_IMAGE_CROSS="${KUBE_BUILD_IMAGE_REPO}:${KUBE_BUILD_IMAGE_CROSS_TAG}"
readonly KUBE_BUILD_GOLANG_VERSION=1.3
# KUBE_BUILD_DATA_CONTAINER_NAME=kube-build-data-<hash>

# Here we map the output directories across both the local and remote _output
# directories:
#
# *_OUTPUT_ROOT    - the base of all output in that environment.
# *_OUTPUT_SUBPATH - location where golang stuff is built/cached.  Also
#                    persisted across docker runs with a volume mount.
# *_OUTPUT_BINPATH - location where final binaries are placed.  If the remote
#                    is really remote, this is the stuff that has to be copied
#                    back.
readonly LOCAL_OUTPUT_ROOT="${KUBE_ROOT}/_output"
readonly LOCAL_OUTPUT_SUBPATH="${LOCAL_OUTPUT_ROOT}/dockerized"
readonly LOCAL_OUTPUT_BINPATH="${LOCAL_OUTPUT_SUBPATH}/bin"
readonly LOCAL_OUTPUT_IMAGE_STAGING="${LOCAL_OUTPUT_ROOT}/images"

readonly REMOTE_OUTPUT_ROOT="/go/src/${KUBE_GO_PACKAGE}/_output"
readonly REMOTE_OUTPUT_SUBPATH="${REMOTE_OUTPUT_ROOT}/dockerized"
readonly REMOTE_OUTPUT_BINPATH="${REMOTE_OUTPUT_SUBPATH}/bin"

readonly DOCKER_MOUNT_ARGS_BASE=(--volume "${LOCAL_OUTPUT_BINPATH}:${REMOTE_OUTPUT_BINPATH}")
# DOCKER_MOUNT_ARGS=("${DOCKER_MOUNT_ARGS_BASE[@]}" --volumes-from "${KUBE_BUILD_DATA_CONTAINER_NAME}")

# We create a Docker data container to cache incremental build artifacts.  We
# need to cache both the go tree in _output and the go tree under Godeps.
readonly REMOTE_OUTPUT_GOPATH="${REMOTE_OUTPUT_SUBPATH}/go"
readonly REMOTE_GODEP_GOPATH="/go/src/${KUBE_GO_PACKAGE}/Godeps/_workspace/pkg"
readonly DOCKER_DATA_MOUNT_ARGS=(
  --volume "${REMOTE_OUTPUT_GOPATH}"
  --volume "${REMOTE_GODEP_GOPATH}"
)


readonly KUBE_RUN_IMAGE_BASE="kubernetes"
readonly KUBE_RUN_IMAGES=(
  kube-apiserver
  kube-controller-manager
  kube-proxy
  kube-scheduler
  kubelet
  bootstrap
)


# This is where the final release artifacts are created locally
readonly RELEASE_STAGE="${LOCAL_OUTPUT_ROOT}/release-stage"
readonly RELEASE_DIR="${LOCAL_OUTPUT_ROOT}/release-tars"

# ---------------------------------------------------------------------------
# Basic setup functions

# Verify that the right utilities and such are installed for building Kube.  Set
# up some dynamic constants.
#
# Vars set:
#   KUBE_ROOT_HASH
#   KUBE_BUILD_IMAGE_TAG
#   KUBE_BUILD_IMAGE
#   KUBE_BUILD_CONTAINER_NAME
#   KUBE_BUILD_DATA_CONTAINER_NAME
#   DOCKER_MOUNT_ARGS
function kube::build::verify_prereqs() {
  if [[ -z "$(which docker)" ]]; then
    echo "Can't find 'docker' in PATH, please fix and retry." >&2
    echo "See https://docs.docker.com/installation/#installation for installation instructions." >&2
    exit 1
  fi

  if kube::build::is_osx; then
    if [[ -z "$(which boot2docker)" ]]; then
      echo "It looks like you are running on Mac OS X and boot2docker can't be found." >&2
      echo "See: https://docs.docker.com/installation/mac/" >&2
      exit 1
    fi
    if [[ $(boot2docker status) != "running" ]]; then
      echo "boot2docker VM isn't started.  Please run 'boot2docker start'" >&2
      exit 1
    else
      # Reach over and set the clock. After sleep/resume the clock will skew.
      echo "+++ Setting boot2docker clock"
      boot2docker ssh sudo date -u -D "%Y%m%d%H%M.%S" --set "$(date -u +%Y%m%d%H%M.%S)" >/dev/null
    fi
  fi

  if ! docker info > /dev/null 2>&1 ; then
    {
      echo "Can't connect to 'docker' daemon.  please fix and retry."
      echo
      echo "Possible causes:"
      echo "  - On Mac OS X, boot2docker VM isn't installed or started"
      echo "  - On Mac OS X, docker env variable isn't set approriately. Run:"
      echo "      \$(boot2docker shellinit)"
      echo "  - On Linux, user isn't in 'docker' group.  Add and relogin."
      echo "    - Something like 'sudo usermod -a -G docker ${USER-user}'"
      echo "    - RHEL7 bug and workaround: https://bugzilla.redhat.com/show_bug.cgi?id=1119282#c8"
      echo "  - On Linux, Docker daemon hasn't been started or has crashed"
    } >&2
    exit 1
  fi

  KUBE_ROOT_HASH=$(kube::build::short_hash "$KUBE_ROOT")
  KUBE_BUILD_IMAGE_TAG="build-${KUBE_ROOT_HASH}"
  KUBE_BUILD_IMAGE="${KUBE_BUILD_IMAGE_REPO}:${KUBE_BUILD_IMAGE_TAG}"
  KUBE_BUILD_CONTAINER_NAME="kube-build-${KUBE_ROOT_HASH}"
  KUBE_BUILD_DATA_CONTAINER_NAME="kube-build-data-${KUBE_ROOT_HASH}"
  DOCKER_MOUNT_ARGS=("${DOCKER_MOUNT_ARGS_BASE[@]}" --volumes-from "${KUBE_BUILD_DATA_CONTAINER_NAME}")
}

# ---------------------------------------------------------------------------
# Utility functions

function kube::build::is_osx() {
  [[ "$(uname)" == "Darwin" ]]
}

function kube::build::clean_output() {
  # Clean out the output directory if it exists.
  if kube::build::build_image_built ; then
    echo "+++ Cleaning out _output/dockerized/bin/ via docker build image"
    kube::build::run_build_command bash -c "rm -rf '${REMOTE_OUTPUT_BINPATH}'/*"
  else
    echo "!!! Build image not built.  Cannot clean via docker build image."
  fi

  echo "+++ Removing data container"
  docker rm -v "${KUBE_BUILD_DATA_CONTAINER_NAME}" >/dev/null 2>&1 || true

  echo "+++ Cleaning out local _output directory"
  rm -rf "${LOCAL_OUTPUT_ROOT}"
}

# Make sure the _output directory is created and mountable by docker
function kube::build::prepare_output() {
  mkdir -p "${LOCAL_OUTPUT_SUBPATH}"

  # On RHEL/Fedora SELinux is enabled by default and currently breaks docker
  # volume mounts.  We can work around this by explicitly adding a security
  # context to the _output directory.
  # Details: https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Resource_Management_and_Linux_Containers_Guide/sec-Sharing_Data_Across_Containers.html#sec-Mounting_a_Host_Directory_to_a_Container
  if which selinuxenabled &>/dev/null && \
      selinuxenabled && \
      which chcon >/dev/null ; then
    if [[ ! $(ls -Zd "${LOCAL_OUTPUT_ROOT}") =~ svirt_sandbox_file_t ]] ; then
      echo "+++ Applying SELinux policy to '_output' directory."
      if ! chcon -Rt svirt_sandbox_file_t "${LOCAL_OUTPUT_ROOT}"; then
        echo "    ***Failed***.  This may be because you have root owned files under _output."
        echo "    Continuing, but this build may fail later if SELinux prevents access."
      fi
    fi
  fi

}

# Detect if a specific image exists
#
# $1 - image repo name
# #2 - image tag
function kube::build::docker_image_exists() {
  [[ -n $1 && -n $2 ]] || {
    echo "!!! Internal error. Image not specified in docker_image_exists." >&2
    exit 2
  }

  # We cannot just specify the IMAGE here as `docker images` doesn't behave as
  # expected.  See: https://github.com/docker/docker/issues/8048
  docker images | grep -Eq "^${1}\s+${2}\s+"
}

# Takes $1 and computes a short has for it. Useful for unique tag generation
function kube::build::short_hash() {
  [[ $# -eq 1 ]] || {
    echo "!!! Internal error.  No data based to short_hash." >&2
    exit 2
  }

  local short_hash
  if which md5 >/dev/null 2>&1; then
    short_hash=$(md5 -q -s "$1")
  else
    short_hash=$(echo -n "$1" | md5sum)
  fi
  echo ${short_hash:0:5}
}

# ---------------------------------------------------------------------------
# Building

function kube::build::build_image_built() {
  kube::build::docker_image_exists "${KUBE_BUILD_IMAGE_REPO}" "${KUBE_BUILD_IMAGE_TAG}"
}

function kube::build::ensure_golang() {
  kube::build::docker_image_exists golang 1.3 || {
    [[ ${KUBE_SKIP_CONFIRMATIONS} =~ ^[yY]$ ]] || {
      echo "You don't have a local copy of the golang docker image. This image is 450MB."
      read -p "Download it now? [y/n] " -r
      echo
      [[ $REPLY =~ ^[yY]$ ]] || {
        echo "Aborting." >&2
        exit 1
      }
    }

    echo "+++ Pulling docker image: golang:${KUBE_BUILD_GOLANG_VERSION}"
    docker pull golang:${KUBE_BUILD_GOLANG_VERSION}
  }
}

# Set up the context directory for the kube-build image and build it.
function kube::build::build_image() {
  local -r build_context_dir="${LOCAL_OUTPUT_IMAGE_STAGING}/${KUBE_BUILD_IMAGE}"
  local -r source=(
    api
    build
    cmd
    examples
    Godeps/_workspace/src
    Godeps/Godeps.json
    hack
    LICENSE
    pkg
    plugin
    README.md
    test
    third_party
  )

  kube::build::build_image_cross

  mkdir -p "${build_context_dir}"
  tar czf "${build_context_dir}/kube-source.tar.gz" "${source[@]}"

  kube::version::get_version_vars
  kube::version::save_version_vars "${build_context_dir}/kube-version-defs"

  cp build/build-image/Dockerfile ${build_context_dir}/Dockerfile
  kube::build::docker_build "${KUBE_BUILD_IMAGE}" "${build_context_dir}"
}

# Build the kubernetes golang cross base image.
function kube::build::build_image_cross() {
  kube::build::ensure_golang

  local -r build_context_dir="${LOCAL_OUTPUT_ROOT}/images/${KUBE_BUILD_IMAGE}/cross"
  mkdir -p "${build_context_dir}"
  cp build/build-image/cross/Dockerfile ${build_context_dir}/Dockerfile
  kube::build::docker_build "${KUBE_BUILD_IMAGE_CROSS}" "${build_context_dir}"
}

# Builds the runtime image.  Assumes that the appropriate binaries are already
# built and in $LOCAL_OUTPUT_BINPATH.
function kube::build::run_image() {
  [[ ${KUBE_BUILD_RUN_IMAGES} =~ ^[yY]$ ]] || return 0

  local -r build_context_base="${LOCAL_OUTPUT_IMAGE_STAGING}/${KUBE_RUN_IMAGE_BASE}"

  # First build the base image.  This one brings in all of the binaries.
  mkdir -p "${build_context_base}"
  tar czf "${build_context_base}/kube-bins.tar.gz" \
    -C "${LOCAL_OUTPUT_BINPATH}/linux/amd64" \
    "${KUBE_RUN_IMAGES[@]}"
  cp -R build/run-images/base/* "${build_context_base}/"
  kube::build::docker_build "${KUBE_RUN_IMAGE_BASE}" "${build_context_base}"

  local b
  for b in "${KUBE_RUN_IMAGES[@]}" ; do
    local sub_context_dir="${build_context_base}-$b"
    mkdir -p "${sub_context_dir}"
    cp -R build/run-images/$b/* "${sub_context_dir}/"
    kube::build::docker_build "${KUBE_RUN_IMAGE_BASE}-$b" "${sub_context_dir}"
  done
}

# Build a docker image from a Dockerfile.
# $1 is the name of the image to build
# $2 is the location of the "context" directory, with the Dockerfile at the root.
function kube::build::docker_build() {
  local -r image=$1
  local -r context_dir=$2
  local -ra build_cmd=(docker build -t "${image}" "${context_dir}")

  echo "+++ Building Docker image ${image}."
  local docker_output
  docker_output=$("${build_cmd[@]}" 2>&1) || {
    cat <<EOF >&2
+++ Docker build command failed for ${image}

${docker_output}

To retry manually, run:

${build_cmd[*]}

EOF
    return 1
  }
}

function kube::build::clean_image() {
  local -r image=$1

  echo "+++ Deleting docker image ${image}"
  docker rmi ${image} 2> /dev/null || true
}

function kube::build::clean_images() {
  kube::build::clean_image "${KUBE_BUILD_IMAGE}"

  kube::build::clean_image "${KUBE_RUN_IMAGE_BASE}"

  local b
  for b in "${KUBE_RUN_IMAGES[@]}" ; do
    kube::build::clean_image "${KUBE_RUN_IMAGE_BASE}-${b}"
  done

  echo "+++ Cleaning all other untagged docker images"
  docker rmi $(docker images -q --filter 'dangling=true') 2> /dev/null || true
}

function kube::build::ensure_data_container() {
  if ! docker inspect "${KUBE_BUILD_DATA_CONTAINER_NAME}" >/dev/null 2>&1; then
    echo "+++ Creating data container"
    local -ra docker_cmd=(
      docker run
      "${DOCKER_DATA_MOUNT_ARGS[@]}"
      --name "${KUBE_BUILD_DATA_CONTAINER_NAME}"
      "${KUBE_BUILD_IMAGE}"
      true
    )
    "${docker_cmd[@]}"
  fi
}

# Run a command in the kube-build image.  This assumes that the image has
# already been built.  This will sync out all output data from the build.
function kube::build::run_build_command() {
  [[ $# != 0 ]] || { echo "Invalid input." >&2; return 4; }

  kube::build::ensure_data_container
  kube::build::prepare_output

  local -a docker_run_opts=(
    "--name=${KUBE_BUILD_CONTAINER_NAME}"
     "${DOCKER_MOUNT_ARGS[@]}"
    )

  # If we have stdin we can run interactive.  This allows things like 'shell.sh'
  # to work.  However, if we run this way and don't have stdin, then it ends up
  # running in a daemon-ish mode.  So if we don't have a stdin, we explicitly
  # attach stderr/stdout but don't bother asking for a tty.
  if [[ -t 0 ]]; then
    docker_run_opts+=(--interactive --tty)
  else
    docker_run_opts+=(--attach=stdout --attach=stderr)
  fi

  local -ra docker_cmd=(
    docker run "${docker_run_opts[@]}" "${KUBE_BUILD_IMAGE}")

  # Remove the container if it is left over from some previous aborted run
  docker rm -v "${KUBE_BUILD_CONTAINER_NAME}" >/dev/null 2>&1 || true
  "${docker_cmd[@]}" "$@"

  # Remove the container after we run.  '--rm' might be appropriate but it
  # appears that sometimes it fails. See
  # https://github.com/docker/docker/issues/3968
  docker rm -v "${KUBE_BUILD_CONTAINER_NAME}" >/dev/null 2>&1 || true
}

# Test if the output directory is remote (and can only be accessed through
# docker) or if it is "local" and we can access the output without going through
# docker.
function kube::build::is_output_remote() {
  rm -f "${LOCAL_OUTPUT_SUBPATH}/test_for_remote"
  kube::build::run_build_command touch "${REMOTE_OUTPUT_BINPATH}/test_for_remote"

  [[ ! -e "${LOCAL_OUTPUT_BINPATH}/test_for_remote" ]]
}

# If the Docker server is remote, copy the results back out.
function kube::build::copy_output() {
  if kube::build::is_output_remote; then
    # When we are on the Mac with boot2docker (or to a remote Docker in any
    # other situation) we need to copy the results back out.  Ideally we would
    # leave the container around and use 'docker cp' to copy the results out.
    # However, that doesn't work for mounted volumes currently
    # (https://github.com/dotcloud/docker/issues/1992).  And it is just plain
    # broken (https://github.com/dotcloud/docker/issues/6483).
    #
    # The easiest thing I (jbeda) could figure out was to launch another
    # container pointed at the same volume, tar the output directory and ship
    # that tar over stdout.

    echo "+++ Syncing back _output/dockerized/bin directory from remote Docker"
    rm -rf "${LOCAL_OUTPUT_BINPATH}"
    mkdir -p "${LOCAL_OUTPUT_BINPATH}"

    # The '</dev/null' here makes us run docker in a "non-interactive" mode. Not
    # doing this corrupts the output stream.
    kube::build::run_build_command sh -c "tar c -C ${REMOTE_OUTPUT_BINPATH} . ; sleep 1" </dev/null \
      | tar xv -C "${LOCAL_OUTPUT_BINPATH}"

    # I (jbeda) also tried getting rsync working using 'docker run' as the
    # 'remote shell'.  This mostly worked but there was a hang when
    # closing/finishing things off. Ug.
    #
    # local DOCKER="docker run -i --rm --name=${KUBE_BUILD_CONTAINER_NAME} ${DOCKER_MOUNT} ${KUBE_BUILD_IMAGE}"
    # DOCKER+=" bash -c 'shift ; exec \"\$@\"' --"
    # rsync --blocking-io -av -e "${DOCKER}" foo:${REMOTE_OUTPUT_BINPATH}/ ${LOCAL_OUTPUT_BINPATH}
  else
    echo "+++ Output directory is local.  No need to copy results out."
  fi
}

# ---------------------------------------------------------------------------
# Build final release artifacts
function kube::release::package_tarballs() {
  # Clean out any old releases
  rm -rf "${RELEASE_DIR}"
  mkdir -p "${RELEASE_DIR}"

  kube::release::package_client_tarballs
  kube::release::package_server_tarballs
  kube::release::package_salt_tarball
  kube::release::package_full_tarball
}

# Package up all of the cross compiled clients.  Over time this should grow into
# a full SDK
function kube::release::package_client_tarballs() {
   # Find all of the built kubecfg binaries
  local platform platforms
  platforms=($(cd "${LOCAL_OUTPUT_BINPATH}" ; echo */*))
  for platform in "${platforms[@]}" ; do
    local platform_tag=${platform/\//-} # Replace a "/" for a "-"
    echo "+++ Building tarball: client $platform_tag"

    local release_stage="${RELEASE_STAGE}/client/${platform_tag}/kubernetes"
    rm -rf "${release_stage}"
    mkdir -p "${release_stage}/client/bin"

    local client_bins=("${KUBE_CLIENT_BINARIES[@]}")
    if [[ "${platform%/*}" == "windows" ]]; then
      client_bins=("${KUBE_CLIENT_BINARIES_WIN[@]}")
    fi

    local bin
    for bin in "${client_bins[@]}"; do
      cp "${LOCAL_OUTPUT_BINPATH}/${platform}/${bin}" \
        "${release_stage}/client/bin/"
    done

    local package_name="${RELEASE_DIR}/kubernetes-client-${platform_tag}.tar.gz"
    kube::release::create_tarball "${package_name}" "${release_stage}/.."
  done
}

# Package up all of the server binaries
function kube::release::package_server_tarballs() {
  local platform
  for platform in "${KUBE_SERVER_PLATFORMS[@]}" ; do
    local platform_tag=${platform/\//-} # Replace a "/" for a "-"
    echo "+++ Building tarball: server $platform_tag"

    local release_stage="${RELEASE_STAGE}/server/${platform_tag}/kubernetes"
    rm -rf "${release_stage}"
    mkdir -p "${release_stage}/server/bin"

    # This fancy expression will expand to prepend a path
    # (${LOCAL_OUTPUT_BINPATH}/${platform}/) to every item in the
    # KUBE_SERVER_BINARIES array.
    cp "${KUBE_SERVER_BINARIES[@]/#/${LOCAL_OUTPUT_BINPATH}/${platform}/}" \
      "${release_stage}/server/bin/"

    # Include the client binaries here too as they are useful debugging tools.
    local client_bins=("${KUBE_CLIENT_BINARIES[@]}")
    if [[ "${platform%/*}" == "windows" ]]; then
      client_bins=("${KUBE_CLIENT_BINARIES_WIN[@]}")
    fi
    cp "${client_bins[@]/#/${LOCAL_OUTPUT_BINPATH}/${platform}/}" \
      "${release_stage}/server/bin/"

    local package_name="${RELEASE_DIR}/kubernetes-server-${platform_tag}.tar.gz"
    kube::release::create_tarball "${package_name}" "${release_stage}/.."
  done
}

# Package up the salt configuration tree.  This is an optional helper to getting
# a cluster up and running.
function kube::release::package_salt_tarball() {
  echo "+++ Building tarball: salt"

  local release_stage="${RELEASE_STAGE}/salt/kubernetes"
  rm -rf "${release_stage}"
  mkdir -p "${release_stage}"

  cp -R "${KUBE_ROOT}/cluster/saltbase" "${release_stage}/"

  local package_name="${RELEASE_DIR}/kubernetes-salt.tar.gz"
  kube::release::create_tarball "${package_name}" "${release_stage}/.."
}

# This is all the stuff you need to run/install kubernetes.  This includes:
#   - precompiled binaries for client
#   - Cluster spin up/down scripts and configs for various cloud providers
#   - tarballs for server binary and salt configs that are ready to be uploaded
#     to master by whatever means appropriate.
function kube::release::package_full_tarball() {
  echo "+++ Building tarball: full"

  local release_stage="${RELEASE_STAGE}/full/kubernetes"
  rm -rf "${release_stage}"
  mkdir -p "${release_stage}"

  # Copy all of the client binaries in here, but not test or server binaries.
  # The server binaries are included with the server binary tarball.
  local platform
  for platform in "${KUBE_CLIENT_PLATFORMS[@]}"; do
    local client_bins=("${KUBE_CLIENT_BINARIES[@]}")
    if [[ "${platform%/*}" == "windows" ]]; then
      client_bins=("${KUBE_CLIENT_BINARIES_WIN[@]}")
    fi
    mkdir -p "${release_stage}/platforms/${platform}"
    cp "${client_bins[@]/#/${LOCAL_OUTPUT_BINPATH}/${platform}/}" \
      "${release_stage}/platforms/${platform}"
  done

  # We want everything in /cluster except saltbase.  That is only needed on the
  # server.
  cp -R "${KUBE_ROOT}/cluster" "${release_stage}/"
  rm -rf "${release_stage}/cluster/saltbase"

  mkdir -p "${release_stage}/server"
  cp "${RELEASE_DIR}/kubernetes-salt.tar.gz" "${release_stage}/server/"
  cp "${RELEASE_DIR}"/kubernetes-server-*.tar.gz "${release_stage}/server/"

  mkdir -p "${release_stage}/third_party"
  cp -R "${KUBE_ROOT}/third_party/htpasswd" "${release_stage}/third_party/htpasswd"

  cp -R "${KUBE_ROOT}/examples" "${release_stage}/"
  cp "${KUBE_ROOT}/README.md" "${release_stage}/"
  cp "${KUBE_ROOT}/LICENSE" "${release_stage}/"
  cp "${KUBE_ROOT}/Vagrantfile" "${release_stage}/"

  local package_name="${RELEASE_DIR}/kubernetes.tar.gz"
  kube::release::create_tarball "${package_name}" "${release_stage}/.."
}

# Build a release tarball.  $1 is the output tar name.  $2 is the base directory
# of the files to be packaged.  This assumes that ${2}/kubernetes is what is
# being packaged.
function kube::release::create_tarball() {
  local tarfile=$1
  local stagingdir=$2

  # Find gnu tar if it is available
  local tar=tar
  if which gtar &>/dev/null; then
    tar=gtar
  fi

  local tar_cmd=("$tar" "czf" "${tarfile}" "-C" "${stagingdir}" "kubernetes")
  if "$tar" --version | grep -q GNU; then
    tar_cmd=("${tar_cmd[@]}" "--owner=0" "--group=0")
  else
    echo "  !!! GNU tar not available.  User names will be embedded in output and"
    echo "      release tars are not official. Build on Linux or install GNU tar"
    echo "      on Mac OS X (brew install gnu-tar)"
  fi

  "${tar_cmd[@]}"
}

# ---------------------------------------------------------------------------
# GCS Release

function kube::release::gcs::release() {
  [[ ${KUBE_GCS_UPLOAD_RELEASE} =~ ^[yY]$ ]] || return 0

  kube::release::gcs::verify_prereqs
  kube::release::gcs::ensure_release_bucket
  kube::release::gcs::push_images
  kube::release::gcs::copy_release_artifacts
}

# Verify things are set up for uploading to GCS
function kube::release::gcs::verify_prereqs() {
  if [[ -z "$(which gsutil)" || -z "$(which gcloud)" ]]; then
    echo "Releasing Kubernetes requires gsutil and gcloud.  Please download,"
    echo "install and authorize through the Google Cloud SDK: "
    echo
    echo "  https://developers.google.com/cloud/sdk/"
    return 1
  fi

  if [[ -z "${GCLOUD_ACCOUNT-}" ]]; then
    GCLOUD_ACCOUNT=$(gcloud auth list 2>/dev/null | awk '/(active)/ { print $2 }')
  fi
  if [[ -z "${GCLOUD_ACCOUNT-}" ]]; then
    echo "No account authorized through gcloud.  Please fix with:"
    echo
    echo "  gcloud auth login"
    return 1
  fi

  if [[ -z "${GCLOUD_PROJECT-}" ]]; then
    GCLOUD_PROJECT=$(gcloud config list project | awk '{project = $3} END {print project}')
  fi
  if [[ -z "${GCLOUD_PROJECT-}" ]]; then
    echo "No account authorized through gcloud.  Please fix with:"
    echo
    echo "  gcloud config set project <project id>"
    return 1
  fi
}

# Create a unique bucket name for releasing Kube and make sure it exists.
function kube::release::gcs::ensure_release_bucket() {
  local project_hash
  project_hash=$(kube::build::short_hash "$GCLOUD_PROJECT")
  KUBE_GCS_RELEASE_BUCKET=${KUBE_GCS_RELEASE_BUCKET-kubernetes-releases-${project_hash}}

  if ! gsutil ls "gs://${KUBE_GCS_RELEASE_BUCKET}" >/dev/null 2>&1 ; then
    echo "Creating Google Cloud Storage bucket: $KUBE_GCS_RELEASE_BUCKET"
    gsutil mb -p "${GCLOUD_PROJECT}" "gs://${KUBE_GCS_RELEASE_BUCKET}"
  fi
}

function kube::release::gcs::ensure_docker_registry() {
  local -r reg_container_name="gcs-registry"

  local -r running=$(docker inspect ${reg_container_name} 2>/dev/null \
    | build/json-extractor.py 0.State.Running 2>/dev/null)

  [[ "$running" != "true" ]] || return 0

  # Grovel around and find the OAuth token in the gcloud config
  local -r boto=~/.config/gcloud/legacy_credentials/${GCLOUD_ACCOUNT}/.boto
  local refresh_token
  refresh_token=$(grep 'gs_oauth2_refresh_token =' "$boto" | awk '{ print $3 }')

  if [[ -z "$refresh_token" ]]; then
    echo "Couldn't find OAuth 2 refresh token in ${boto}" >&2
    return 1
  fi

  # If we have an old one sitting around, remove it
  docker rm ${reg_container_name} >/dev/null 2>&1 || true

  echo "+++ Starting GCS backed Docker registry"
  local -ra docker_cmd=(
    docker run -d "--name=${reg_container_name}"
    -e "GCS_BUCKET=${KUBE_GCS_RELEASE_BUCKET}"
    -e "STORAGE_PATH=${KUBE_GCS_DOCKER_REG_PREFIX}"
    -e "GCP_OAUTH2_REFRESH_TOKEN=${refresh_token}"
    -p 127.0.0.1:5000:5000
    google/docker-registry
  )

  "${docker[@]}"

  # Give it time to spin up before we start throwing stuff at it
  sleep 5
}

function kube::release::gcs::push_images() {
  [[ ${KUBE_BUILD_RUN_IMAGES} =~ ^[yY]$ ]] || return 0

  kube::release::gcs::ensure_docker_registry

  # Tag each of our run binaries with the right registry and push
  local b image_name
  for b in "${KUBE_RUN_IMAGES[@]}" ; do
    image_name="${KUBE_RUN_IMAGE_BASE}-${b}"
    echo "+++ Tagging and pushing ${image_name} to GCS bucket ${KUBE_GCS_RELEASE_BUCKET}"
    docker tag "${KUBE_RUN_IMAGE_BASE}-$b" "localhost:5000/${image_name}"
    docker push "localhost:5000/${image_name}"
    docker rmi "localhost:5000/${image_name}"
  done
}

function kube::release::gcs::copy_release_artifacts() {
  # TODO: This isn't atomic.  There will be points in time where there will be
  # no active release.  Also, if something fails, the release could be half-
  # copied.  The real way to do this would perhaps to have some sort of release
  # version so that we are never overwriting a destination.
  local -r gcs_destination="gs://${KUBE_GCS_RELEASE_BUCKET}/${KUBE_GCS_RELEASE_PREFIX}"
  local gcs_options=()

  if [[ ${KUBE_GCS_NO_CACHING} =~ ^[yY]$ ]]; then
    gcs_options=("-h" "Cache-Control:private, max-age=0")
  fi

  echo "+++ Copying release artifacts to ${gcs_destination}"

  # First delete all objects at the destination
  if gsutil ls "${gcs_destination}" >/dev/null 2>&1; then
    echo "!!! ${gcs_destination} not empty."
    read -p "Delete everything under ${gcs_destination}? [y/n] " -r || {
      echo "EOF on prompt.  Skipping upload"
      return
    }
    [[ $REPLY =~ ^[yY]$ ]] || {
      echo "Skipping upload"
      return
    }
    gsutil -m rm -f -R "${gcs_destination}"
  fi

  # Now upload everything in release directory
  gsutil -m "${gcs_options[@]+${gcs_options[@]}}" cp -r "${RELEASE_DIR}"/* "${gcs_destination}"

  # Having the "template" scripts from the GCE cluster deploy hosted with the
  # release is useful for GKE.  Copy everything from that directory up also.
  gsutil -m "${gcs_options[@]+${gcs_options[@]}}" cp \
    "${RELEASE_STAGE}/full/kubernetes/cluster/gce/templates/*.sh" \
    "${gcs_destination}extra/gce-templates/"

  # Upload the "naked" binaries to GCS.  This is useful for install scripts that
  # download the binaries directly and don't need tars.
  local platform platforms
  platforms=($(cd "${RELEASE_STAGE}/client" ; echo *))
  for platform in "${platforms[@]}"; do
    local src="${RELEASE_STAGE}/client/${platform}/kubernetes/client/bin/*"
    local dst="${gcs_destination}bin/${platform/-//}/"
    # We assume here the "server package" is a superset of the "client package"
    if [[ -d "${RELEASE_STAGE}/server/${platform}" ]]; then
      src="${RELEASE_STAGE}/server/${platform}/kubernetes/server/bin/*"
    fi
    gsutil -m "${gcs_options[@]+${gcs_options[@]}}" cp \
      "$src" "$dst"
  done

  # TODO(jbeda): Generate an HTML page with links for this release so it is easy
  # to see it.  For extra credit, generate a dynamic page that builds up the
  # release list using the GCS JSON API.  Use Angular and Bootstrap for extra
  # extra credit.

  if [[ ${KUBE_GCS_MAKE_PUBLIC} =~ ^[yY]$ ]]; then
    echo "+++ Marking all uploaded objects public"
    gsutil acl ch -R -g all:R "${gcs_destination}" >/dev/null 2>&1
  fi

  gsutil ls -lhr "${gcs_destination}"
}
