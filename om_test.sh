#! /bin/bash -x

om -u pivotalcf -p snowflakes-throwdown -k -t https://pcf.monkeyface.sandbox.releng.cf-app.com configure-director -i '{"iaas_confiation": { "project": "some-wrong-project", "default_deployment_tag": "my-vms","auth_json": "{\"some-auth-field\": \"some-service-key\",\"some-private_key\": \"some-key\"}"}}'
