# ip2country service 
Service receive IP address and return country and country code

## How to use with docker (dockerize service)
Download GeoIP2-Country.mmdb or GeoLite2-Country.mmdb file to ./GeoIP2-Country.mmdb,

Build:
```
docker build -t ip2country .
```
Run container: 
```
docker run -p 8080:8080 ip2country
```

