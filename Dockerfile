
FROM plugins/base:multiarch

ADD /root/gce /bin/
ENTRYPOINT ["/bin/gce"]