# log-parser
Log and metric ingestion using fluent-bit and go-statsd to New Relic.
1. Install required dependencies using script.sh in fluent-bit and statsd folders. Make sure the New Relic API key (license key) is already defined.
2. Define input and output log files in script/main.go.
3. Run script/main.go. The output file will be created and have JSON formatted logs. The script will also send relevant metrics to New Relic.
4. Fluent-bit will consume the output file. Make sure to set the correct input path in fluent-bit.conf.
5. Verify the log and metric results in New Relic.

Further Documentation: https://docs.google.com/document/d/1oceGtIJEaaG_QDFa1o9Mb9OyoY2roq6T_LXbwsErmeQ/edit?usp=sharing
