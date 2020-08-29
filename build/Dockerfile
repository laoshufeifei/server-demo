FROM docker.navicore.cn/ncbase/ubuntu:1.0.0

COPY _build/.	 	/etc/ncserver/

ENTRYPOINT ["/tini", "--"]
CMD ["/bin/bash", "-c", "/etc/ncserver/start.sh"]
