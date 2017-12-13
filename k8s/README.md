# diff-demo

```console
$ go install -v

$ diff-demo diff -s ./examples/voyager/src.yaml -d ./examples/voyager/on-master.yaml
spec:
  template:
    spec:
      nodeSelector:
        node-role.kubernetes.io/master: ""
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoSchedule
        key: node-role.kubernetes.io/master
        operator: Exists



$ diff-demo diff -s ./examples/voyager/src.yaml -d ./examples/voyager/custom-tpl.yaml

spec:
  template:
    spec:
      $setElementOrder/containers:
      - name: voyager
      $setElementOrder/volumes:
      - name: cloudconfig
      - name: templates
      containers:
      - $setElementOrder/volumeMounts:
        - mountPath: /etc/kubernetes
        - mountPath: /srv/voyager/custom
        args:
        - run
        - --v=3
        - --cloud-provider=$CLOUD_PROVIDER
        - --cloud-config=$CLOUD_CONFIG
        - --ingress-class=$INGRESS_CLASS
        - --custom-templates=/srv/voyager/custom/*.cfg
        name: voyager
        volumeMounts:
        - mountPath: /srv/voyager/custom
          name: templates
          readOnly: true
      volumes:
      - configMap:
          name: voyager-templates
        name: templates
```

