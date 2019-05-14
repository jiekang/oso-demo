# oso-demo

## docker

```
docker build . [-t tag]
docker run --rm -p 8080:8080 [--name name] <id/tag>
docker exec -it <id/name> bash
```

## openshift

```
oc create -f <template.json>
oc new-app <template-name>
```

## oc

```
oc cluster up --public-hostname <ip> --routing-suffix <ip>.nip.io
```

### Clean Up Workaround

```
for i in $(mount | grep openshift | awk '{ print $3}'); do sudo umount "$i"; done && sudo rm -rf ./openshift.local.clusterup
```
