# Ethereum Swarm Proxy for SPAs

Frameworks such Vuejs, Angular etc. rely on absolute paths to resolve static assets.

### The Problem

When static assets(folders) are deployed to Swarm, a hash is produced and represents the entry point of the SPA.

```
http://localhost:8500/bzz:/f47eb1b09c76b55f85893502c46de5db5e38caeddcae1fb0eec2a98c8625cacb/
```

The problem is that your SPA tries to load resources from `http://localhost:8500/` and therefore not found by Swarm http server.


### One of the Solutions

```
go get github.com/iNDicat0r/ethereum-swarm-proxy
```

Compile, then run: 
```
./ethereum-swarm-proxy --swarm-hash f47eb1b09c76b55f85893502c46de5db5e38caeddcae1fb0eec2a98c8625cacb --local-port 8080 --swarm-ip localhost
```

Open in browser:
```
http://localhost:8080
```

### Use ENS

Use ENS and set the base url of your app to `swarm-gateways.net/bzz:/{ENS}`
