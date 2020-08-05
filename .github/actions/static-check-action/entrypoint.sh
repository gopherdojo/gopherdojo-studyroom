#!/bin/sh

set -e

# -------------
# Environments
# -------------

RUN=$1
WORKING_DIR=$2
SEND_COMMENT=$3
GITHUB_TOKEN=$4
COMMENT=""
SUCCESS=0

# if not set, assign default value
if [ "$2" = "" ]; then
    WORKING_DIR="nn"
fi
if [ "$3" = "" ]; then
    SEND_COMMENT="true"
fi

cd ${WORKING_DIR}
PKGNAME=$(go list ./...)

# ------------
# Functions
# ------------

# send_comment sends ${COMMENT} to pull request
# this function uses ${GITHUB_TOKEN}, ${COMMENT} and ${GITHUB_EVENT_PATH}
send_comment() {
    PAYLOAD=$(echo '{}' | jq --arg body "${COMMENT}" '.body = $body')
    COMMENTS_URL=$(cat ${GITHUB_EVENT_PATH} | jq -r .pull_request.comments_url)
    curl -s -S -H "Authorization: token ${GITHUB_TOKEN}" --header "Content-Type: application/json" --data "${PAYLOAD}" "${COMMENTS_URL}" > /dev/null
}

# module_download prepares depending modules
module_download() {
    # if not exist go.mod
    if [ ! -e go.mod ]; then
	go mod init
    fi

    # if finished in error status, exit 1
    go mod download
    if [ $? -ne 0 ]; then
	exit 1;
    fi
}

# run_gofmt runs gofmt and generate ${COMMENT} and ${SUCCESS}
run_gofmt() {
    set +e

    NG_FILE_LIST=$(sh -c "gofmt -l . $*" 2>&1)
    test -z "${NG_FILE_LIST}"
    SUCCESS=$?
    
    set -e
    # exit successfully
    if [ ${SUCCESS} -eq 0 ]; then
	return
    fi

    if [ "${SEND_COMMENT}" = "true" ]; then
	FMT_OUTPUT=""
	for file in ${NG_FILE_LIST}; do
	    mdfile=${file%.go*}.go
	    if [ -e ${mdfile} ]; then
		echo ${mdfile}
		
		# display all errors instead of rewriting file
		FILE_DIFF=$(gofmt -d -e "${mdfile}" | sed -n '/@@.*/,//{/@@.*/d;p}')
		FMT_OUTPUT="${FMT_OUTPUT}
<details><summary><code>${mdfile}</code></summary>

\`\`\`
${FILE_DIFF}
\`\`\`
</details>
	
"
	    fi
	done
	COMMENT="## gofmt failed
${FMT_OUTPUT}
"
    fi
}

# run_goimports runs goimports and generate ${COMMENT} and ${SUCCESS}
run_goimports() {
    set +e

    NG_FILE_LIST=$(sh -c "goimports -l . $*" 2>&1)
    test -z "${NG_FILE_LIST}"
    SUCCESS=$?

    set -e
    
    # exit successfully
    if [ ${SUCCESS} -eq 0 ]; then
	return
    fi

    if [ "${SEND_COMMENT}" = "true" ]; then
	FMT_OUTPUT=""
	for file in ${NG_FILE_LIST}; do
	    # display all errors instead of rewriting file
	    # and delete unnecessary zone
	    mdfile=${file%.go*}.go
	    if [ -e ${mdfile} ]; then
		echo ${mdfile}
		
		FILE_DIFF=$(goimports -d -e "${mdfile}" | sed -n '/@@.*/,//{/@@.*/d;p}')
		FMT_OUTPUT="${FMT_OUTPUT}
<details><summary><code>${mdfile}</code></summary>

\`\`\`				
${FILE_DIFF}
\`\`\`
</details>

"
	    fi
	done
	COMMENT="## goimports failed
${FMT_OUTPUT}
"
    fi
}

# run_golint for checking coding style
run_golint() {
    set +e

    OUTPUT=$(sh -c "golint -set_exit_status ./... $*" 2>&1)
    SUCCESS=$?
    
    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0 ]; then
	return
    fi

    if [ "${SEND_COMMENT}" = "true" ]; then
	COMMENT="## golint failed
$(echo "${OUTPUT}" | awk 'END{print}')
<details><summary>Show Detail</summary>

\`\`\`
$(echo "${OUTPUT}" | sed -e '$d')
\`\`\`
</details>
	
"
    fi
}

# run_gsc for static analysis
run_gsc() {
    set +e

    OUTPUT=$(sh -c "gsc ./... $*" 2>&1)
    SUCCESS=$?

    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0 ]; then
	return
    fi

    if [ "${SEND_COMMENT}" = "true" ]; then
	COMMENT="## gsc failed
<details><summary>Show Detail</summary>

\`\`\`
$(echo "${OUTPUT}" | sed -e '$d')
\`\`\`
</details>

"
    fi
}

# run_gosec executes gosec
run_gosec() {
    set +e
    
    gosec -out result.txt ./...
    SUCCESS=$?

    set -e

    if [ ${SUCCESS} -eq 0 ]; then
	return
    fi

    if [ "${SEND_COMMENT}" = "true" ]; then
	COMMENT="## gosec failed

\`\`\`
$(tail -n 6 result.txt)
\`\`\`

<details><summary>Show Detail</summary>

\`\`\`
$(head -n -7 result.txt)
\`\`\`
[Code Reference](https://github.com/securego/gosec#available-rules)

</details>
"
    fi
    rm result.txt
}

# staticcheck is like a go vet
run_staticcheck() {
    set +e

    OUTPUT=$(sh -c "staticcheck ./... $*" 2>&1)
    SUCCESS=$?

    set -e

    # exit successfully
    if [ ${SUCCESS} -eq 0 ]; then
	return
    fi

    if [ "${SEND_COMMENT}" = "true" ]; then
	COMMENT="## staticcheck failed
<details><summary>Show Detail</summary>

\`\`\`
$(echo "${OUTPUT}" | sed -e '$d')
\`\`\`
</details>

"
    fi
}

# -------------
# Main
# ------------
case ${RUN} in
    "fmt" )
	module_download
	run_gofmt
	;;
    "imports" )
	module_download
	run_goimports
	;;
    "lint" )
	module_download
	run_golint
	;;
    "gsc" )
	module_download
	run_gsc
	;;
    "sec" )
	module_download
	run_gosec
	;;
    "staticcheck" )
	module_download
	run_staticcheck
	;;
    * )
	echo "Invalid command."
	exit 1
esac

if [ ${SUCCESS} -ne 0 ]; then
    echo "Check failed."
    echo "${COMMENT}"
    if [ "${SEND_COMMENT}" = "true" ]; then
	send_comment
    fi
fi

exit ${SUCCESS}
