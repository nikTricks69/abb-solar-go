# abb-solar-go
An attempt at pushing ABB solar inverter feed to prometheus and grafana

```
$env:abb_user = 'user21'
# real password pls
$env:abb_password ='password'
$env:abb_ip ='192.168.80.172'
cd ABB-SOLAR-GO
go run .\abb-solar-go.go
...
...
...
http://localhost:8080/metrics TADAA


```

```
Iin1 4.325991630554199
Vin1 299.4452819824219
Pin1 1291.75048828125
Iin2 4.326968193054199
Vin2 331.92730712890625
Pin2 1444.34423828125
Pin 2726.287353515625
Igrid 10.671629905700684
Pgrid 2653.68505859375
Vgrid 247.66421508789062
```

# TODO

Nice to now have a docker image
and maybe https://pkg.go.dev/github.com/prometheus/client_golang/prometheus/promauto has its merits
