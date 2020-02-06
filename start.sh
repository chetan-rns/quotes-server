oc new-project s2i-test

oc create secret generic ci-secret --from-file=.dockerconfigjson=${XDG_RUNTIME_DIR}/containers/auth.json --type=kubernetes/dockerconfigjson

oc secrets link pipeline ci-secret

oc apply -f pipelines/ci

oc apply -f pipelines/rbac/


oc apply -f pipelines/triggers/

# Add route in webhook task run

# Add Github token in webhook-secret
oc apply -f pipelines/webhook/