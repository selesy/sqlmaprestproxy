# sqlmap-restproxy

Proxies REST server calls so they appear to be form data and wild-carded
URLs to work with SQLMap.  This project is an experiment based on a
feature request submitted to the SQLMap project - see
<https://github.com/sqlmapproject/sqlmap/issues/3140#issuecomment-578484549.>

## Operation

The basic idea is to create a sitemap.xml document that can be read by
SQLMap to determine what REST resources are available, what HTTP verbs
should be used for each resource and, assuming input should be sent,
convert the request JSON document to an HTTP form so that SQLMap can
attempt SQL injection.
