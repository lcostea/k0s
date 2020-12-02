## k0s etcd leave

Sign off a given etc node from etcd cluster

```
k0s etcd leave [flags]
```

### Options

```
  -h, --help                  help for leave
      --peer-address string   etcd peer address
```

### Options inherited from parent commands

```
  -c, --config string            config file (default: ./k0s.yaml)
      --data-dir string          Data Directory for k0s (default: /var/lib/k0s). DO NOT CHANGE for an existing setup, things will break!
  -d, --debug                    Debug logging (default: false)
  -l, --logging stringToString   Logging Levels for the different components (default [kube-apiserver=1,kube-controller-manager=1,kube-scheduler=1,kubelet=1,etcd=info,containerd=info,konnectivity-server=1])
```

### SEE ALSO

* [k0s etcd](k0s_etcd.md)	 - Manage etcd cluster

