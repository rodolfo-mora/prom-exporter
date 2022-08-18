# prom-exporter

Simple example of a web API service that receives a payload in key:value format.

Example.
```
{
  "name": "some_text"
}
```

And translates such information into a Prometheus gauge with label Hostcheck.
