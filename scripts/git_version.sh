#!/bin/bash

GIT_TAG=$(awk '/src\/k8s.io\/client-go/ {print $NF}' go.mod | awk -F- '{print $1}')+${OEM:-vip}
TREE_STATE=clean
COMMIT=$(git rev-parse HEAD)
DIRTY=""

if [ -d .git ]; then
    if [ -z "$GIT_TAG" ]; then
        GIT_TAG=$(git tag -l --contains HEAD | head -n 1)
    fi
    if [ -n "$(git status --porcelain --untracked-files=no)" ]; then
        DIRTY="-dirty"
        TREE_STATE=dirty
    fi

    COMMIT=$(git log -n3 --pretty=format:"%H %ae" | grep -v ' drone@localhost$' | cut -f1 -d\  | head -1)
    if [ -z "${COMMIT}" ]; then
    COMMIT=$(git rev-parse HEAD || true)
    fi
fi

export GIT_TAG
export TREE_STATE
export COMMIT
export DIRTY
