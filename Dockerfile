FROM scratch
MAINTAINER Stephen Price <stephen@stp5.net>
ADD instarss instarss
EXPOSE 80
ENTRYPOINT ["/instarss"]
