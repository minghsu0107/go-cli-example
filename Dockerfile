
FROM plugins/base:multiarch

ADD /tmp/gce /bin/
ENTRYPOINT ["/bin/gce"]