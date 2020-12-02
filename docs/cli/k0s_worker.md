## k0s worker

Run worker

```
k0s worker [join-token] [flags]
```

### Examples

```
	Command to add worker node to the master node:
	CLI argument:
	$ k0s worker [token]

	or CLI flag:
	$ k0s worker --token-file [path_to_file]
	Note: Token can be passed either as a CLI argument or as a flag
```

### Options

```
      --cri-socket string       contrainer runtime socket to use, default to internal containerd. Format: [remote|docker]:[path-to-socket]
      --enable-cloud-provider   Whether or not to enable cloud provider support in kubelet
  -h, --help                    help for worker
      --profile string          worker profile to use on the node (default "default")
      --token-file string       Path to the file containing token.
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

