#!/bin/bash

set -e

mkdir -p ~/.kube

export KUBECONFIG=~/.kube/config

eval $(gp env -e | grep OKTETO_KUBECONFIG)

echo ${OKTETO_KUBECONFIG} | base64 --decode > ${KUBECONFIG}

okteto context use https://cloud.okteto.com
okteto up
