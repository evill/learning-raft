#!/bin/bash
set -e

# Set variables
BIN_DIR="bin"
OPENAPI_VERSION="7.11.0"
JAR_NAME="openapi-generator-cli-${OPENAPI_VERSION}.jar"
DOWNLOAD_URL="https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/${OPENAPI_VERSION}/${JAR_NAME}"
WRAPPER_NAME="openapi-generator"

# Create tools directory if it doesn't exist
mkdir -p "${BIN_DIR}"

# Download JAR if it doesn't exist
if [ ! -f "${BIN_DIR}/${JAR_NAME}" ]; then
    echo "Downloading OpenAPI Generator ${OPENAPI_VERSION}..."
    curl -L "${DOWNLOAD_URL}" -o "${BIN_DIR}/${JAR_NAME}"
    echo "Download complete"
fi

# Create wrapper script
cat > "${BIN_DIR}/${WRAPPER_NAME}" << EOF
#!/bin/bash
java -jar \$(dirname \$0)/${JAR_NAME} "\$@"
EOF

# Make wrapper executable
chmod +x "${BIN_DIR}/${WRAPPER_NAME}"

echo "OpenAPI Generator ${OPENAPI_VERSION} installed successfully"
echo "You can now use it with: ${BIN_DIR}/${WRAPPER_NAME}"
