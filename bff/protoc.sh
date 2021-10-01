#!/bin/bash

SHELL_NAME=$(basename $0)

for OPT in $@; do
    case $OPT in
        --out)
            if [[ ! -z "$2" ]] && [[ ! "$2" =~ ^-+ ]]; then
                unset FLAG_SRC
                OUT=$2
                shift 2
            fi
            ;;
        --src)
            FLAG_SRC=1
            shift
            ;;
        -*)
            echo "$SHELL_NAME: unknown option -- '$(echo $1 | sed 's/^-*//')'" 1>&2
            exit 1
            ;;
        *)
            if [[ ! -z "$1" ]] && [[ ! "$1" =~ ^-+ ]] && [[ ! -z "$FLAG_SRC" ]]; then
                PROTO_SRC_LIST+=( "$1" )
                shift
            fi
            ;;
    esac
done

if [ -z "$OUT" ] || [ -z "$PROTO_SRC_LIST" ]; then
    if [ -z "$OUT" ]; then
        echo "$SHELL_NAME: option requires an argument -- --out" 1>&2
    fi
    if [ -z "$PROTO_SRC_LIST" ]; then
        echo "$SHELL_NAME: option requires at least one argument -- --src" 1>&2
    fi
    exit 1
fi

export PATH="$PATH:$(npm bin)"
mkdir -p ${OUT}

for PROTO_SRC in "${PROTO_SRC_LIST[@]}"; do
    PROTO_FILES="${PROTO_SRC}/*.proto"
    IMPORT_OPTS="-I node_modules/google-proto-files -I ${PROTO_SRC}"

    grpc_tools_node_protoc \
      --js_out=import_style=commonjs,binary:${OUT} \
      --grpc_out=${OUT} \
      --plugin=protoc-gen-grpc=$(which grpc_tools_node_protoc_plugin) \
      ${IMPORT_OPTS} \
      ${PROTO_FILES}

    grpc_tools_node_protoc \
      --plugin=protoc-gen-ts=$(npm bin)/protoc-gen-ts \
      --ts_out=${OUT} \
      ${IMPORT_OPTS} \
      ${PROTO_FILES}
done
