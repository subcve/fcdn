# fcdn

Fcdn is just a simple go script for reading ips from stdin and exclude ips which belongs to the cdns (cloudflare, Akamai and etc.)

# Install
for installing fcdn, just hit this command in your terminal
```
go install github.com/subcve/fcdn@latest
```
# Usage

Single IP:
```
echo "ip" | fcdn
```

List of IP's:
```
cat ips | fcdn
```
