echo $INFLUXDB_DATA_DIR
mkdir -p /data/influx
echo "Starting InfluxDB..."
systemctl start influxdb
# if no INITSYSTEM
#service influxdb start
modprobe i2c-dev

chmod +x goapp.o
# we use exec so SIGTERM is propagated correctly to goapp
exec ./goapp.o
