# install docker
snap install docker

# install newrelic statsd agent
docker run -d --restart unless-stopped --name newrelic-statsd -h $(hostname) -e NR_ACCOUNT_ID={{ NR_ACCOUNT_ID }} -e NR_API_KEY={{ NR_API_KEY }} -p 8125:8125/udp newrelic/nri-statsd:latest
