lke-ssh-key
===========

A small utility for Linode Kubernetes Engine beta testers. It grabs all SSH keys from Linode account and add them to `authorized_keys` of the Linode instances created by Linode Kubernetes Engine.

## Why

Typically we don't need to manually log into LKE-managed Linodes, `kubectl` can apply all our deployments. However some use cases do requires manual changes in the host nodes, such as using a NFS volume.

We want to automate this process with provisioning tools like Ansible or Terraform, so instead of manually reset root password of each nodes and then set up services, it's better to import SSH keys into all these nodes. As of Jan 2020, LKE-managed Linodes can't directly import SSH keys of Linode account (like newly created Linodes) yet, so I created this utility.

## Steps

1. Add your SSH public keys to [Linode profile](https://cloud.linode.com/profile/keys).

2. Get a Linode [Personal Access Token](https://cloud.linode.com/profile/tokens).

2. Download the binary:

```bash
go get github.com/mudkipme/lke-ssh-key
```

3. Run the binary:

```bash
LINODE_TOKEN=<your-linode-personal-access-token> lke-ssh-key
```

## Warning

This program will **restart** all your LKE-managed Linodes and **reset** the root password and authorized keys. As [linodego](https://github.com/linode/linodego) doesn't support LKE APIs yet, this program selects LKE-managed Linodes by matching labels with *"lke"* prefix and image names containing *"kube"*. This might result in unexpected behaviors.

## License

[MIT](LICENSE)
