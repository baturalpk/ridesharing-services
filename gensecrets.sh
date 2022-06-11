#!/bin/bash

echo ❓ Please, insert your MongoDB® URI:
read MONGO_URI

echo ""
echo ❓ Please, insert your Redis® URI:
read REDIS_URI

cat > manifests/secrets.yml << EOL
apiVersion: v1
kind: Secret
metadata:
  name: tripsecret
type: Opaque
data:
  mongo_uri: $(echo -n $MONGO_URI | base64)
  redis_uri: $(echo -n $REDIS_URI | base64)
EOL

echo ""
echo ⚙️ Generated secrets: '../manifests/secrets.yml'
