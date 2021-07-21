# terraform-provider-bridgecrew

You need to add an env var - BRIDGECREW_API or it won't work.
To build and run:

```bash
make check
```

This will build and install the provider locally, and run a test template.

If you're not using a Mac you will have to change OS_ARCH=darwin_amd64 to your value.

The example tf gets all the repositories you have in Bridgecrew and lists them.
