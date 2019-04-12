# oso-demo-2018

## docker

```
docker build . [-t tag]
docker run --rm -p 8080:8080 [--name name] <id/tag>
docker exec -it <id/name> bash
```

## openshift

```
oc create -f <template>
oc new-app <template>
```