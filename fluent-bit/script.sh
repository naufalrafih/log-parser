# install fluent-bit
curl https://raw.githubusercontent.com/fluent/fluent-bit/master/install.sh | sh

# install golang and make
apt-get install golang-go make

# install newrelic fluent-bit plugin
cd /tmp
git clone https://github.com/newrelic/newrelic-fluent-bit-output.git
cd newrelic-fluent-bit-output/
make linux/amd64
mv out_newrelic-linux-amd64-dev.so /etc/fluent-bit

# copy plugins config
cp plugins.conf /etc/fluent-bit/plugins.conf

# copy fluent-bit config
cp fluent-bit.conf /etc/fluent-bit/fluent-bit.conf

# start and enable fluent-bit
systemctl start fluent-bit
systemctl enable fluent-bit
