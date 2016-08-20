# systemctl start influxdb
# if no INITSYSTEM
service influxdb start
modprobe i2c-dev
# we use exec so SIGTERM is propagated correctly to goapp
exec ./goapp.o
