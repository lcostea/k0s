## k0s etcd

Manage etcd cluster

### Options

```
  -h, --help   help for etcd
```

### Options inherited from parent commands

```
  -c, --config string            config file (default: ./k0s.yaml)
      --data-dir string          Data Directory for k0s (default: /var/lib/k0s). DO NOT CHANGE for an existing setup, things will break!
  -d, --debug                    Debug logging (default: false)
  -l, --logging stringToString   Logging Levels for the different components (default [kube-apiserver=1,kube-controller-manager=1,kube-scheduler=1,kubelet=1,etcd=info,containerd=info,konnectivity-server=1])
```

### SEE ALSO

* [k0s](k0s.md)	 - k0s - Zero Friction Kubernetes
* [k0s etcd leave](k0s_etcd_leave.md)	 - Sign off a given etc node from etcd cluster
* [k0s etcd member-list](k0s_etcd_member-list.md)	 - Returns etcd cluster members list

